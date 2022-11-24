# Conclusion

When you know the exact number of tasks to be executed, this pool will be faster.

Because it reduces the allocation of goroutines.

```
BenchmarkNaive
BenchmarkNaive-8   	   10000	      1136 ns/op	     225 B/op	       2 allocs/op
BenchmarkPool
BenchmarkPool-8    	   10000	       727.9 ns/op	     107 B/op	       2 allocs/op
```

However, if you don't know, or you have a number of task to be executed.

And you can't stand the performance of waiting, or you have enough resources to waste. Using go is the best.

```
BenchmarkNaive
BenchmarkNaive-8   	 2458435	       421.6 ns/op	      96 B/op	       2 allocs/op
BenchmarkPool
BenchmarkPool-8    	 1739314	       708.9 ns/op	      16 B/op	       1 allocs/op
```

# About the sequence

Any "simulated" tasks in go will be out-of-order possibly.

So if the order must be sorted, limit the number of workers to 1.

That will keep the order.

Otherwise, like that

```
=== RUN   TestSeq
    worker_test.go:58: 1
    worker_test.go:58: 2
    worker_test.go:58: 3
    worker_test.go:58: 4
    worker_test.go:58: 5
    worker_test.go:58: 6
    worker_test.go:58: 13
    worker_test.go:58: 7
    worker_test.go:58: 8
    worker_test.go:58: 9
    worker_test.go:58: 47
    worker_test.go:58: 10
    worker_test.go:58: 11
    worker_test.go:58: 12
    worker_test.go:58: 60
    worker_test.go:58: 14
    worker_test.go:58: 15
    worker_test.go:58: 76
    worker_test.go:58: 16
    worker_test.go:58: 17
    worker_test.go:58: 85
    worker_test.go:58: 18
    worker_test.go:58: 19
```

It's a horrible case when you are handling the packet.

# Finally

Thanks to @gobwas 