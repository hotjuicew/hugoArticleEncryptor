# hugoArticleEncryptor

[English](https://github.com/hotjuicew/hugoArticleEncryptor/blob/master/README.md) | 简体中文

hugoArticleEncryptor 是一个 hugo 文章加密工具。是 ⭐[Hugo Encryptor](https://github.com/Li4n0/hugo_encryptor)的 go 版本。
使用 AES-GCM 算法对整个 hugo 文章进行加密，在用户输入正确的口令之后解密内容。
配置非常简单。并且当你输入过一次正确密码后，下次访问加密页面就不需重复输入密码，会直接呈现解密后的内容。

hugoArticleEncryptor 只对你 content 文件夹下的 posts（或 post）文件夹起效果

## [DEMO](https://juicebar-demo.add1.dev/)

这篇文章的密码是 password

## 安装与使用

### Option A: 使用自动发布的二进制文件

请注意，发布页面提供的二进制文件是通过 GitHub Actions 根据公开的源代码自动生成的。这些二进制文件提供了方便和简便的使用方式，特别适用于在 Vercel 或 Netlify 等平台上部署。

如果你愿意，你也可以直接从源代码构建并使用该工具。有关从源代码构建的说明，请参阅本 README 中的相应部分。

#### 本地运行

1.下载：下载 [hugoArticleEncryptor](https://github.com/hotjuicew/hugoArticleEncryptor/releases/latest) 到你的博客项目文件夹，

2.在你的文章中做加密标记:用{{< secret "password" >}}和{{< /secret >}}包裹住你要加密的帖子。{{< secret "password" >}}前面需要有<!--more-->
例如：

```markdown
---
title: "example"
date: 2023-07-11T01:53:48+08:00
---

<!--more-->

{{< secret "password" >}}

## hi

### hugoArticleEncryptor is a hugo article encryption tool!

Let's try it.

> hugoArticleEncryptor was inspired by the hugo_encryptor project

{{< /secret >}}
```

3.运行命令：进入你的博客项目，运行你之前下载的二进制文件

```bash
$ .\hugoArticleEncryptor-windows-amd64.exe
```

4.查看效果：如果你安装了 python，可以运行以下命令后，在浏览器打开`http://localhost:1313/` 查看效果

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
```

2.命令行输入

```bash
$ cd hugoArticleEncryptor && go build -o ../hugoArticleEncryptor.exe && cd ..
$ .\hugoArticleEncryptor.exe
```

3.如果你安装了 python，可以运行以下命令后，在浏览器打开`http://localhost:1313/` 查看效果

```bash
$ python3 -m http.server -b 0.0.0.0 -d public 1313
```

## 注意

⚠️ 重要安全提示：保护您的博客代码 ⚠️

为了确保您博客内容的安全性和隐私性，强烈建议将您的博客代码保存在**私人存储库**中。
