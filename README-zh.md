# WebSocks

[教程资源](https://zhuji.lu/tags/websocks)
[TG群](https://t.me/websocks)

一个基于 WebSocket 的代理工具

本项目目前还在开发中，更多功能仍在完善中。如果你对这个项目感兴趣，请star它来支持我，蟹蟹

有任何问题或建议可以直接发issue或者联系我 [@halulu](https://t.me/halulu)，也可以来[TG群](https://t.me/websocks)水一水

开发记录可以看[我的博客](https://halu.lu/post/websocks-development/)

友人正在[一键脚本以及GUI](https://zhuji.lu/topic/15/websocks)

优点
 - 使用WS+TLS，十分安全且不易被检测，和普通HTTPS网站一样
 - 可以搭配使用cloudflare这类cdn，完全不怕被墙！

缺点就是刚刚开始开发，没有GUI客户端，功能也比较少，如果你能来帮我那就太好了！

To-Do: WebSocks mux

## 示例

### Caddy TLS

#### 服务端
```
./websocks -mode server -listen :2333 -path /password
```

#### 客户端
```
./websocks -mode client -listen :1080 -server wss://server.com/password
```

#### Caddyfile
```
https://server.com {
  proxy /password localhost:2333 {
    websocket
  }
}
```

### 内置 TLS 混淆域名并反向代理

#### 服务端
```
./websocks -mode cert
./websocks -mode server -listen :2333 -path /password -proxy https://www.centos.org/ -tls
```

#### 客户端
```
./websocks -mode client -listen :1080 -server wss://server.com:2333/password -name www.centos.com -insecure
```
