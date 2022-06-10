package requests

import (
	"fmt"
	"math"
	"sync"
)

func InfoCurrentItem(links []string) []FloatInfo {
	var wg sync.WaitGroup

	flCh := make(chan FloatInfo)
	var floatInfoList []FloatInfo

	start := 0

	countOfGoRoutines := int(math.Ceil(float64(len(links)) / 100))
	fmt.Println("Count of goroutines: ", countOfGoRoutines)

	if len(links) > 100 {
		for i := 0; i < countOfGoRoutines; i++ {
			wg.Add(1)
			count := len(links) - start
			if count <= 100 {
				go func(urls []string, ch chan FloatInfo) {

					GetExtraInfo(urls, ch)
					wg.Done()
					fmt.Println("Done goroutine")

				}(links[start:start+count], flCh)
			} else {
				go func(urls []string, ch chan FloatInfo) {

					GetExtraInfo(urls, ch)
					wg.Done()
					fmt.Println("Done goroutine")

				}(links[start:start+100], flCh)
			}
			start += 100
		}
	} else {
		wg.Add(1)
		go func(urls []string, ch chan FloatInfo) {

			GetExtraInfo(urls, ch)
			wg.Done()

		}(links, flCh)
	}

	go func() {
		wg.Wait()
		fmt.Println("Goroutines done")
		close(flCh)
		fmt.Println("Channel closed")
	}()

	for v := range flCh {
		floatInfoList = append(floatInfoList, v)
	}

	return floatInfoList
}