package config

type Config struct {
	Postgres      Postgres      `mapstructure:"postgres"`
	Server        Server        `mapstructure:"server"`
	ContainerRepo ContainerRepo `mapstructure:"containerRepo"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type ContainerRepo struct {
	URL      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
