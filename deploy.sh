#!/bin/bash
path=`pwd`

i=$(docker images | grep "demo-server" | awk '{print $1}')
if test -z $i; then
echo "not found the docker image, start build image..."
docker build -f ./DockerFile -t demo-server:v0.9.0 ../go-web-demo
fi

i=$(docker images | grep "demo-server" | awk '{print $1}')
if test -z $i; then
echo "build image error, exit shell!"
exit
fi

c=$(docker ps -a | grep "demo-mysql" | awk '{print $1}')
if test -z $c; then
echo "not found the mysql server, start mysql server..."
docker run -d \
    -p 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=123456 \
    -e MYSQL_DATABASE=demo \
    --name demo-mysql \
    --restart always \
    mysql:8.0
echo "waiting for database initialization..."
sleep 20s
docker logs --tail=10 demo-mysql
fi

i=$(docker ps -a | grep "demo-server" | awk '{print $1}')
if test ! -z $i; then
echo "the server container already exists, delete..."
docker rm -f demo-server
fi

echo "start demo-server..."
docker run -d \
-p 8086:8086 \
-w /go-web-demo \
-v $path/conf:/go-web-demo/conf \
-v $path/log:/go-web-demo/log \
--name demo-server \
--restart always \
demo-server:v0.9.0 \
bash -c "cd src&&./demo -config ../conf/system_config.yaml"
sleep 2s
docker logs demo-server
echo "the demo-server has been started!"


i=$(docker ps -a | grep "demo-web" | awk '{print $1}')
if test ! -z $i; then
echo "the web container already exists, delete..."
docker rm -f demo-web
fi

echo "start demo web server..."
chmod -R 777 $path/web/
docker run -d \
-p 80:80 \
-v $path/web/conf.d:/etc/nginx/conf.d \
-v $path/web/resources:/usr/share/nginx/resources \
--name demo-web \
--restart always \
nginx:1.23.3
echo "the demo web server has been started!"