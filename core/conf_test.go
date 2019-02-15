package core_test

import (
	"encoding/json"
	"testing"

	"github.com/wujie1993/rpm-mirrors/core"
)

func TestLoadConf(t *testing.T) {
	conf, err := core.LoadConf("/root/go/src/github.com/wujie1993/rpm-mirrors/" + core.DefaultConfPath)
	if err != nil {
		t.Fatal(err)
	}
	datas, err := json.MarshalIndent(conf, "", "\t")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(datas))
}
