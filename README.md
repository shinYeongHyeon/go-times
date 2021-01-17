# go-times

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
```go
func GetNthWeekOfMonth(t time.Time) int
```