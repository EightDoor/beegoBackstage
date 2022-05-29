# beegoBackstage
- vue3 beego 后台管理系统
# 访问
- 访问地址: beegobackstage.start6.cn
- 账号: admin 密码: 123456

# nginx配置
```nginx
server
{
    listen 80;
    server_name beegobackstage.start6.cn;
    root /www/wwwroot/beegobackstage.start6.cn/;
    location / {
       # 需要指向下面的@router否则会出现vue的路由在nginx中刷新出现404
       try_files $uri $uri/ @router;
       index index.html;
    }
    location @router {
           #对应上面的@router，主要原因是路由的路径资源并不是一个真实的路径，所以无法找到具体的文件
           #因此需要rewrite到index.html中，然后交给路由在处理请求资源
           rewrite ^.*$ /index.html last;
    }
    location /api {
      proxy_pass http://localhost:8098/api/v1/;
    }
    location /static {
      proxy_pass http://localhost:8098/static/;
    }
}
```
# License
[MIT](https://opensource.org/licenses/MIT)

Copyright (c) 2022, EightDoor
