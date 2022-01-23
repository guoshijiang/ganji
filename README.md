### 1.项目概述

基于 beego 的市集商城 APP 后端代码。演示后台：http://60.205.1.144:8000/
用户名：admin
密码：asd..123

### 2.项目部署

第一步：克隆代码

```bigquery
git clone git@github.com:guoshijiang/ganji.git
cd blockshop_service
```

第二步：数据库 migrate

去 models 下面的 base.go 里面打开 `RunSyncdb` 函数，自动生成数据库，注意配置 false 和 true 项

第三步：运行开发
```bigquery
bee run 即可运行代码进行开发
```

第三步：代码部署

```bigquery
go build 完成之后，使用进程管理管理启动服务即可
```

如果您使用这套代码，开发搭建过程中有任何问题，可以去问我学院（www.wenwoha.com） 上面找联系方式联系我们，也可以直接加我的微信：LGZAXE

### 3.关联项目

网站：https://github.com/guoshijiang/shiji_web
APP 端代码：请去问我学院（www.wenwoha.com）或者椭圆曲线科技官网上联系我们获取

