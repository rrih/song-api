# song-api

[![library-test](https://github.com/rrih/song-api/actions/workflows/build.yml/badge.svg)](https://github.com/rrih/song-api/actions/workflows/build.yml)

API server for song management written in Go

## About API

TBD

|Content-Type|
|-|
|application/json|

|                            |   method   |          endpoint          | body | header |
|----------------------------|------------|-----------------------|------|-|
|Register An User Account    |    POST    | `/api/v1/auth/signup/`|`{"name":"test user", "email": "foo@bar.com", "password": "testtest"}`||
|Log-in as a registered user |    POST    | `/api/v1/auth/login/` |`{"email":"foo@bar.com", "password":"testtest"}`||
|Log-out as a registered user |    POST    | `/api/v1/auth/login/` ||`{"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzE3MjE5OTMsInVzZXIiOiJ0ZXN0QGdtYWlsLmNvbSJ9.IybN5Wrq8s1SqNSKK5C00hoZtc2Qgy_aCgNZS4oVOdo"}`|
|Get An User Account       |GET|`/api/v1/users/view/:id/`|||
|Update An User Account    |PUT|`/api/v1/users/view/:id/`|||
|Delete An User Account    |DELETE|`/api/v1/users/view/:id/`|||
|Create Song Record|POST|`/api/v1/songs/add/`|||


The author is [rrih_dev](https://twitter.com/rrih_dev)