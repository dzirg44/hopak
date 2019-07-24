package driver

import (
	"fmt"

	"github.com/hopak/server/config"
	"github.com/hopak/server/models"
	_ "github.com/go-sql-driver/mysql" // this is import required for gorm
	"github.com/jinzhu/gorm"
)

// ConnectSQL connects to a database
func ConnectSQL(*config.DBConfig) *gorm.DB {
	Db2Config := config.GetConfig()
	// username:password@tcp(host:port)/dbname
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", Db2Config.Username, Db2Config.Password, Db2Config.Host, Db2Config.Port, Db2Config.DBName)
	db, err := gorm.Open(Db2Config.Driver, dbSource)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	db.Model(&models.User{}).Related(&models.Post{})
	db.Debug().AutoMigrate(&models.User{})
	db.Debug().AutoMigrate(&models.Post{})
	db.Model(&models.Post{}).AddForeignKey("user_id", "users(user_id)", "RESTRICT", "RESTRICT")
	return db
}
