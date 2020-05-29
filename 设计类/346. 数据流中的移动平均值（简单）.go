package 设计类

type MovingAverage struct {
	Size    int   //滑动窗口大小
	Element []int //存储元素
	Total   int   //当前元素和
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		Size:    size,
		Element: make([]int, 0),
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.Element) < this.Size {
		this.Element = append(this.Element, val)
		this.Total += val

	} else {
		this.Total -= this.Element[0]
		this.Total += val
		this.Element = this.Element[1:]
		this.Element = append(this.Element, val)
	}

	return float64(this.Total) / float64(len(this.Element))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
