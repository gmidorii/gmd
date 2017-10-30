package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

func cmdHist(c *cli.Context) error {
	cfg, err := loadCfg()
	if err != nil {
		return err
	}
	cmd := exec.Command("tail", "-n", "30", cfg.HistoryFile)

	var buf bytes.Buffer
	cmd.Stderr = os.Stderr
	cmd.Stdout = &buf
	if err := cmd.Run(); err != nil {
		return err
	}

	histories := strings.Split(buf.String(), "\n")
	cmds := make([]string, len(histories))
	for i, history := range histories {
		if strings.Contains(history, ";") {
			cmds[i] = strings.Split(history, ";")[1]
			continue
		}
		cmds[i] = history
	}

	command, err := selectPeco(cmds)
	title, err := scan("Alias")
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
