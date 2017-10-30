package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func cmdDel(c *cli.Context) error {
	cfg, err := loadCfg()
	if err != nil {
		return err
	}

	var file string
	switch len(c.Args()) {
	case 0:
		f, err := os.Open(cfg.SaveDir)
		if err != nil {
			return err
		}

		files, err := f.Readdirnames(-1)
		if err != nil {
			return err
		}
		file, err = selectPeco(files)
		if err != nil {
			return err
		}
	default:
		file = c.Args().First()
	}

	err = os.Remove(filepath.Join(cfg.SaveDir, file))
	if err != nil {
		return fmt.Errorf("failed delete command: %v", err)
	}
	fmt.Printf("Success Delete Command: %s\n", file)
	return nil
}
