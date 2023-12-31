# hugoArticleEncryptor

English | [简体中文](https://github.com/hotjuicew/hugoArticleEncryptor/blob/master/README-zh_CN.md)

hugoArticleEncryptor is a tool for encrypting Hugo articles. It is a Golang version of ⭐[Hugo Encryptor](https://github.com/Li4n0/hugo_encryptor).
It uses the AES-GCM algorithm to encrypt the entire Hugo article, and decrypts the content after the user enters the correct password.
The configuration is very simple. Once you have entered the correct password, you won't need to enter it again when accessing encrypted pages. The decrypted content will be directly displayed.

hugoArticleEncryptor only works on the "posts" folder (or "post" folder) under your "content" directory.

## [DEMO](https://juicebar-demo.add1.dev/)

(The password for this post is password)

## Installation and Usage

### Option A: Using automatic Binary Release

Please note that the binary files available in the releases section are automatically generated by GitHub Actions based on the publicly available source code. These binary files are provided for convenience and ease of use, especially when deploying on platforms like Vercel or Netlify.

If you prefer, you can also build and use the tool from the source code directly. Instructions for building from source can be found in the respective section of this README.

#### Local Execution

1.Download: Download [hugoArticleEncryptor](https://github.com/hotjuicew/hugoArticleEncryptor/releases/latest) to your blog project folder.，

2.In your article, you can use encryption tags as follows: enclose the content you want to encrypt with {{< secret "password" >}} and {{< /secret >}}. The {{< secret "password" >}} tag should be preceded by <!--more-->.

example:

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

3.Run the command: Navigate to your blog project and run the binary file you downloaded.

```bash
$ .\hugoArticleEncryptor-windows-amd64.exe
```

4.Check the result: If you have Python installed, you can run the following command and open`http://localhost:1313/` in your browser to see the result.

```bash
$ python3 -m http.server -b 0.0.0.0 -d public 1313
```

#### Configuration on Platforms like Vercel and Netlify

1.Copy [build.sh](https://github.com/hotjuicew/hugoArticleEncryptor/blob/master/exampleSite/build.sh)to your blog project folder.

2.Set the build command to `sh build.sh`

### Option B: Building from Source Code

1.Navigate to your blog directory and clone this project.

```bash
$ git clone https://github.com/hotjuicew/hugoArticleEncryptor.git
```

2.Run the following commands in the terminal:

```bash
$ cd hugoArticleEncryptor && go build -o ../hugoArticleEncryptor.exe && cd ..
$ .\hugoArticleEncryptor.exe <your-theme-name>
```

3.If you have Python installed, you can run the following command and open `http://localhost:1313/` in your browser to see the result.

```bash
$ python3 -m http.server -b 0.0.0.0 -d public 1313
```

## Attention

⚠️ Important Security Note: Protecting Your Blog Code ⚠️

To ensure the security and privacy of your blog content, we highly recommend keeping your blog code in a **private repository**.
