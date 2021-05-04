all: build install
build:
	go build
install:
	cp -pf git-user-switch ~/.local/bin/
uninstall:
	rm -f ~/.local/bin/git-user-switch
clean:
	rm -f ./git-user-switch
test:
	go test ./... -count=1 -v #"..."で再帰的に実行する