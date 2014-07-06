![asink](https://raw.githubusercontent.com/GroundSix/asink/master/images/asink.png)

Asink is an concurrent task runner! It allows you to organise tasks
in a particular order and for certain ones to run in groups concurrently.

These tasks can be ran on your local machine. However, Asink is also able to
start up SSH sessions on multiple remote machines at one time to also
run tasks on there.

[![Build Status](https://travis-ci.org/GroundSix/asink.svg?branch=master)](https://travis-ci.org/GroundSix/asink)

## Features

* Written in [Go](http://golang.org)
* Very easy to get started with
* Comes with 3 different ways to use
  * Via local configuration file
  * Remote configuration file
  * Small internal server
* Good speed and performance
* Public API for Go developers
* Can automate SSH sessions

## Getting Started

Building from source requires:

* Go (tested on 1.2+)
* Git
* Make

```bash
$ git clone https://github.com/GroundSix/asink.git
$ cd asink
$ make
$ sudo make install
```

Run `asink help` for list of available commands.

### Configuring

Asink requires JSON to be configured. Here is an example JSON
file that could be used:

`my-tasks.json`
```json
{
  "tasks" : {
      "clone-asink" : {
          "command" : "git",
          "args"    : [
              "clone",
              "https://github.com/groundsix/asink"
          ]
      },
      "build-asink" : {
          "dir"     : "asink",
          "command" : "make"
      }
  }
}
```

In the example above there are two tasks being ran. `clone-asink` and
`build-asink`. To run this we just run the start sub command.

```bash
$ asink start my-tasks.json
```

By default all tasks will be ran chronologically, so from the top down.
After the first task has been ran, it will then move onto the second.


There are only three examples of keys being used in this example, these
are the available keys that can be used:


| Key      | Description                         | Usage                       |
|----------|-------------------------------------|-----------------------------|
| command  | This is the root command            | `"command" : "git"`         |
| args     | An array of command arguments       | `"args" : ["status"]`       |
| count    | The asynchronous and relative count | `"count" : [2, 6]`          |
| require  | The required command is ran first   | `"require" : "my-other-cmd"`|
| group    | Groups are ran at the same time     | `"group" : "my-new-group"`  |
| remote   | The remote machine to run on        | `"remote" : "vagrant"`      |
| dir      | The directory to be in when running | `"dir" : "/var/www"`        |
