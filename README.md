# managedby

### Run database server and web api by `air`.
#### run api server
```bash
touch .env
cp .env.example .env
go install github.com/cosmtrek/air
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
air -v
air
```

#### run db
```bash
docker-compose build && docker-compose up
docker exec -it managedby_db bash -c 'mysql -u root -ppassword'
```

#### reset container
```bash
docker-compose down --rmi all --volumes --remove-orphans
docker-compose up --build
docker exec -it managedby_db bash -c 'mysql -u root -ppassword'
```

### author
[@rrih_dev](https://twitter.com/rrih_dev)
