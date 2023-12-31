

# GO语言简单整合ChatGPT的API

通过对OpenAI官方对外开放的接口的调用，实现简单的整合，目前只整合了聊天的API。后续也可以整合画图等API。

<!-- PROJECT SHIELDS -->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />

<p align="center">
  <a href="https://github.com/langwudong/openai/">
    <img src="logo.png" alt="img" width="120" height="120">
  </a>


  <h3 align="center">GO语言简单整合ChatGPT的API</h3>
  <p align="center">
    欢迎测试、学习以及交流
    <br />
    <a href="https://github.com/langwudong/openai"><strong>探索本项目的文档 »</strong></a>
    <br />
    <br />
    <a href="https://github.com/langwudong/openai">查看Demo</a>
    ·
    <a href="https://github.com/langwudong/openai/issues">报告Bug</a>
    ·
    <a href="https://github.com/langwudong/openai/issues">提出新特性</a>
  </p>


</p>

## 目录

- [上手指南](#上手指南)
  - [安装](#安装)
- [文件目录说明](#文件目录说明)
- [开发的架构](#开发的架构)
- [部署](#部署)
- [使用到的框架](#使用到的框架)
- [贡献者](#贡献者)
  - [如何参与开源项目](#如何参与开源项目)
- [版本控制](#版本控制)
- [作者](#作者)
- [鸣谢](#鸣谢)

### 上手指南

在调用接口前，只需要构建OpenAI结构体，然后附带请求地址和请求消息体调用CallCompletion方法即可。OpenAI结构体只需要你的OpenAI的Key。你也可以将请求地址换成中转点的请求地址。

#### 安装

克隆本仓库

```sh
git clone https://github.com/langwudong/openai.git
```

### 文件目录说明

eg:

```
filetree 
├─.idea
└─openai

```





### 开发的架构 

暂无

### 部署

暂无

### 使用到的框架

暂无

### 贡献者

- [朗 吾 東](https://github.com/langwudong)

#### 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



### 版本控制

该项目使用Git进行版本管理。您可以在repository参看当前可用版本。

### 作者

langwudong@yeah.net

知乎: [朗 吾 東](https://www.zhihu.com/people/---60-9-44)  &ensp; qq: 63265742

*欢迎添加我的联系方式交流学习*

*您也可以在贡献者名单中参看所有参与该项目的开发者。*

### 版权说明

该项目由个人原创，你可以随意进行**下载**、**测试**以及**二次开发**。

### 鸣谢


暂无

<!-- links -->

[your-project-path]:langwudong/openai
[contributors-shield]: https://img.shields.io/github/contributors/langwudong/openai.svg?style=flat-square
[contributors-url]: https://github.com/langwudong/openai/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/langwudong/openai.svg?style=flat-square
[forks-url]: https://github.com/langwudong/openai/network/members
[stars-shield]: https://img.shields.io/github/stars/langwudong/openai.svg?style=flat-square
[stars-url]: https://github.com/langwudong/openai/stargazers
[issues-shield]: https://img.shields.io/github/issues/langwudong/openai.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues/langwudong/openai.svg
[license-shield]: https://img.shields.io/github/license/shaojintian/Best_README_template.svg?style=flat-square
[license-url]: https://github.com/langwudong/openai/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/langwudong
