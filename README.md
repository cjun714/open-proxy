## Compile
``` shell
$ go build -o open_proxy ./src
```

## How to use
``` shell
$ nohup ./open_proxy 8080 > ./open.log 2>&1 &
```
