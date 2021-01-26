# go-times
[![Go Report Card](https://goreportcard.com/badge/github.com/shinYeongHyeon/go-times)](https://goreportcard.com/report/github.com/shinYeongHyeon/go-times)
[![codecov](https://codecov.io/gh/shinYeongHyeon/go-times/branch/master/graph/badge.svg?token=2A399RA25X)](https://codecov.io/gh/shinYeongHyeon/go-times)
[![GoDoc](https://pkg.go.dev/badge/github.com/shinYeongHyeon/go-times?status.svg)](https://pkg.go.dev/github.com/shinYeongHyeon/go-times)

It will save your time such as coding a calendar program.

## Installation
To install `go-times` package, you need to install Go and set your Go workspace first.
```sh
$ go get -u github.com/shinYeongHyeon/go-times
```

## Contents

- [Contents](#contents)
  - [GetDateOfFirstMondayOfMonth](#getDateOfFirstMondayOfMonth)
  - [GetDateOfFirstTuesdayOfMonth](#getDateOfFirstTuesdayOfMonth)
  - [GetDateOfFirstWednesdayOfMonth](#getDateOfFirstWednesdayOfMonth)
  - [GetDateOfFirstThursdayOfMonth](#getDateOfFirstThursdayOfMonth)
  - [GetDateOfFirstFridayOfMonth](#getDateOfFirstFridayOfMonth)
  - [GetDateOfFirstSaturdayOfMonth](#getDateOfFirstSaturdayOfMonth)
  - [GetDateOfFirstSundayOfMonth](#getDateOfFirstSundayOfMonth)
  - [GetNthWeekOfMonth](#getNthWeekOfMonth)
    
## Dependency
 - [time](https://golang.org/pkg/time/)

## Description

#### GetDateOfFirstMondayOfMonth
This will get date of first monday of month
```go
func GetDateOfFirstMondayOfMonth(t time.Time) int
```

#### GetDateOfFirstTuesdayOfMonth
This will get date of first tuesday of month
```go
func GetDateOfFirstTuesdayOfMonth(t time.Time) int
```

#### GetDateOfFirstWednesdayOfMonth
This will get date of first wednesday of month
```go
func GetDateOfFirstWednesdayOfMonth(t time.Time) int
```

#### GetDateOfFirstThursdayOfMonth
This will get date of first thursday of month
```go
func GetDateOfFirstThursdayOfMonth(t time.Time) int
```

#### GetDateOfFirstFridayOfMonth
This will get date of first friday of month
```go
func GetDateOfFirstFridayOfMonth(t time.Time) int
```

#### GetDateOfFirstSaturdayOfMonth
This will get date of first saturday of month
```go
func GetDateOfFirstSaturdayOfMonth(t time.Time) int
```

#### GetDateOfFirstSundayOfMonth
This will get date of first sunday of month
```go
func GetDateOfFirstSundayOfMonth(t time.Time) int
```

#### ~~GetDaysOfFirstSaturdayOfMonth~~
~~This will get day of first saturday of month~~  
deprecated, recommend [GetDateOfFirstSaturdayOfMonth](#getDateOfFirstSaturdayOfMonth)
```go
func GetDaysOfFirstSaturdayOfMonth(t time.Time) int
```

#### GetNthWeekOfMonth
This will get `n`th Week of Month  
(Regard the start of the week as Sunday)
```go
func GetNthWeekOfMonth(t time.Time) int
```
