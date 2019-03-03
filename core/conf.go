package core

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	DefaultConfPath = "/etc/rpm-mirrors/rpm-mirrors.conf"

	StorageTypeLocal = "local"
	StorageTypeS3    = "s3"
)

type Conf struct {
	SyncInterval int64        `yaml:"sync_interval" json:"sync_interval"`
	Mirrors      []ConfMirror `yaml:"mirrors" json:"mirrors"`
	Storage      ConfStorage  `yaml:"storage" json:"storage"`
}

type ConfMirror struct {
	Target string `yaml:"target" json:"target"`
	Source string `yaml:"source" json:"source"`
}

type ConfStorage struct {
	Type     string        `yaml:"type" json:"type"`
	LocalDir string        `yaml:"local_dir" json:"local_dir"`
	S3       ConfStorageS3 `yaml:"s3" json:"s3"`
}

type ConfStorageS3 struct {
	Ssl       bool   `yaml:"ssl" json:"ssl"`
	Endpoint  string `yaml:"endpoint" json:"endpoint"`
	AccessKey string `yaml:"access_key" json:"access_key"`
	SecretKey string `yaml:"secret_key" json:"secret_key"`
}

var globalConf Conf

func LoadConf(confPath string) (Conf, error) {
	if confPath == "" {
		confPath = DefaultConfPath
	}
	datas, err := ioutil.ReadFile(confPath)
	if err != nil {
		return globalConf, err
	}
	if err := yaml.Unmarshal(datas, &globalConf); err != nil {
		return globalConf, err
	}
	return globalConf, nil
}

func GetConf() Conf {
	return globalConf
}
