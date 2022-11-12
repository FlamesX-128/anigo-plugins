<div align="center">
    <h1>Anigo models</h1>
    <p>
        Anigo models to make your plugins.
    </p>
</div>

---

<div align="center">
    <img src="https://github.com/FlamesX-128/anigo/blob/main/assets/image.jpg" height="200" />
</div>

---

<div align="center">
    <p>
        <a href="https://github.com/FlamesX-128/anigo-plugins/graphs/contributors">
            <img src="https://img.shields.io/github/contributors/FlamesX-128/anigo-plugins" alt="contributors" />
        </a>
        <a href="">
            <img src="https://img.shields.io/github/last-commit/FlamesX-128/anigo-plugins" alt="last update" />
        </a>
        <a href="https://github.com/FlamesX-128/anigo-plugins/network/members">
            <img src="https://img.shields.io/github/forks/FlamesX-128/anigo-plugins" alt="forks" />
        </a>
        <a href="https://github.com/FlamesX-128/anigo-plugins/stargazers">
            <img src="https://img.shields.io/github/stars/FlamesX-128/anigo-plugins" alt="stars" />
        </a>
        <a href="https://github.com/FlamesX-128/anigo-plugins/issues/">
            <img src="https://img.shields.io/github/issues/FlamesX-128/anigo-plugins" alt="open issues" />
        </a>
        <a href="https://github.com/FlamesX-128/anigo-plugins/blob/master/LICENSE">
            <img src="https://img.shields.io/github/license/FlamesX-128/anigo-plugins.svg" alt="license" />
        </a>
    </p>
</div>

<br />

## Usage

### Dependencies

- Go 1.19

### How to

```sh
# Create a gomodule
go mod init <name>
```

```sh
# Install dependencies.
go get github.com/FlamesX-128/anigo-plugins/models
```

```go
// Create de plugin.
// Currently only plugins from providers like gogoanime.
package main

import "github.com/FlamesX-128/anigo-plugins/models"

type PackageModel struct{}

// Here you must return the list of animes that were found based on certain keywords.
type (p PackageModel) Search(query string) []models.Anime {}

// here you must return the information of a specific anime.
type (p PackageModel) Info(id string) models.Info {}

// here you must return the urls to be able to see it.
type (p PackageModel) Watch(query string) []models.Source {}

// Create the plugin symbol.
var Plugin = models.Plugin{
	Providers: map[string]models.Provider{
		"myplugin": PackageModel{},
	},
}
```

```sh
# Build you plugin.
go build -buildmode=plugin -o "myplugin@1.0.0.so" main.go
```

## Plugin directory

Anigo will look for the `anigo-plugins` folder in the current directory from where it is called.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Resources

- [Consumet](https://github.com/consumet/api.consumet.org)
- [Shields.io](https://shields.io/)

## License

[MIT](https://opensource.org/licenses/MIT)
