# F1 Ergast Services
Module for consume data from Ergast API. This public web service provides a historical record of Formula One data.

Installation
============

To install F1-Ergast-Go, use `go get`:

    go get github.com/ebcardoso/f1-ergast-go

This will then make the following packages available to you:

    github.com/ebcardoso/f1-ergast-go/modules/f1_services

Usage
=====

## Drivers
```go
//List of all drivers
result, err := f1_services.ListDrivers(offset int, limit int)
//List of all drivers within a year
result, err := f1_services.ListDriversBySeason(year int, offset int, limit int)
//List of all drivers within a race
result, err := f1_services.ListDriversByRace(year int, round int, offset int, limit int)
//Driver Information
result, err := f1_services.FindDriverById(driverId string)
```

## Constructors
```go
//List of all constructors
result, err := f1_services.List(offset int, limit int)
//List of all constructors within a year
result, err := f1_services.BySeason(year int, offset int, limit int)
//List of all constructors within a race
result, err := f1_services.ByRace(year int, round int, offset int, limit int)
//Constructor Information
result, err := f1_services.GetByConstructorId(constructorId string)
```

## Circuits
```go
//List of all circuit
result, err := f1_services.ListCircuits(offset int, limit int)
//List of all circuit within a year
result, err := f1_services.ListCircuitsBySeason(year int, offset int, limit int)
//List of all circuit within a race
result, err := f1_services.ListCircuitsByRace(year int, round int, offset int, limit int)
//Circuit Information
result, err := f1_services.FindCircuitById(circuitId string)
```

## Seasons
```go
//Seasons List
result, err := f1_services.ListSeasons(offset int, limit int)
```

## Results
```go
//Most recent race result
result, err := f1_services.MostRecentResults()
//Race Result
result, err := f1_services.ListResultsByRace(year int, round int)
```

## Qualifying
```go
//Qualifying Results
result, err := f1_services.ListQualifyingByRace(year int, round int, offset int, limit int)
```

## Schedules
```go
//Schedule of races for a season
result, err := f1_services.ListSchedulesBySeason(year int, offset int, limit int)
//Schedule of the current season
result, err := f1_services.ListSchedulesByCurrentSeason(offset int, limit int)
//Schedule of a specific race
result, err := f1_services.ListSchedulesByRace(year int, round int)
```

## Standings
```go
//Current drivers' standing
result, err := f1_services.DriversCurrentStandingss()
//Season and driver standings
result, err := f1_services.DriversStandingsBySeason(year int, offset int, limit int)
//Driver standings after a race
result, err := f1_services.DriversStandingsAfterRace(year int, round int, offset int, limit int)
//Driver standings by specifying the driver
result, err := f1_services.StandingsHistoryByDriver(driverId string, offset int, limit int)
//All winners of drivers' championships
result, err := f1_services.ChampionsHistoryByDrivers(offset int, limit int)
//Current constructor's standing
result, err := f1_services.ConstructorsCurrentStandings()
//Season and constructor standing
result, err := f1_services.ConstructorsStandingsBySeason(year int, offset int, limit int)
//Constructor standings after a race
result, err := f1_services.ConstructorsStandingsAfterRace(year int, round int, offset int, limit int)
//Constructor standings by specifying the constructor
result, err := f1_services.StandingsHistoryByConstructor(constructorId string, offset int, limit int)
//All winners of constructors' championships
result, err := f1_services.ChampionsHistoryByConstructors(offset int, limit int)
```

## Finishing Status
```go
//List of all finishing status codes
result, err := f1_services.ListFinishingStatus(offset int, limit int)
//List of finishing status for a specific season
result, err := f1_services.ListFinishingStatusBySeason(year int, offset int, limit int)
//List of finishing status for a specific round in a season
result, err := f1_services.ListFinishingStatusByRace(year int, round int, offset int, limit int)
```

## Laps
```go
//Lap time
result, err := f1_services.ListLapTimes(year int, round int, lap int, offset int, limit int)
```

## PitStops
```go
//Pit stop data for a race
result, err := f1_services.ListPitstopsByRace(year int, round int, offset int, limit int)
//Information for a specific pit stop
result, err := f1_services.ListByPitstopNumber(year int, round int, number int, offset int, limit int)
```
