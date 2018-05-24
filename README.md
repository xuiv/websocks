# WebSocks

[教程资源](https://zhuji.lu/tags/websocks)
[中文说明](https://github.com/lzjluzijie/websocks/blob/master/README-zh.md)
[Telegram Group](https://t.me/websocks)

A secure proxy based on websocket.

This project is still working in progress, more features are still in development. If you are interested in this project, please star this project in order to support me. Thank you.

If you have any problems or suggestions, please do not hesitate to submit issues or contact me [@halulu](https://t.me/halulu). We also have a [telegram group](https://t.me/websocks) (mostly Chinese, English is ok).

Advantages:

- Using WebSocket and TLS which are very secure and difficult to be detected, same as regular HTTPS websites
- Can be used with cdn such as cloudflare, not afraid of gfw at all!

The disadvantage is that I have just started development, there is no GUI client, and features are not enough. I will appreciate if you can help me!

To-Do: WebSocks mux

## Example

### Built-in TLS with fake server name and reversing proxy

#### Server
```
./websocks -mode cert
./websocks -mode server -listen :2333 -path /password -proxy https://www.centos.org/ -tls
```

#### Local
```
./websocks -mode client -listen :1080 -server wss://the-real-server.com:2333/password -name www.centos.com -insecure
```

### Caddy TLS

#### Server
```
./websocks -mode server -listen :2333 -path /password
```

#### Local
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
