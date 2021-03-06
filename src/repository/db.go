package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rmitsubayashi/bodyweight-server/main/config"
	"google.golang.org/appengine"
)

var (
	conn *sqlx.DB
	err  error
)

func NewDBConnection() (*sqlx.DB, error) {
	if conn != nil || err != nil {
		return conn, err
	}

	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	conn, err := sqlx.Open("mysql", formatConnectionString(cfg))
	if err != nil {
		return nil, fmt.Errorf("could not get connection %v", err)
	}
	if err := conn.Ping(); err != nil {
		conn.Close()
		return nil, fmt.Errorf("could not establish good connection: %v", err)
	}
	return conn, nil
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

	timeMapping := "?parseTime=true"
	if appengine.IsDevAppServer() {
		return fmt.Sprintf("%stcp([%s]:%d)/%s%s", cred, "localhost", 3306, cfg.DB.Schema, timeMapping)
	}

	return fmt.Sprintf("%sunix(/cloudsql/%s)/%s%s", cred, cfg.DB.Instance, cfg.DB.Schema, timeMapping)
}
