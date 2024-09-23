#! /bin/bash

echo "=====>>>>> 准备安装并启动 mysql " && sleep 1 && echo ""
docker-compose -f scripts/env/mysql.yml up -d

sleep 3 && docker ps && echo ""

echo "=====>>>>> mysql启动成功 "
echo "=====>>>>> 浏览器访问8880端口，数据库密码请查询 mysql.yml " && sleep 2


sleep 3 && docker ps && echo ""

echo "=====>>>>> 准备安装并启动 redis " && sleep 1 && echo ""

docker-compose -f scripts/env/redis.yml up -d

echo "=====>>>>> redis启动成功" && sleep 1 && echo ""

echo "=====>>>>> 准备安装并启动 nacos " && sleep 1 && echo ""

docker-compose -f scripts/env/nacos.yml up -d

echo "=====>>>>> nacos启动成功" && sleep 1 && echo ""
echo "=====>>>>> 浏览器访问8848端口，用户和密码：nacos/nacos " && sleep 2