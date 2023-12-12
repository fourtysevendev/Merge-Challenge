# Merge-Challenge

merge-challenge is a programming task where you have to implement a function that can merge a list of intervals. Merging intervals means inserting intervals that have an intersection or fit into each other and assembling them. All intervals that do not fit together remain untouched and are added. 

### Env

`INTERVALS` : `string` for example: `[25,30] [2,19] [14,23] [4,8]`

**Important to know an interval must be  *a ≤ x ≤ b* and the pattern shows like \`\[\d+,\d+\]\`**

### Run and Build

you can start the program as docker as a scratch image or compile directly via go.

Compile the program

`go build main.go`

and run the Binary `INTERVALS="[25,30] [2,19] [14,23] [4,8]" ./main` 

or 

`INTERVALS="[25,30] [2,19] [14,23] [4,8]" go run main.go`

and with Docker:

**first you need to compile on amd64**

`GOOS=linux GOARCH=amd64 go build main.go`

`docker build -t mergeIntervals --build-arg INTERVALS="[25,30] [2,19] [14,23] [4,8]" . && docker run mergeIntervals`

### Program Runtime

in `program_duration_test.go` you can test the merge function with random intervals and you can also specify how many random Intervals you want to merge. With `10` random Intervals are approx. `100µs` with `100` approx. `200µs` and with 1000 it was just under `1ms`.
So the runtime function is approximately linear.
The first part is but not linear, when the list is sorted by intervals. This is logarithmic (`O(nlogn)`).
therefore the step from 100 to 1000 is bigger than from 10 to 100, although the factor is the same. This is due to the logarithmic function. 

### Memory Usage

Go has an automatic garbage collector and we do not need to manage or share memory. We have no cache processes in the program, the memory usage is also linear. When merging, we do that inplace, so this is also linear to the input.