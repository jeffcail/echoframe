#!/bin/bash

# 检测操作系统
OS=$(cat /etc/os-release | grep ^ID= | cut -d'=' -f2 | tr -d '"')

# 百度镜像
DOCKER_MIRROR="https://mirror.baidu.com/registry-mirrors"

install_docker_ubuntu() {
    echo "正在安装 Docker 及 Docker Compose (Ubuntu)..."

    # 更新包索引
    sudo apt-get update

    # 安装必要的依赖包
    sudo apt-get install -y \
        apt-transport-https \
        ca-certificates \
        curl \
        software-properties-common

    # 添加 Docker 的官方 GPG 密钥
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

    # 添加 Docker 仓库
    sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

    # 更新包索引并安装 Docker CE
    sudo apt-get update
    sudo apt-get install -y docker-ce

    # 设置 Docker 使用百度镜像加速
    sudo mkdir -p /etc/docker
    echo "{ \"registry-mirrors\": [\"$DOCKER_MIRROR\"] }" | sudo tee /etc/docker/daemon.json

    # 启动 Docker 并设置开机自启
    sudo systemctl start docker
    sudo systemctl enable docker

    # 安装 Docker Compose
    sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose

    # 验证安装
    sudo docker run hello-world
    docker-compose --version

    echo "Docker 和 Docker Compose 安装完成!"
}

install_docker_centos() {
    echo "正在安装 Docker 及 Docker Compose (CentOS)..."

    # 安装必要的依赖包
    sudo yum install -y yum-utils

    # 添加 Docker 的官方仓库
    sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

    # 安装 Docker CE
    sudo yum install -y docker-ce docker-ce-cli containerd.io

    # 设置 Docker 使用百度镜像加速
    sudo mkdir -p /etc/docker
    echo "{ \"registry-mirrors\": [\"$DOCKER_MIRROR\"] }" | sudo tee /etc/docker/daemon.json

    # 启动 Docker 并设置开机自启
    sudo systemctl start docker
    sudo systemctl enable docker

    # 安装 Docker Compose
    sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    sudo chmod +x /usr/local/bin/docker-compose

    # 验证安装
    sudo docker run hello-world
    docker-compose --version

    echo "Docker 和 Docker Compose 安装完成!"
}

# 根据操作系统类型安装 Docker
case $OS in
    "ubuntu")
        install_docker_ubuntu
        ;;
    "centos")
        install_docker_centos
        ;;
    *)
        echo "不支持的操作系统: $OS"
        exit 1
        ;;
esac
