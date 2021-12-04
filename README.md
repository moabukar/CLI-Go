# Creating a simple CLI tool in the Go Programming Language for personal learning and fun

Open to feedback :) 

## Build docker dev environment

```sh

docker build --target dev . -t go
docker run -it -v ${PWD}:/work go sh
go version

```