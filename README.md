# scrapquiz

## backend

### 起動

```shell
docker-compose up --build
```

### API 確認

```shell
curl --include -X POST -H "Content-Type: application/json" -d @user.json "http://localhost:8080/v1/users"
curl --include -X PUT -H "Content-Type: application/json" -d '{"name":"二郎", "avatar_url":"https://example.com/avatar.png"}' "http://localhost:8080/v1/users/01FVSHW3SER8977QCJBYZD9HAW"
curl --include "http://localhost:8080/v1/users"
curl --include "http://localhost:8080/v1/users/01FVSHW3SER8977QCJBYZD9HAW"
curl --include -X DELETE "http://localhost:8080/v1/users/01FVSHW3SER8977QCJBYZD9HAW"

curl --include -X POST -H "Content-Type: application/json" -d @quiz.json "http://localhost:8080/v1/quizzes"
curl --include "http://localhost:8080/v1/quizzes?limit=10&offset=0"
curl --include "http://localhost:8080/v1/quizzes/01J62HMVN1DARH0YR7SR996QRP"
curl --include -X DELETE "http://localhost:8080/v1/quizzes/01J62HMVN1DARH0YR7SR996QRP"
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
mockgen -package quiz -source domain/quiz/quiz_repository.go -destination domain/quiz/mock_quiz_repository.go
mockgen -package quiz -source usecase/quiz/quiz_query_service.go -destination usecase/quiz/mock_quiz_query_service.go
mockgen -package answer -source domain/answer/answer_repository.go -destination domain/answer/mock_answer_repository.go
mockgen -package answer -source usecase/answer/answer_query_service.go -destination usecase/answer/mock_answer_query_service.go
```

### テスト実行

```shell
go test -v ./...
```
