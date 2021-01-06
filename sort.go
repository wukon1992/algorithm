package main

//排序算法太多了，有很多可能你连名字都没听说过，比如猴子排序、睡眠排序、面条排序等。我只讲众多排序算法中的一小撮，也是最经典的、最常用的
//：冒泡 排序、插入排序、选择排序、归并排序、快速排序、计数排序、基数排序、桶排序
// 冒泡，插入，选择： O(n^2)
// 快排，并归：O(nlogn)
// 桶计数，基数：O(n)

import(
	"fmt"
)
var p = fmt.Println
var arr = []int{2,3,1,89,3,44,23,47,99,67}
var arr1 = []int{2,3,1,89,3,44,23,47,99,67}
var count = len(arr)

//冒泡排序只会操作相邻的两个数据。每次冒泡操作都会对相邻的两个元素进行比较，看是否满足大小关系要求。如果不满足就让它俩互换。一次冒泡会让至少一
//个元素移动到它应该在的位置，重复n次，就完成了n个数据的排序工作
func BubbleSort(arr []int) []int {
	if(count <= 1 || arr == nil) {
		return arr
	}

	for i := 0; i < count - 1; i++{
		for j := 0; j < count - i - 1; j++{
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		} 
	}
	return arr
}

//插入排序
//首先，我们将数组中的数据分为两个区间，已排序区间和未排序区间。初始已排序区间只有一个元素，就是数组的第一个元素。插入算法的核心思想是取未排序
//区间中的元素，在已排序区间中找到合适的插入位置将其插入，并保证已排序区间数据一直有序
func InsertSort(arr []int) []int{
	count = len(arr)
	if count == 1{
		return arr
	}
	for i := 1; i < count; i++ {
		j := i-1;
		v := arr[i]
		for  j >= 0 && arr[j] < v {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = v
	}
	return arr
}

//选择排序
//选择排序空间复杂度为O(1)，是一种原地排序算法
//选择排序每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾

func SelectSort(arr []int)[]int {
	for i := 0; i < count - 1; i++ {
		min := i //最小值的坐标从最左边开始依次向右

		for j := count - 1; j > i; j-- { //从右边查找最小值的坐标。然后与i替换
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}	
	return arr
}

//快速排序
//快速排序是C.R.A.Hoare于1962年提出的一种划分交换排序。它采用了一种分治的策略，通常称其为分治法(Divide-and-ConquerMethod)。
//1．先从数列中取出一个数作为基准数。 
//2．分区过程，将比这个数大的数全放到它的右边，小于或等于它的数全放到它的左边。 
//3．再对左右区间重复第二步，直到各区间只有一个数。
func QuickSort(arr []int) []int{
	count = len(arr)
	if count <= 1 {
		return arr
	}
	left := make([]int, 0, count)
	right := make([]int, 0, count)
	middle := make([]int, 0, count)
	mid := arr[0]
	for _, v := range arr {
		tmp := v //因为是递归分治， range 中的V是地址引用，如果直接用v。则只能取到v指向地址的最后一个值
		if mid > tmp {
			left = append(left,tmp)
		} else if mid == tmp {
			middle = append(middle, mid)
		} else if mid < tmp {
			right = append(right, tmp)
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)
	data := append(left,middle...)
	data = append(data,right...)
	return data
}

//归并排序
//归并排序是建立在归并操作上的一种有效的排序算法。该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。归并排序是一种稳定的排序方法。将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为2-路归并
func MergeSort(arr []int) []int{
	count = len(arr)
	if count <= 1 {
		return arr
	}
	mid := count/2

	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	result := merge(left,right)
	return result
}

func merge(left,right []int) []int{
	result := make([]int, 0)
	m,n := 0, 0
	l, r := len(left), len(right)
	for m < l && n < r {
		if left[m] < right[n] {
			result = append(result, left[m])
			m++
		}else{
			result = append(result, right[n])
			n++
		}
	}
	result = append(result, left[m:]...)
	result = append(result, right[n:]...)
	return result
}

//计数排序
//计数排序的核心在于将输入的数据值转化为键存储在额外开辟的数组空间中
//讲序列中的值作为健，值的次数作为值放入数组中。
//找去序列中的最大值，和最小值。从最小值遍历到最大值。依次将原序列的值放入排序数组中
func CountSort(arr []int) []int{

	var getMaxAndMinVal = func(arr []int)(max,min int, res map[int]int) {
		max,min = arr[0], arr[0]
		res = make(map[int]int,0)
		for _,v := range arr {
			res[v]++
			if v > max {
				max = v
			}
			if v < min {
				min = v
			}
		}
		return 
	}
	index := 0
	result := make([]int, len(arr))
	max, min, res := getMaxAndMinVal(arr)
	for i := min; i <= max; i++ {
		for res[i] > 0 {
			result[index] = i
			index++
			res[i]--
		}
	}
	return result
}
//桶排序
// 桶排序的基本思路是将数据根据计算规则来分组，并将数据依次分配到对应分组中。
// 分配时可能出现某分组里有多个数据，那么再进行分组内排序。
// 然后得到了有序分组，最后把它们依次取出来放到数组中即实现了整体排序。
func BucketSort(arr []int) []int{

	 var getMax = func(arr []int)(max int){
	 	max = arr[0]
	 	for _, v := range arr {
	 		if max < v {
	 			max = v
	 		}
	 	}
	 	return 
	 }
	 //桶数
	 count = len(arr)
	 //二维切片
	 buckets := make([][]int, count)
	 //k（数组最大值）
	 max := getMax(arr)
	 for i := 0 ; i < count; i++ {
	 	index := arr[i] * (count-1) /max//分配桶index = value * (n-1) /k
	 	buckets[index] = append(buckets[index], arr[i])
	 }
	 //桶内排序
	tmpPos := 0
	for i := 0; i < count; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0{
			InsertSort(buckets[i])
			copy(arr[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
	return arr
}

//堆排序
//先对随机数组堆化
//构造堆
//如果给两个已构造好的堆添加一个共同父节点，
//将新添加的节点作一次下沉将构造一个新堆，
//由于叶子节点都可看作一个构造好的堆，所以
//可以从最后一个非叶子节点开始下沉，直至
//根节点，最后一个非叶子节点是最后一个叶子
//节点的父节点，角标为N/2
func HeapSort(arr []int) {
	N := len(arr)-1
	Sink(arr)
	//下沉排序
    //将堆顶元素放到数组最后面然后末尾放到堆顶堆化长度减一，逐一堆化
    for N > 1 {
        swap(arr, 1, N) //将大的放在数组后面，升序排序
        N--
        Sink(arr[:N+1]) //重新堆化
    }
	p(arr[1:])
}


//堆化 自下而上的堆化 /叶子节点k,与父节点k/2,以及k+1节点比较取大的值来替换，当k无替换则交换结束,将元素构建堆化
func Sink(arr []int){
	for i := 0; i < len(arr); i++{
		if i <= 0 {
			continue
		}	
		j := i
		for j/2>0 && arr[j/2] < arr[j] {
			swap(arr, j/2, j)
			j = j/2
		}
	}
}
func swap(arr []int, k, i int) {
	arr[k], arr[i] = arr[i], arr[k]
}

func HeapSorts(s []int) {
    N := len(s) - 1 //s[0]不用，实际元素数量和最后一个元素的角标都为N
    //构造堆
    //如果给两个已构造好的堆添加一个共同父节点，
    //将新添加的节点作一次下沉将构造一个新堆，
    //由于叶子节点都可看作一个构造好的堆，所以
    //可以从最后一个非叶子节点开始下沉，直至
    //根节点，最后一个非叶子节点是最后一个叶子
    //节点的父节点，角标为N/2
    for k := N / 2; k >= 1; k-- {
        sink(s, k, N)
    }
    p(s)
    //下沉排序
    for N > 1 {
        swap(s, 1, N) //将大的放在数组后面，升序排序
        N--
        sink(s, 1, N)
    }
    p(s)
}
//自上而下堆化
func sink(s []int, k, N int) {
    for {
        i := 2 * k
        if i > N { //保证该节点是非叶子节点
            break
        }
        if i < N && s[i+1] > s[i] { //选择较大的子节点
            i++
        }
        if s[k] >= s[i] { //没下沉到底就构造好堆了
            break
        }
        swap(s, k, i)
        k = i
    }
}

func main() {
	// p(BubbleSort(arr))
	// p(InsertSort(arr))
	// p(SelectSort(arr))
	// p(QuickSort(arr))
	// p(MergeSort(arr))
	// CountSort(arr)
	// p(CountSort(arr))
	// p(BucketSort(arr))
	// HeapSort(arr)
	HeapSorts(arr1)
}