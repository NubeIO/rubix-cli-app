# rubix-edge

## install

need to run as sudo to install apps

```
go mod tidy
go build main.go && sudo ./main server
```

## product file (this is for hardware identification)

https://github.com/NubeIO/rubix-edge/blob/master/service/product/funcs.go#L28

`sudo nano /data/product.json`

```
{
    "version": "v1.1.1",
    "type": "status"
}
```

## command docs

[CLI](docs/api.md)
[CLI](docs/cli.md)
