### db migration 

`migrate create -ext sql -dir db/migration -seq init_schema`

### create db after execing into postgres container

```
  createdb --username=postgres --owner=postgres simple_bank
  su - postgres
  psql simple_bank
```

### create db from outside the contianer using exec command.

```
  docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank
  docker exec -it postgres psql -U postgres simple_bank
```


### Used sqlc for database interaction in project

```
brew install sqlc
sqlc version
sqlc help
sqlc init
```

#### Added sqlc code into makefile : `make sqlc`

#### psql in docker contianer 
```
(base) ➜  simplebank git:(master) ✗ docker exec -it postgres psql -U postgres
psql (13.16)
Type "help" for help.

postgres=# \c product
You are now connected to database "product" as user "postgres".
product=# \dt
          List of relations
 Schema |  Name   | Type  |  Owner   
--------+---------+-------+----------
 public | product | table | postgres
(1 row)
```




