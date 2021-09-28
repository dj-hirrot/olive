package infrastructure

import (
	"fmt"
	"time"

	"olive/interfaces/database"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewMySqlDb() database.SqlHandler {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"password",
		"db",
		"3306",
		"olive_db",
	)

	fmt.Println("---------- connectionString----------")
	fmt.Println(connectionString)
	fmt.Println("---------- connectionString----------")

	conn, err := open(connectionString, 30)

	if err != nil {
		panic(err.Error())
	}

	conn.LogMode(true)

	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn

	return sqlHandler
}

func open(path string, count uint) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", path)
	if err != nil {
		if count == 0 {
			return nil, fmt.Errorf("retry count over")
		}
		time.Sleep(time.Second)
		count--
		return open(path, count)
	}
	return db, nil
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

func (handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

func (handler *SqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}
