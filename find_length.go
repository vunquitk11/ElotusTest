package main

func findLength(nums1 []int, nums2 []int) int {
	var result int
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				count := 1
				for (i+count < len(nums1)) && (j+count < len(nums2)) {
					if nums1[i+count] == nums2[j+count] {
						count++
					}
				}
				if count > result {
					result = count
				}
			}
		}
	}

	return result
}
