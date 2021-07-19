# managedby

## Run database server and web api by `air`.
```bash
touch .env
cp .env.example .env
# run api server
go get -u github.com/cosmtrek/air
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
air -v
air
```

```bash
# run db
docker-compose build && docker-compose up
docker exec -it managedby_db bash -c 'mysql -u root -ppassword'
```

## author
[@rrih_dev](https://twitter.com/rrih_dev)