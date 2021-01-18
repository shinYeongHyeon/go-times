# go-times
[![Go Report Card](https://goreportcard.com/badge/github.com/shinYeongHyeon/go-times)](https://goreportcard.com/report/github.com/shinYeongHyeon/go-times)
[![GoDoc](https://pkg.go.dev/badge/github.com/shinYeongHyeon/go-times?status.svg)](https://pkg.go.dev/github.com/shinYeongHyeon/go-times)

It will save your time such as coding a calendar program.

## Installation
To install `go-times` package, you need to install Go and set your Go workspace first.
```sh
$ go get -u github.com/shinYeongHyeon/go-times
```

## Contents

- [Contents](#contents)
  - [GetDaysOfFirstSaturdayOfMonth](#getDaysOfFirstSaturdayOfMonth)
  - [GetNthWeekOfMonth](#getNthWeekOfMonth)
    
## Dependency
 - [time](https://golang.org/pkg/time/)

## Description

#### GetDaysOfFirstSaturdayOfMonth
This will get day of first saturday of month
```go
func GetDaysOfFirstSaturdayOfMonth(t time.Time) int
```

#### GetNthWeekOfMonth
This will get `n`th Week of Month  
(Regard the start of the week as Sunday)
```go
func GetNthWeekOfMonth(t time.Time) int
```
