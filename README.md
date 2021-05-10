# go errors

![GitHub go.mod Go version of a Go module](https://img.shields.io/github/v/tag/claudiumocanu/golab20errors?color=%2300AA00&include_prereleases&label=Version&logo=Go&style=for-the-badge)

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
