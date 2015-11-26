package merge

import (
	"reflect"
	"sync"
)

// ReflectMerge is for merging channels of arbitrary type using reflection
func ReflectMerge(o interface{}, is ...interface{}) {
	var wg sync.WaitGroup
	wg.Add(len(is))

	ro := reflect.ValueOf(o)

	for _, i := range is {
		go func(ri reflect.Value) {
			defer wg.Done()
			for v, ok := ri.Recv(); ok; v, ok = ri.Recv() {
				ro.Send(v)
			}
		}(reflect.ValueOf(i))
	}

	go func() {
		wg.Wait()
		ro.Close()
	}()
}

// StaticIntMerge is for merging channels of integer type without reflection
func StaticIntMerge(o chan<- int, is ...chan int) {
	var wg sync.WaitGroup
	wg.Add(len(is))

	for _, i := range is {
		go func(i <-chan int) {
			defer wg.Done()
			for v := range i {
				o <- v
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(o)
	}()
}
