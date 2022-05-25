# 运行
- 安装bee
  - bee run
- 进程内监控 访问: [http://localhost:8088/](http://localhost:8088/)
- 注解路由生成 `bee generate routers`
# 打包
- `bee pack -be GOOS=linux -be GOARCH=amd64` 打包资源为tar.gz, 服务器使用ubuntu
## 注意

- 注解路由 controller 下创建文件`// @router / [get]` 前后不能存在别的内容,否则无法生成路由
- controller 统一继承 `utils.BaseController`
## 使用

- 表单校验
  - **注意**需要在格式化json之后执行

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
          Data:      &dataInfo
      })
  ```

  - 读取body参数
  ```go
  var userPassword BaseModels.SysUserUpdatePassword
  // 读取body参数
  c.GetBodyToJson(&userPassword)
  ```
