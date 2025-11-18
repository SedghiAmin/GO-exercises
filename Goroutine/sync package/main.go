package main

import(
	"fmt"
	"sync"
)

type SafeCounter struct{
	mu sync.Mutex
	val map[string]int
}

func (sc *SafeCounter) Inc(key string){
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.val[key]++
}

func (sc *SafeCounter) Value(key string) int{
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.val[key]
}

func main(){
	sc:= SafeCounter{val: make(map[string]int)}
	var wg sync.WaitGroup

	for i := 1; i <= 1000; i++{
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			sc.Inc("index")
		}()
	}

	wg.Wait()
	fmt.Println(sc.Value("index"))

}