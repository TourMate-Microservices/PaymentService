package dbserver

type ISQLServer interface {
	GetSQLServer() string
	GetCnnStr() string
}
