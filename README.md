# PlayFab SDK in Go (unofficial)

## Usage
Check [examples](examples) folder. You'll need to `go get github.com/mitchellh/mapstructure`.

## How was this generated?
Check [this](https://github.com/dgkanatsios/SDKGenerator/tree/golang) branch.

## Points

- Why `Model` at the end of every type definition?
Because there was a name clash between a func and a struct that shared the same name.

---
This is not an official Microsoft/PlayFab product.