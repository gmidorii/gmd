package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func cmdSave(c *cli.Context) error {
	cfg, err := loadCfg()
	if err != nil {
		return err
	}

	title, err := scan("Alias")
	if err != nil {
		return err
	}
	command, err := scan("Command")
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(cfg.SaveDir, title))
	if err != nil {
		return err
	}
	fmt.Fprint(file, command)

	return nil
}
