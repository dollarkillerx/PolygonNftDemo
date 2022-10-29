package utils

import (
	"fmt"

	"github.com/dollarkillerx/PolygonNftDemo/internal/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPgSQL(pgConf *conf.PgSQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", pgConf.Host, pgConf.User, pgConf.Password, pgConf.DbName, pgConf.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
