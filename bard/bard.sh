#!/bin/sh


function status_bard() {
  noecho=''
  if [ $# -ne 0 ]
  then
    noecho=$1
  fi

  bard_pid=`cat run.pid`
  #没有run.pid这个文件就认为没有开启程序
  if [ $? = 1 ]
  then
    return 0        # cat执行粗无， bard理论上一次也运行
  fi

  if [ `ps ${bard_pid} | grep -c ${bard_pid}` -ne 0 ]
    then
      if [ $noecho = 'no-echo' ]
      then
        return 1
      fi
      echo "程序运行中 (Program is running)"
      return 1          # bard运行中
    else
      if [ $noecho = 'no-echo' ]
      then
        return 0
      fi
      echo "程序没有运行 (The program is not running)"
      return 0          # bard没有运行
  fi
}

## 开启bard程序
# shellcheck disable=SC2112
function start_bard() {
  # 先判定run.pid记录的服务是否开启
  status_bard no-echo
  flag=$?
  if [ $flag -eq 1 ]
  then
    echo "[1]程序在此之前已运行 (The program was already running before that.)"
  else
    nohup ./bard-client > ./nohup.out 2>&1 & echo $! > ./run.pid
    status_bard no-echo
    flag=$?
    if [ $flag -eq 1 ]
    then
      echo "[0]程序开启成功 (Successful Opening of Program)"
    else
      echo "[1]程序开启失败 (Program failed to open)"
    fi
  fi
}

# 关闭bard程序
function stop_bard() {
  cat run.pid | xargs kill > /dev/null 2>&1
  if [ $? -eq 0 ]
  then
    echo "[0]程序关闭 (Program shutdown)"
  else
    echo "[1]程序并未运行 (The program shutdown program did not run)"
  fi
}

# 重启
function restart_bard() {
  # stop提示内容不显示
  stop_bard > /dev/null 2>&1
  start_bard
}

# 安装bard程序 不安装依赖环境
function install_bard() {
  # 不需要config
  ./install-bard.sh $1
}

## 主程序
case $1 in
  "start")
    start_bard
  ;;
  "stop")
    stop_bard
  ;;
  "restart")
    restart_bard
  ;;
  "status")
    status_bard
  ;;
  "install")
    install_bard "noconfig"
  ;;
  *)
    echo "[1]输入指令错误"
esac

