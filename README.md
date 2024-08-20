# scrapquiz

## backend

### 起動

```shell
docker-compose up --build
```

### API 確認

```shell
curl --include http://localhost:8080/v1/users

curl --include -X POST -H "Content-Type: application/json" -d '{"id": "01FVSHW3SER8977QCJBYZD9HAW", "name":"太郎", "avatar_url":"https://example.com/avatar.png"}' http://localhost:8080/v1/users

curl --include -X PUT -H "Content-Type: application/json" -d '{"name":"二郎", "avatar_url":"https://example.com/avatar.png"}' http://localhost:8080/v1/users/01FVSHW3SER8977QCJBYZD9HAW

curl --include http://localhost:8080/v1/users/01FVSHW3SER8977QCJBYZD9HAW

curl --include -X DELETE http://localhost:8080/v1/users/01FVSHW3SER8977QCJBYZD9HAW
```

### DB 接続

```shell
docker exec -it mysql /bin/sh
mysql -h mysql -P 3306 -u user -ppassword db
```

### マイグレーション

```shell
docker exec -it server /bin/sh
mysqldef -h mysql -p 3306 -u user -p password db < infrastructure/mysql/db/schema/schema.sql
```

### sqlc コード生成

```shell
sqlc generate
```

### モックコード生成

```shell
mockgen -package user -source domain/user/user_repository.go -destination domain/user/mock_user_repository.go
```

### テスト実行

```shell
go test -v ./...
```
