
# coopgo/gtfs

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/coopgo/gtfs.svg)](https://github.com/coopgo/gtfs)
[![GoReportCard example](https://goreportcard.com/badge/github.com/coopgo/gtfs)](https://goreportcard.com/report/github.com/coopgo/gtfs)

Package `coopgo/gtfs` implements a GTFS parser to load and exploit GTFS feeds.

The main features of this package are: 

- Two different datastructures to work with GTFS data.
- A "serializable" datastructure which is a one to one mapping of the [GTFS reference](https://developers.google.com/transit/gtfs/reference). This version should be used when persisting or serializing GTFS data.
- A "usable" version which is using pointers to link the different GTFS objects. This version should be used to work with GTFS data in algorithm.
- Fast loading of GTFS file. Be it a zip file or an uncompressed folder.
- Total validation of a feed by following the [GTFS reference specification](https://developers.google.com/transit/gtfs/reference).

## Install
You'll need a [correctly configured Go toolchain](https://golang.org/doc/install). You can then use:

```
go get github.com/coopgo/gtfs
```

This will install the `coopgo/gtfs` to your `$GOPATH/bin` directory.

You can then use this library by importing `github.com/coopgo/gtfs` into your application.

## Usage

### Parsing a GTFS file

Let's start by parsing the sample [feed provided](https://developers.google.com/transit/gtfs/examples/gtfs-feed) by Google.

```go
package main

import (
	"github.com/coopgo/gtfs"
	"fmt"
)

func main() {
	parser := gtfs.NewParser()
	feed, err := parser.Load("sample-feed.zip")
	if err != nil {
		panic(err)	
	}

	fmt.Println(feed)
}
```

Here we use the `Load` method of our parser that gives as the "usable" feed. 

Our parser, will start by loading the given file, `unmarshal` it (this give the serializable datastructure) then `Link` the `FeedSerializable` to transforms it into a `Feed`. Our parser then do a total validation to check if the feed is in valid in the eyes of the [GTFS reference](https://developers.google.com/transit/gtfs/reference).

If you want the "serializable" version of our feed, you can do it like this:

```go
package main

import (
	"github.com/coopgo/gtfs"
	"fmt"
)

func main() {
	parser := gtfs.NewParser()
	feed, err := parser.Unmarshal("sample-feed.zip")
	if err != nil {
		panic(err)	
	}

	fmt.Println(feed)
}
```

A  `FeedSerializable` can also be transformed into a `Feed` like this:
```go
func main() {
	parser := gtfs.NewParser()
	feedSerializable, err := parser.Unmarshal("sample-feed.zip")
	if err != nil {
		panic(err)
	}
	feed, err := feedSerializable.Link()
	if err != nil {
		panic(err)
	}
}
```

### Configuring the Parser

The `Parser` should **always** be created with `gtfs.NewParser()`. This allow the library to set the default value of the different field of our `Parser`. You can give `func(* gtfs.Parser)` to the `gtfs.NewParser()` function to configure it. This is part of the [self referential functions design](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html).
```go
func main() {
	noValidation := func(p *gtfs.Parser){
		p.Validation = false
	}
	parser := gtfs.NewParser(noValidation)
	...
}
```
This function `noValidation` is already implemented in the package so you don't need to re implement it.
```go
func main() {
	parser := gtfs.NewParser(gtfs.NoValidation)
	...
}
```
## Project Status

This library is still in development. 
Encoding and Testing should be coming soon.

## Contributing


We welcome any contributions following theses guidelines :
- Write simple, clear and maintainable code and avoid technical debt. 
- Leave the code cleaner than when you started.
- Refactoring existing code for better performance, better readability or better testing wins over creating a new feature.

If you want to contribute, you can fork the repository and create a pull request.

## Bug report

For reporting a bug, you can open an issue using the **Bug Report** template. Try to write a bug report that is easy to understand and explain how to reproduce the bug. 
Do not duplicate an existing issue and keep each issue specific to an individual bug.

## License

`coopgo/gtfs` is under the Apache 2.0 license. Please refer to the [LICENSE](LICENSE) file for details.
