# F1 Ergast Services
Module for consume data from Ergast API. This public web service provides a historical record of Formula One data.

Installation
============

To install F1-Ergast-Go, use `go get`:

    go get github.com/ebcardoso/f1-ergast-go

This will then make the following packages available to you:

    github.com/ebcardoso/f1-ergast-go/modules/drivers
    github.com/ebcardoso/f1-ergast-go/modules/constructors
    github.com/ebcardoso/f1-ergast-go/modules/circuits
    github.com/ebcardoso/f1-ergast-go/modules/seasons
    github.com/ebcardoso/f1-ergast-go/modules/results
    github.com/ebcardoso/f1-ergast-go/modules/qualifying
    github.com/ebcardoso/f1-ergast-go/modules/schedules
    github.com/ebcardoso/f1-ergast-go/modules/standings
    github.com/ebcardoso/f1-ergast-go/modules/finishing_status
    github.com/ebcardoso/f1-ergast-go/modules/laps
    github.com/ebcardoso/f1-ergast-go/modules/pitstops

Usage
=====

## Drivers
```go
//List of all drivers
result, err := drivers.List(offset int, limit int)
//List of all drivers within a year
result, err := drivers.BySeason(year int, offset int, limit int)
//List of all drivers within a race in a year
result, err := drivers.ByRace(year int, round int, offset int, limit int)
//Driver Information
result, err := drivers.GetByDriverId(driverId string)
```

## Constructors
```go
//List of all constructors
result, err := constructors.List(offset int, limit int)
//List of all constructors within a year
result, err := constructors.BySeason(year int, offset int, limit int)
//List of all constructors within a race in a year
result, err := constructors.ByRace(year int, round int, offset int, limit int)
//Constructor Information
result, err := constructors.GetByConstructorId(constructorId string)
```

## Circuits
```go
//List of all circuit
result, err := circuits.List(offset int, limit int)
//List of all circuit within a year
result, err := circuits.BySeason(year int, offset int, limit int)
//List of all circuit within a race in a year
result, err := circuits.ByRace(year int, round int, offset int, limit int)
//Circuit Information
result, err := circuits.GetByCircuitId(circuitId string)
```

## Seasons
```go
//Seasons List
result, err := seasons.List(offset int, limit int)
```

## Results
```go
//Race Result
result, err := results.ByRace(year int, round int)
//Most recent race result
result, err := results.MostRecent()
```

## Qualifying
```go
//Qualifying Results
result, err := qualifying.ByRace(year int, round int, offset int, limit int)
```

## Schedules
```go
//Schedule of races for a season
result, err := schedules.BySeason(year int, offset int, limit int)
//Schedule of the current season
result, err := schedules.ByCurrentSeason(offset int, limit int)
//Info about a specific race
result, err := schedules.ByRace(year int, round int)
```

## Standings
```go
//Driver standings after a race
result, err := standings.DriversAfterRace(year int, round int, offset int, limit int)
//Constructor standings after a race
result, err := standings.ConstructorsAfterRace(year int, round int, offset int, limit int)
//Season end driver standings
result, err := standings.DriversBySeason(year int, offset int, limit int)
//Season end constructor standing
result, err := standings.ConstructorsBySeason(year int, offset int, limit int)
//Current drivers' standing
result, err := standings.DriversCurrent()
//Current constructor's standing
result, err := standings.ConstructorsCurrent()
//All winners of drivers' championships
result, err := standings.WinnersByDrivers(offset int, limit int)
//All winners of constructors' championships
result, err := standings.WinnersByConstructors(offset int, limit int)
//Driver standings by specifying the driver
result, err := standings.HistoryByDriver(driverId string, offset int, limit int)
//Constructor standings by specifying the constructor
result, err := standings.HistoryByConstructor(constructorId string, offset int, limit int)
```

## Finishing Status
```go
//List of all finishing status codes
result, err := finishing_status.List(offset int, limit int)
//List of finishing status for a specific season
result, err := finishing_status.BySeason(year int, offset int, limit int)
//List of finishing status for a specific round in a season
result, err := finishing_status.ByRace(year int, round int, offset int, limit int)
```

## Laps
```go
//Lap time
result, err := laps.GetLapTimes(year int, round int, lap int, offset int, limit int)
```

## PitStops
```go
//Pit stop data for a race
result, err := pitstops.ByRace(year int, round int, offset int, limit int)
//Information for a specific pit stop
result, err := pitstops.GetByNumber(year int, round int, number int, offset int, limit int)
```
