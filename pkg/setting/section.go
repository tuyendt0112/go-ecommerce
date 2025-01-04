package setting

type Config struct {
}

type MysqlSetting struct {
	User string `mapstructure:"user"`
}
