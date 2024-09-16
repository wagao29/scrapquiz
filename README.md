# scrapquiz

## backend

### 起動

```shell
docker-compose up --build
```

### API 確認

```shell
curl --include -X POST -H "Content-Type: application/json" -H "x-api-key:local-api-key" -d @user.json "http://localhost:8080/v1/users"
curl --include -X PUT -H "Content-Type: application/json" -H "x-api-key:local-api-key" -d '{"name":"二郎", "avatarUrl":"https://example.com/avatar.png"}' "http://localhost:8080/v1/users/123456789012345678901"
curl --include -H "x-api-key:local-api-key" "http://localhost:8080/v1/users"
curl --include -H "x-api-key:local-api-key" "http://localhost:8080/v1/users/123456789012345678901"
curl --include -X DELETE -H "x-api-key:local-api-key" "http://localhost:8080/v1/users/123456789012345678901"

curl --include -X POST -H "Content-Type: application/json" -H "x-api-key:local-api-key" -d @quiz.json "http://localhost:8080/v1/quizzes"
curl --include -H "x-api-key:local-api-key" "http://localhost:8080/v1/quizzes?limit=10&offset=0"
curl --include -H "x-api-key:local-api-key" "http://localhost:8080/v1/quizzes/counts"
curl --include -H "x-api-key:local-api-key" "http://localhost:8080/v1/quizzes/01J7ECNB9SH0YYJJK6YB6F9V9P"
curl --include -X DELETE -H "x-api-key:local-api-key" "http://localhost:8080/v1/quizzes/01J7ECNB9SH0YYJJK6YB6F9V9P"

curl --include -X POST -H "Content-Type: application/json" -H "x-api-key:local-api-key" -d @answer.json "http://localhost:8080/v1/quizzes/01J7ECNB9SH0YYJJK6YB6F9V9P/answers"
curl --include -H "x-api-key:local-api-key" "http://localhost:8080/v1/quizzes/01J7ECNB9SH0YYJJK6YB6F9V9P/answer_counts"
```

### DB 接続

```shell
# MySQL
docker exec -it mysql /bin/sh
mysql -h mysql -P 3306 -u user -ppassword db

# PostgreSQL
docker exec -it postgres /bin/sh
psql -U user
```

### マイグレーション

```shell
# MySQL
mysqldef -h mysql -p 3306 -u user -p password db < infrastructure/mysql/db/schema/schema.sql

# PostgreSQL
migrate create -ext sql -dir infrastructure/postgresql/db/migrations -seq create_users_tables
migrate --path infrastructure/postgresql/db/migrations --database 'postgresql://user:password@localhost:5432/db?sslmode=disable' -verbose up
```

### sqlc コード生成

```shell
sqlc generate
```

### モックコード生成

```shell
mockgen -package user -source domain/user/user_repository.go -destination domain/user/mock_user_repository.go
mockgen -package quiz -source domain/quiz/quiz_repository.go -destination domain/quiz/mock_quiz_repository.go
mockgen -package quiz -source application/quiz/quiz_query_service.go -destination application/quiz/mock_quiz_query_service.go
mockgen -package answer -source domain/answer/answer_repository.go -destination domain/answer/mock_answer_repository.go
mockgen -package answer -source application/answer/answer_query_service.go -destination application/answer/mock_answer_query_service.go
```

### テスト実行

```shell
go test -v ./...
```

## frontend

### 起動

```shell
npm install
npm run dev
```

### shadcn/ui コンポーネントのインストール

```shell
npx shadcn-ui@latest add button
```
