# The Rick and Morty API Go client

This is a Go wrapper to use the [The Rick and Morty API](https://rickandmortyapi.com)

**To get started check the documentation on [rickandmortyapi.com](https://rickandmortyapi.com/documentation)**

## Installation

`go get -u github.com/pitakill/rickandmortyapigowrapper`

## Basic Usage
```go
// Import and rename the package for a easy handle
import rnm "github.com/pitakill/rickandmortyapigowrapper"

// To get the character with id 1
character, err := rnm.GetCharacter(1)
```

## Full Usage

The complete usage is in the [example directory](https://github.com/pitakill/rickandmortyapigowrapper/tree/master/examples)
