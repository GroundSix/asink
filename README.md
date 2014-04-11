asink
=====

[![Build Status](https://travis-ci.org/GroundSix/asink.svg?branch=master)](https://travis-ci.org/GroundSix/asink)

asink is a tool that allows you to concurrently
run a command a number of times very quickly. It
can be used independently as a tool or within
your own Go programs.

### Standalone Usage

#### Install

You will need:

  - Git
  - Go (tested with versions 1, 1.1 and 1.2)
  - Make (tested with version 3.81)

```bash
$ git clone https://github.com/GroundSix/asink.git
$ cd asink
$ make
$ sudo make install
```

#### Configuration

asink requires one configuration file written in JSON. An example
looks like this:

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
$ asink config.json
```

You may call this file what you wish. What the example above will do,
is run two batches of `ls -a` running five times. So the two batches will
both run concurrently and in each batch it will execute five times.

Multiplying the two numbers together will give you the total number of
times the command will run.

So if your config file looked like this:

```json
{
    "command" : "php",
    "args" : [
        "index.php",
        "hello"
    ],
    "count" : [10, 10]
}
```

It will run `php index.php hello` ten times (concurrently) in batches of ten, so one hundred
in total.

### Integrating asink

You may integrate asink into your own Go programs like so:

```bash
$ go get github.com/groundsix/asink/asink
```

#### Example

```go
package main

import (
    "fmt"
    "github.com/groundsix/asink/asink"
)

func main() {
    command := "ls"
    args    := []string{"-la"}

    asyncCount := 2
    syncCount  := 2

    if (asink.ExecuteCommand(command, args, asyncCount, syncCount) == true) {
        fmt.Println("Done!")
    }
}
```

### Running Tests

```bash
$ make test
```

### License

[MIT](https://github.com/GroundSix/asink/blob/master/LICENSE)