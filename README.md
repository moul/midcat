# midcat

:smile: midcat

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/moul.io/midcat)
[![License](https://img.shields.io/badge/license-Apache--2.0%20%2F%20MIT-%2397ca00.svg)](https://github.com/moul/midcat/blob/main/COPYRIGHT)
[![GitHub release](https://img.shields.io/github/release/moul/midcat.svg)](https://github.com/moul/midcat/releases)
[![Docker Metrics](https://images.microbadger.com/badges/image/moul/midcat.svg)](https://microbadger.com/images/moul/midcat)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)

[![Go](https://github.com/moul/midcat/workflows/Go/badge.svg)](https://github.com/moul/midcat/actions?query=workflow%3AGo)
[![Release](https://github.com/moul/midcat/workflows/Release/badge.svg)](https://github.com/moul/midcat/actions?query=workflow%3ARelease)
[![PR](https://github.com/moul/midcat/workflows/PR/badge.svg)](https://github.com/moul/midcat/actions?query=workflow%3APR)
[![GolangCI](https://golangci.com/badges/github.com/moul/midcat.svg)](https://golangci.com/r/github.com/moul/midcat)
[![codecov](https://codecov.io/gh/moul/midcat/branch/main/graph/badge.svg)](https://codecov.io/gh/moul/midcat)
[![Go Report Card](https://goreportcard.com/badge/moul.io/midcat)](https://goreportcard.com/report/moul.io/midcat)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/midcat/badge)](https://www.codefactor.io/repository/github/moul/midcat)

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/moul/midcat)

## Usage

[embedmd]:# (.tmp/usage.txt console)
```console
foo@bar:~$ midcat -h
_     _            _
 _ __  (_) __| | __  __ _ | |_
| '  \ | |/ _` |/ _|/ _` ||  _|
|_|_|_||_|\__,_|\__|\__,_| \__|
4 CPUs, /Users/moul/.local/bin/midcat, manfred-spacegray, go1.16.5


USAGE
  midcat [FLAGS] <ADDRESS>[,OPTS] <ADDRESS>[,OPTS]

FLAGS
  -debug false  debug mode

ADDRESS
  midi        midi port id=FIRST
  -           stdio
  pipe        echo/fifo
  tcp         ...
  tick        ...
  rand        ...
  udp         ...
  websocket   ...

OPTS
  debug       ...
  reconnect   ...
  bpm         ...
  quantify    ...
  filter      ...

HARDWARE
  IN:  id=0 is-open=false name="IAC Driver Bus 1"
  OUT: id=1 is-open=false name="IAC Driver Bus 1"
```

## Install

### Using go

```sh
go get moul.io/midcat
```

### Releases

See https://github.com/moul/midcat/releases

## Contribute

![Contribute <3](https://raw.githubusercontent.com/moul/moul/main/contribute.gif)

I really welcome contributions.
Your input is the most precious material.
I'm well aware of that and I thank you in advance.
Everyone is encouraged to look at what they can do on their own scale;
no effort is too small.

Everything on contribution is sum up here: [CONTRIBUTING.md](./.github/CONTRIBUTING.md)

### Dev helpers

Pre-commit script for install: https://pre-commit.com

### Contributors ✨

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg)](#contributors)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="http://manfred.life"><img src="https://avatars1.githubusercontent.com/u/94029?v=4" width="100px;" alt=""/><br /><sub><b>Manfred Touron</b></sub></a><br /><a href="#maintenance-moul" title="Maintenance">🚧</a> <a href="https://github.com/moul/midcat/commits?author=moul" title="Documentation">📖</a> <a href="https://github.com/moul/midcat/commits?author=moul" title="Tests">⚠️</a> <a href="https://github.com/moul/midcat/commits?author=moul" title="Code">💻</a></td>
    <td align="center"><a href="https://manfred.life/moul-bot"><img src="https://avatars1.githubusercontent.com/u/41326314?v=4" width="100px;" alt=""/><br /><sub><b>moul-bot</b></sub></a><br /><a href="#maintenance-moul-bot" title="Maintenance">🚧</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors)
specification. Contributions of any kind welcome!

### Stargazers over time

[![Stargazers over time](https://starchart.cc/moul/midcat.svg)](https://starchart.cc/moul/midcat)

## License

© 2021   [Manfred Touron](https://manfred.life)

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)
([`LICENSE-APACHE`](LICENSE-APACHE)) or the [MIT license](https://opensource.org/licenses/MIT)
([`LICENSE-MIT`](LICENSE-MIT)), at your option.
See the [`COPYRIGHT`](COPYRIGHT) file for more details.

`SPDX-License-Identifier: (Apache-2.0 OR MIT)`
