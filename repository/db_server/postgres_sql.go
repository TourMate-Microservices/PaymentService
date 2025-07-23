package dbserver

import (
	"os"
	"tourmate/payment-service/constant/env"
)

type postgresSql struct{}

func InitializeMsSQL() ISQLServer {
	return &postgresSql{}
}

// GetCnnStr implements ISQLServer.
func (m *postgresSql) GetCnnStr() string {
	return os.Getenv(env.POSTGRE_DB_CNNSTR)
}

// GetSQLServer implements ISQLServer.
func (m *postgresSql) GetSQLServer() string {
	return env.POSTGRE_SERVER
}
