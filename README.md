# Dockerdiff
Docker image export zip tool. Delete redundant layers.

## Build && install
````
# ENV
linux
go
````
```bash
make && make install
```

## use
```
Usage:
  dockerdiff IMAGE1 IMAGE2 [flags]
  dockerdiff [command]

Examples:
dockerdiff IMAGE1 IMAGE2 -o dst.tar

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Show the Docker version information

Flags:
  -h, --help            help for dockerdiff
  -o, --output string   Write to a file, instead of STDOUT

Use "dockerdiff [command] --help" for more information about a command.
```