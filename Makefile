BIN_PATH=bin/asink
INSTALL_PATH=/usr/local/bin/asink

all: asink

dependencies: vendor/
	git submodule init
	git submodule update
	git submodule foreach git pull origin master

asink: dependencies main.go progress.go command.go validator.go
	go build -o ${BIN_PATH} main.go progress.go command.go validator.go

install: ${BIN_PATH}
	@cp ${BIN_PATH} ${INSTALL_PATH}
	@echo "Installed asink!"

uninstall:
	@rm -f ${INSTALL_PATH}
	@echo "Uninstalled asink"

test: asink_test.go
	go test asink_test.go task_test.go

clean:
	@echo "Deleting ${BIN_PATH}."
	rm -f ${BIN_PATH}