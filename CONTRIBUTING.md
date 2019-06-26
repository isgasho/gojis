# How to contribute

1. Create an issue where you tell me that you want to contribute (and what, Runtime, Parser or API, or combinations of them)
2. I will add you to the repository and the respective team(s)

Alternatively, you can

1. Fork the repository
2. Make changes
3. Open a Pull Request against `develop`

# Code requirements

* No test cases can be removed and/or changed (if a test fails, but ran green in the past, this is a regression)
  * If you think a test is incorrect, open an issue, stating why the test should be changed/removed
* To be consistent, new objects are created with constructor methods (`foo := NewFoo()` instead of `foo := &Foo{...}`)
  * Within constructor functions, the keyword `new` shall be used (`f := new(Foo); f.x = y` instead of `f := &Foo{x: y}`)
* [![Reviewed by Hound](https://img.shields.io/badge/Reviewed_by-Hound-8E64B0.svg)](https://houndci.com)
* Readability > Performance (there **are** exceptions, but don't optimize prematurely)

# Getting started

* Checkout the repository
* Run `make shorttest` or `go run -short ./...` to ensure everything works
* To test everything, run `make test` (includes parser and conformance tests, takes a while)
* To run the application, run one of
    * `go run ./cmd`
    * `make run` (does exactly what the line above does)
    * `make start` (builds into `/bin`, then executes built file)
* To build the application, run `make build`. It will build an executable file into `/bin`