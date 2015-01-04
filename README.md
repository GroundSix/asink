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

* Go (tested on 1.2+)
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

Let's take a look at a slightly more involved example:

```yaml
---
tasks:
  clone-asink:
    command: git
    args:
      - clone
      - https://github.com/GroundSix/asink.git
    group: clones

  clone-martini:
    command: git
    args:
      - clone
      - https://github.com/go-martini/martini.git
    group: clones

  do-ls:
    command: ls
    args:
      - -la
    dir: asink
    require: clone-asink
```

Here we are running 3 tasks. Asink will clone itself and
[Martini](https://github.com/go-martini/martini.git) at the same time. This
is because we have defined `group: clones`. Any tasks using the same group
key will be executed at the same time. We then want to run `ls -la` in
the `asink` directory. There is a require key on here which will wait for
`clone-asink` to finish running first. As it's part of a group, `do-ls` will
wait for all tasks in that group to execute first, before executing.

Here is a list of keys that can be used for each task, with a small example of each
case:

| Key      | Description                         | Usage                       | 
|----------|-------------------------------------|-----------------------------|
| command  | This is the root command            | `"command" : "git"`         |
| args     | An array of command arguments       | `"args" : ["status"]`       |
| count    | The asynchronous and relative count | `"count" : [2, 6]`          |
| require  | The required command is ran first   | `"require" : "my-other-cmd"`|
| group    | Groups are ran at the same time     | `"group" : "my-new-group"`  |
| dir      | The directory to be in when running | `"dir" : "/var/www"`        |
| remote   | The remote machine to run on        | `"remote" : "vagrant"`      |

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

```yaml
---
tasks:
  # Clones the Asink repo
  clone-asink:
    command: git
    args:
      - clone
      - https://github.com/groundsix/asink.git
    # We define a new group here called repos
    group: repos

  # Clones the PHP client libary repo
  clone-client:
    command: git
    args:
      - clone
      - https://github.com/asink/asink-php.git
    group: repos
```

Here we are cloning two repos, Asink and a client library. A `group` has been defined.
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
This is done using the `remotes` key outside of the `tasks` scope.

Here is an example of running a command on a vagrant box:

```yaml
---
remotes:
  # We define a remote machine here
  vagrant:
    host: localhost
    port: 2222
    user: vagrant
    key: ~/.ssh/id_rsa

tasks:
  # Clones the Asink repo
  clone-asink:
    # We tell the task which remote to use
    remote: vagrant
    command: git
    args:
      - clone
      - https://github.com/groundsix/asink.git
```

Multiple remotes can be specified under the `ssh` key and then are
accessed in tasks using the `remote` key. You can name this remotes
whatever you like. In the example it's been named `vagrant`.

In your output when this is being ran you'll be able to see which remote
it is being ran on as it will be highlighted blue and show the name of
the remote.

More to come...

### License

GNU GPL v2.0