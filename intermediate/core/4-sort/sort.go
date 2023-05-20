/*
$ go doc sort
package sort // import "sort"
func IntsAreSorted(x []int) bool
func IsSorted(data Interface) bool
func Search(n int, f func(int) bool) int
func SearchFloat64s(a []float64, x float64) int
func SearchInts(a []int, x int) int
func SearchStrings(a []string, x string) int
func Slice(x interface{}, less func(i, j int) bool)
func SliceIsSorted(x interface{}, less func(i, j int) bool) bool
func SliceStable(x interface{}, less func(i, j int) bool)
func Sort(data Interface)
func Stable(data Interface)
func Strings(x []string)
func StringsAreSorted(x []string) bool
type Float64Slice []float64
type IntSlice []int
type Interface interface{ ... }
    func Reverse(data Interface) Interface
type StringSlice []string
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{8, 7, 34, 9, 5}
	sort.Ints(numbers) // sorting numbers
	fmt.Println(numbers)

	texts := []string{"vijay", "aruna", "kumar", "bala", "vikram"}
	sort.Strings(texts) // sorting texts
	fmt.Println(texts)
}
