#  目录结构
* go_blog_service
  * configs: 配置文件
  * docs: 文档集合
  * global: 全局变量
  * internal: 内部模块
    * dao: 数据访问层(Database Access Object), 所有与数据相关的操作都会在dao层进行，例如MySQL，Elasticsearch...
    * middleware: HTTP中间件
    * model: 模型层，用于存放model对象
    * routers: 路由相关的逻辑
    * services: 项目核心业务逻辑
  * pkg: 项目相关的模块包
  * storage: 项目生成的临时文件
  * scripts: 各类构建、安装、分析等操作的脚本
  * third_party: 第三方资源工具，如Swagger UI
