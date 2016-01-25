# fg-frames.go
A golang based, http web service that provides fighting game character
frame data. A proof of concept service built as a means to explore
golang and various http tool kits. This service only hosts Ultra Street
Fighter 4 version 1.05 data at the moment, but will be expanded to
include other games in the future.

## Data sources
- [eventhubs](http://www.eventhubs.com/moves/sf4/)
- [wiki.shoryuken](http://wiki.shoryuken.com/Ultra_Street_Fighter_IV)

The compiled data used by this service can be found here
- [source](https://github.com/jpgnotgif/frame-data)

## Bootstrap
- [Install golang](https://golang.org/doc/install)
- Clone this repo
- Generate the binary

```
$ go install github.com/jpgnotgif/fgframes
$ GOPATH/bin/fgframes
```

## JSON Schema
```
{
  normal-attack: {
    standing: { startup: num, active: num, recovery: num,
block-advantage: num, hit-advantage: num }
    close standing: {...}
    crouching: {...}
    neutral jump: {...}
    diagonal jump: {...}
  }
  unique-attack: {...}
}
```

## Attack Legend
Notation|Expanded
--------|--------
lp|Light Punch / Jab
mp|Medium Punch / Strong
hp|Hard Punch / Fierce
lk|Light Kick / Short
mk|Medium Kick / Forward
hk|Hard Kick / Roundhouse

## Character State Legend
Notation|Expanded
--------|-------
s|Standing
cs|Close Standing
cr|Crouching
nj|Neutral Jump / Jump Up
dj|Diagonal Jump / Jump Toward

## Frame Data Legend
Notation|Expanded
--------|--------
s|Startup
a|Active
r|Recovery
ba|Block Advantage
ha|Hit Advantage

## Tech
- [golang](https://golang.org/)
- [gin gonic](https://gin-gonic.github.io/gin/)
