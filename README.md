# gmd

## Overview
simple execute saved command.

## Usage

### Install
```
go get github.com/midorigreen/gmd
```

*requirement*  
[peco](https://github.com/peco/peco)  

## Command
### Save
```sh
% gmd save
```
register alias
```sh
% gmd save
Alias: db_hoge
```
register command
```sh
% gmd save
Alias: db_hoge
Command: mysql -h xxx -u xxx -p
```

### Save from history
```sh
% gmd hist
```

history latest 30
```
QUERY>                                 IgnoreCase [30 (1/1)]
git mv hsave.go hist.go
vi main.go
```

register alias
```
% gmd hist
Alias: main open
```

### Execute
```sh
% gmd exec
```

select alias 
```sh
QUERY>                                  IgnoreCase [5 (1/1)]
build
db_hoge
..
```

execute
```
% gmd exec
[CMD]: mysql -h xxx -u xxx -p
Enter password:
```

### Preview
```
% gmd list
+---------+------------------------+
|  ALIAS  |        COMMAND         |
+---------+------------------------+
| build   | go build               |
| db_hoge | mysql -h xxx -u xxx -p |
| log     | git log                |
| ls      | ls -l                  |
+---------+------------------------+
```


### Help
```sh
NAME:
   gcmd - simple exec saved cmd

USAGE:
   gmd [global options] command [command options] [arguments...]

VERSION:
   v1.0

COMMANDS:
     exec, e  execute saved command
     save, s  saved new command
     hist, h  saved new command from hisory
     list, l  output cmd list
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Todo
- [ ] delete command
- [ ] config edit command
