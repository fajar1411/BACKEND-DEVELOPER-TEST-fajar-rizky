package prostgestsql

// import (
// 	"fmt"
// 	"log"
// 	"test/config"
// 	"test/migrate"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func InitDB(cfg *config.AppConfig) *gorm.DB {
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Cannot connect to DB")
// 	}
// 	migrate.MigrateDB(db)
// 	return db
// }
