package utils

import (
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

func Reverse(s string) string {
	var ret strings.Builder
	r := []rune(s)
	for i := len(r) - 1; i >= 0; i-- {
		ret.WriteRune(r[i])
	}
	return ret.String()
}

func FoundString(a []string, s string) bool {
	for _, v := range a {
		if strings.HasPrefix(v, s) {
			return true
		}
	}
	return false
}

func FindInSlice[T comparable](S []T, s T) bool {
	for _, v := range S {
		if v == s {
			return true
		}
	}
	return false
}

func GetLargest[T Number](n []T) (key int, result T) {
	for k, v := range n {
		if v > result {
			result = v
			key = k
		}
	}
	return
}

func Abs[T Number](i T) T {
	if i < 0 {
		return -i
	}
	return i
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	if n == 0 {
		return 0
	}
	if m == 1 {
		return n
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result

}

func MinMax[T Number](nums []T) (min, max T) {
	min, max = nums[0], 0
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return
}

func SortUniqNum[T Number](s []T) []T {
	slices.Sort(s)
	j := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i+1] {
			continue
		}
		s[j] = s[i]
		j++
	}
	return s[:j]
}

func SumArray[T Number](n []T) (res T) {
	for _, v := range n {
		res += v
	}
	return
}

func MultiplyArray[T Number](n []T) (res T) {
	res = 1
	for _, v := range n {
		res *= v
	}
	return
}

func Biggest[T Number](a, b T) T {
	return Ter(a < b, b, a)
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func DiffNum[T Number](a, b T) (res T) {
	res = b - a
	if res < 0 {
		return -res
	}
	return
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) (res int) {
	res = a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		res = LCM(res, integers[i])
	}
	return
}

func Ter[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func Select[T any](in []T, f func(i T) bool) (res []T) {
	res = make([]T, 0)
	for _, v := range in {
		if f(v) {
			res = append(res, v)
		}
	}
	return
}

func CopySlice[V any](s []V) (res []V) {
	res = make([]V, len(s))
	copy(res, s)
	return
}

func Combo[T any](iterable []T, r int) chan []T {
	ch := make(chan []T)

	go func() {
		l := len(iterable)
		for combo := range GenCombo(l, r) {
			res := make([]T, r)
			for i, v := range combo {
				res[i] = iterable[v]
			}
			ch <- res
		}
		close(ch)
	}()
	return ch
}

func GenCombo(n, r int) <-chan []int {
	if r > n {
		panic("invalid argument")
	}
	ch := make(chan []int)

	go func() {
		res := make([]int, r)
		for i := range res {
			res[i] = i
		}
		t := make([]int, r)
		copy(t, res)
		ch <- t
		for {
			for i := r - 1; i >= 0; i-- {
				if res[i] < i+n-r {
					res[i]++
					for j := 1; j < r-i; j++ {
						res[i+j] = res[i] + j
					}
					t := make([]int, r)
					copy(t, res)
					ch <- t
					break
				}
			}
			if res[0] >= n-r {
				break
			}
		}
		close(ch)
	}()
	return ch
}

func ChunkSlice[T any](input []T, size int) (res [][]T) {
	res = make([][]T, 0, len(input)/size)
	for i := 0; i < len(input); i += size {
		chunk := make([]T, 0, size)
		for j := 0; j < size; j++ {
			chunk = append(chunk, input[i+j])
		}
		res = append(res, chunk)
	}
	return
}
