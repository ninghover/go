server 服务端
web    客户端 （相对于微服务是客户端，相对于浏览器是服务端）

```
客户端通过rpc调用 得到 rsp和err
服务端最好只要在调用失败或者异常的时候才返回err，否则都返回nil
如果数据操作不成功，但是不是异常，不要返回err，而是把errno设置为不ok

客户端的ctx.Json(http.StatusOk,rsp)，这个ok是说浏览器对我客户端的访问是ok的
```