# wpkgtcli
![GitHub](https://img.shields.io/github/license/henior2/wpkgtcli)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/henior2/wpkgtcli)

Open source terminal based WPKG Dashboard CLI written in Go

# Installation

## Requirements
- Go - get it fom [here](https://go.dev/dl/) or through your package manager

On Linux, you might have to add the go packages' binaries directory to PATH and set GOPATH by adding these lines to `~/.bashrc` or `~/.zshrc`:
```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

You can install the program by running:
```sh
go install github.com/henior2/wpkgtcli@latest
```

# Running
Execute the program with
```sh
wpkgtcli
```

# Config
The config files are stored in `~/.config/wpkg2-cli` for Linux and `%AppData%\.config\wpkg2-cli` for Windows, same as the official WPKG CLI. The config files are fully compatible with the official WPKG CLI.
