package postgres

import (
	"context"
	"fmt"
	"log"
	"proxyStoreServer/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	pool *pgxpool.Pool
}

func New(databaseConfig *config.DatabaseConfig) (*Postgres, error) {
	ctx := context.Background()

	connPool, err := pgxpool.NewWithConfig(ctx, databaseConfig.ConnectOptions())
	if err != nil {
		return nil, err
	}

	connection, err := connPool.Acquire(ctx)
	if err != nil {
		log.Fatal("Error while acquiring connection from the database pool!!")
	}
	defer connection.Release()

	err = connection.Ping(ctx)
	if err != nil {
		log.Fatal("Could not ping database")
	}

	fmt.Println("Connected to the database!!")

	return &Postgres{
		pool: connPool,
	}, nil
}

func (db *Postgres) Close() {
	db.Close()
}
