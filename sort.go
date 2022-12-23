package sleepsortgo

import "time"

const defaultDuration = time.Millisecond * 100

func SleepSort(input []int, baseSleep time.Duration) []int {
	output := make([]int, len(input))
	receiver := make(chan int)
	for _, v := range input {
		go func(sleep int) {
			time.Sleep(time.Duration(sleep) * baseSleep)
			receiver <- sleep
		}(v)
	}
	for i := 0; i < len(input); i++ {
		output[i] = <-receiver
	}
	return output
}

func SleepSortDefault(input []int) []int {
	return SleepSort(input, defaultDuration)
}
