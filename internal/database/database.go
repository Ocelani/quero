package database

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connector interface {
	GetDatabase() *gorm.DB
}

type Connection struct {
	Host     string
	Port     string
	User     string
	Password string
	NameDB   string
	SSL      string
	TimeZone string
}

func NewConnection(host, port, user, password, namedb, ssl, timezone string) *Connection {
	return &Connection{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		NameDB:   namedb,
		SSL:      ssl,
		TimeZone: timezone,
	}
}

func (c *Connection) GetDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(c.String()), &gorm.Config{})
	if err != nil {
		time.Sleep(time.Millisecond * 200)
		c.GetDatabase()
	}
	return db
}

func (c *Connection) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s  sslmode=%s TimeZone=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_SSL"),
		os.Getenv("DATABASE_TZ"),
	)
}
