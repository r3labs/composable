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

To generate a docker-compose file (default: docker-compose.yml), run the following command:

```
$ composable gen definition.yml template.yml
```

This will deploy git repos to the deployment directory (default: /tmp/composable). Please note, this directory must exist!


To run override a branch specified in the definition yaml, you can run:

```
$ composable gen -b REPONAME:BRANCH definition.yml template.yml
```

Substitutiong `REPONAME` for the name of the repo and `BRANCH` for the desired branch you want to use.

For further options, you can run:
```
$ composable --help
```


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
