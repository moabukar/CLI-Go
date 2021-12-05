# Creating a simple CLI tool in the Go Programming Language for personal learning and fun

Open to feedback :) 

## Build docker dev environment

```sh

docker build --target dev . -t go
docker run -it -v ${PWD}:/work go sh
go version

```

## Test CLI

```sh

docker run -it -v ${PWD}:/work go sh
cd src
go build main.go

./videos get --all (gets all videos

./videos add -id test -imageurl hello -title vidtitle -url testurl -desc test

./videos get --all

| ID   |    Title |  URL  |   ImageURL    |    Description |
| -----| --------- | -------- | ----- | ------ |
| test  |   vidtitle  |      testurl    |     hello  | test |
)
``` 