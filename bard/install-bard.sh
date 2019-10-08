#!/bin/sh


# 本脚本的运行前提： 需要装git和golang的编译环境
# 这个文件只适用于mac和linux

trap "printf [1]执行终止" INT

# 仅测试时运行
function test_dir() {
  mkdir bard
  cd bard
}

function clone() {
  git clone https://github.com/GuoYuefei/bard.git
  cd bard
  git checkout v0.4.4
  git submodule init
  git submodule update
}

function compile() {
  # 编译主程序
  go build -o bard-client client/client.go
  mv bard-client ../
  cp -rf client ../

  # 编译并放置默认插件
  cd bard-plugin
  go build -buildmode=plugin -o get_des_cfb_client.so base/client/get_des_cfb_client.go
  mv get_des_cfb_client.so ../../client/debug/plugins/
}

function delete_src() {
  # 删除源码文件
  cd ../..
  rm -rf bard
}

function config() {
  if [ "$1" != "noconfig" ]
  then
    vim client/debug/config/config.yml
  fi
}

function main() {
  # 测试时进入这个文件夹测试
  test_dir

  # 这个错误不知道怎么trap，不输出。。。go程序中只要判定是否为[0]就行了
  set -e

  # 获取源码
  clone

  # 编译主程序
  # 编译并放置默认插件
  compile

  # 删除源码文件
  delete_src

  # 若携带第一个参数为noconfig 让用户编写配置文件   $"$1" -a "$1" = "config"
  config $1
}

# 如果不需要配置那么就不在终端中显示安装过程
if [ $"$1" -a "$1" = "noconfig" ]
then
  main $1 > /dev/null 2>&1
else
  main $1
fi

printf "[0]安装成功 (Successful installation)"

exit 0

