# Creating a simple CLI tool in the Go Programming Language for personal learning and fun

Open to feedback :) 

## Build docker dev environment

```sh

docker build --target dev . -t go
docker run -it -v ${PWD}:/work go sh
go version

```

## Test CLI using Docker dev env

```sh

docker run -it -v ${PWD}:/work go sh
cd src
go build main.go

./videos get --all (gets all videos)

./videos add -id test -imageurl hello -title vidtitle -url testurl -desc test

./videos get --all

```

| ID   |    Title |  URL  |   ImageURL    |    Description |
| -----| --------- | -------- | ----- | ------ |
| test  |   vidtitle  |      testurl    |     hello  | test |


## Test CLI using local Go build

```sh

git clone https://github.com/moabukar/CLI-Go.git
cd src
go build main.go

./videos get --all (gets all videos)

./videos add -id test -imageurl hello -title vidtitle -url testurl -desc test (adds an inout with all the flags)

./videos get --all (returns all video lists added)

```