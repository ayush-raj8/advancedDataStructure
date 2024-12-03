# go-collection
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/anchore/go-collections.svg)](https://github.com/anchore/go-collections)
[![License: Apache-2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/anchore/go-collections/blob/main/LICENSE)

A Go library providing a collection of data structures. The library currently includes a thread-safe, generic `Set` implementation.

## Features

- Thread-safe `Set` implementation using `sync.RWMutex`.
- Supports generics for any comparable type.
- Includes utility operations like Union, Intersection, and Difference.

## Installation

```bash
go get github.com/ayush-raj8/advancedDataStructure
