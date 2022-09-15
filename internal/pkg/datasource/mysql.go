package datasource

import (
	"context"
	"fmt"
	"log"

	"github.com/elysiumyun/elysium/internal/app/model"
	"github.com/elysiumyun/elysium/internal/pkg/middleware"
	"github.com/elysiumyun/elysium/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlClient struct {
	ctx    context.Context
	client *gorm.DB
}

func (m *mysqlClient) GetClient() (*gorm.DB, context.Context) {
	return m.client, m.ctx
}

func (m *mysqlClient) openMysql() {
	m.ctx = context.Background()

	options := config.Cfg.Get()

	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true",
		options.Mysql.User,
		options.Mysql.Passwd,
		options.Mysql.Addr,
		options.Mysql.Database,
		options.Mysql.Charset,
	)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{
		Logger: middleware.GormLogger(),
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Mysql Use %v\n", options.Mysql.Addr)

	m.client = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci")

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
