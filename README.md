- [Getting Started](#getting-started)
  + [Install](#install)
  + [Building the Application](#building-the-application)
  + [Usage](#usage)
- [License](#license)
- [Preface](#preface)

## Getting Started

### Install

```shell
go get -u install github.com/GuoYuefei/bard-gui
```

或者

```git
git clone https://github.com/GuoYuefei/bard-gui.git
```

### Building the Application

在运行程序之前，你应当确保有**go1.11以上的编译环境**和**git**工具。

在满足条件的情况下，可以直接执行bard-gui文件夹下的main.* (mac和linux为main.sh， windows为main.cmd).

#### Mac and Linux

ps. Linux未测试

```shell
cd bard-gui
# 首先确保main.sh和main_stop.sh有运行权限， ！！如果没有请执行_一次_以下命令
chmod +x main.sh main_stop.sh

# 运行程序   ps. 运行结束可关闭命令框
./main.sh		
# 退出程序 	ps. 同时推出gui和bard程序
./main_stop.sh
```



#### Windows X

ps. 未测试其他环境，照理说只要支持cmd的windows都可以

**开启**：直接双击main.cmd文件直接运行

**关闭**： 直接双击main_stop.cmd



### Usage

程序开启后会自动打开游览器进入localhost:2019, 在本页面配置bard的连接信息。保存后点击右侧开关即可进行连接。

由于未实现透明代理，所以现在只能在游览器中手动设置代理，推荐 SwitchyOmega。firefox和chrome都有相应插件。。由于现在chrome只能从商店下载插件使用，同时商店在墙外。所以推荐使用firefox. 如果要使用chrome可以通过启动开发者模式，从github或者官网下载SwitchyOmega后进行安装，然后在进入商店安装。

## License

GNU AFFERO GENERAL PUBLIC LICENSE (AGPLv3.0)

## Preface

未来会在windows上先实现透明代理。