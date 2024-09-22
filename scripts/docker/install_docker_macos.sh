#!/bin/bash

# 检查是否已经安装 Homebrew
if ! command -v brew &> /dev/null; then
    echo "Homebrew 未安装，正在安装..."
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
fi

# 更新 Homebrew
brew update

# 安装 Docker Desktop
brew install --cask docker

# 提示用户手动启动 Docker
echo "请手动启动 Docker Desktop 应用程序!"

# 验证安装
# 此步骤在 Docker 启动后可以运行
echo "请在 Docker 启动后运行 'docker run hello-world' 验证安装."

echo "Docker 安装完成!"
