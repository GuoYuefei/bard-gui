#!/bin/sh

case $1 in
  "start")
    nohup ./bard-client > ./nohup.out 2>&1 & echo $! > ./run.pid
  ;;
  "stop")
    cat run.pid | xargs kill
  ;;
  *)
    echo "输入指令错误s"
esac