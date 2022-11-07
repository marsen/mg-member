# Start Your Develop Environment 

## init your go module

> go mod download

## migrate test data

```go=
o := access.DB{}
o.Migrate()
o.Seed()
```

(fin)