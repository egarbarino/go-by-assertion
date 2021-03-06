---
title: Go By Assertion
author: Ernesto Garbarino
date: 2019-11-26
abstract: A Go reference using unit testing assertions as opposed to print statements. 
image: go-by-assertion.png
---

# Introduction

This is a compact Go reference based on _unit testing_ assertions as opposed to print statements---usually the norm. Comments are kept to a minimum; the reader is assumed to be familiarised with C-style languages (Java, C#, C++, etc.). 

The order of topics is relatively bottom-up (basic topics appear before complex ones) but there are some circular dependencies. For example, covering _control flow_ requires some understanding of basic types.

For convenience, the assertions are stylised as follows:

* `assert.Equal(t, a, b)` becomes `a ⇔ b`
* `assert.NotEqual(t, a, b)` becomes `a ⇎ b` 

All source code snippets are linked to original files on GitHub at [https://github.com/egarbarino/go-by-assertion/](https://github.com/egarbarino/go-by-assertion/)


