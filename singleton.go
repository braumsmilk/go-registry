package registry

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var r *Registry
var m sync.Mutex = sync.Mutex{}

func Set(reg Registry) {
	m.Lock()
	defer m.Unlock()

	r = &reg
}

func Get() *Registry {
	m.Lock()
	defer m.Unlock()
	if r == nil {
		panic("please init registry before using it")
	}

	return r
}

func InitDefault() error {
	m.Lock()
	defer m.Unlock()

	p := os.Getenv("REGISTRY_PATH")
	if p == "" {
		p = "/opt/config/registry.yaml"
	}

	return Init(p)
}

func Init(path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to init registry: %v", err)
	}

	var reg Registry
	err = yaml.Unmarshal(b, &reg)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config into registry: %v", err)
	}

	r = &reg

	return nil
}
