## 项目结构

-   blog-service
    -   configs
    -   docs
    -   global
    -   internal
        -   dao
        -   middleware
        -   model
        -   routers
        -   service
    -   scripts
    -   storage
    -   pkg
    -   third_party

configs : 配置文件

docs : 文档集合

global : 全局变量

internal : 内部模块

dao : 数据访问层

middleware : HTTP 中间件

model : 模型层

routers : 路由逻辑

service : 业务层

scripts : 脚本

storage : 项目生成的临时文件

pkg : 项目依赖包

third_party : 第三方资源工具

---

## 业务模块

该博客根据 **RESTful API**的基本规范，设计业务模块的路由规则

1.  标签管理

|     功能     | HTTP 方法 |   路径    |
| :----------: | :-------: | :-------: |
|   新增标签   |   POST    |   /tags   |
| 删除指定标签 |  DELETE   | /tags/:id |
| 更新指定标签 |    PUT    | /tags/:id |
| 获取标签列表 |    GET    |   /tags   |

2. 文章管理

|     功能     | HTTP 方法 |     路径      |
| :----------: | :-------: | :-----------: |
|   新增文章   |   POST    |   /articles   |
| 删除指定文章 |  DELETE   | /articles/:id |
| 更新指定文章 |    PUT    | /articles/:id |
| 获取指定文章 |    GET    | /articles/:id |
| 获取文章列表 |    GET    |   /articles   |
