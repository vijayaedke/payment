package mysql

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBParams struct {
	mysqlClient *gorm.DB
}

type MysqlClientService interface {
	FindOne(data interface{}, key any) (interface{}, error)
	Find(data interface{}) (interface{}, error)
	Create(interface{}) (interface{}, error)
	Exists(data interface{}, paramName string, id any) bool
}

func InitMysqlClient() MysqlClientService {
	port := os.Getenv("MYSQL_PORT")
	host := os.Getenv("MYSQL_HOST")
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting DB : ", err)
	}
	log.Println("Database connection established.")

	return &DBParams{
		mysqlClient: db,
	}
}

func (m *DBParams) Find(data interface{}) (interface{}, error) {
	result := m.mysqlClient.Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return result, nil
}

func (m *DBParams) FindOne(data interface{}, key any) (interface{}, error) {
	result := m.mysqlClient.First(&data, key)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.Error == gorm.ErrRecordNotFound {
		return nil, gorm.ErrRecordNotFound
	}
	return data, nil
}

func (m *DBParams) Create(data interface{}) (interface{}, error) {
	result := m.mysqlClient.Create(data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func (m *DBParams) Exists(data interface{}, paramName string, id any) bool {
	query := fmt.Sprintf("`%s` = ?", paramName)
	err := m.mysqlClient.Where(query, id).Limit(1).Take(&data).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}
