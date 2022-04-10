<p align="center">
  <img src="https://user-images.githubusercontent.com/20264712/162619417-d310506e-9ed1-4453-98fa-5bf28b83abfe.png"/>
  <h3 align="center">F1 CLI</h3>
  <p align="center">A helper CLI for broadcasting and recording F1 game data</p>
  <p align="center">
  <a href="https://opensource.org/licenses/Apache-2.0"><img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg" alt="Apache 2.0"></a>
  <a href="https://goreportcard.com/report/github.com/racing-telemetry/f1-cli"><img src="https://goreportcard.com/badge/github.com/racing-telemetry/f1-cli" alt="Go Report"></a>
  <a href="https://github.com/racing-telemetry/f1-cli/actions?workflow=test"><img src="https://img.shields.io/github/workflow/status/racing-telemetry/f1-cli/release" alt="Build Status"></a>
  <a href="https://github.com/racing-telemetry/f1-cli/releases/latest"><img src="https://img.shields.io/github/release/racing-telemetry/f1-cli.svg" alt="GitHub release"></a>
  <a href="https://github.com/racing-telemetry/f1-cli/"><img src="https://img.shields.io/github/go-mod/go-version/racing-telemetry/f1-cli" alt="Go Mod"></a>
</p>

## Table of Contents

- [F1 CLI](#f1-cli)
    - [Introduction](#introduction)
    - [Installation](#installation)
      - [Go](#go)
      - [Homebrew](#homebrew)
    - [Quick Start](#quick-start)
      - [Recording](#recording)
      - [Broadcasting](#broadcasting)

## Introduction

We made this tool thinking that we could take the F1 game data from UDP and simulate that race or laps again using this
data.

## Installation

### Go

If you have Go 1.17+, you can directly install by running:

```shell
$ go install github.com/racing-telemetry/f1-dump@latest
```

and the resulting binary will be placed at **_$HOME/go/bin/f1_**.

### Homebrew

If you have brew installed, then you can easily download this with the following commands:

```shell
brew tap racing-telemetry/homebrew-tap
brew install f1-cli
```

# Quick Start
```shell
$ f1 --help
A helper CLI for broadcasting and recording F1 game data

Usage:
  f1 [command]

Available Commands:
  broadcast   Start broadcasting
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List F1 packet types
  record      Start recording packets from UDP source
  version     Prints the CLI version

Flags:
  -h, --help      help for f1dump
  -v, --verbose   verbose output

Use "f1 [command] --help" for more information about a command.
```

## Recording

You can record the packets from the F1 game with UDP socket.

```shell
# Default options
$ f1 record

# Specify the output file name
$ f1 record -f f1-cli-out.bin

# Different IP and Port
$ f1 record --ip 192.168.1.41 -p 20777

# Ignore Packets
$ f1 record --ignore 1,2,3,4,5,6
```

## Broadcasting

You can broadcast the data you have recorded with using UDP. 

```shell
# Recorded binary file
$ f1 broadcast -f f1-cli-out.bin

# Ignore Packets
$ f1 broadcast -f f1-cli-out.bin --ignore 1,2,3,4,5,6

# Broadcast all files instantly
$ f1 broadcast -f f1-cli-out.bin -i

# Different IP and Port
$ f1 broadcast -f f1-cli-out.bin --ip 192.168.1.41 -p 20777
```
