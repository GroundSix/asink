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

For remote access there are two ways you can connets. Either with a
password of with a key.

```json
{
  "ssh" : {
    "vagrant" : {
      "host" : "127.0.0.1",
      "port" : "2222",
      "user" : "vagrant",
      "key"  : "/path/to/key"
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


