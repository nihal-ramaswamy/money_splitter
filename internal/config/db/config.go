package db_config

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func New(options ...func(*Config)) *Config {
	config := &Config{}
	for _, option := range options {
		option(config)
	}
	return config
}

func WithHost(host string) func(*Config) {
	return func(c *Config) {
		c.Host = host
	}
}

func WithPort(port string) func(*Config) {
	return func(c *Config) {
		c.Port = port
	}
}

func WithUser(user string) func(*Config) {
	return func(c *Config) {
		c.User = user
	}
}

func WithPassword(password string) func(*Config) {
	return func(c *Config) {
		c.Password = password
	}
}

func WithDbname(dbname string) func(*Config) {
	return func(c *Config) {
		c.Dbname = dbname
	}
}
