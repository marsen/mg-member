# Developer Local Database 

## premise
- Familiar with docker, git

## Postgres DB

1. pull image 
    > docker pull postgres

2. run postgres database

   perhaps you'll get an error message like this:

   > Error: Database is uninitialized and superuser password is not specified.  
   > You must specify POSTGRES_PASSWORD to a non-empty value for the  
   > superuser. For example, "-e POSTGRES_PASSWORD=password" on "docker run".  
   > 
   > You may also use "POSTGRES_HOST_AUTH_METHOD=trust" to allow all
   > connections without a password. This is *not* recommended.  
   > See PostgresSQL documentation about "trust":
   >  https://www.postgresql.org/docs/current/auth-trust.html

   to fix this problem, you have to provide password , and default user will be `postgres`  
   but you can change it too
   
   ```shell
    docker run --name=pg-member-db -p 5432:5432 \
    -e POSTGRES_PASSWORD= \
    -e POSTGRES_USER= \
    -d postgres
   ```

3. create database

```sql
create database member
```

### NEXT? MIGRATION OR SQL COMMAND ?