# Start Your Develop Environment 

## init your go module

> go mod download

## Database

### postgresSQL

### start with docker 

```shell
docker run -d \
	--name mg-member \
	-e POSTGRES_PASSWORD=p@ssw0rd \
	-e POSTGRES_DB=member \
	-e POSTGRES_USER=admin \
	-p 5432:5432 \
	postgres
```

#### migrate test data

```go=
o := access.DB{}
o.Migrate()
o.Seed()
```

## DI 使用 wire

安裝 CLI Tool
```shell
go install github.com/google/wire/cmd/wire@latest
```
專案安裝 package

```terminal
go get github.com/google/wire/cmd/wire@latest
```

執行 CLI

```shell
wire 
```


(fin)