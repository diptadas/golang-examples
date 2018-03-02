# Simple CLI

Simple CLI using [Cobra](https://github.com/spf13/cobra).

```console
$ go build -o my-cli main.go
```

```
$ ./my-cli --help

Simple CLI using Cobra

Usage:
  my-cli [command]

Examples:
my-cli logo

Available Commands:
  echo        Echo anything to the screen
  help        Help about any command
  logo        Print the logo

Flags:
  -h, --help      help for my-cli
  -v, --verbose   verbose output

Use "my-cli [command] --help" for more information about a command.
```

```console
$ ./my-cli echo hello world

Echo: hello world
```

```console
$ ./my-cli echo hello world -t 3

Echo: hello world
Echo: hello world
Echo: hello world
```

```console
$ ./my-cli echo upper hello world

Echo Upper: HELLO WORLD
```

```console
$ ./my-cli echo upper hello world -v

Info: Echo in upper case
Echo Upper: HELLO WORLD
```