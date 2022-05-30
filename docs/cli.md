## cmd cli

add github token

## download and install an app cli

```
cd cmd
```
first add a store, the store is used to store how an app is to be installed and its name
```
go build main.go && sudo ./main apps --store-add=true --app=flow-framework  --version=latest --download-path=/home/aidan/downloads 
```

Once a store is added you can install a new app, you need to pass in the


```
go build main.go && sudo ./main apps --store-add=false --install=true  --app=flow-framework  --version=latest --download-path=/home/aidan/downloads --token= TOKEN
```


## REST

### add an app to the store

- `http://{{rubix_apps_ip}}:{{rubix_apps_port}}/api/stores`
- `POST`

```
{
    "name": "flow-framework",
    "allowable_products": [
        "RubixCompute",
        "AllLinux"
    ],
    "Port": 1660,
    "app_type_name": "Go",
    "repo": "flow-framework",
    "service_name": "nubeio-flow-framework",
    "RubixRootPath": "/data",
    "apps_path": "/data/rubix-apps/installed",
    "app_path": "/data/flow-framework",
    "download_path": "/home/aidan/downloads",
    "asset_zip_name": "",
    "owner": "NubeIO",
    "run_as_user": "root",
    "data_dir": "",
    "config_dir": "",
    "config_file_name": "",
    "description": "flow-framework",
    "service_working_directory": "/data/rubix-apps/installed/flow-framework",
    "service_exec_start": "/data/rubix-apps/installed/flow-framework/app-amd64 -p 1660 -g /data/flow-framework -d data -prod",
    "product_type": "status",
    "arch": ""
}

```


### install an app

- `http://{{rubix_apps_ip}}:{{rubix_apps_port}}/api/apps`
- `POST`

```
{
    "app_name": "flow-framework",
    "version":"latest",
    "token":""
}
```