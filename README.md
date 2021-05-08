# go errors

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)
[![GitHub release](https://img.shields.io/github/release/Naereen/StrapDown.js.svg)](https://GitHub.com/Naereen/StrapDown.js/releases/)
[![GitHub tag](https://img.shields.io/github/tag/Naereen/StrapDown.js.svg)](https://GitHub.com/Naereen/StrapDown.js/tags/)

> Just playing atm...trying to handle and trace stuff

## How to run

Export the environment variable expected by the goapp

```sh
. exportEnvs.sh
```

Run the app

```sh
go run main.go
```

## CI

Any push to the `release` branch generates an auto-bump for the `#patch` version by default.  
Include `#major`, `#minor` or `#patch` in the commit message to change the default behavior.
