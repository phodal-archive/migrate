# 1. migration dsl design

日期: 2019-12-09

## 状态

2019-12-09 提议

## 背景

为了做到 DDD 自动化重构，那么首先我们要能识别现有的工程及目录结构。从当前的情况来看，一个传统的三层分层架构，往往由两种形式组成：

 - 单体三层分层架构。
 - 模块化分层架构。其中又分为 ``maven`` 和 ``gradle``

单体三层分层架构示例：

示例代码见：https://github.com/Pragmatists/DDD-ADM

```
├── Application.java
├── controller
│   ├── IssueController.java
│   └── json
│       ├── IssueCommentJson.java
│       ├── IssueJson.java
│       ├── NewIssueCommentJson.java
│       ├── NewIssueJson.java
│       └── UpdateIssueJson.java
├── model
│   ├── Issue.java
│   ├── IssueComment.java
│   └── IssueStatus.java
├── repository
│   ├── IssueCommentRepository.java
│   └── IssueRepository.java
└── service
    └── IssueService.java
```

模块化分层架构：

示例代码见：https://github.com/shuzheng/zheng

```
├── README.md
├── pom.xml
├── zheng-api-common
│   ├── pom.xml
│   └── src
├── zheng-api-rpc-api
│   ├── pom.xml
│   └── src
├── zheng-api-rpc-service
│   ├── pom.xml
│   └── src
├── zheng-api-server
│   ├── pom.xml
│   └── src
└── zheng-api.iml
```

对于单体而言，我们需要做的是：

1. 重新划分包。即在保持业务不中断的情况下重构，以让新的代码运行在新的架构上。
2. 分析抽象领域模型
3. 编写 API 测试，保证现有的功能
4. 编写抽象接口，进行依赖反转
5. 拆分 service 层，重构代码。将行为绑定于是领域对象上。

而模块化的分层架构，则还需要：

 - 0. 分析包依赖
 - ...
 - 6. 生成包依赖


为此，我们需要这么一些架构信息：

 - DDD 架构模式
   - Command / Event / Representation
   - Request / Response 
 - 依赖信息生成与分析 
 - 服务调用链
 - 生成接口
 - ……
 
## 决策

一个简单的 DSL 设计 / Config：

```migration
project: DDD ADM
pattern: clean
tools: gradle / maven # 临时 DSL
```


## 后果

在这里记录结果...
