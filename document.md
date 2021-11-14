# 開発環境構築
フロントエンドの開発環境は[こちら](https://github.com/rrih/songscoreonline-next#songscoreonline-next)  
バックエンドの開発環境はdbサーバーとapiサーバーをそれぞれ起動させる必要がある。  
それぞれ以下に記述。

#### backend起動
```bash
touch .env
cp .env.example .env
go install github.com/cosmtrek/air
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
air -v
air
```

#### dbサーバ起動
cliでmysqlコンテナに入る
```bash
docker-compose build && docker-compose up
docker exec -it managedby_db bash -c 'mysql -u root -ppassword'
```

#### コンテナの初期化
```bash
docker-compose down --rmi all --volumes --remove-orphans
docker-compose up --build
docker exec -it managedby_db bash -c 'mysql -u root -ppassword'
```

## api使用例

ユーザー登録
```
curl -H 'Content-Type:application/json' -d '{"name":"test user", "email":"foo@bar.com", "password":"testtesttest"}' http://localhost:8080/api/v1/users/signup/
```

## deploy
事前ビルド、Procfile等は heroku の master への push のみで完結する
```
# deploy to heroku
git push heroku master
```

ログイン
```
curl -i -H 'Content-Type:application/json' -d '{"email":"test@gmail.com", "password":"testtesttest"}' http://localhost:8080/api/v1/auth/login/

HTTP/1.1 200 OK
Date: Sun, 05 Sep 2021 15:32:10 GMT
Content-Length: 145
Content-Type: text/plain; charset=utf-8

{"Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzA4NTk1MzAsInVzZXIiOiJmb29AYmFyLmNvbSJ9.9G3gPXQ2ZzZ5pa7qMN1md8h55I5RDNFbWZfZOqEBgzE"}
```

ログアウト
```
curl -i -H 'Content-Type:application/json' -d '{"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzE3MjE5OTMsInVzZXIiOiJ0ZXN0QGdtYWlsLmNvbSJ9.IybN5Wrq8s1SqNSKK5C00hoZtc2Qgy_aCgNZS4oVOdo"}' http://localhost:8080/api/v1/auth/logout/
```

ユーザー詳細取得
get /api/v1/users/view/1/
```
curl -i http://localhost:8080/api/v1/users/view/3 \
-H "accept: application/json" \
-H "Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzQwNTQ2NTQsInVzZXIiOiJ0ZXN0QGdtYWlsLmNvbSJ9.4K2vZk8E7likU-ipzatTwwGCoyI9O---hMMBBaysSDc"
```

ログインユーザー情報取得(上記ログイン後に取得するTokenを使用)
```bash
curl -i http://localhost:8080/api/v1/mypage/ \
-H "accept: application/json" \
-H "Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzQwNTQ2NTQsInVzZXIiOiJ0ZXN0QGdtYWlsLmNvbSJ9.4K2vZk8E7likU-ipzatTwwGCoyI9O---hMMBBaysSDc"

HTTP/1.1 200 OK
Access-Control-Allow-Credentials: true
Access-Control-Allow-Headers: Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization
Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE
Access-Control-Allow-Origin: http://localhost:3000
Content-Type: application/json
Date: Tue, 07 Sep 2021 18:53:25 GMT
Content-Length: 228

{"data":{"ID":4,"Name":"test user","Email":"foo@bar.com","Password":"$2a$10$ZQZd4OeWAkRqUesnw7LsUeFcDrcIhDWpxbgjvbMvZcE.RFpqQpZny","IsAdmin":false,"Deleted":null,"Created":"2021-09-07 18:11:49","Modified":"2021-09-07 18:11:49"}}
```

曲作成
```
curl http://localhost:8080/api/v1/songs/add/ \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzE3MTg3ODAsInVzZXIiOiJ0ZXN0QGdtYWlsLmNvbSJ9.JtLRyFfVg_Teqrjw_xICzxsTZKJL17DGSCsxA3lbgZQ" -d '{"registered_user_id":"2", "category_id":"2", "name":"hageの歌", "singer_name":"某兄貴", "composer_name":"ベートーベン", "source":"名探偵コナンのOP", "url":"http://ssss.ss", "is_anime_video_dam":true, "is_anime_video_joy":true, "is_official_video_dam":true, "is_official_video_joy":false, "start_singing":"YOSEI 夏が君としたくなる", "deleted":null, "created":"2000-02-02 00:00:00", "modified":"2000-02-02 00:00:00"}'

{"data":{"code":400,"message":"invalid character '\\x06' in string literal"}}
```