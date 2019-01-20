package repository

import (
	"database/sql"
	"fmt"

	"github.com/rmitsubayashi/bodyweight-server/main/config"
	"google.golang.org/appengine"
)

type DB struct {
	conn *sql.DB
}

var (
	db  *DB
	err error
)

func NewDBConnection() (*DB, error) {
	if db != nil || err != nil {
		return db, err
	}

	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	conn, err := sql.Open("mysql", formatConnectionString(cfg))
	if err != nil {
		return nil, fmt.Errorf("could not get connection %v", err)
	}
	if err := conn.Ping(); err != nil {
		conn.Close()
		return nil, fmt.Errorf("could not establish good connection: %v", err)
	}

	db := &DB{
		conn: conn,
	}
	err = nil
	return db, nil
}

func formatConnectionString(cfg *config.Config) string {
	var cred string
	// the logic copied from GCP sample
	// https://github.com/GoogleCloudPlatform/golang-samples/blob/master/getting-started/bookshelf/db_mysql.go#L68
	if cfg.DB.Username != "" {
		cred = cfg.DB.Username
		if cfg.DB.Password != "" {
			cred = cred + ":" + cfg.DB.Password
		}
		cred = cred + "@"
	}

	if appengine.IsDevAppServer() {
		return fmt.Sprintf("%stcp([%s]:%d)/%s", cred, "localhost", 3306, cfg.DB.Schema)
	}

	return fmt.Sprintf("%sunix(/cloudsql/%s)/%s", cred, cfg.DB.Instance, cfg.DB.Schema)
}
