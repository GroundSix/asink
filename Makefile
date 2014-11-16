SHELL := /bin/bash
BIN_PATH=build/asink
INSTALL_PATH=/usr/local/bin/asink
ALL_GO_SRC := $(wildcard *.go)
GO_SRC := $(filter-out %_test.go, $(ALL_GO_SRC))

all: deps asink

deps: vendor/
	@if [ -a vendor/.deps ] ; \
	then \
		echo "Skipping dependencies - run 'make clean' first if needed" ; \
	else \
		git submodule init ; \
		git submodule update ; \
		git submodule foreach git pull origin master ; \
		touch vendor/.deps ; \
	fi;

.PHONY: asink
asink: ${GO_SRC}
	@mkdir -p build
	go build -o ${BIN_PATH} $^
	@echo "asink has been built in '${BIN_PATH}'"

install: ${BIN_PATH}
	@cp ${BIN_PATH} ${INSTALL_PATH}
	@echo "Installed asink!"

uninstall:
	@rm -f ${INSTALL_PATH}
	@echo "Uninstalled asink"

test: command_test.go block_test.go task_test.go module_test.go
	go test -v $^

.PHONY: clean
clean:
	rm -f ${BIN_PATH}
	rmdir build
	rm -f vendor/.deps
	@echo "Deleting ${BIN_PATH}."