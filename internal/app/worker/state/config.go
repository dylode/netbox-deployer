package state

import "github.com/Khan/genqlient/graphql"

type Config struct {
	client graphql.Client
}

func NewConfig() Config {
	return Config{}
}

func (c Config) WithClient(client graphql.Client) Config {
	c.client = client
	return c
}
