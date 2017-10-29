package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli"
)

func cmdExec(c *cli.Context) error {
	cfg, err := loadCfg()
	if err != nil {
		return err
	}
	f, err := os.Open(cfg.SaveDir)
	if err != nil {
		return err
	}
	defer f.Close()

	files, err := f.Readdirnames(-1)
	if err != nil {
		return err
	}
	cmdFile, err := selectPeco(files)
	if err != nil {
		return err
	}

	cmdStr, err := ioutil.ReadFile(filepath.Join(cfg.SaveDir, cmdFile))
	if err != nil {
		return err
	}
	cmd := exec.Command("sh", "-c", string(cmdStr))
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
