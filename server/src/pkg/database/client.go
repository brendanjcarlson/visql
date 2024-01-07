package database

import (
	"fmt"
	"log"

	"github.com/brendanjcarlson/visql/server/src/pkg/config"
	"github.com/brendanjcarlson/visql/server/src/pkg/domains/common"
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

	newClient.mustSync()

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

func (c *Client) mustSync() {
	c.mustSyncAccountsTable()
}

func (c *Client) checkTableExists(tableName string) (bool, error) {
	row := c.db.QueryRowx(`
        SELECT EXISTS (
            SELECT 1
            FROM information_schema.tables
            WHERE table_schema = 'public'
            AND table_name = $1
        )`,
		tableName,
	)
	if row.Err() != nil {
		return false, row.Err()
	}

	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		fmt.Println("error scanning in check table exists")
		return false, err
	}

	if !exists {
		return false, nil
	}

	return true, nil
}

func (c *Client) mustSyncAccountsTable() {
	exists, err := c.checkTableExists("accounts")
	if err != nil {
		log.Fatalf("failed to check if accounts table exists: %v\n", err.Error())
	}

	if !exists {
		_, cancel := common.NewQueryContext()
		defer cancel()

		c.db.MustExec(`
            CREATE TABLE accounts (
                id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                updated_at TIMESTAMP WITH TIME ZONE,
                last_login_at TIMESTAMP WITH TIME ZONE,
                login_count INTEGER DEFAULT 0,
                full_name TEXT NOT NULL,
                email TEXT NOT NULL UNIQUE,
                password TEXT NOT NULL
            )
        `)
	}
}
