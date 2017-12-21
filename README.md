![](logo.png)

# afto [alpha]

_/αυτο/_


[![Build Status](https://travis-ci.org/hako/afto.svg)](https://travis-ci.org/hako/afto)
![](http://goreportcard.com/badge/hako/afto)
![](http://img.shields.io/status/alpha.png?color=yellow)
![](https://img.shields.io/badge/version-0.2-yellow.png)

afto is an automated command-line cydia repo generator/builder and server for Cydia tweak developers.

_The name 'afto' comes from the word 'automatic' in greek._

### features
+ Automatic Cydia repo generation.
+ Automatic Cydia repo updating.
+ Cydia repo server testing.
+ Many more.

### install

If you have Go installed you can install `afto` by typing:
`go get github.com/hako/afto/cmd/afto`

This will install the project source code and the `afto` binary.

Release binaries will be provided in the near future.

### requirements:
The only thing you need on your system is:

`dpkg`

You also need at least 1 or more `.deb` files. So that you can test or host your repo.

### usage
```
Usage:
  afto new <name>
  afto serve <dir> [-w | --watch] [-p <port> | --port <port>]
  afto update -r <name> [-f <file> | --file <file>] 
  afto [-c <file> | --control <file>]
  afto [-s <dir> | --sign <dir>]

options:
  -c, --control  Specify control file to use.
  -p, --port     Specify port number for afto.
  -h, --help     Show this screen.
  --version      Show version.

commands:
  new             Generate a new Cydia repo.
  serve           Serve the Cydia repo.
```

### example

Below is a basic example of how to use afto:

```
afto new example_repo # Generate a new Cydia repo.
afto serve example_repo # Serve the Cydia repo.
or
afto serve -w example_repo # Serve the Cydia repo and watch for changes.
```

###### tip:
You can serve your repo within the same directory without giving it a name:

```
afto new . # Generate a new Cydia repo.
afto serve . # Serve the Cydia repo.
or
afto serve -w . # Serve the Cydia repo and watch for changes.
```

### roadmap
see [AFTODO.md](AFTODO.md)

### contributing:
Contributions are welcome. Fork this repo and add your changes and submit a PR. 

If you would like to fix a bug, add a feature or provide feedback you can do so in the issues section.

You can run afto tests by using `make test`
and make afto builds by doing `make build`

Make sure you run `make test ; make scrutinise` so that your changes do not cause [go lint](https://github.com/golang/lint) and [go vet](https://golang.org/cmd/vet/) to scream errors at you.

### special thanks:
[@saurik](http://twitter.com/saurik)

[@return](https://github.com/return)

[you???](https://github.com/hako/afto#contributing)

[Haiku Project](https://haiku-os.org)

(_The original package icon used at the top_)

### license

MIT
