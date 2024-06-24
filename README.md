# FilePatrol

A lightweight, real-time file system watcher and static HTTP file server. It triggers custom user commands upon detecting changes in directories or files. It is written in Go and has 0 dependencies.

## Why FilePatrol?

While developing my portfolio's API, I continuously restarted my [portfolio server](https://github.com/sifatulrabbi/sifatul-api) to reflect my code changes. Also, as I build my [Portfolio WebApp](https://github.com/sifatulrabbi/sifatulrabbi.github.io) for testing I need to install an HTTP server or configure my laptop's Nginx to serve the html file created after the build process. So, I've built this CLI tool to automate these two most common tasks.

## Installation

Currently `filepatrol` is only available if you have Go installed in your system APT, AUR, and Brew versions are coming soon.

### Install and set up Go

To install Go on your system follow the [official instructions](https://go.dev/doc/install). After installing Go on your system follow these following instructions to make sure your system is able to use any installed go packages from it's terminal (bash/zsh/fish).

If you are using `bash`

```bash
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
echo 'export PATH=$PATH:$GOBIN' >> ~/.bashrc
source ~/.bashrc
```

If you are using `bash`

```bash
echo 'export GOPATH=$HOME/go' >> ~/.zsh
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zsh
echo 'export PATH=$PATH:$GOBIN' >> ~/.zsh
source ~/.zshrc
```

### Install filepatrol as a CLI tool

```bash
go install github.com/sifatulrabbi/filepatrol@latest
```

## Usage

### 1. As a file watcher

This CLI can watch files of a selected directory and execute one or many valid terminal commands. If the command/s fail it will print out the errors and gracefully exit. This file watcher stores the watching files list in the memory using a Map.

**Basic syntax**

```bash
filepatrol --path [target dir path] --cmd [the command/s it should run]
```

**Example 1:** It is better to wrap the command/s with `''` or `""`

```bash
filepatrol --path ./sifatul-api --cmd 'make run'
```

**Example 2:** If it's a valid terminal command then it will execute

```bash
filepatrol --cmd "jq '.items[-1]' ./logs/errors.json ; jq '. | length' ./logs/errors.json" --path ./logs
```

### 2. As a static file server

**Basic syntax**

```bash
filepatrol --exec http --path [target dir path]

# Example
filepatrol --exec http --path ./portfolio/build
```
