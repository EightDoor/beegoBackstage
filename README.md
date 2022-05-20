# 运行
- 安装bee
  - bee run
  - 初始化swagger文档 `bee run -downdoc=true -gendoc=true`
- swagger访问: http://localhost:8080
- 注解路由生成 `bee generate routers`

## 注意
- 注解路由 controller 下创建文件`// @router / [get]` 前后不能存在别的内容,否则无法生成路由
