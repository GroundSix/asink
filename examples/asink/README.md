### Installing asink, using asink

Steps:

  - Clones asink from GitHub
  - Runs `make` in the asink directory
  - Runs `make install` in the `asink` directory

Requires:

  - Git
  - Go (1.0+)
  - Make

So far this has only been tested on OS X.

For this particualr example it must be ran as root, so:

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
				"https://github.com/GroundSix/asink.git"
			],
			"output" : true
		},
		"build-asink" : {
			"command" : "make",
			"args"    : [
				"-C",
				"asink"
			],
			"output" : true
		},
		"install-asink" : {
			"command" : "make",
			"args"    : [
				"-C",
				"asink",
				"install"
			],
			"output" : true
		}
	}
}
```