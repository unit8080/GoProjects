// 346. Moving Average from Data Stream
// https://leetcode.com/problems/moving-average-from-data-stream/

type MovingAverage struct {
	size  int // maximum last size to average
	sum   int // current sum
	count int // running count so far
	nums  []int // ring buffer to hold values
}

func Constructor(size int) MovingAverage {
	var avg MovingAverage

	avg.size = size
	avg.nums = make([]int, size)
	return avg
}

func (this *MovingAverage) Next(val int) float64 {
	index := this.count % this.size

	this.sum += val - this.nums[index]
	this.nums[index] = val
	this.count += 1    // crucial to take min as below
	avg := float64(this.sum) / float64(min(this.size, this.count))
	return float64(avg)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
