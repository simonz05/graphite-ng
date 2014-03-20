package stores

import (
	"github.com/triplesek/graphite-ng/chains"
	"github.com/triplesek/graphite-ng/config"
	"github.com/triplesek/graphite-ng/metrics"
	"errors"
)

var InitFn = make(map[string]func(config config.Main) Store)

type Store interface {
	Add(metric metrics.Metric) (err error)
	Get(name string) (our_el *chains.ChainEl, err error)
	Has(name string) (found bool, err error)
	List() (list []string, err error)
}

var List = make(map[string]*Store)

func Init(config config.Main) (err error) {
	for _, key := range config.Stores {
		if constructor, ok := InitFn[key]; ok {
			store := constructor(config)
			List[key] = &store
		} else {
			return errors.New("no such store: " + key)
		}
	}
	return
}
