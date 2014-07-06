### Installing asink, using asink

Asink doesn't have any specific code in place to
update itself or a particualr command that will
do this but due to the nature of what it does,
it can actually update itself using a
configuration file. 

Steps:

  - Clones asink from GitHub
  - Runs `make` in the asink directory
  - Runs `make install` in the `asink` directory
  - Cleans up after itself

Requires:

  - Git
  - Go (1.2+)
  - Make

So far this has only been tested on OS X.

For this particualr example it must be ran as root. Since the
configuration file is actually already on GitHub you could use
the `get` sub command in asink like so:

```bash
$ asink get https://raw.githubusercontent.com/GroundSix/asink/master/examples/asink/conf.json
```
Or if you grab a copy of the configuration you could run it
using `start` from your machine like so:

```bash
$ sudo asink start conf.json
```

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
		"checkout-stable" : {
			"dir"     : "asink",
			"command" : "git",
			"args"    : [
				"checkout",
				"current-stable"
			],
			"require" : "clone-asink"
		},
		"build-asink" : {
			"command" : "make",
			"args"    : [
				"-C",
				"asink"
			],
			"require" : "checkout-stable"
		},
		"install-asink" : {
			"command" : "make",
			"args"    : [
				"-C",
				"asink",
				"install"
			],
		"require" : "build-asink"
		},
		"clean" : {
			"command" : "rm",
			"args"    : [
				"-rf",
				"asink"
			],
		"require" : "install-asink"
		}
	}
}
```

* * *

![Ground Six](https://raw.githubusercontent.com/GroundSix/asink/master/images/groundsix.jpg)