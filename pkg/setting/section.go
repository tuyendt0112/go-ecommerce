package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Mysql  MysqlSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
	Redis  RedisSetting `mapstructure:"redis"`
	
}

type ServerSetting struct {
	Port int `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
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

type LoggerSetting struct {
	Log_level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_size      int    `mapstructure:"max_size"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_age       int    `mapstructure:"max_age"`
	Compress      bool   `mapstructure:"compress"`
}


type RedisSetting struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int `mapstructure:"database"`
}