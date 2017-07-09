# Flipped

A backend program for [Hack.init](http://hackinit.org/) project.

## Prerequisite
- Golang
- Govendor
- Makefile

## Quick start
- config
```
cp config.example.json config.json
change config in config.json
```
- pull pkgs
```
govender sync
```
- build
```
make dev-build
```
- run 
```
make dev-run
```