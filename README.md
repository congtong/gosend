# gosend

在windows下打包linux下的包
在cmd中执行

GOARCH 包括 386，amd64，arm
GOOS 包括 darwin，freebsd，linux，windows

```cmd
set GOARCH=amd64
set GOOS=linux
go build
```