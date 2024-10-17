package leetcode

func maximumSwap(num int) int {
	var data []int
	for num > 0 {
		r := num % 10
		num /= 10
		data = append(data, r)
	}

	var ans int = 0
	n := len(data) - 1
	var idx int = -1
	var idx2 int = -1
	var val int = -1
	for i := n - 1; i >= 0 && idx == -1; i-- {
		if data[i] > data[i+1] {
			idx = i + 1
			idx2 = i
			val = data[i]
		}
	}

	for i := idx - 1; i >= 0 && idx >= 0; i-- {
		if data[i] >= val {
			idx2 = i
			val = data[i]
		}
	}

	if idx >= 0 {
		for i := idx; i <= n; i++ {
			if data[i] < val {
				idx = i
			}
		}
		data[idx2], data[idx] = data[idx], data[idx2]
	}

	for i := n; i >= 0; i-- {
		ans *= 10
		ans += data[i]
	}

	return ans
}
