package maple

import (
	"time"
)

type config struct {
	// Server
	Server_Port			string
	Server_Timeout 	time.Duration  // in secs

	Db_MaxIdle			int
	Db_MaxConn			int
	Db_Alias				string
	Db_Auth					string

	DebugMode bool
}

func NewConfig() *config {
	return &config {
		Server_Port: ":3001",
		Server_Timeout: 10 * time.Second,

		Db_MaxIdle: 30,
		Db_MaxConn: 30,
		Db_Alias: "default",
		Db_Auth: "jenesis:rakaboon88@/baboo?charset=utf8",

		DebugMode: true,
	}
}

