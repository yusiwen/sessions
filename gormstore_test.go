package sessions

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func createSQLite() (*gorm.DB, error) {
	f := "test.db"
	log.Println("db file -->", f)
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s?loc=Asia/Shanghai", f)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",                              // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
		Logger: logger.New(
			log.New(os.Stdout, "[GORM] ", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,       // Disable color
			}),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createMySQL() (*gorm.DB, error) {
	dsn := "db_test_user:db_test_user_password@tcp(lattepanda:3306)/db_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",                              // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
		Logger: logger.New(
			log.New(os.Stdout, "[GORM] ", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,       // Disable color
			}),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func TestGormStore_SQLite(t *testing.T) {
	db, err := createSQLite()
	if err != nil {
		t.Error(err)
	}
	store := NewGormStoreWithOptions(db, GormStoreOptions{}, []byte("EasyDarwin@2018"))
	store.Cleanup()
}

func TestGormStore_MySQL(t *testing.T) {
	db, err := createMySQL()
	if err != nil {
		t.Error(err)
	}
	store := NewGormStoreWithOptions(db, GormStoreOptions{}, []byte("EasyDarwin@2018"))
	store.Cleanup()
}
