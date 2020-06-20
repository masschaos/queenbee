---
title: 运作方式
slug: how-it-works
weight: 20
description: >
  系统中一些重要节点的运作方式。
---

## 工作设备和控制中心是怎么联系的

![工作设备原理](/images/worker.png)

- 由于工作设备数量巨大，WorkerBee 被安装到工作设备时一般是批量安装或用户自行安装的，
  在安装时安装包内就预置了控制中心的地址和一个 `分组 Token` 。
- WorkerBee 在每次启动时，都会用 `分组 Token` 去调用登录接口，换取独属于自己的 `Token` 。
- WorkerBee 会周期性的主动向控制中心请求任务，控制中心依靠 `Token` 识别出工作设备后，分配给它合适的任务。
- 在完成任务之后会向控制中心报告，并继续请求任务。
- WorkerBee 随时有可能因各种原因退出，不用担心未完成的任务，控制中心在任务超时后重新将任务分给别的工作设备。


## 任务在控制中心内是怎么流转的

![生命周期](/images/lifecycle.jpg)

## 任务分配过程

![分配过程](/images/dispatch.jpg)

