# 运行
- 安装bee
  - bee run
  - 初始化swagger文档 `bee run -downdoc=true -gendoc=true`
- swagger访问: [http://localhost:8080](http://localhost:8080)
- 进程内监控 访问: [http://localhost:8088/](http://localhost:8088/)
- 注解路由生成 `bee generate routers`
# 接口文档
- https://www.apifox.cn/apidoc/shared-9c86940d-fb4a-4db3-a5dc-bc411f75c0a7
## 注意

- 注解路由 controller 下创建文件`// @router / [get]` 前后不能存在别的内容,否则无法生成路由
- controller 统一继承 `utils.BaseController`
## 使用
- 开发一个接口
  - controllers 新建文件
  - models 新建文件
  - models/RegisterModels 填写新建的model
    - model名称对应数据库表 驼峰写法，自动转换为小写_
      - 例如: SysUser -> sys_user
  - routers/router 填写新建的controller
  - 执行 `bee generate routers` 在routers/commentsRouter 文件里面生成对应路由

- 表单校验
  - **注意 **需要在格式化json之后执行

```go
    // 在struct定义时添加 valid,label显示输出的提示
    type SysUserUpdatePassword struct {
        Id          int    `json:"id" valid:"Required" label:"主键id"`
        Password    string `json:"password" valid:"Required" label:"旧密码"`
        NewPassword string `json:"newPassword" valid:"Required" label:"新密码"`
    }
```

```go
	var userPassword BaseModels.SysUserUpdatePassword
	var user BaseModels.SysUser
	o := orm.NewOrm()
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &userPassword)
	// 表单校验
	c.CustomValid(userPassword)
```

- 统一返回
  - Controller统一继承utils.BaseController
  - 成功返回

```go
	// orm查询数据
	o := orm.NewOrm()
	var t []*BaseModels.Test
	qs := o.QueryTable("test")
	qs.All(&t)
	c.RSuccess(utils.R{Data: t})
```

- 失败返回

```go
	// orm查询数据
	o := orm.NewOrm()
	var t []*BaseModels.Test
	qs := o.QueryTable("test")
	qs.All(&t)
	c.RError(utils.R{Data: t})
```

- 自动处理异常返回

```go
	var user []*BaseModels.SysUser
	o := orm.NewOrm()
	result, err := o.QueryTable(BaseModels.SysUser{}).All(&user)
	c.RBack(utils.R{Data: result}, err)
```

- 分页返回

```go
	// orm查询数据
o := orm.NewOrm()
var dataInfo []*BaseModels.Test
qs := o.QueryTable("test")
c.RPaging(utils.RPage{
OrmResult: qs,
Data:      &dataInfo})
```
