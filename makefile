SHELL=/bin/bash
all: build install bash-comp
build:
	go build
install:
	sudo cp -pf git-user-switch /usr/local/bin/
uninstall:
	sudo rm -f /usr/local/bin/git-user-switch
	sudo rm -f /etc/bash_completion.d/git-user-switch
bash-comp:
	sudo bash -c "git-user-switch completion bash > /etc/bash_completion.d/git-user-switch"
	source /etc/bash_completion.d/git-user-switch
clean:
	rm -f ./git-user-switch
test:
	go test ./... -count=1 -v #"..."で再帰的に実行する