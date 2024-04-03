package database

type (
	Options struct {
		Sqlite       *Sqlite   `json:"sqlite,omitempty"`
		Driver       string    `json:"driver,omitempty"`
		MySQL        *Mysql    `json:"mysql,omitempty"`
		Postgres     *Postgres `json:"postgres,omitempty"`
		disableWrite bool      `json:"-"`
		Mode         *Mode     `json:"mode,omitempty"`
	}

	Mysql struct {
		Host       string `json:"host"`
		User       string `json:"user"`
		Password   string `json:"password"`
		DBName     string `json:"db_name"`
		Parameters string `json:"parameters,omitempty"`
		Charset    string `json:"charset,omitempty"`
		Port       int    `json:"port"`
	}

	Postgres struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"db_name"`
		SSLMode  string `json:"ssl_mode"`
		Port     int    `json:"port"`
	}

	Sqlite struct {
		Path       string `json:"path"`
		Parameters string `json:"parameters,omitempty"`
	}

	Mode struct {
		DelteColumn bool `json:"delete_column,omitempty"`
	}
)

func (Options) ConfKey() string {
	return "database"
}

func (o Options) DisableWrite() bool {
	return o.disableWrite
}
