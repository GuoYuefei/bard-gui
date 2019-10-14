:: cmd /c go build cmd/main.go && start bard/hideexec.exe main.exe

cmd /c
if not exist main.exe (
    go build cmd/main.go
)

start bard/hideexec.exe main.exe