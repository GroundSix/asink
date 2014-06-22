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


More to come...


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
