package config

import (
	"arbitrage_watcher/internal/model"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Interval      time.Duration `yaml:"interval"`
	PrometheusApi string        `yaml:"prometheus_api"`
	Webhooks      struct {
		LineNotifyToken string `yaml:"line_notify_token"`
	} `yaml:"webhooks"`
	Pairs []model.Pair `yaml:"pairs"`
}

func Load(path string) *Config {
	var config Config
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(dat, &config)
	if err != nil {
		log.Fatal(err)
	}
	return &config
}
