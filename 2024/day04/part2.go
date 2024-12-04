package main

import (
	"aocli/utils/reader"
	"image"
	"sync"
)

var mdelta = []image.Point{
	{-1, -1}, {-1, 1},
}

var aletters = [][][]rune{
	{{'M', 'S'}, {'S', 'M'}},
	{{'M', 'S'}, {'S', 'M'}},
}

func doPartTwo(input string) int {
	var mapper Mapper
	mapper, allA := makeMap(reader.FileLineByLine(input), 'A')
	var wg sync.WaitGroup
	var ch = make(chan int, len(allA))
	for _, i := range allA {
		wg.Add(1)
		func(i image.Point, wg *sync.WaitGroup, ch chan int) {
			defer wg.Done()
			ch <- mapper.checkMAS(i)
		}(i, &wg, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	var res int
	for out := range ch {
		res += out
	}
	return res
}

func (m Mapper) checkMAS(XY image.Point) int {
	for _, a := range aletters {
		var succ int
		for i, d := range mdelta {
			if (m[XY.Add(d)] == a[i][0] && m[XY.Sub(d)] == a[i][1]) || (m[XY.Add(d)] == a[i][1] && m[XY.Sub(d)] == a[i][0]) {
				succ++
			}
		}
		if succ == 2 {
			return 1
		}
	}

	return 0
}
