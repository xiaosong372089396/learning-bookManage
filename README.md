# learning-bookManage 个人学习项目
---- controller                   // CLD: 服务入口，复杂处理路由、参数校验、请求转发

---- dao                          // 负责数据与存储相关功能(mysql、redis、es等)
       --- mysql
             --- mysq.go

--- main.go                       // 项目启动入口
--- middleware                    // 中间件: tolen验证
       ---- auth.go           

--- model                         // 模型
    --- book.go
    --- user.go
    --- user_m2m_book.go

router                            // 路由
     --- api_router.go
     --- init_router.go
     --- test_router.go