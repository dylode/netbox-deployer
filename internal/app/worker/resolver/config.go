package resolver

type Config struct {
	graphqlURL string
}

func NewConfig() Config {
	return Config{
		graphqlURL: "",
	}
}

func (c Config) WithGraphqlURL(url string) Config {
	c.graphqlURL = url
	return c
}
