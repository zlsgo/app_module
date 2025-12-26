package database

type (
	// Options 定义模块配置
	Options struct {
		Sqlite       *Sqlite   `json:"sqlite,omitempty"`
		MySQL        *Mysql    `json:"mysql,omitempty"`
		Postgres     *Postgres `json:"postgres,omitempty"`
		Mode         *Mode     `json:"mode,omitempty"`
		Driver       string    `json:"driver,omitempty"`
		disableWrite bool      `json:"-"`
	}

	// Mysql 定义 MySQL 配置
	Mysql struct {
		Host       string `json:"host"`
		User       string `json:"user"`
		Password   string `json:"password"`
		DBName     string `json:"db_name"`
		Parameters string `json:"parameters,omitempty"`
		Charset    string `json:"charset,omitempty"`
		Port       int    `json:"port"`
	}

	// Postgres 定义 PostgreSQL 配置
	Postgres struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"db_name"`
		SSLMode  string `json:"ssl_mode"`
		Port     int    `json:"port"`
	}

	// Sqlite 定义 SQLite 配置
	Sqlite struct {
		Path       string `json:"path"`
		Parameters string `json:"parameters,omitempty"`
	}

	// Mode 定义模式配置
	Mode struct {
		DelteColumn bool `json:"delete_column,omitempty"`
	}
)

// ConfKey 返回配置键
func (Options) ConfKey() string {
	return "database"
}

// DisableWrite 返回是否禁写
func (o Options) DisableWrite() bool {
	return o.disableWrite
}
