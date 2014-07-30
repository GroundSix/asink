![asink](https://raw.githubusercontent.com/GroundSix/asink/master/images/asink.png)

Asink is an concurrent task runner! It allows you to organise tasks
in a particular order and for certain ones to run in groups concurrently.

These tasks can be ran on your local machine. However, Asink is also able to
start up SSH sessions on multiple remote machines at one time to also
run tasks on there.

What can it be used for? Loads! You could configure and deploy a project, build / install
sortware from source, run one task lots and lots of times, check up on the status of
a remote machine, install dot files and anything you find yourself doing manually time and time again.

[![Build Status](https://travis-ci.org/GroundSix/asink.svg?branch=master)](https://travis-ci.org/GroundSix/asink)

## Features

* Written in [Go](http://golang.org)
* Very easy to get started with
* Comes with 3 different ways to use
  * Via local configuration file
  * Remote configuration file
  * Small internal server
* Good speed and performance
* [Public API](https://github.com/GroundSix/asink/tree/v0.0.2-dev#public-go-api) for Go developers
* Can automate [SSH sessions](https://github.com/GroundSix/asink/tree/v0.0.2-dev#remote-access-ssh)

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

\* NOTE: This has only been tested on OS X and Ubuntu.
Everything works fine, it could work on Windows, however
the `make install` won't. The binary should still build
and can be found in the `bin` directory.
If someone has a chance to test this for me I'd be very
grateful.

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
| dir      | The directory to be in when running | `"dir" : "/var/www"`        |
| remote   | The remote machine to run on        | `"remote" : "vagrant"`      |

See the [examples](https://github.com/GroundSix/asink/tree/master/examples)
for more.

##### Command
This must just be the root command. So in this example it is `git`. It could
be anything, however you can't pass any arguments or flags in this.

##### Args
In args you may pass all arguments and flags. For example if your command
was `ls`, your args could be just `["-la"]`. These are just comma-seperated
values and there is no limit to how many you can use.

##### Count
You'll notice that the `count` key has requires two numbers. This is because
it can run the same command lots of times in sets. In the example above
it has been set to `[2, 6]`. This means that it will run 2 batches of
the command 6 times, concurrently. So 12 times in total. This can be useful
if you have a command you need to run lots of times very quickly.

##### Require
Sometimes you'll have a command that first requires another one to be ran
first. Be default commands are ran chronologically, but if the order becomes
mixed up or you have a fairly complex configuration you can define the key
of another command in here and that will be ran first.

##### Group
Groups allow you to take advantage of Asink's concurrency. Here is a small
example:

```json
{
  "tasks" : {
    "clone-asink" : {
      "command" : "git",
      "args"    : [
        "clone",
        "https://github.com/groundsix/asink"
      ],
      "group" : "repos"
    },
    "clone-mux" : {
      "command" : "git",
      "args"    : [
        "clone",
        "https://github.com/GroundSix/mux"
      ],
      "group" : "repos"
    }
  }
}
```

Here we are cloning two repos, asink and mux. A `group` has been defined.
This means that both of these commands will run concurrently. You can
add groups to as many commands as you like. It plays well with `require`.

##### Dir
This is where you can speicfy the directory for the command to be ran in.
It may be a relative one to where you are running asink from, or an absolute
path.

##### Remote
The remote key allows you to specify a remote machine for the command
to be ran on. See below for how this can be set up.

### Remote Access (SSH)

As well as being able to run commands locally, Asink can start up
SSH sessions and run commands on another machine at the same time.
This is done using a special `ssh` key outside of the `tasks` scope.

Here is an example of running a command on a vagrant box:

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
    "clone-asink" : {
      "remote"  : "vagrant",
      "command" : "git",
      "args"    : [
        "clone",
        "https://github.com/groundsix/asink"
      ]
    }
  }
}
```
Multiple remotes can be specified under the `ssh` key and then are
accessed in tasks using the `remote` key. You can name this remotes
whatever you like. In the example it's been named `vagrant`.

In your output when this is being ran you'll be able to see which remote
it is being ran on as it will be highlighted blue and show the name of
the remote.

#### Authentication

For remote access there are two ways you can connect. Either with a
password of with a key.

```json
{
  "ssh" : {
    "vagrant" : {
      "host" : "127.0.0.1",
      "port" : "2222",
      "user" : "vagrant",
      "key"  : "~/.ssh/id_rsa"
    },
    "my-other-vagrant" : {
      "host"     : "127.0.0.1",
      "port"     : "1234",
      "user"     : "vagrant",
      "password" : "vagrant"
    }
  }
}
```

And yes using `~` will work, however only for the `key` key within
`ssh` and the `dir` key within `tasks`.

### Execution Methods

As stated above there are 3 ways to start up Asink.

* `asink start my-conf.json`
* `asink get http://example.com/my-conf.json`
* `asink server`

The first is just using a local config file, the second is a remote one
and the third is by starting Asink's server.

#### Using the Server

It is started using:

```bash
$ asink server
```

The defalt port is `9000`. All you need to do is send your JSON
configuration as the POST body. Currently there are no specific
routes, so it could be just a request to `http://127.0.0.1:9000`.

\* NOTE: The server is very much experimental and has not had much
time worked on it. It will work, however is not yet recommended to
do so for production.

## Public Go API

For any Go developers, Asink can also be used as a package in your
own programs. This does not provide everything you get in the
program itself. Currently the public API supports the legacy
JSON interface to run commands, and the tasks interface. Here are
some examples to get you started:

```bash
$ go get github.com/groundsix/asink/asink
```

Running JUST the command lots of times (legacy):

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
function which allows you to just pass all the params through as an
alternate method.

Since version 0.0.2, tasks are now part of the public API. A task
consists of a command, on top of various other aspects. Here is an example:

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
However doing it like this means you can add as many tasks to run as you
wish. The 3rd and 4th arguments of `AddTask` are for the use of `require`
and `group` which are described above in the usage of asink.

## Tests

Tests may be ran using make

```bash
$ make test
```

## Contributors

  - [@harry4_](http://twitter.com/harry4_)

## Contributing

Contributions would be great, so do feel free to make a pull request!

1. Fork Asink
2. Create a feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to your feature branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License

[MIT](https://github.com/GroundSix/asink/blob/master/LICENSE)

* * *

![Ground Six](https://raw.githubusercontent.com/GroundSix/asink/master/images/groundsix.jpg)
