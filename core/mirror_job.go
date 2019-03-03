package core

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
)

type MirrorJob struct {
	Mirror  ConfMirror
	Storage ConfStorage
}

func (job *MirrorJob) Start() error {
	// Init storage
	switch job.Storage.Type {
	case StorageTypeLocal:
		targetDir := path.Dir(job.Storage.LocalDir + "/" + job.Mirror.Target)
		if err := os.MkdirAll(targetDir, 0774); err != nil {
			return err
		}
		cmd := exec.Command("/usr/bin/rsync", "-avrtH", "--delete", job.Mirror.Source, targetDir)
		log.Println(cmd.Args)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}
		if err := cmd.Start(); err != nil {
			return err
		}
		reader := bufio.NewReader(stdout)
		// Read and print standard output
		for {
			datas, _, err := reader.ReadLine()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			log.Println(string(datas))
		}
		if err := cmd.Wait(); err != nil {
			return err
		}
		if err := cmd.Process.Kill(); err != nil {
			return err
		}
	case StorageTypeS3:
		return errors.New("storage type not implemented: " + job.Storage.Type)
	default:
		return errors.New("storage type unsupported: " + job.Storage.Type)
	}

	return nil
}
