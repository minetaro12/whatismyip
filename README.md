# whatismyip
```
#ビルド
$ go build

#実行
$ ./whatismyip

$ PORT=8080 ./whatismyip

#リバースプロキシの後ろの場合はクライアントのIPアドレスが入っているヘッダーを指定する
$ REAL_IP_HEADER="X-Real-IP" ./whatismyip
```

## 使い方
```
$ curl localhost:8000

$ curl localhost:8000/json
```