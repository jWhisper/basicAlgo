package schedule

import (
	"sort"
	"sync"
	"testing"
)

func TestWaterFactory(t *testing.T) {
	waterCount := 10

	waterBucket := make(chan string, waterCount*3)
	produceH := func() error {
		waterBucket <- "H"
		return nil
	}
	produceO := func() error {
		waterBucket <- "O"
		return nil
	}

	h := NewH2O()
	var wg sync.WaitGroup
	wg.Add(waterCount * 3)
	//暂时没有考虑生产失败的情况
	for i := 0; i < waterCount*2; i++ {
		go func() {
			h.produceH(produceH)
			wg.Done()
		}()
	}

	for i := 0; i < waterCount; i++ {
		go func() {
			h.produceO(produceO)
			wg.Done()
		}()
	}

	wg.Wait()

	if len(waterBucket) != waterCount*3 {
		t.Error("xxxxxxxxx")
	}

	s := make([]string, 3)
	for i := 0; i < waterCount; i++ {
		s[0] = <-waterBucket
		s[1] = <-waterBucket
		s[2] = <-waterBucket
		sort.Strings(s)
		if water := s[0] + s[1] + s[2]; water != "HHO" {
			t.Error("product pipeline error")
		}
	}
}
