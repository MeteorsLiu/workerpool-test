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

And you can't stand the performance of waiting. Using go is the best.

```

```