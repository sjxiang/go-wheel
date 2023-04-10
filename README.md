
# 造轮子，自娱自乐

web 框架 
ORM 框架（object relational mapping, 对象关系映射）
缓存客户端




面试技巧

精髓 - 把面试官拉到自己擅长领域，拖过 30 分钟。


框架基准测试 Benchmark

    开发者十分清楚，什么样的测试条件对自己是有利的

    利益无关，设计好测试。




性能测试 - 看谁测，要什么参数
 

100 mistakes

职业发展

    - 技术专家 
    - 领域专家（金融、医疗 ... 具体的非 IT 行业，变革慢，职业生涯较漫长，所谓复合型人才） 
    - 技术管理（人多坑少，PUA 较恶心 ... 脱离技术，很难回头）





3. 掌握 HTTP Server 和 Context 的设计，并且提供丰富 API

4. 掌握 HTTP 中 Session 的设计和实现

5. 掌握 Web 框架中 AOP 的解决方案

6. 设计并实现简单的静态资源服务器


学习和工作中的痛点

用过很多 Web 框架，但是并不了解 Web 框架的原理，并不知道怎么注册路由，怎么执行路由匹配
面试的时候无法清晰阐述前缀路由树的原理，在实际开发中，不知道如何快速定位 404 之类的错误
不知道如何设计统一的 Session 抽象，支持 Session 运行在本地内存或者 Redis 上
无法灵活运用 Web 框架提供的 AOP 方案解决登录校验、鉴权、日志、tracing、logging 等问题


实践练习
设计一个 HTTP Server，该 HTTP Server 将会基于前缀路由树支持路由通配符匹配、路径参数、正则匹配
为 HTTP Server 添加静态资源支持（CSS、JS 等），并且提供缓存和内存控制功能
为 HTTP Server 添加 Session 功能
为 HTTP Server 添加模板引擎功能，并提供基于 Go 模板的默认实现
为 HTTP Server 设计 AOP 方案
利用该 HTTP Server 实现简单的用户 API，支持注册、登录等
利用该 HTTP Server 的 AOP 方案解决登录校验、日志、tracing 和 metric 问题
利用该 HTTP Server 的 Session 功能，维护登录态




第二周
支持 AOP 方案：

责任链模式、洋葱模式、拦截器详解

开源实例：

    Gin 中 Handler 的设计与实现
    Beego 中 Filter 的设计与实现
    Kratos 中 Middleware 的设计与实现


代码演示：为 Web 框架提供 AOP 支持，并且支持

    access log（访问日志）
    tracing
    metric
    recovery：从 panic 中恢复过来
    错误处理：支持返回特定错误页面或者重定向等

面试要点



第三周

1. 文件上传与下载

    开源实例：Gin、Beego、Iris 和 Echo 中的文件上传和下载功能

    代码演示：实现文件上传和下载功能

    进阶语法：文件操作


2. 页面渲染

模板引擎设计

进阶语法：Template 编程

代码演示：在 Web 框架中支持页面渲染

Option 设计模式

代码演示：使用泛型设计通用的 Option 模式

静态资源支持

开源实例：Gin、Beego、Iris 和 Echo 中的静态资源支持

代码演示：设计并实现一个可配置、可扩展的静态资源处理器

缓存静态资源
控制缓存消耗
大文件的解决方案

支持 Session：

开源实例：Beego 中的 Session 模块、Gin 的 Session 扩展

代码演示：Session API 设计与实现

基于内存的实现
基于 Redis 的实现

面试要点