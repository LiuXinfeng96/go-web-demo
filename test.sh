#!/bin/bash

loginJSON="{\"userName\":\"exec\",\"userPwd\":\"1234\"}"
tokenJSON=$(curl -H "Content-Type: application/json" \
-d "${loginJSON}" \
http://localhost:8086/satellitebc/user/login)

token=$(echo "${tokenJSON}" | jq -r '.data.token')

i=1
while true 
do
debrisId=$(uuidgen | sed 's/-//g')
debrisName="碎片${i}"
angle=`expr $RANDOM / 10`
speed=`expr $RANDOM / 10`
height=`expr $RANDOM / 10`
volume=`expr $RANDOM / 10`
type="Small"

JSON="{\"debrisId\":\"${debrisId}\",\"debrisName\":\"${debrisName}\",\"angle\":$angle,\"speed\":$speed,\"height\":$height,\"volume\":$volume,\"type\":\"${type}\"}"
curl -H "Content-Type: application/json" \
-H "token: ${token}" \
-d "${JSON}" \
http://localhost:8086/satellitebc/control/adddebris

i=`expr $i + 1`
sleep 10s
done

	