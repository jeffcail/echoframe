#! /bin/bash

echo "=====>>>>> 准备安装并启动 mongodb " && sleep 1 && echo ""

docker-compose -f scripts/env/mongo.yml up -d

sleep 3 && docker ps && echo ""

echo "=====>>>>> mongodb启动成功 "
