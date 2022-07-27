# rubix-edge

## install

need to run as sudo to install apps

```
go mod tidy
go build main.go && sudo ./main server
```

## product file (this is for hardware identification)

https://github.com/NubeIO/lib-command/blob/master/product/product.go#L7

`sudo nano /data/product.json`



```
{
    "version": "v1.1.1",
    "type": "AllLinux"
}
```

## command docs

[CLI](docs/api.md)
[CLI](docs/cli.md)
