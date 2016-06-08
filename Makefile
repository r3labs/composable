install:
	go install .

deps:
	go get gopkg.in/yaml.v2
	go get github.com/fsouza/go-dockerclient
