package api

type Config struct {
	host string
	port int
}

func NewConfig() Config {
	return Config{
		host: "0.0.0.0",
		port: 8080,
	}
}

func (c Config) WithHost(host string) Config {
	c.host = host
	return c
}

func (c Config) WithPort(port int) Config {
	c.port = port
	return c
}
