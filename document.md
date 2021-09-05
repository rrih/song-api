## example

ユーザー登録
```
curl -H 'Content-Type:application/json' -d '{"name":"test user", "email":"foo@bar.com", "password":"testtesttest"}' http://localhost:8080/api/v1/users/signup/
```

ログイン
```
curl -i -H 'Content-Type:application/json' -d '{"email":"foo@bar.com", "password":"testtesttest"}' http://localhost:8080/api/v1/users/signin/
HTTP/1.1 200 OK
Date: Sun, 05 Sep 2021 15:32:10 GMT
Content-Length: 145
Content-Type: text/plain; charset=utf-8

{"Token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzA4NTk1MzAsInVzZXIiOiJmb29AYmFyLmNvbSJ9.9G3gPXQ2ZzZ5pa7qMN1md8h55I5RDNFbWZfZOqEBgzE"}
```

ユーザー詳細取得
get /api/v1/users/view/1/
```
curl -i http://localhost:8080/api/v1/users/view/1/
```