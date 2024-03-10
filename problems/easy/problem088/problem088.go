package problem088

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	if m < 0 || m > 200 || n < 0 || n > 200 || m+n < 1 || m+n > 200 {
		return nums1
	}

	if len(nums1) != m+n || len(nums2) != n {
		return nums1
	}

	index1, index2 := m-1, n-1

	for i := m + n - 1; i >= 0; i-- {
		if index2 < 0 {
			break
		}

		if index1 >= 0 && nums1[index1] > nums2[index2] {
			nums1[i] = nums1[index1]
			index1--
		} else {
			nums1[i] = nums2[index2]
			index2--
		}
	}

	return nums1
}

func mergeUsingSort(nums1 []int, m int, nums2 []int, n int) []int {
	if m < 0 || m > 200 || n < 0 || n > 200 || m+n < 1 || m+n > 200 {
		return nums1
	}

	if len(nums1) != m+n || len(nums2) != n {
		return nums1
	}

	for i := 0; i < len(nums2); i++ {
		nums1[m+i] = nums2[i]
	}

	sort.Ints(nums1)

	return nums1
}
