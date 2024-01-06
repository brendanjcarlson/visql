package database

import (
	"fmt"
	"log"

	"github.com/brendanjcarlson/visql/server/src/pkg/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Client struct {
	db *sqlx.DB
}

type dbConfig struct {
	driver string
	dsn    *dsn
}

type dsn struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
	sslmode  string
}

func (d *dsn) string() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.host, d.port, d.user, d.password, d.dbname, d.sslmode)
}

func loadDbConfig() *dbConfig {
	return &dbConfig{
		driver: config.MustGet("DB_DRIVER"),
		dsn: &dsn{
			host:     config.MustGet("DB_HOST"),
			port:     config.MustGet("DB_PORT"),
			user:     config.MustGet("DB_USER"),
			password: config.MustGet("DB_PASSWORD"),
			dbname:   config.MustGet("DB_NAME"),
			sslmode:  config.GetOrDefault("DB_SSLMODE", "disable"),
		},
	}
}

func MustConnect() *Client {
	cfg := loadDbConfig()

	db, err := sqlx.Connect(cfg.driver, cfg.dsn.string())
	if err != nil {
		log.Fatalf("failed to connect to database: %v\n", err.Error())
	}

	newClient := &Client{
		db: db,
	}

	return newClient
}

func (c *Client) MustClose() {
	err := c.db.Close()
	if err != nil {
		log.Fatalf("failed to close database connection: %v\n", err.Error())
	}
}

func (c *Client) DB() *sqlx.DB {
	return c.db
}
