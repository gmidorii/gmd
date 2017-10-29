package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

func cmdList(c *cli.Context) error {
	cfg, err := loadCfg()
	if err != nil {
		return err
	}

	sd, err := os.Open(cfg.SaveDir)
	if err != nil {
		return err
	}
	defer sd.Close()

	files, err := sd.Readdirnames(-1)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Alias", "Command"})

	for _, file := range files {
		cmd, err := ioutil.ReadFile(filepath.Join(cfg.SaveDir, file))
		if err != nil {
			return err
		}
		table.Append([]string{file, string(cmd)})
	}
	table.Render()
	return nil
}
