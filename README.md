# FilePatrol

A two in one CLI tool written in Go with 0 extra dependencies. Run a terminal executable command or serve your static files over HTTP.

## Why FilePatrol?

While I was delving deep into Go echo system and enjoying the simplicity of Go programming language. I found myself continuously restarting my [portfolio server](https://github.com/sifatulrabbi/sifatul-api) to reflect the code changes I made. But don't like using nodemon since I don't have a super computer to backup it's CPU and RAM consumption. Also, as I build my [Portfolio WebApp](https://github.com/sifatulrabbi/sifatulrabbi.github.io) for testing I need to install a HTTP server or configure my laptop's Nginx to serve the html file created after the build process. So, I've built this CLI tool to automate these two most common tasks.

## Installation

Preparing...

## Usage

### 1. As a file watcher

This CLI can watch files of a selected directory and execute one or many valid terminal command. If the command/s fail it will print out the errors and gracefully exit. This file watcher stores the watching files list in the memory using a Map.

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

# Todo

- [ ] ignores user specified files or .gitignore files if a .gitignore is present in the dir.
- [x] upgrade the cli to support static server feature.
- [x] watches for changes in a dir or in a file.
  - [x] also watches for changes in the dirs within the parent dir.
- [x] runs a user command when any changes are noticed.
- [x] uses a cli to get user input
