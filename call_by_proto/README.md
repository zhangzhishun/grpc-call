# Call GRPC By Proto

## 1. Generate GRPC Code
```shell
cd protowire
protoc --go_out=plugins=grpc:. messages.proto
protoc --go_out=plugins=grpc:. p2p.proto 
```

# 2. Call GRPC Server
```shell
go run main.go
```