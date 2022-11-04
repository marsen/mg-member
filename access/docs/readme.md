# Developer Local Database 

## premise
- Familiar with docker, git

## Postgres DB

pull image 
> docker pull postgres

run database
> docker run postgres

perhaps you'll get an error message like this:

>Error: Database is uninitialized and superuser password is not specified.  
> You must specify POSTGRES_PASSWORD to a non-empty value for the  
> superuser. For example, "-e POSTGRES_PASSWORD=password" on "docker run".  
> 
> You may also use "POSTGRES_HOST_AUTH_METHOD=trust" to allow all
> connections without a password. This is *not* recommended.  
> See PostgreSQL documentation about "trust":
>  https://www.postgresql.org/docs/current/auth-trust.html


```
docker run --name=pg-member-db -p 5432:5432 -e POSTGRES_PASSWORD=password postgres
```