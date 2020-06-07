---
title: 总览
weight: 1
description: >
  包含项目的概述，总体架构，以及核心概念。
---

## BeeFog 是什么？

BeeFog 是一个开源的基于雾计算的分布式爬虫方案，如下图所示：

![原理示意图](/images/beefog.jpg)

一个爬虫任务往往需要抓取一个站点的若干页面，`QueenBee`会将每个页面分配给目标站点附近的边缘节点`Worker`。
边缘节点会通过预设的任务模版，将页面抓取后再将其中有用的信息组合成预设的结构化数据，返回给`QueenBee`。
`QueenBee`再将结果数据根据`User`预设的行为配置进行相应的处置。

## 和其他爬虫框架的区别？
由于爬虫框架太多，不在这里一一对比，列举一些我们的特性：
- 大部分的网络流量和计算任务发生在边缘节点，可以显著节约服务器的流量和计算资源。
- 不必使用代理服务器，请求会被自动分布至所有合适的边缘节点，可以为限定单个IP的请求速率。
- 支持各种类型的目标内容，包括包含js的页面、单页应用、XML、JSON等。系统会选择更加合适的`Worker`类型去执行任务。
- 通过 HTTP API 配置，业务端可以选用任何技术栈和语言。

## 我需要 BeeFog 吗？
虽然 BeeFog 可以适应任何场景的爬虫任务，但有的场景会比竞争对手有显著优势：

### 抓取分布在多个国家的内容
BeeFog 可以总是让距离目标最近的节点去抓取和处理数据，将传输往云端的数据量减少到最少，避免网络带来的问题。

### 需要抓取的结构化内容只占页面很小的比例
通常爬虫提取的内容都不会高于原始内容的10%，但有的场景，需要提取的内容甚至不到1%。这个比例越悬殊，将计算放在边缘节点会越有优势。

### 目标站点对游客有较为严格的速率限制
分布式的抓取将天然避免绝大部分的服务端限制策略，比代理服务器更为有效。

### 目标站点对于浏览器或请求者有严格的要求。
我们有好几种`Worker`来适应各种对于客户端的要求，如果有更为特殊的要求，可以根据我们开源的协议，自己实现一个`Worker`去执行任务，
这比写一套完整的爬虫项目更加敏捷和节约成本。

### 我的爬虫策略在时间上分布不均或很随机，浪费服务器资源
加入我们的互助联盟，您只需要部署几个`Worker`节点在服务器空闲时持续的执行任务，就可以换来等价的资源在您需要时使用。大大节约了成本。

## 如何开始？
您至少要阅读本章的概述文档了解基本的概念，然后阅读"快速开始"章节就可以上手使用了。在需要更高级的特性时，有选择的阅读剩余的文档。