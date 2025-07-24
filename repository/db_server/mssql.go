package dbserver

import (
	"os"
	"tourmate/payment-service/constant/env"
)

type msSQL struct{}

func InitializeMsSQL() ISQLServer {
	return &msSQL{}
}

// GetCnnStr implements ISQLServer.
func (m *msSQL) GetCnnStr() string {
	return os.Getenv(env.POSTGRE_DB_CNNSTR)
}

// GetSQLServer implements ISQLServer.
func (m *msSQL) GetSQLServer() string {
	return env.POSTGRE_SERVER
}
