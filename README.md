# go errors

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
