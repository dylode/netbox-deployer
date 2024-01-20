package state

import (
	"github.com/Khan/genqlient/graphql"
	"github.com/luthermonson/go-proxmox"
)

type Config struct {
	client        graphql.Client
	proxmoxClient *proxmox.Client
}

func NewConfig() Config {
	return Config{}
}

func (c Config) WithClient(client graphql.Client) Config {
	c.client = client
	return c
}

func (c Config) WithProxmoxClient(client *proxmox.Client) Config {
	c.proxmoxClient = client
	return c
}
