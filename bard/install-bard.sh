#!/bin/sh

trap "echo 执行终止" INT

# 测试时进入这个文件夹测试
mkdir bard
cd bard

# 前提需要装git和golang的编译环境 该脚本不支持该操作
# 这个文件只适用于mac和linux

set -e

# 获取源码
git clone https://github.com/GuoYuefei/bard.git
cd bard 
git checkout v0.4.4
git submodule init
git submodule update

# 编译主程序
go build -o bard-client client/client.go
mv bard-client ../
cp -rf client ../

# 编译并放置默认插件
cd bard-plugin
go build -buildmode=plugin -o get_des_cfb_client.so base/client/get_des_cfb_client.go
mv get_des_cfb_client.so ../../client/debug/plugins/

# 删除源码文件
cd ../..
rm -rf bard

# 若携带第一个参数为config 让用户编写配置文件
if [ $"$1" -a "$1" = "config" ]
then
  vim client/debug/config/config.yml
fi

echo "安装成功 (Successful installation)"

exit 0

