# crunch

a number cruncher. A service to manage a generic counter.


## design
the requirements are quite contradictory. As in, latency is inversely related to consistency. i.e, the counter can be either maintained accurately at the cost of latency or at low latency at the cost of accuracy.

latency is favored as per the requirements (2) below.
i.e, the order of api's do not influence the counter value. For practical use cases, order of concurrent requests cannot be controlled. If the order is a concern, it must be the responsibility of the client. 

```
4. Server must send increment response – a response with the new global counter value
(1) behave correctly 
(2) complete the most request-response cycles possible in one minute.

```

ex: given the current counter value is 10 and two concurrent requests update the counter by 2 and 5 respectively,
- the api can return 12, 15, 17
- at the end of the two api calls, the counter value will be 17

The counter service is built on the HTTP protocol.

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
- integration tests and load tests to measure performance
- api documentation
- add persistence to counters