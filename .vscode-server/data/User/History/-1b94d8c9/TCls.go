package util

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ElasticsearchInfo struct {
	Address string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type GCSInfo struct