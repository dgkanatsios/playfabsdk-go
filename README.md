[![unofficial Google Analytics for GitHub](https://gaforgithub.azurewebsites.net/api?repo=PlayFabSDKGo)](https://github.com/dgkanatsios/gaforgithub)
# PlayFab SDK in Go (unofficial)

## Usage

Check [e2e](e2e) or [examples](examples) folders.

## How is this SDK generated?

Check [this](https://github.com/dgkanatsios/SDKGenerator/tree/golang) branch which is a fork of [the PlayFab SDKGenerator repo](https://github.com/PlayFab/SDKGenerator) with extra files to facilate the generation of a Go client.

## How can I contribute?

If you want to propose a big change, I'd appreciate if you opened an [issue](https://github.com/dgkanatsios/playfabsdk-go/issues) on this repo so we can discuss about it. In any case, contributions should be made to the SDK Generator `golang` branch [here](https://github.com/dgkanatsios/SDKGenerator/tree/golang).

## FAQ

- Why `Model` at the end of every type definition?
Because there was a name clash between a func and a struct that shared the same name.

---
This is **not** an official Microsoft/PlayFab product.