package main

import (
	"fmt"
//	"slices"
)

//slice ->dynamic and most used construct in go
//A slice in Go is a dynamically-sized,
// flexible view into the elements of an array. Unlike arrays, slices can grow and shrink in size, making them more versatile.

// + useful methods
func main(){
	/*
	//uninitialized slice in nil
	var nums []int
	fmt.Println(nums) //output: []
	fmt.Println(nums == nil) //output: true
	fmt.Println(len(nums)) //output : 0
	*/

/*
	var nums []int
	var nums = make([]int, 2)
	fmt.Println(nums) //output: [0 0]
	fmt.Println(nums == nil) //output: false
*/

/*
	var nums = make([]int, 2)
	//capacity -> maximum numbers of elements can fit
	fmt.Println(cap(nums)) //output: 2
*/

/*
	var nums = make([]int, 2, 5)
	// Initial capacity -> maximum numbers of elements can fit
	nums = append(nums, 1)
	fmt.Println(nums)
	fmt.Println(cap(nums)) 
output: [0 0 1]
5
*/

/*
//increase the capacity automatically, it double the cabacity 
	var nums = make([]int, 0, 5)
	// Initial capacity -> maximum numbers of elements can fit
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)

	fmt.Println(nums)
	fmt.Println(cap(nums)) 
/*output: [1 2 3 4 5]
10
*/

/*
	
	nums := []int{}
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)

	fmt.Println(nums) //output: [1 2 3 4 5]
	fmt.Println(cap(nums)) //output: 8
	fmt.Println(len(nums)) //output: 5
*/
	
/* // without enter element
	nums := []int{}

	fmt.Println(nums) //output: []
	fmt.Println(cap(nums)) //output: 0
	fmt.Println(len(nums)) //output: 0
*/

/*
//append element in array
	nums := []int{}
	nums = append(nums, 1)
	nums = append(nums, 2)
	fmt.Println(nums) //output: [1 2]
	fmt.Println(cap(nums)) //output: 2
	fmt.Println(len(nums)) //output: 2

*/	

/*
// slice change the element
	
	var nums = make([]int, 2, 5) // 0 to 2 change because increase the lenthg
	//nums = append(nums, 1)
	//nums = append(nums, 2)
	nums[0]=3
	nums[1]=4
	fmt.Println(nums) //output: [3 4]
	fmt.Println(cap(nums)) //output: 5
	fmt.Println(len(nums)) //output: 2
*/

/*
	//copy function
	var nums = make([]int, 0, 5)
	nums = append(nums, 2)
	var nums2 = make([]int, len(nums))
	
	copy(nums2, nums)
	fmt.Println(nums, nums2) //output: [2] [2]
*/

/*
	//slice operator
	var nums = []int{1, 2, 3}
	fmt.Println(nums[0:2]) // output: [1 2]
	fmt.Println(nums[:1]) // output: [1]
	fmt.Println(nums[:2]) // output: [1 2]
	fmt.Println(nums[1:]) // output: [2 3] // exclude the element after mention 1:
*/

/*
//slice package
	var nums1 = []int{1, 2}
	var nums2 = []int{1,2}
	fmt.Println(slices.Equal(nums1, nums2)) //output: true
*/

//making 2d slices
	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums) //output: [[1 2 3] [4 5 6]]






}
