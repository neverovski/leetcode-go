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

	i, j := m-1, n-1
	k := m + n - 1

	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}

	for ; j >= 0; k-- {
		nums1[k] = nums2[j]
		j--
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
