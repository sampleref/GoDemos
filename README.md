
set GOROOT=<Go Installation> # Not /bin
set PATH=PATH + $GOROOT/bin

mkdir ../../GoDemos => For new Projects

git clone 

set GOPATH=../../GoDemos

set PATH="%PATH%;%GOPATH%\bin"

set http_proxy=<if needed>
set https_proxy=<if needed>

-> Install go build tool glide manually or through https://github.com/Masterminds/glide

mkdir  ../../GoDemos/src/<project Name> => For new projects

************* In <project Name> ***************
create main.go with package main and func main.

Set Settings>Lang & Frameworks>Go>Go Libraries> User Project Libraries to ../../GoDemos - *** If using Intellij Idea

go build -o HelloDemo.exe main.go

glide init -> to generate glide.yaml file
glide update -> to update later on

To use protoc/grpc
- Creates extra proto folder within proto => protoc --go_out=plugins=grpc:./proto/ ./proto/*.proto
- Creates .go in same proto file folder => protoc --go_out=plugins=grpc:.\ .\proto\*.proto
                                 linux => protoc --go_out=plugins=grpc:./ ./proto/*.proto