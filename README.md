# qiniu-manager

自学go的一个练手小Demo(也是一个轮子), 为客户端上传文件到七牛空间生成上传凭证

## 部署信息

<http://qiniu-manager.herokuapp.com/>

## 获取凭证接口

```bash
# 请求参数
#       key: 上传文件保存的文件名，可任意指定，但是不能是七牛空间里已经存在的文件名，目前不允许覆盖同名文件
# 返回值
#       上传的token，供客户端上传使用

$ curl http://qiniu-manager.herokuapp.com/uptoken?key=test
{"uptoken":"XVp1PMBWtvo5HVLa3sKS28KS87Yggj1cXtAQhpAk:vMEkUTLWXP3pVItIZGuj1SES_7w=:eyJzY29wZSI6ImJsb2ctYXNzZXQ6dGVzdCIsImRlYWRsaW5lIjoxNDQ1NjE5Nzc5LCJpbnNlcnRPbmx5IjoxfQ=="}
```

## 客户端代码

<https://github.com/crazygit/qiniu-client>

