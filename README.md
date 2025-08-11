# GavinKit
[![Go Version](https://img.shields.io/github/go-mod/go-version/GavinHaydy/GavinKit)](https://go.dev/)
[![License](https://img.shields.io/github/license/GavinHaydy/GavinKit)](LICENSE)
[![Release](https://img.shields.io/github/v/release/GavinHaydy/GavinKit)](https://github.com/GavinHaydy/GavinKit/releases)
[![Stars](https://img.shields.io/github/stars/GavinHaydy/GavinKit?style=social)](https://github.com/GavinHaydy/GavinKit/stargazers)
[![Forks](https://img.shields.io/github/forks/GavinHaydy/GavinKit?style=social)](https://github.com/GavinHaydy/GavinKit/network/members)
[![Contributors](https://img.shields.io/github/contributors/GavinHaydy/GavinKit)](https://github.com/GavinHaydy/GavinKit/graphs/contributors)

[中文](README_zh.md)

## Introduction
This is a cross-platform desktop application developed using [Fyne](https://fyne.io), supporting both Windows and Linux.

## Requirements
### Windows
- Windows 10 or later (64-bit recommended)
- Go 1.18 or later
- GCC or MinGW (for Cgo support)
- Git (for source code management)

#### Environment Setup
- Install [MSYS2](https://www.msys2.org/) following the official instructions.
```shell
    # Open the MSYS terminal and update packages
    pacman -Syu
    # After the update is complete, restart MSYS2 and run:
    pacman -S --needed base-devel mingw-w64-x86_64-toolchain mingw-w64-x86_64-go
    # Once installation is done, add the installation path to your system environment variables:
    path:\msys64\mingw64\bin  (e.g. E:\tools\msys64\mingw64\bin)
    # Verify installation by opening any terminal and running:
    gcc --version
```

### Linux
- A modern Linux distribution (e.g., Ubuntu 20.04+, Fedora 36+, etc.)
- Go 1.18 or later
- GCC and related build tools (such as `build-essential`)
- Git

#### Environment Setup (Ubuntu)
For other distributions, please refer to the [Fyne documentation](https://docs.fyne.io/started).

- Install GCC and other build dependencies:
```shell
sudo apt install gcc libgl1-mesa-dev xorg-dev libxkbcommon-dev
```

### Run and Build
- Install [Go](https://go.dev/doc/install) following the official instructions.
```shell
# Install dependencies
go mod tidy
# Run
go run .
# Build
# Windows
build.bat
# Linux
go build .
```
