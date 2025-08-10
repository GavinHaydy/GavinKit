@echo off
REM 编译当前目录下的 main.go，生成 GavinKit.exe，无控制台窗口
go build -ldflags="-H=windowsgui" -o GavinKit.exe main.go

echo 编译完成，生成 GavinKit.exe（无控制台窗口）
pause
