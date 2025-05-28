package data

import (
	"fmt"
	"github.com/go-kratos/kratos-layout/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	gormhelper "github.com/nuominmin/gorm-helper"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db: db,
	}, cleanup, nil
}

func NewDB(c *conf.Data) (*gorm.DB, func(), error) {
	db, cleanup, err := newDB(c.Database)
	if err != nil {
		return nil, nil, err
	}

	// Automatic migration
	if err = db.AutoMigrate(); err != nil {
		return nil, nil, err
	}

	return db, cleanup, err
}

func newDB(database *conf.Data_Database) (*gorm.DB, func(), error) {
	var db *gorm.DB
	var cleanup func()
	var err error
	opts := []gormhelper.ConnOption{
		gormhelper.WithConnLogLevel(logger.LogLevel(database.LogLevel)),
		gormhelper.WithMaxIdleConns(int(database.MaxIdleConns)),
		gormhelper.WithMaxOpenConns(int(database.MaxOpenConns)),
		gormhelper.WithConnMaxLifetime(database.ConnMaxLifetime.AsDuration()),
	}

	switch database.Driver {
	case "mysql":
		db, cleanup, err = gormhelper.ConnectMysql(database.Source, opts...)
	case "sqlite":
		db, cleanup, err = gormhelper.ConnectSqlite(database.Source, opts...)
	default:
		return nil, func() {}, fmt.Errorf("unsupported driver %s", database.Driver)
	}
	if err != nil {
		return nil, nil, err
	}

	return db, cleanup, err
}
