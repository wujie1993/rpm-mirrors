package core_test

import (
	"testing"

	"github.com/wujie1993/rpm-mirrors/core"
)

func TestMirrorJob(t *testing.T) {
	mirrorJob := &core.MirrorJob{
		Mirror: core.ConfMirror{
			Target: "centos/7.6.1810/os/x86_64/",
			Source: "rsync://rsync.mirrors.ustc.edu.cn/repo/centos/7.6.1810/os/x86_64/",
		},
		Storage: core.ConfStorage{
			Type:     core.StorageTypeLocal,
			LocalDir: "../repos/",
		},
	}
	if err := mirrorJob.Start(); err != nil {
		t.Fatal(err)
	}
}
