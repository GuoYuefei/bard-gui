#!/bin/sh

go build cmd/main.go

# 不论之前有没有运行，先杀gui程序再启动以保证只有一个gui跑起来，也防止记录错误pid
cat run.pid | xargs kill > /dev/null 2>&1
nohup ./main > ./nohup.out 2>&1 & echo $! > ./run.pid