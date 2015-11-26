# Merging inbox channels to one outbox channel

Generic merge with reflection VS Static int-typed merge:

```
merge|⇒ go test -bench=.
testing: warning: no tests to run
PASS
BenchmarkReflectMerge1InputSmallAmountValues-4   	   20000	     66923 ns/op
BenchmarkStaticIntMerge1InputSmallAmountValues-4 	   50000	     29042 ns/op
BenchmarkReflectMerge5InputSmallAmountValues-4   	   10000	    325126 ns/op
BenchmarkStaticIntMerge5InputSmallAmountValues-4 	   10000	    108772 ns/op
BenchmarkReflectMerge50InputSmallAmountValues-4  	   10000	   3741938 ns/op
BenchmarkStaticIntMerge50InputSmallAmountValues-4	    5000	   1389956 ns/op
BenchmarkReflectMerge1InputLargeAmountValues-4   	   10000	    581571 ns/op
BenchmarkStaticIntMerge1InputLargeAmountValues-4 	   10000	    279794 ns/op
BenchmarkReflectMerge5InputLargeAmountValues-4   	   10000	   3065348 ns/op
BenchmarkStaticIntMerge5InputLargeAmountValues-4 	    5000	   1398173 ns/op
BenchmarkReflectMerge50InputLargeAmountValues-4  	   10000	  33131181 ns/op
BenchmarkStaticIntMerge50InputLargeAmountValues-4	    5000	  13828888 ns/op
ok  	github.com/waterlink/workbench/go/channels/merge	506.745s
```

## Where

Benchmark name scheme:

```
Benchmark    ReflectMerge     5 Input       Small Amount Values
Benchmark    StaticIntMerge   5 Input       Small Amount Values
             ^func name ^    ^ inputs ^     ^ amount of values ^
```

Benched functions:
- `merge.ReflectMerge`
- `StaticIntMerge`

Benched input sizes (count of input channels to merge):
- 1
- 5
- 50

Benched data sizes (count of values flowing through EACH single input channel):
- Small = 50
- Large = 500

## Outcome

`reflect` version is about 3 times slower. Which can affect apps that mostly
spend time sending data through channels. Apps that have a lot of IO and
sometimes send data through channel, can benefit from that.

## TODO

- [ ] Bench with bigger type (struct with some strings and ints).
