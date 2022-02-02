# Goodwe monitor
This application retrieves data from a Goodwe Inverter, store it and 
generate stats and graphs.

## Installation
```
docker compose up -d
```

With `docker ps` you should have a couple of containers:
- `goodwe_monitor_web`
- `goodwe_monitor_db`

## How to execute it
1- Access to the `goodwe_monitor_web` container

2- Execute the following command:
```
go run main.go
```

## Troubleshooting
1- ``` $GOPATH/go.mod exists but should not ```

``` unset GOPATH ```
