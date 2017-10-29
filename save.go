package main

import (
	"bufio"
	"errors"
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

	title, err := scan("Title")
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

func scan(description string) (string, error) {
	fmt.Printf("%s: ", description)
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return "", errors.New("canceld")
	}
	if scanner.Err() != nil {
		return "", scanner.Err()
	}
	return scanner.Text(), nil
}
