# go-web-demo

## docker 一键部署

### 在线环境准备

- docker

### 离线环境依赖

- docker离线安装包
- mysql: 8.0 镜像包
- nginx: 1.23.3 镜像包
- demo 镜像包



### 部署前必要配置修改

#### 数据库配置

配置目录：conf/system_config.yaml

配置内容：

```yaml
db_config:
  user: root
  password: 123456
  ip: 172.16.2.225
  port: 13306
  dbname: demo
```

- user：数据库用户名
- password：数据库密码
- ip：数据库IP地址
- port：数据库端口
- dbname：数据库名称



#### Nginx配置文件

配置目录：web/conf.d/nginx.conf

配置内容：

```yaml
server {
    listen       80;
    listen  [::]:80;
    server_name  _;
    #access_log  /var/log/nginx/host.access.log  main;
    location / {
        root   /usr/share/nginx/resources/;
        index  index.html index.htm;
    }
    location /satellitebc/ {
      proxy_pass http://172.16.2.225:8086/;
    }
    #error_page  404              /404.html;
    # redirect server error pages to the static page /50x.html
    #
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }
}
```

- proxy_pass：demo服务的地址和端口

### 使用部署脚本启动

在项目主目录下，运行部署脚本：

```shell
$ ./deploy.sh
```











