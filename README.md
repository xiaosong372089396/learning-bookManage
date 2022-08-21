# learning-bookManage 个人学习项目
├── controller      // CLD:服务入口，负责处理路由、参数校验、请求转发
│   ├── book.go 
│   └── user.go 
├── dao             // CLD:负责数据与存储相关功能(mysql、redis、ES等)
│ └── mysql
│   └── mysql.go
├── main.go          // 项目启动入口
├── middleware       // 中间件： token验证
│ └── auth.go
├── model            // 模型
│ ├── book.go
│ ├── user.go
│ └── user_m2m_book.go 
└── router              // 路由
   ├── api_router.go 
   ├── init_router.go 
   └── test_router.go

// 
常用项目结构
├── Readme.md   // 项目说明(帮助你快速的属性和了解项目)
├── config      // 配置文件(mysql配置 ip 端口 用户名 密码，不能写死到代码中) 
├── controller // CLD:服务入口，负责处理路由、参数校验、请求转发
├── service    // CLD:逻辑(服务)层，负责业务逻辑处理
├── dao        // CLD:负责数据与存储相关功能(mysql、redis、ES等)
│ ├── mysql
├── model      // 模型
├── logging    // 日志处理
├── main.go   // 项目启动入口
├── middleware // 中间件
├── pkg       // 公共服务(所有模块都能访问的服务)
├── router    // 路由(路由分发)

创建数据库
mysql> create database books charset utf8;

当前项目结构
go mod init bookManage