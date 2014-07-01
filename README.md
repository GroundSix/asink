![asink](https://raw.githubusercontent.com/GroundSix/asink/master/images/asink.png)

[![Build Status](https://travis-ci.org/GroundSix/asink.svg?branch=master)](https://travis-ci.org/GroundSix/asink)

Asink is an concurrent task runner! It allows you to organise tasks
in a particular order and for certain ones to run in groups concurrently.
These can be ran on your local machine, however Asink is also able to
start up SSH sessions on multiple remote machines at one time to also
run tasks on there.

These tasks can be created using configuration files written in JSON.
Asink will allow you to use a local task file, a remote task file or
you can even start up it's small internal server to POST your
configuration across and Asink will run it all for you.

##### Why would you use it?

It can be used to do an awful lot of things. You could use it
to locally / remotely set up and deploy a project, configure a
remote server, run lots and lots of commands concurrently,
install software and pretty much do all the things you find
yourself repeating a lot.

* * *

### Standalone Usage

#### Install

You will need:

  - Git
  - Go (1.0+)
  - Make

```bash
$ git clone https://github.com/GroundSix/asink.git
$ cd asink
$ make deps
$ make
$ sudo make install
```

Run `asink help` for list of available commands.

#### Basic Usage

##### Commands

There are two main ways of using asink. It can either be
used to to execute one command in multiple sets concurrently,
or used to execute a number of tasks in groups and / or a
particular order.

A configuration file is needed for asink. The simplest way
to configure a single command to be executed lots of times
could look like the following: 

```json
{
    "command" : "ls",
    "args" : [
        "-a"
    ],
    "count" : [2, 5]
}
```

Then you simply pass through that file as the param when you run asink

```bash
$ asink start config.json
```

You may call this file what you wish. What the example above will do,
is run two batches of `ls -a` running five times. So the two batches will
both run concurrently and in each batch it will execute five times.

##### Tasks

Tasks can be ran using a similar kind of configuration. There are various
keys that can currently be used which are as follows:

  - `command`
  - `args`
  - `count`
  - `require`
  - `group`
  - `dir`
  - `remote`


Here is an example of three tasks running, two of which are executed
concurrently and one of wich requires another one to run first:

```json
{
  "tasks" : {
    "do-ls" : {
      "command" : "ls",
      "args"    : [
        "-la"
      ],
      "count"   : [1, 1],
      "require" : "make-text-file"
    },
    "make-text-file" : {
      "command" : "touch",
      "args"    : [
        "file.txt"
      ],
      "count"  : [1, 1],
      "group"  : "create-files"
    },
    "make-json-file" : {
      "command" : "touch",
      "args"    : [
        "file.json"
      ],
      "count"  : [1, 1],
      "group"  : "create-files"
    }
  }
}
```

The `create-files` group tells asink that these commands
are to be ran at the same time. By default asink will
always initially run chronologically, so from the top
down.

In the example above `ls` will not actually run first. This
is because it requires `make-text-file` (the key of a different
task) to run. However, `make-text-file` is in a group
along-side `make-json-file`, so both of these will run first
at the same time, then our `do-ls` task will run afterwards.

NOTE: `count` will always default to `[1, 1]`, so each command
will only run once. It is specified in the example above
however this is not required.

You can use the `dir` key to specify a directory for the command
to be ran in. This directory change is relative to that particular
task and not all that will run after it. e.g.

```json
{
  "tasks" : {
    "do-ls" : {
      "dir"     : "/var",
      "command" : "ls"
    }
  }
}
```

You can either use an absolute path like the example listed above,
or a path that is relative to where you are when you run asink.

See the examples directory for more.

##### Remote Access (SSH)

You can execute commands on a remote machine if you wish
by listing them in the `ssh` key like so:

```json
{
  "ssh" : {
    "vagrant" : {
      "host"     : "127.0.0.1",
      "port"     : "2222",
      "user"     : "vagrant",
      "password" : "vagrant"
    }
  },
  "tasks" : {
    "do-ls" : {
      "remote"  : "vagrant",
      "command" : "ls",
      "args"    : [
        "-la"
      ]
    }
  }
}
```

You may list multiple boxes and then just use the `remote` key
in your task to choose where that command will be executed. In
the example above, `vagrant` has been used as the remote key. This
may be called anything you like as long as you reference it by
the same name when you choose which `remote` you'd like it to run on.

Vagrant makes a good example as you could essentially `vagrant up` using
asink and then remotely execute any extra commands on that box as
soon as it's available. The same rules apply with using the `group` and
`require` keys on remote machines. This means you could execute multiple
commands at once on both your host machine and your remote one if you
so wish to.

##### Ways of starting tasks

There a 3 different ways to start a task or set of commands:

Run with a local configuration file using `start`.

```bash
$ asink start conf.json
```

Run with a remote file using `get`.

```bash
$ asink get http://example.com/conf.json
```

Start a small Http server!

```bash
$ asink server
```

##### Http API

If you use asink as a small Http server, currently, all you have to
do is to send a `POST` request to `127.0.0.1:9000`. The request body
just has to be the raw JSON that would normally just go into your
configuration file.

### Integrating asink

You may integrate asink into your own Go programs like so:

```bash
$ go get github.com/groundsix/asink/asink
```

#### Example

```go
package main

import (
    "github.com/groundsix/asink/asink"
)

func main() {
    command := asink.New()

    command.Name          = "ls"
    command.AsyncCount    = 2
    command.RelativeCount = 2
    command.Args          = []string{"-la"}

    command.Execute()
}
```

See `asink/asink.go` for full API. You may also use `ExecuteCommand`
function which allows you to just pass all the params through as
an alternate method.

Since version 0.0.2, tasks are now part of the public API. A task
consists of a command, on top of various other aspects. Here is an
example:

```go
package main

import (
    "github.com/groundsix/asink/asink"
)

func main() {
    task    := asink.NewTask()
    command := asink.New()

    command.Name          = "ls"
    command.AsyncCount    = 2
    command.RelativeCount = 2
    command.Args          = []string{"-la"}

    task.AddTask("task-name", command, "", "")
    task.Execute()
}
```

This example is similar to the initial example with using commands.
However doing it like this means you can add as many tasks to run
as you wish. The 3rd and 4th arguments of `AddTask` are for the
use of `require` and `group` which are described above in the usage
of asink.

### Running Tests

```bash
$ make test
```

### Contributors

  - [@harry4_](http://twitter.com/harry4_)

### Contributing

Contributions would be great, so do feel free to make a pull request!

1. Fork asink
2. Create a feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to your feature branch (`git push origin my-new-feature`)
5. Create new Pull Request

### License

[MIT](https://github.com/GroundSix/asink/blob/master/LICENSE)

* * *

![Ground Six](https://raw.githubusercontent.com/GroundSix/asink/master/images/groundsix.jpg)
