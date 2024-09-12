package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"scrapquiz/config"
	"scrapquiz/infrastructure/postgresql/db/dbgen"

	_ "github.com/lib/pq"
)

const (
	maxRetries = 5
	delay      = 5 * time.Second
)

var (
	once      sync.Once
	query     *dbgen.Queries
	readQuery *dbgen.Queries
	dbcon     *sql.DB
)

func GetQuery(ctx context.Context) *dbgen.Queries {
	txq := getQueriesWithContext(ctx)
	if txq != nil {
		return txq
	}
	return query
}

func GetReadQuery() *dbgen.Queries {
	return readQuery
}

func SetQuery(q *dbgen.Queries) {
	query = q
}

func SetReadQuery(q *dbgen.Queries) {
	readQuery = q
}

func GetDB() *sql.DB {
	return dbcon
}

func SetDB(d *sql.DB) {
	dbcon = d
}

func NewMainDB(cnf config.DBConfig) {
	once.Do(func() {
		dbcon, err := connect(
			cnf.User,
			cnf.Password,
			cnf.Host,
			cnf.Port,
			cnf.Name,
		)
		if err != nil {
			panic(err)
		}
		q := dbgen.New(dbcon)
		SetQuery(q)
		SetDB(dbcon)
	})
}

func connect(user string, password string, host string, port string, name string) (*sql.DB, error) {
	for i := 0; i < maxRetries; i++ {
		// connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, name)
		connect := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, name)
		db, err := sql.Open("postgres", connect)
		if err != nil {
			return nil, fmt.Errorf("could not open db: %w", err)
		}

		err = db.Ping()
		if err == nil {
			return db, nil
		}

		log.Printf("could not connect to db: %v", err)
		log.Printf("retrying in %v seconds...", delay/time.Second)
		time.Sleep(delay)
	}

	return nil, fmt.Errorf("could not connect to db after %d attempts", maxRetries)
}

type CtxKey string

const (
	QueriesKey CtxKey = "queries"
)

func WithQueries(ctx context.Context, q *dbgen.Queries) context.Context {
	return context.WithValue(ctx, QueriesKey, q)
}

func getQueriesWithContext(ctx context.Context) *dbgen.Queries {
	queries, ok := ctx.Value(QueriesKey).(*dbgen.Queries)
	if !ok {
		return nil
	}
	return queries
}
