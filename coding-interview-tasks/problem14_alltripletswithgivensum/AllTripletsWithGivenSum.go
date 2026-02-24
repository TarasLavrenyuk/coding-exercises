package problem14_alltripletswithgivensum

func AllTripletsWithGivenSum(nums []int, target int) [][]int {
	result := make([][]int, 0)

	numsSet := make(map[int]int) // number to index
	for index, num := range nums {
		numsSet[num] = index
	}

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			difference := target - nums[i] - nums[j]
			if index, ok := numsSet[difference]; ok && index > j {
				result = append(result, []int{nums[i], nums[j], difference})
			}
		}
	}

	return result
}
