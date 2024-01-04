# Description
Golang API wikipedia

## Tech Stack
- Golang
- Mysql (Database)

## Framework & Library
- fiber
- gorm
- viper
- logrus
- validators

## Configuration
Check .env.example

## API SPEC
check `/api/api-spec.json`

## Run Migration
[dbmate](https://github.com/amacneil/dbmate)
```sh
make dbmate.up
```

## Run web Server
```sh
go run cmd/main.go
```