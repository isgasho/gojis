# How to contribute

1. Fork the repository
2. Make changes
3. Open a Pull Request

# Code requirements

* For every change, there **MUST** be a test case that proves, that your change did not break anything
* All tests must pass successfully
* No test cases can be removed and/or changed (if a test fails, but ran green in the past, this is a regression)
  * If you think a test is incorrect, open an issue, stating why the test should be changed/removed
* To be consistent, new objects are created with constructor methods (`foo := NewFoo()` instead of `foo := &Foo{...}`)
  * Within constructor functions, the keyword `new` shall be used (`f := new(Foo); f.x = y` instead of `f := &Foo{x: y}`)
* The formatter to be used is `gopls`'s default formatter, **NOT** `gofmt`
* Code should be [fail-fast](https://en.wikipedia.org/wiki/Fail-fast)
* Readability > Performance (there **are** exceptions, but don't optimize prematurely)
