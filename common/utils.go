package common

import (
	"github.com/jinzhu/configor"
	"sync"
)

const (
	DiscoveryTopic = "sonmDicso.v1"
	SeekerPort     = ":12345"
	CasterPort     = ":12346"
)

type bootNodeConfig struct {
	Bootnodes []string `required:"true" yaml:"bootnodes"`
}

// LoadBootnodes loads yaml config with bootnodes, returns bootnodes enode add
func LoadBootnodes() []string {
	cfg := &bootNodeConfig{}
	err := configor.Load(cfg, "./bootnodes.yaml")
	if err != nil {
		panic(err.Error())
	}

	return cfg.Bootnodes
}

// counter store anonymous and signed message count
type counter struct {
	sync.Mutex
	signed uint64
	anon   uint64
}

func NewCounter() *counter {
	return &counter{signed: 0, anon: 0}
}

func (c *counter) GetAnon() uint64 {
	c.Lock()
	defer c.Unlock()
	return c.anon
}

func (c *counter) GetSigned() uint64 {
	c.Lock()
	defer c.Unlock()
	return c.signed
}

func (c *counter) AddAnon() {
	c.Lock()
	defer c.Unlock()
	c.anon++
}

func (c *counter) AddSigned() {
	c.Lock()
	defer c.Unlock()
	c.signed++
}
