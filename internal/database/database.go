package database

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/thiagocprado/golang-api-structure/internal/env"
)

var MySqlDB *sql.DB

func NewMySqlConn() error {
	mysqlConnString := getMySqlConnString()

	db, err := sql.Open(env.DBKindMySql, mysqlConnString)
	if err != nil {
		slog.Error("Erro ao conectar no mysql!", slog.String("err", err.Error()))
		return err
	}

	err = db.Ping()
	if err != nil {
		slog.Error("Erro ao pingar mysql!", slog.String("err", err.Error()))
		return err
	}

	MySqlDB = db

	return nil
}

func getMySqlConnString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		env.MySqlDbUser,
		env.MySqlDbPass,
		env.MySqlDbHost,
		env.MySqlDbPort,
		env.MySqlDbName,
		"multiStatements=true",
	)
}
