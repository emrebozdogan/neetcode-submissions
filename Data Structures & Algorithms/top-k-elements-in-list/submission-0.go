func topKFrequent(nums []int, k int) []int {
	frequentMap := make(map[int]int)

	for _, v := range nums {
		frequentMap[v]++
	}

	buckets := make([][]int, len(nums)+1)

	for num, freq := range frequentMap {
		buckets[freq] = append(buckets[freq], num)
	}

	tops := make([]int, 0, k)

	for i := len(buckets)-1; i>=1; i-- {
		for _, num := range buckets[i] {
			tops = append(tops, num)

			if len(tops) == k {
				return tops
			}
		}
	}

	return tops
}
