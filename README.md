# Merge-Challenge

merge-challenge is a programming task where you have to implement a function that can merge a list of intervals. Merging intervals means inserting intervals that have an intersection or fit into each other and assembling them. All intervals that do not fit together remain untouched and are added. 

### Env

`INTERVALS` : `string` for example: `[25,30] [2,19] [14,23] [4,8]`

### Run and Build

you can start the program as docker as a scratch image or compile directly via go.

Compile the program

`go build main.go`

and run the Binary `./main` 

or 

`go run main.go`

and with Docker:

`docker build -t mergeIntervals --build-arg INTERVALS="[25,30] [2,19] [14,23] [4,8]" .`

