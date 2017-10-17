# composable

## Synopsis

A tool for generating a docker compose file based on git repo branches

## Installing

```
$ make deps
$ make install
```

## Tests

Running the tests:
```
make test
```

## Usage

### Generating a docker-compose file for testing

To generate a docker-compose file (default: docker-compose.yml), run the following command:

```
# community edition
$ composable generate definition.yml template.yml

# enterprise edition
$ composable generate --edition enterprise definition.yml template.yml
```

This will deploy git repos to the deployment directory (default: /tmp/composable). Please note, this directory must exist!


To override a specific branch in the definition yaml, you can run:

```
$ composable generate -b REPONAME:BRANCH definition.yml template.yml
```

Substitutiong `REPONAME` for the name of the repo and `BRANCH` for the desired branch you want to use.

For further options, you can run:
```
$ composable --help
```

### Running a release

To create a release and publish images to a container registry, you can run:

```
$ composable release -v 1.x.x -L quay.io -O r3labs -U john definition.yml template.yml
```

This will:
- Clone all repositories
- Build all docker images, tagged as the release version
- Upload all docker images to docker hub
- Create a release docker-compose file with images tagged to the release version

Please note, you will be asked for your docker credentials upon running this command

## Contributing

Please read through our
[contributing guidelines](CONTRIBUTING.md).
Included are directions for opening issues, coding standards, and notes on
development.

Moreover, if your pull request contains patches or features, you must include
relevant unit tests.

## Versioning

For transparency into our release cycle and in striving to maintain backward
compatibility, this project is maintained under [the Semantic Versioning guidelines](http://semver.org/).

## Copyright and License

Code and documentation copyright since 2015 r3labs.io authors.

Code released under
[the Mozilla Public License Version 2.0](LICENSE).
