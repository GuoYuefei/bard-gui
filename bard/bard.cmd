@echo off

call:main %1

goto end

:main
	:: 转换成utf-8字符集
	chcp 65001 > nul
	cd /d %~dp0

	if "%1" == "test" (
	    call:test
	)
	if "%1" == "status" (
		call:status_bard
	)
	if "%1" == "start" (
		call:start_bard
	)
	if "%1" == "stop" (
	    call:stop_bard
	)
	if "%1" == "restart" (
	    call:restart_bard
	)
	if "%1" == "install" (
	    call:install_bard
	)
goto:eof
	

:status_bard
	:: windows 上我们不用pid控制，我实在不知道怎么在程序执行的时候获取到pid,所以直接用镜像名控制
    set running=0
	:: 同tasklist找到bard-client.exe是否在运行
	tasklist /fi "IMAGENAME eq bard-client.exe" | find /c "bard-client.exe" > temp && set /p running=<temp

	if not "%running%"=="0" (
	    if not "%1" == "no-echo" (
            set /p t=[0]程序已开启 ^(Program is running^) < nul
	    )
	    :: 返回1表示在运行
	    exit /b 1
	) else (
        if not "%1" == "no-echo" (
    	    set /p t=[1]程序没有运行 ^(The program is not running^) < nul
    	)
    	exit /b 0
	)
goto:eof
	
:start_bard
	call:status_bard no-echo
	if "%errorlevel%" == "1" (
		echo [1]程序在此之前已运行 ^(The program was already running before that.^)
	) else (
		:: 证明没运行呢，可以运行此程序 windows下利用hideexec工具来实现后台运行，并且无法捕捉输出
		cmd /c start hideexec.exe bard-client.exe

		:: 这一部分在golang程序中其实是执行不到的，所以需要在golang中手动调用status
        call:status_bard no-echo
        if errorlevel 1 (
        :suss
            set /p = [0]程序开启成功 ^(Successful Opening of Program^) < nul
        ) else (
            set /p = [1]程序开启失败 ^(Program failed to open^) < nul
        )
	)
	exit /b 0
goto:eof

:stop_bard
    call:status_bard no-echo
    if %errorlevel% == 0 (
        set /p = [1]程序并未运行 ^(The program shutdown program did not run^) < nul
        exit /b 0
    )
    taskkill /im bard-client.exe /f > nul

    call:status_bard no-echo
    if errorlevel 1 (
        set /p = [1]程序关闭失败 (Program shutdown error) < nul
    ) else (
        set /p = [0]程序关闭 (Program shutdown) < nul
    )

    exit /b 0
goto:eof

:restart_bard
    call:stop_bard > nul 2>&1
    call:start_bard
    exit /b 0
goto:eof

:: 安装bard程序 不安装依赖环境
:install_bard
    call install_bard.cmd %1
    exit /b 0
goto:eof

:test
    echo this is test
    exit /b 0
goto:eof
	
:end
