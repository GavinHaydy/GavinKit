# GavinKit
[![Go 版本](https://img.shields.io/github/go-mod/go-version/GavinHaydy/GavinKit)](https://go.dev/)
[![许可](https://img.shields.io/github/license/GavinHaydy/GavinKit)](LICENSE)
[![发布版本](https://img.shields.io/github/v/release/GavinHaydy/GavinKit)](https://github.com/GavinHaydy/GavinKit/releases)
[![Star 数](https://img.shields.io/github/stars/GavinHaydy/GavinKit?style=social)](https://github.com/GavinHaydy/GavinKit/stargazers)
[![Fork 数](https://img.shields.io/github/forks/GavinHaydy/GavinKit?style=social)](https://github.com/GavinHaydy/GavinKit/network/members)
[![贡献者](https://img.shields.io/github/contributors/GavinHaydy/GavinKit)](https://github.com/GavinHaydy/GavinKit/graphs/contributors)

[English](README.md)

## 简介
这是一个使用 [Fyne](https://fyne.io) 开发的跨平台桌面应用程序，支持 Windows 和 Linux。

## 环境要求
### Windows
- Windows 10 或更高版本（推荐 64 位）
- Go 1.18 或更高版本
- GCC 或 MinGW（用于 Cgo 支持）
- Git（用于管理源代码）
#### 环境安装
 - [MSYS](https://www.msys2.org/)根据说明下载安装即可
 ```shell
    # 打开MSYS终端输入命令更新
    pacman -Syu
    # 等待更新完成后重启 MSYS2 再执行  
    pacman -S --meeded base-devel mingw-w64-x86_64-toolchain mingw-w64-x86_64-go
    # 等待安装完成后 将你的安装路径添加到环境变量中
    path:\msys64\mingw64\bin 例如 E:\tools\msys64\mingw64\bin
    # 验证 新开任意命令行输入以下命令，如果有版本信息则成功
    gcc --version
 ```

### Linux
- 较新的 Linux 发行版（Ubuntu 20.04+、Fedora 36+ 等）
- Go 1.18 或更高版本
- GCC 及相关构建工具（如 `build-essential`）
- Git
#### 环境安装(ubuntu),其他发行版请查看[fyne文档](https://docs.fyne.io/started)
- gcc等构建工具
```shell
sudo apt install gcc libgl1-mesa-dev xorg-dev libxkbcommon-dev
```

### 运行及编译
- [go](https://go.dev/doc/install)根据说明下载安装即可
 ```shell
# 依赖安装
 go mod tidy
 # 运行
 go run .
 # 编译
 # windows
 build.bat
 # linux
 go build . 
 ```