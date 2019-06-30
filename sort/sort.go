package sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
	"unsafe"
)

func quicksort(array []SS, begin, end int) {
	var i, j int
	if begin < end {
		i = begin + 1 // 将array[begin]作为基准数，因此从array[begin+1]开始与基准数比较！
		j = end       // array[end]是数组的最后一位

		for {
			if i >= j {
				break
			}
			if lt(&array[i] , &array[begin]) {
				array[i], array[j] = array[j], array[i]
				j = j - 1
			} else {
				i = i + 1
			}
		}

		if let(&array[i] , &array[begin]) { // 这里必须要取等“>=”，否则数组元素由相同的值时，会出现错误！
			i = i - 1
		}

		array[begin], array[i] = array[i], array[begin]
		quicksort(array, begin, i)
		quicksort(array, j, end)
	}
}

func main() {
	var s string = "husu849Vbzhh0000000000000000000000hhhhhhhhhhhhhhhFcHy9999999999999999999XxB7AIpfTNqBnTcOcb5AZzYhfJquK9qz61tGFphXtdwuARlM-DiPf2QvauTXw1ccdHTsHuO8AvkdNcl8OyIM4PZqsSbP0ouIJwgdmFScQ7qlPXeHRaPTPxSj7AiUSnXzh43XF1nEI4k_ywzd-U_XYaQeOxchI.cv1"

	s = "ahusu849Vbzhh0000000000000000000000hhhhhhhhhhhhhhhFcHy9999999999999999999XxB7AIpfTNqBnTcOcb5AZzYhfJquK9qz61tGFphXtdwuARlM-DiPf2QvauTXw1ccdHTsHuO8AvkdNcl8OyIM4PZqsSbP0ouIJwgdmFScQ7qlPXeHRaPTPxSj7AiUSnXzh43XF1nEI4k_ywzd-U_XYaQeOxchI.cv1husu849Vbzhh0000000000000000000000hhhhhhhhhhhhhhhFcHy9999999999999999999XxB7AIpfTNqBnTcOcb5AZzYhfJquK9qz61tGFphXtdwuARlM-DiPf2QvauTXw1ccdHTsHuO8AvkdNcl8OyIM4PZqsSbP0ouIJwgdmFScQ7qlPXeHRaPTPxSj7AiUSnXzh43XF1nEI4k_ywzd-U_XYaQeOxchI.cv1husu849Vbzhh0000000000000000000000hhhhhhhhhhhhhhhFcHy9999999999999999999XxB7AIpfTNqBnTcOcb5AZzYhfJquK9qz61tGFphXtdwuARlM-DiPf2QvauTXw1ccdHTsHuO8AvkdNcl8OyIM4PZqsSbP0ouIJwgdmFScQ7qlPXeHRaPTPxSj7AiUSnXzh43XF1nEI4k_ywzd-U_XYaQeOxchI.cv1husu849Vbzhh0000000000000000000000hhhhhhhhhhhhhhhFcHy9999999999999999999XxB7AIpfTNqBnTcOcb5AZzYhfJquK9qz61tGFphXtdwuARlM-DiPf2QvauTXw1ccdHTsHuO8AvkdNcl8OyIM4PZqsSbP0ouIJwgdmFScQ7qlPXeHRaPTPxSj7AiUSnXzh43XF1nEI4k_ywzd-U_XYaQeOxchI.cv1"


	var n int = 5e6
	z := make(map[int]string,n)
	for i:=0; i<n; i++ {
		z[i] = s
	}

	//fmt.Println(len(z))
	//fmt.Println(fmt.Sprintf("%d", int(10e1)))

	fmt.Println(reflect.TypeOf(z).Size())
	fmt.Println(unsafe.Sizeof(s))

	var ch = make(chan int,1)
	time.Sleep(time.Hour)
	<- ch

	n = 1000
	fmt.Println("quick")
	testQuick(n)
	fmt.Println("merge")
	testMerge(n)
}

func let(a,b *SS) bool {
	if a.B > b.B {
		return true
	}

	if a.B == b.B{
		return a.A >= b.A
	}

	return false
}

func lt(a,b *SS) bool {
	if a.B > b.B {
		return true
	}

	if a.B == b.B{
		return a.A > b.A
	}

	return false
}

type SS struct {
	A int
	B int
}

func testMerge(n int)  {
	var m []SS
	var i int
	var o = map[int]int{}

	for i=0;i<n;i++{
		o[i]=i
	}

	for _,i = range o{
		m = append(m,SS{A:i})
	}

	for _,i = range o{
		m = append(m,SS{A:i,B:rand.Intn(n)})
	}
	nums := m
	t := time.Now()

	MergeSort(nums, 0, len(nums)-1)

	l:= nums[:20]
	fmt.Println(l)
	fmt.Println(time.Since(t))

	//for _,s := range nums{
	//	fmt.Println(s)
	//}

}

func merge(arr []SS, l int, mid int, r int) {
	temp := make([]SS, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-l] = arr[i]
	}

	left := l
	right := mid + 1

	for i := l; i <= r; i++ {
		if left > mid {
			arr[i] = temp[right-l]
			right++
		} else if right > r {
			arr[i] = temp[left-l]
			left++
		} else if lt(&temp[left - l] , &temp[right - l]) {
			arr[i] = temp[right - l]
			right++
		} else {
			arr[i] = temp[left - l]
			left++
		}
	}
}

func MergeSort(arr []SS, l int, r int) {
	// 第二步优化，当数据规模足够小的时候，可以使用插入排序
	if r - l <= 15 {
		// 对 l,r 的数据执行插入排序
		for i := l + 1; i <= r; i++ {
			temp := arr[i]
			j := i
			for ; j > 0 && lt(  &arr[j-1],&temp); j-- {
				arr[j] = arr[j-1]
			}
			arr[j] = temp
		}
		return
	}

	mid := (r + l) / 2
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)

	// 第一步优化，左右两部分已排好序，只有当左边的最大值大于右边的最小值，才需要对这两部分进行merge操作
	if lt(&arr[mid] , &arr[mid + 1]) {
		merge(arr, l, mid, r)
	}
}

func testQuick(n int)  {
	var m []SS
	var o = map[int]int{}

	for i:=0;i<n;i++{
		o[i]=i
	}

	for _,i := range o{
		m = append(m,SS{A:i})
	}

	for _,i := range o{
		m = append(m,SS{A:i,B:rand.Intn(n)})
	}
	nums := m
	t := time.Now()

	quicksort(nums, 0, len(nums)-1)
	l:= nums[:20]

	var idx int
	var node SS
	var ids = make([]int, len(l))
	for idx,node = range l{
		ids[idx] = node.A
	}

	fmt.Println(time.Since(t))
	fmt.Println(ids)

	//for _,s := range nums{
	//	fmt.Println(s)
	//}
}