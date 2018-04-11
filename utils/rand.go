package utils

import (
	"math/rand"
	"time"
)

//UniqRand 创建唯一随机数数组，返回一个数组序列，没有重复数
func UniqRand(l int, n int) []int {

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	set := make(map[int]struct{})

	nums := make([]int, 0, l)
	for {
		num := rnd.Intn(n)
		if _, ok := set[num]; !ok {
			set[num] = struct{}{}
			nums = append(nums, num)
		}
		if len(nums) == l {
			return nums
		}
	}

}

//Rand 创建随机数数组，返回随机数序列，可能会有重复
func Rand(l int, n int) []int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	nums := make([]int, 0, l)
	for {
		num := rnd.Intn(n)
		nums = append(nums, num)
		if len(nums) == l {
			return nums
		}
	}

}
