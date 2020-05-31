# 项目说明文档

## 路由
1. `/`: 根路径，用于渲染`static/index.html`；
2. `/upload`: 用户上传路径；
3. `/download`: 用户文件下载路径，文件名通过`file`参数传入，如`http://localhost:8080/download?file=test.txt`；

## 项目的核心功能
- 客户端可以上传文件到服务器上
- 客户端可以获取文件列表，从中选择文件下载
