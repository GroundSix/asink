![asink](https://raw.githubusercontent.com/GroundSix/asink/master/images/asink.png)

Asink is both a [Go](http://golang.org) package that allows you to execute
code and / or commands concurrently and a command line tool that harnesses
all the functionality from the package itself to help you create and automate tasks.

What can it be used for? Loads! You could configure and deploy a project, build / install
software from source, provision a machine, run one task lots and lots of times, check up on the status of a remote machine and anything you find yourself doing manually time and time again.

## Features

* Written in [Go](http://golang.org)
* Can automate [SSH sessions](https://github.com/GroundSix/asink#remote-access-ssh)
* Very easy to get started with
* Comes with 3 main different ways to use
  * Via local configuration file
  * Remote configuration file
  * Small internal server
* Excellent speed and performance
* Public API for Go developers
* Client libraries for other languages

## New in v0.1.1

The main focus for this release is the public Go API for using Asink as part of your own program.
Due to this, the vast bulk has been compleatly re-written with a much cleaner API for you to
take advantage of concurrently running tasks or code within your program. Refer to the API usage
section for more info on this.

In terms of Asink, the program, there has been added support for YAML configuration as appose to
just JSON, the same command can be ran on multiple remote machines at the same time and
installing software has become either locally or remotly has become much easier.

## Getting Started

Building from source requires:

* Go >= v1.2
* Git
* Make

```bash
$ git clone https://github.com/GroundSix/asink.git
$ cd asink
$ export GOPATH=$PWD/vendor
$ make
$ sudo make install
```

Run `asink version` to check the install, or `asink help` to see a full list of commands.

### Configuring

Asink can be configured to run your tasks using either YAML or JSON. Here is a small
Hello World example:

```yaml
---
tasks:

  # All tasks are named
  # You can name them whatever you like
  hello-world:

    # Define the root command
    command: echo

    # And any args and / or flags here
    args:
      - 'Hello, World!'
```

And then once we are all configured, running is as simple as this:

```bash
$ asink start config.yml
```

You can find more information and examples [on the wiki](https://github.com/GroundSix/asink/wiki/Basic-Usage)

### License

GNU GPL v2.0