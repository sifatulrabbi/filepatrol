# FilePatrol

A lightweight, real-time file system watcher. It triggers custom user commands upon detecting changes in directories or files, respects .gitignore rules, and is easily configurable via CLI

# Todo

- [ ] ignores user specified files or .gitignore files if a .gitignore is present in the dir.
- [x] watches for changes in a dir or in a file.
    - [x] also watches for changes in the dirs within the parent dir.
- [x] runs a user command when any changes are noticed.
- [x] uses a cli to get user input
