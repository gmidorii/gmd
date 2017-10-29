package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
)

const (
	VERSION = "0.1"
)

var cmds = []cli.Command{
	{
		Name:    "exec",
		Aliases: []string{"e"},
		Usage:   "execute saved command",
		Action:  cmdExec,
	},
	{
		Name:    "save",
		Aliases: []string{"s"},
		Usage:   "saved new command",
		Action:  cmdSave,
	},
	{
		Name:    "hsave",
		Aliases: []string{"h"},
		Usage:   "saved new command from hisory",
		Action:  cmdHSave,
	},
	{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "output cmd list",
		Action:  cmdList,
	},
}

type Config struct {
	SaveDir     string
	HistoryFile string
}

func loadCfg() (Config, error) {
	// TODO: fix for windows
	dir := filepath.Join(os.Getenv("HOME"), ".config", "gmd")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return Config{}, fmt.Errorf("failed create dir: %v", err)
	}

	var cfg Config
	file := filepath.Join(dir, "config.toml")
	_, err := os.Stat(file)
	// exist file
	if err == nil {
		_, err := toml.DecodeFile(file, &cfg)
		if err != nil {
			return Config{}, err
		}
		return cfg, nil
	}

	// init config file
	cfg.SaveDir = filepath.Join(dir, "_saved")
	err = os.Mkdir(cfg.SaveDir, 0700)
	if err != nil {
		return Config{}, err
	}
	cfg.HistoryFile = filepath.Join(os.Getenv("HOME"), ".zsh_history")
	f, err := os.Create(file)
	if err != nil {
		return Config{}, err
	}
	toml.NewEncoder(f).Encode(cfg)
	return cfg, nil
}

func selectPeco(list []string) (string, error) {
	sreader := strings.NewReader(strings.Join(list, "\n"))

	cmd := exec.Command("sh", "-c", "peco")
	var buf bytes.Buffer
	cmd.Stderr = os.Stderr
	cmd.Stdin = sreader
	cmd.Stdout = &buf
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(buf.String()), err
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

func run() int {
	app := cli.NewApp()
	app.Name = "gcmd"
	app.Usage = "simple exec saved cmd"
	app.Version = VERSION
	app.Commands = cmds
	return msg(app.Run(os.Args))
}

func msg(err error) int {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
