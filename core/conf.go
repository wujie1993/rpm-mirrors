package core

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const (
	// DefaultConfPath is the default path to find configuration file
	DefaultConfPath = "/etc/rpm-mirrors/rpm-mirrors.conf"
)

const (
	// StorageTypeLocal indicates that the file is stored to the local path
	StorageTypeLocal = "local"
	// StorageTypeS3 indicates that the file is stored to the s3
	StorageTypeS3 = "s3"
)

// Conf defines the configuration struct
type Conf struct {
	SyncInterval int64        `yaml:"sync_interval" json:"sync_interval"`
	Mirrors      []ConfMirror `yaml:"mirrors" json:"mirrors"`
	Storage      ConfStorage  `yaml:"storage" json:"storage"`
}

// ConfMirror defines the configuration of mirror
type ConfMirror struct {
	Target string `yaml:"target" json:"target"`
	Source string `yaml:"source" json:"source"`
}

// ConfStorage defines the configuration of storage
type ConfStorage struct {
	Type     string        `yaml:"type" json:"type"`
	LocalDir string        `yaml:"local_dir" json:"local_dir"`
	S3       ConfStorageS3 `yaml:"s3" json:"s3"`
}

// ConfStorageS3 defines the configuration of s3 storage
type ConfStorageS3 struct {
	Ssl       bool   `yaml:"ssl" json:"ssl"`
	Endpoint  string `yaml:"endpoint" json:"endpoint"`
	AccessKey string `yaml:"access_key" json:"access_key"`
	SecretKey string `yaml:"secret_key" json:"secret_key"`
}

var globalConf Conf

// LoadConf read configuration from confPath and save in globalConf
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

// GetConf return the globalConf
func GetConf() Conf {
	return globalConf
}
