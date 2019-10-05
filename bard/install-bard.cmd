@echo off

if "%1" == "noconfig" (
	call:main %1 > nul 2>&1
) else (
	call:main %1
)

echo [0]安装成功 (Successful installation)
pause
goto end


:test_dir
	mkdir bard
	cd bard
	exit /b 0
goto:eof

:clone
	git clone https://github.com/GuoYuefei/bard.git
	cd bard
	git checkout v0.4.4
	git submodule init
	git submodule update
	exit /b 0
goto:eof

:compile
	:: 编译主程序
	go build -o bard-client client/client.go
	move /y bard-client ..
	:: cp -rf client ..
	:: shell同上作用
	echo D | xcopy client ..\client /e /y
	
	exit /b 0
goto:eof
	
:delete_src
	:: 删除源码文件
	cd ..
	
	:: rd 删除目录 
	rd bard /S /Q
	exit /b 0
goto:eof

:config
	:: 编辑配置文件
	if not "%1" == "noconfig" (
		:: notepad记事本命令
		start /wait notepad client\debug\config\config.yml
	)
	exit /b 0
goto:eof

:main
	:: 转换成utf-8字符集
	chcp 65001
	::call:test_dir
	
	call:clone
	
	call:compile
	
	call:delete_src
	
	call:config %1
	
	exit /b 0
goto:eof

:end

	
	
	
	