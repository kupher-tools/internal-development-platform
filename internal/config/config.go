package config

type Config struct {
	Server        Server        `mapstructure:"server"`
	ContainerRepo ContainerRepo `mapstructure:"containerRepo"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type ContainerRepo struct {
	URL      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
