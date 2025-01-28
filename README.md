# Todo


## Root command

```shell
This CLI tool is a simple todo list application.

Usage:
  todo [flags]
  todo [command]

Available Commands:
  add         Add a new task to todo list.
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  show        Show all tasks in the todo list.
  update      Update a task status to 'Done' in the todo list.

Flags:
  -h, --help      help for todo
  -v, --version   Show version

Use "todo [command] --help" for more information about a command.
```

## Command

### add

```shell
Usage:
  todo add [flags]

Flags:
  -h, --help          help for add
  -n, --name string   Task name
```

### show

```shell
Usage:
  todo show [flags]

Flags:
  -h, --help   help for show
```

### update

```shell
Usage:
  todo update [flags]

Flags:
  -h, --help          help for update
  -n, --name string   Name of the task to update
```