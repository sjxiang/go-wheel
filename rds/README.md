
## "快速开发的脚手架" -> "Rapid Development Scaffold" 


日志


中间件
    可观测性
        logging 访问日志（一环）
        tracing 链路追踪（分布式环境下，多环串起来）
        metric 
        

    从 panic 中恢复

    

HTTP 无状态


session 

    服务端维护用户状态的机制

    cookie 客户端维护

    源码
        // 塞进去
        func Sessions(name string, store Store) gin.HandlerFunc {
            return func(c *gin.Context) {
                s := &session{name, c.Request, store, nil, false, c.Writer}
                c.Set(DefaultKey, s)
                defer context.Clear(c.Request)
                c.Next()
            }
        }

        // 取出来
        func Default(c *gin.Context) Session {
            return c.MustGet(DefaultKey).(Session)
        }


    Session 接口
        - Get、Set、Save（有批量操作的，刷新） 操作键值对



    设计
        对 gollira session 的封装


        type Session struct {
            Values map[interface{}]interface{}  // 临时存放内存
            store  Store                        // 刷新到真实存储的地方
        }        

        数据一开始放在内存中，即 Values 里面；
        后来调用 Save，刷新到存储里面，例如刷新到 Redis 中


    
    安全
        较弱鸡，只认 session id 不认人（被劫持）

        保护措施
            1. 在使用 cookie 时，同时设置 http_only 和 secure 选项，限制 cookie 只能在 HTTPS 协议里面传输

            2. 在 session id 编码的时候，带上一些客户端信息（例如 agent 信息、mac 地址 之类的）
               服务端检测到 session id 所携带的信息发生了变化，就要求用户重新登录


        