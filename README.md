[![unofficial Google Analytics for GitHub](https://gaforgithub.azurewebsites.net/api?repo=PlayFabSDKGo)](https://github.com/dgkanatsios/gaforgithub)
# PlayFab SDK in Go (unofficial)

## Usage

Check [examples](examples) or [e2e](e2e) folder. You'll need to `go get github.com/mitchellh/mapstructure`.

## How is this SDK generated?

Check [this](https://github.com/dgkanatsios/SDKGenerator/tree/golang) branch which is a fork of [the PlayFab SDKGenerator repo](https://github.com/PlayFab/SDKGenerator).

## How can I contribute?

If you want to propose a big change, I'd appreciate if you opened an issue on this repo so we can discuss about it. In any case, contributions should be made to the SDK Generator golang branch [here](https://github.com/dgkanatsios/SDKGenerator/tree/golang).

## FAQ

- Why `Model` at the end of every type definition?
Because there was a name clash between a func and a struct that shared the same name.

---
This is not an official Microsoft/PlayFab product.