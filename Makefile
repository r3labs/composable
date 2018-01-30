install:
	go install .

deps:
	go get gopkg.in/yaml.v2
	go get github.com/howeyc/gopass
	go get github.com/docker/docker/client
	go get github.com/r3labs/libcompose
	go get github.com/spf13/cobra
	go get github.com/spf13/viper
	go get github.com/mitchellh/go-homedir

dev-deps: deps
	go get github.com/smartystreets/goconvey
	go get github.com/alecthomas/gometalinter
	gometalinter --install
	
lint:
	gometalinter --config .linter.conf

test:
	go test -v ./... --cover
