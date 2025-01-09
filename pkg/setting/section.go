package setting

type Config struct {
	Mysql MysqlSetting `mapstructure:"mysql"`
}

type MysqlSetting struct {
	Username        string `mapstructure:"user"`
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbName"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifetime"`
}
