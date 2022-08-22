package silde_window

import (
	"math"
	"time"
)

type Window struct {
	Size         int64   //窗口大小
	Limit        int64   //窗口内限制的大小
	SplitNum     int64   //切分小窗口的数目大小
	Counters     []int64 //当每隔小窗口的计数数组
	currentIndex int64   //当前窗口时间
	startTime    int64   //窗口开始时间
}

func NewWindow(size, limit, split int64) *Window {
	return &Window{
		Size:      size,
		Limit:     limit,
		SplitNum:  split,
		startTime: time.Now().UnixNano(),
	}
}

func (w *Window) TryAcquire() bool {

	curTime := time.Now().UnixNano()

	windowsNum := max(curTime-w.Size-w.startTime, 0) / (w.Size / w.SplitNum) //计算滑动小窗口的数量

	w.slideWindow(windowsNum) //滑动窗口

	var count int64
	var i int64
	for ; i < w.SplitNum; i++ {
		count += w.Counters[i]
	}
	if count >= w.Limit {
		return false
	} else {
		w.Counters[w.currentIndex]++
		return true
	}
}

func (w *Window) slideWindow(windowsNum int64) {

	if windowsNum == 0 {
		return
	}

	slideNum := min(windowsNum, w.SplitNum)

	var i int64

	for ; i < slideNum; i++ {
		w.currentIndex = (w.currentIndex + 1) % w.SplitNum
		w.Counters[w.currentIndex] = 0
	}

	w.startTime = w.startTime + windowsNum*(w.Size/w.SplitNum) //更新滑动窗口时间
}

func max(nums ...int64) int64 {
	var maxNum int64 = math.MinInt64
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func min(nums ...int64) int64 {
	var maxNum int64 = math.MaxInt64
	for _, num := range nums {
		if num < maxNum {
			maxNum = num
		}
	}
	return maxNum
}
