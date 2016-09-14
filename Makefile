install:
	go install .

deps:
	go get gopkg.in/yaml.v2
	go get github.com/howeyc/gopass
	go get github.com/fsouza/go-dockerclient

test:
	go test -v ./... --cover
