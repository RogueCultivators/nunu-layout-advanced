package repository

import (
	"context"
	"fmt"
	"github.com/go-nunu/nunu-layout-advanced/pkg/config"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
	"time"
)

const ctxTxKey = "TxKey"

type Repository struct {
	db     *gorm.DB
	rdb    *redis.Client
	logger *log.Logger
}

func NewRepository(db *gorm.DB, rdb *redis.Client, logger *log.Logger) *Repository {
	return &Repository{
		db:     db,
		rdb:    rdb,
		logger: logger,
	}
}

type Transaction interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
}

func NewTransaction(r *Repository) Transaction {
	return r
}

// DB return tx
// If you need to create a Transaction, you must call DB(ctx) and Transaction(ctx,fn)
func (r *Repository) DB(ctx context.Context) *gorm.DB {
	v := ctx.Value(ctxTxKey)
	if v != nil {
		if tx, ok := v.(*gorm.DB); ok {
			return tx
		}
	}
	return r.db.WithContext(ctx)
}

func (r *Repository) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ctxTxKey, tx)
		return fn(ctx)
	})
}

func NewDB(conf *config.Config, l *log.Logger) *gorm.DB {
	logger := zapgorm2.New(l.Logger)
	logger.SetAsDefault()
	var db *gorm.DB
	var err error
	switch conf.Data.Db.Type {
	case "postgres":
		db, err = gorm.Open(postgres.Open(conf.Data.Db.Dsn), &gorm.Config{})
	case "mysql":
		db, err = gorm.Open(mysql.Open(conf.Data.Db.Dsn), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(conf.Data.Db.Dsn), &gorm.Config{})
	default:
		panic(fmt.Sprintf("%s error: 数据库类型不支持", conf.Data.Db.Type))
	}
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	return db
}
func NewRedis(conf *config.Config) *redis.Client {
	println(conf.Data.Redis.Addr)
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Data.Redis.Addr,
		Password: conf.Data.Redis.Password,
		DB:       conf.Data.Redis.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}

	return rdb
}
