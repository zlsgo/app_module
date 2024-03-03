package database

type (
	Options struct {
		Sqlite       Sqlite   `json:"sqlite"`
		Driver       string   `json:"driver"`
		MySQL        Mysql    `json:"mysql"`
		Postgres     Postgres `json:"postgres"`
		disableWrite bool     `json:"-"`
		Model        Model    `json:"model"`
	}

	Mysql struct {
		Host       string `json:"host"`
		User       string `json:"user"`
		Password   string `json:"password"`
		DBName     string `json:"db_name"`
		Parameters string `json:"parameters"`
		Charset    string `json:"charset"`
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
		Parameters string `json:"parameters"`
	}
	Model struct {
		DelteColumn bool `json:"delete_column"`
	}
)

func (Options) ConfKey() string {
	return "database"
}

func (o Options) DisableWrite() bool {
	return o.disableWrite
}
