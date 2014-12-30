SHELL := /bin/bash
BIN_PATH=build/asink
INSTALL_PATH=/usr/local/bin/asink

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
asink:
	@mkdir -p build
	go build -v -o ${BIN_PATH}
	@echo "asink has been built in '${BIN_PATH}'"

cross:
	@mkdir -p build
	go build -v -o ${BIN_PATH}-${GOOS}-${GOARCH}${EXT}
	@echo "asink has been built in '${BIN_PATH}-${GOOS}-${GOARCH}${EXT}'"

install: ${BIN_PATH}
	@cp ${BIN_PATH} ${INSTALL_PATH}
	@echo "Installed asink!"

uninstall:
	@rm -f ${INSTALL_PATH}
	@echo "Uninstalled asink"

.PHONY: clean
clean:
	rm -rf build
	rm -f vendor/.deps
	@echo "Deleting ${BIN_PATH}."
