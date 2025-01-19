# crunch

a number cruncher. A service to manage a generic counter.


## design
the requirements are quite contradictory. As in, latency is inversely related to consistency. i.e, the counter can be either maintained accurately at the cost of latency or at low latency at the cost of accuracy.
```
4. Server must send increment response – a response with the new global counter value
(1) behave correctly 
(2) complete the most request-response cycles possible in one minute.

```
latency is favored as per the requirements (2) above. The counter is updated asynchronously to avoid api contention. The requests are buffered and eventually updated.
i.e, the order of api's do not influence the counter value. For practical use cases, order of concurrent requests cannot be controlled. 

ex: given the current counter value is 10 and two concurrent requests update the counter by 2 and 5 respectively,
- the api's can return 10 or 12 and 15 or 17 respectively.
- at the end of the two api calls, the counter value will be 17, after a delay.

downside:
the counter is not updated realtime. 
maximum parallelism depend on the number of cores. Are concurrent atomic updates faster than pushing the updates to a channel at scale?
todo: at scale, monitor contention to update the counter value concurrently. If atomic operations do not affect latency, perform synchronous updates.


The counter service is built on the HTTP protocol, with json data format.

## prerequisites
- go 1.22+

## build
```
go build  -o ./build/
```

## test
```
curl -H 'Content-Type: application/json' \
      -d '{ "count":10}' \
      -X POST \
      http://localhost:8080/counter
```

## run
```
./build/crunch
```

## Future improvements
- integration tests and load tests to measure performance
- api documentation
- add persistence to counters