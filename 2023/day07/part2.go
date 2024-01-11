package main

import (
	"reflect"
	"slices"
)

func doPartTwo(input string) int {
	hands := getHands(input, true)
	Counter := func(H Hand) int {
		C := map[int]int{}
		for _, c := range H.hand {
			C[c]++
		}
		T := reflect.ValueOf(C).MapKeys()[0].Interface()
		for k := range C {
			if k != 1 {
				if C[k] > C[T.(int)] || T == 1 {
					T = k
				}
			}
		}
		if _, ok := C[1]; ok && T.(int) != 1 {
			C[T.(int)] += C[1]
			delete(C, 1)
		}
		res := []int{}
		for _, v := range C {
			res = append(res, v)
		}
		slices.Sort(res)
		for _, S := range handscores {
			if reflect.DeepEqual(res, S.hand) {
				return S.strength
			}
		}
		return 0
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		Ca := Counter(a)
		Cb := Counter(b)
		if Ca < Cb {
			return -1
		}
		if Ca > Cb {
			return 1
		}
		// If we get here, they have the same score, and so need to iterate through the cards
		for i := 0; i < len(a.hand); i++ {
			if a.hand[i] < b.hand[i] {
				return -1
			}
			if a.hand[i] > b.hand[i] {
				return 1
			}
		}
		return 0
	})
	res := 0
	for i, H := range hands {
		res += (i + 1) * H.bid
	}
	return res
}
