install:
	go install .

deps:
	go get gopkg.in/yaml.v2
	go get github.com/howeyc/gopass
	go get github.com/fsouza/go-dockerclient
	go get github.com/spf13/cobra
	go get github.com/docker/libcompose

test:
	go test -v ./... --cover
