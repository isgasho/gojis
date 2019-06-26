# Gojis

[![Build Status](https://travis-ci.org/TimSatke/gojis.svg?branch=develop)](https://travis-ci.org/TimSatke/gojis)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/dd5507e3d34543e3a526b05aaea3eba8)](https://www.codacy.com/app/gojisvm/gojis?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gojisvm/gojis&amp;utm_campaign=Badge_Grade)
[![Reviewed by Hound](https://img.shields.io/badge/Reviewed_by-Hound-8E64B0.svg)](https://houndci.com)

<a href="https://www.buymeacoffee.com/timsatke" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

Gojis is an implementation of ECMAScript 2018 (ES 9). It basically is a JavaScript VM, just like
[Goja](https://github.com/dop251/goja) or [Otto](https://github.com/robertkrimen/otto).
The documentation can be found [here](https://gojisvm.github.io).

## Why
[Goja](https://github.com/dop251/goja) and [Otto](https://github.com/robertkrimen/otto) are both stuck at implementing _most_ features of ES 5.1. This implementation aims to support ES 9, and, after that maybe even ES 10 (not released as of 2019-06-27) and later.

## What is this Gojis VM good for?
The Gojis VM can be run as a standalone, or you can embed it into your Go application by using the API. Go get it with
```bash
go get -u github.com/gojisvm/gojis
```
and start using it with
```go
// FIXME: design an API, this is just a draft how it COULD look like

vm := gojis.NewVM()

proc := vm.EvaluateFiles("./*.js")
<-proc.Done()
```

For more documentation, please have a look at the [API documentation](https://gojisvm.github.io/api.html).

## What are the goals?
The primary goal of this project is to have fun coding, as I love to code, but thinking about system designs and architectures is difficult. The ECMAScript language specification (which can be found in `/docs`), takes care of most of these things already, so a contributor can really focus on implementation and optimization.

Another goal I am trying to achieve is, to provide the Go community with a JavaScript VM that supports at least ES 6 features.
[Goja](https://github.com/dop251/goja) and [Otto](https://github.com/robertkrimen/otto) are both stuck at implementing ES 5.1, but this implementation does exactly that.

## Current status
There is a [milestone](https://github.com/gojisvm/gojis/milestone/1) to keep track of the implementation progress of ES 9.