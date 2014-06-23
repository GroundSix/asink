![asink](https://raw.githubusercontent.com/GroundSix/asink/master/images/asink.png)

[![Build Status](https://travis-ci.org/GroundSix/asink.svg?branch=master)](https://travis-ci.org/GroundSix/asink)
[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/GroundSix/asink/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

asink lets you run tasks concurrently...

![example](https://raw.githubusercontent.com/GroundSix/asink/master/images/screenshots/example2.gif)

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
$ make
$ sudo make install
```

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
    "count" : [2, 5],
    "output" : true
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
  - `output`
  - `require`
  - `group`


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
      "output"  : true,
      "require" : "make-text-file"
    },
    "make-text-file" : {
      "command" : "touch",
      "args"    : [
        "file.txt"
      ],
      "count"  : [1, 1],
      "output" : true,
      "group"  : "create-files"
    },
    "make-json-file" : {
      "command" : "touch",
      "args"    : [
        "file.json"
      ],
      "count"  : [1, 1],
      "output" : true,
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

NOTE: `count` will always default to [1, 1], so each command
will only run once. It is specified in the example above
however this is not required.

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
  }
  "tasks" : {
    "do-ls" : {
      "remote"  : "vagrant",
      "command" : "ls",
      "args"    : [
        "-la"
      ],
      "output" : true
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
    command.Output        = true

    command.Execute()
}
```

See `asink/asink.go` for full API. You may also use `ExecuteCommand`
function which allows you to just pass all the params through as
an alternate method.

### Running Tests

```bash
$ make test
```

### License

[MIT](https://github.com/GroundSix/asink/blob/master/LICENSE)

* * *

![Ground Six](https://raw.githubusercontent.com/GroundSix/asink/master/images/groundsix.jpg)
