# bit-oom
Project for OOM course in BIT.

## Explanation
本项目是北京理工大学研究生课程《面向对象技术与方法》课程的课程作业项目，主题为"设计与开发一个网络应用程序"，选择的课题为设计与开发一个简单的文件上传与下载程序。

## Documentation
项目设计与开发文档详见`docs`目录下相关说明。

## Setup
本项目包含的程序比较简单，使用的语言为 Golang，后端程序的主体位于`main.go`，前端页面位于`ugc`目录下。本项目引用的开源库仅包含日志框架 [logrus](github.com/sirupsen/logrus)，其余部分仅包含 Go 的官方包。自行编译请按照如下步骤进行：
```shell script
# 依赖下载
$ go mod vendor

# MacOS
$ make darwin

# Linux
$ make linux

# Windows
$ make windows
```
也可直接使用`bin`目录下的已经编译好的可执行文件：
```shell script
$ tree -L 1 bin
bin
├── main        # for Linux 
├── main-mac    # for MacOS
└── main.exe    # for Windows
```

## Contribution
- [logrus](github.com/sirupsen/logrus)

