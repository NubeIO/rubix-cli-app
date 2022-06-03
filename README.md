# rubix-cli-app

## install

```
go mod tidy
go run ./cmd/main.go server
```

## product file (this is for hardware identification)
https://github.com/NubeIO/rubix-cli-app/blob/master/service/product/funcs.go#L28

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
