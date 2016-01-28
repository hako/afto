![](logo.png)

# afto [alpha]

_/αυτο/_


[![Build Status](https://travis-ci.org/hako/afto.svg)](https://travis-ci.org/hako/afto)
![](http://goreportcard.com/badge/hako/afto)
![](http://img.shields.io/status/alpha.png?color=yellow)
![](https://img.shields.io/badge/version-0.2-yellow.png)

afto is an automated command-line cydia repo generator/builder and hoster for Cydia tweak developers.

### features
+ automatic Cydia repo generation.
+ automatic Cydia repo updating.
+ Cydia repo server testing.
+ Many more.

### install

If you have Go installed you can install the project by using:
`go get github.com/hako/afto`

And you can install `afto` by using:
`go get github.com/hako/afto/cmd/afto`

Single binaries will be provided in the near future.

### requirements:
The only thing you need on your system is:

`dpkg`

### roadmap
see [AFTODO.md](AFTODO.md)

### contributing:
Contributions are welcome. Fork this repo and add your changes and submit a PR. 

If you would like to fix a bug, add a feature or provide feedback you can do so in the issues section.

You can run afto tests by using `make test`
and make afto builds by doing `make build`

Make sure you run `make scrutinise` so that your changes do not cause [go lint](https://github.com/golang/lint) and [go vet](https://golang.org/cmd/vet/) to scream errors at you.

### special thanks:
[@saurik](http://twitter.com/saurik)

[you???](https://github.com/hako/afto#contributing)

### license

MIT

_The name 'afto' comes from the word 'automatic' in greek._
