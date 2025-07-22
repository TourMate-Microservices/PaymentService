package noti

const (
	INTERNALL_ERR_MSG string = "There is something wrong in the system during the process. Please try again later."

	REPO_ERR_MSG string = "Error in %s repository at "

	MAIL_ERR_MSG string = "Error while generating mail at %s -  "

	GIN_ERR_MSG string = "Error while starting gin server in %s service - "

	NET_LISTENING_ERR_MSG string = "Error while listening on port %s - "
)

// Env
const (
	ENV_LOAD_ERR_MSG string = "Error while loading .env file in %s service - "

	ENV_SET_ERR_MSG string = "Error while setting environment variable %s with value %s - "
)

// Database
const (
	DB_CONNECTION_ERR_MSG string = "Error while connecting to database - "

	DB_MIGRATION_ERR_MSG string = "Error while migrating database - "

	DB_SET_CONNECTION_STRING_ERR_MSG string = "Error while setting database connection string - "
)

// gRPC
const (
	GRPC_CONNECTION_ERR_MSG string = "Error while connection to %s service gRPC - "
)
