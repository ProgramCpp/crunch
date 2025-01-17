# crunch

a number cruncher. A service to manage a generic counter.

## prerequisites
- go 1.22+

## build
```
go build  -o ./build/
```

## run
```
./build/crunch
```

## Future improvements
- add persistence to counters
- add api for async counters for use cases that do not need accuracy. latency vs consistency