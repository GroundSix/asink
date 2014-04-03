BIN_PATH=bin/asink
INSTALL_PATH=/usr/local/bin/asink

all: asink

dependencies: vendor/
	git submodule init
	git submodule update

asink: dependencies main.go
	go build -o ${BIN_PATH} main.go

install: ${BIN_PATH}
	@cp ${BIN_PATH} ${INSTALL_PATH}
	@echo "Installed asink!"

uninstall:
	@rm -f ${INSTALL_PATH}
	@echo "Uninstalled asink"

clean: ${BIN_PATH}
	rm ${BIN_PATH}