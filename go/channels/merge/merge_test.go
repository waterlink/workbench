package merge_test

import (
	"testing"

	"github.com/waterlink/workbench/go/channels/merge"
)

const (
	small = 50
	large = 500
)

func BenchmarkReflectMerge1InputSmallAmountValues(b *testing.B) {
	for n := 0; n < b.N; n++ {
		is := giveMeSomeChannelsI(1)
		o := make(chan int)

		merge.ReflectMerge(o, is...)

		sendIntsToChanI(small, is)
		go consumeSomeInts(small, o)
	}
}

func BenchmarkStaticIntMerge1InputSmallAmountValues(b *testing.B) {
	for n := 0; n < b.N; n++ {
		is := giveMeSomeChannels(1)
		o := make(chan int)

		merge.StaticIntMerge(o, is...)

		sendIntsToChan(small, is)
		go consumeSomeInts(small, o)
	}
}

func BenchmarkReflectMerge5InputSmallAmountValues(b *testing.B) {
	for n := 0; n < b.N; n++ {
		is := giveMeSomeChannelsI(5)
		o := make(chan int)

		merge.ReflectMerge(o, is...)

		sendIntsToChanI(small, is)
		go consumeSomeInts(small*5, o)
	}
}

func BenchmarkStaticIntMerge5InputSmallAmountValues(b *testing.B) {
	for n := 0; n < b.N; n++ {
		is := giveMeSomeChannels(5)
		o := make(chan int)

		merge.StaticIntMerge(o, is...)

		sendIntsToChan(small, is)
		go consumeSomeInts(small*5, o)
	}
}

func BenchmarkReflectMerge50InputSmallAmountValues(b *testing.B) {
	for n := 0; n < b.N; n++ {
		is := giveMeSomeChannelsI(50)
		o := make(chan int)

		merge.ReflectMerge(o, is...)

		sendIntsToChanI(small, is)
		go consumeSomeInts(small*50, o)
	}
}

func BenchmarkStaticIntMerge50InputSmallAmountValues(b *testing.B) {
	for n := 0; n < b.N; n++ {
		is := giveMeSomeChannels(50)
		o := make(chan int)

		merge.StaticIntMerge(o, is...)

		sendIntsToChan(small, is)
		go consumeSomeInts(small*50, o)
	}
}

func BenchmarkReflectMerge1InputLargeAmountValues(b *testing.B) {
	for n := 0; n < b.N; n++ {
		is := giveMeSomeChannelsI(1)
		o := make(chan int)

		merge.ReflectMerge(o, is...)

		sendIntsToChanI(large, is)
		go consumeSomeInts(large, o)
	}
}

func BenchmarkStaticIntMerge1InputLargeAmountValues(b *testing.B) {
	for n := 0; n < b.N; n++ {
		is := giveMeSomeChannels(1)
		o := make(chan int)

		merge.StaticIntMerge(o, is...)

		sendIntsToChan(large, is)
		go consumeSomeInts(large, o)
	}
}

func BenchmarkReflectMerge5InputLargeAmountValues(b *testing.B) {
	for n := 0; n < b.N; n++ {
		is := giveMeSomeChannelsI(5)
		o := make(chan int)

		merge.ReflectMerge(o, is...)

		sendIntsToChanI(large, is)
		go consumeSomeInts(large*5, o)
	}
}

func BenchmarkStaticIntMerge5InputLargeAmountValues(b *testing.B) {
	for n := 0; n < b.N; n++ {
		is := giveMeSomeChannels(5)
		o := make(chan int)

		merge.StaticIntMerge(o, is...)

		sendIntsToChan(large, is)
		go consumeSomeInts(large*5, o)
	}
}

func BenchmarkReflectMerge50InputLargeAmountValues(b *testing.B) {
	for n := 0; n < b.N; n++ {
		is := giveMeSomeChannelsI(50)
		o := make(chan int)

		merge.ReflectMerge(o, is...)

		sendIntsToChanI(large, is)
		go consumeSomeInts(large*50, o)
	}
}

func BenchmarkStaticIntMerge50InputLargeAmountValues(b *testing.B) {
	for n := 0; n < b.N; n++ {
		is := giveMeSomeChannels(50)
		o := make(chan int)

		merge.StaticIntMerge(o, is...)

		sendIntsToChan(large, is)
		go consumeSomeInts(large*50, o)
	}
}

func giveMeSomeChannels(n int) []chan int {
	cs := make([]chan int, n)
	for j := 0; j < n; j++ {
		cs[j] = make(chan int)
	}
	return cs
}

func giveMeSomeChannelsI(n int) []interface{} {
	cs := make([]interface{}, n)
	for j := 0; j < n; j++ {
		cs[j] = make(chan int)
	}
	return cs
}

func sendIntsToChan(n int, os []chan int) {
	for j := 0; j < len(os); j++ {
		go sendSomeInts(n, os[j])
	}
}

func sendIntsToChanI(n int, os []interface{}) {
	for j := 0; j < len(os); j++ {
		go sendSomeInts(n, os[j].(chan int))
	}
}

func sendSomeInts(n int, o chan<- int) {
	for j := 0; j < n; j++ {
		o <- j
	}
	close(o)
}

func consumeSomeInts(n int, i <-chan int) {
	for j := 0; j < n; j++ {
		<-i
	}
}
