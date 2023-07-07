# hugoArticleEncryptor
[English](https://github.com/hotjuicew/hugoArticleEncryptor/blob/master/README.md) | 简体中文

hugoArticleEncryptor是一个hugo文章加密工具，使用AES-GCM算法对整个hugo文章进行加密，将javascript代码插入到加密的文章中，在用户输入正确的口令之后解密内容。
配置非常简单。并且当你输入过一次正确密码后，下次访问加密页面就不需重复输入密码，会直接呈现解密后的内容
## [DEMO](https://juicebar-demo.add1.dev/)
## 安装与使用
### Option A: 使用二进制文件（推荐）
#### 本地运行
1.下载：下载 [hugoArticleEncryptor](https://github.com/hotjuicew/hugoArticleEncryptor/releases/latest) 到你的博客项目文件夹，

2.做加密标记：在你想要加密的文章的元信息中加入两个键值对
`protected: true`和 `password: 'your_password'`
例如：
```yaml
---
title: "Secret Post"
date: 2023-02-20T01:02:08+08:00
categories: ["Guide"]
protected: true
password: 'password'
---
```
3.运行命令：进入你的博客项目，运行你之前下载的二进制文件
```bash
$ .\hugoArticleEncryptor-windows-amd64.exe 
```
4.查看效果：如果你安装了python，可以运行以下命令后，在浏览器打开`http://localhost:1313/` 查看效果
```bash
$ python3 -m http.server -b 0.0.0.0 -d public 1313
```
#### Vercel、Netlify 等平台配置
1.将[build.sh](https://github.com/hotjuicew/hugoArticleEncryptor/blob/master/exampleSite/build.sh)复制到你的博客项目文件夹下
2.Build command: `sh build.sh`
### Option B: 使用源码构建
1.进入博客所在目录，克隆本项目
```bash
$ git clone https://github.com/hotjuicew/hugoArticleEncryptor.git
````
2.命令行输入
```bash
$ cd hugoArticleEncryptor ; go build -o ../hugoArticleEncryptor.exe ; cd ..
$ .\hugoArticleEncryptor.exe 
```

3.如果你安装了python，可以运行以下命令后，在浏览器打开`http://localhost:1313/` 查看效果
```bash
$ python3 -m http.server -b 0.0.0.0 -d public 1313
```