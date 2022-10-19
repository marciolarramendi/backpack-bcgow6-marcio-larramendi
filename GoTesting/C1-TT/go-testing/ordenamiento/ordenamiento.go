package ordenamiento

import (
	"sort"
)

/*func main() {
	a := []int{5, 3, 4, 7, 8, 9}
	//var b []int
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	for _, v := range a {
		b = append(b, v)
		//fmt.Println(v)
	}
	fmt.Println(a)
}*/

func OrdenarSliceAsc(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return a
}
