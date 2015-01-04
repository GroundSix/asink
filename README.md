![asink](https://raw.githubusercontent.com/GroundSix/asink/master/images/asink.png)

Asink is command line tool that harnesses the functionality from 
[libasink](https://github.com/asink/libasink) to help you create
and automate tasks.

What can it be used for? Loads! You could configure and deploy a project, build / install
software from source, provision a machine, run one task lots and lots of times, check up on the status of a remote machine and anything you find yourself doing manually time and time again.

## Features

* Written in [Go](http://golang.org)
* Can automate [SSH sessions](https://github.com/GroundSix/asink#remote-access-ssh)
* Very simple to get started with
* Excellent speed and performance
* Client libraries and support for other languages

## New in v0.1.1

The main focus for v0.1.1 was to have a complete rewrite of the public Go
API that Asink has internally depended on since the project started. This
has since been moved over to [asink/libasink](https://github.com/asink/libasink)
and development will continue on both the independant package and this
command line tool.

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

Alternativly Asink can be built using [Docker](https://www.docker.com/):

```bash
$ git clone https://github.com/GroundSix/asink.git
$ cd asink
$ docker build -t asink .
$ ./build.sh
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
$ asink run config.yml
```

You can find more information and examples [on the wiki](https://github.com/GroundSix/asink/wiki/Basic-Usage)

### License

GNU GPL v2.0
