## example

sign up
```
curl -H 'Content-Type:application/json' -d '{"name":"test user", "email":"foo@bar.com", "password":"testtesttest"}' http://localhost:8080/api/v1/users/signup/
```

sign in
```
curl -H 'Content-Type:application/json' -d '{"email":"foo@bar.com", "password":"testtesttest"}' http://localhost:8080/api/v1/users/signin/
```