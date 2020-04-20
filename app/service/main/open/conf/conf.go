package conf

import (
	"errors"
	"flag"

	"go-common/library/cache/memcache"
	"go-common/library/conf"
	"go-common/library/database/orm"
	"go-common/library/log"
	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/rpc/warden"
	"go-common/library/net/trace"
	xtime "go-common/library/time"

	"github.com/BurntSushi/toml"
)

// Config .
type Config struct {
	App    *bm.App
	Log    *log.Config
	Tracer *trace.Config
	ORM    *orm.Config
	// Uname load ticker
	AppTicker xtime.Duration
	// mc
	Memcache     *memcache.Config
	WardenServer *warden.ServerConfig
}

var (
	confPath string
	client   *conf.Client
	// Conf config
	Conf = &Config{}
)

// init() .
func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

// Init .
func Init() (err error) {
	if confPath != "" {
		return local()
	}
	return remote()
}

// local .
func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

// remote .
func remote() (err error) {
	if client, err = conf.New(); err != nil {
		return
	}
	if err = load(); err != nil {
		return
	}
	go func() {
		for range client.Event() {
			log.Info("config reload")
		}
	}()
	return
}

// load .
func load() (err error) {
	var (
		s       string
		ok      bool
		tmpConf *Config
	)
	if s, ok = client.Toml2(); !ok {
		return errors.New("load config center error")
	}
	if _, err = toml.Decode(s, &tmpConf); err != nil {
		return errors.New("could not decode config")
	}
	*Conf = *tmpConf
	return
}
