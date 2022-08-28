# rubix-edge

## install

need to run as sudo to install apps

```
go mod tidy
go build main.go && sudo ./main server --auth=false
go build main.go && sudo ./main server --auth
```

## product file (this is for hardware identification)

https://github.com/NubeIO/lib-command/blob/master/product/product.go#L7

`sudo nano /data/product.json`

```
{
    "image_version": "v1.1.1",
    "arch": "amd64",
    "product": "Server"
}
```

```
{
    "image_version": "v1.1.1",
    "arch": "armv7",
    "product": "RubixComputeIO"
}
```

## command docs

[CLI](docs/api.md)
[CLI](docs/cli.md)
