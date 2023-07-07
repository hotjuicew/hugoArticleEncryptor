# hugoArticleEncryptor
English | [简体中文](https://github.com/hotjuicew/hugoArticleEncryptor/blob/master/README-zh_CN.md)

hugoArticleEncryptor is a tool for encrypting Hugo articles. It uses the AES-GCM algorithm to encrypt the entire Hugo article and inserts JavaScript code into the encrypted article. The content can be decrypted by entering the correct passphrase.

The configuration is very simple. Once you enter the correct password, you don't need to re-enter it when accessing the encrypted page again. The decrypted content will be directly displayed.
## [DEMO](https://juicebar-demo.add1.dev/)
## Installation and Usage
### Option A: Using Binary File (Recommended)
#### Local Execution
1.Download: Download  [hugoArticleEncryptor](https://github.com/hotjuicew/hugoArticleEncryptor/releases/latest) to your blog project folder.，


2.Add encryption markers: Add two key-value pairs to the metadata of the article you want to encrypt: `protected: true` and `password: "your_password"`. For example:

example:
```yaml
---
title: "Secret Post"
date: 2023-02-20T01:02:08+08:00
categories: ["Guide"]
protected: true
password: 'password'
---
```
4.Run the command: Navigate to your blog project and run the binary file you downloaded.
```bash
$ .\hugoArticleEncryptor-windows-amd64.exe
```

5.Check the result: If you have Python installed, you can run the following command and open`http://localhost:1313/` in your browser to see the result.
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

4.If you have Python installed, you can run the following command and open `http://localhost:1313/` in your browser to see the result.
```bash
$ python3 -m http.server -b 0.0.0.0 -d public 1313
```