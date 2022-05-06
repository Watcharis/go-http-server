package gopointer

func ShiftLeftLoop(n int, data []int) []int {
	i := 1
	for i <= n {
		ShiftLeftArray(data)
		i += 1
	}
	return data
}

func ShiftLeftArray(data []int) []int {
	temp := data[0]
	n := len(data)
	for i := 0; i < len(data)-1; i++ {
		data[i] = data[i+1]
	}
	data[n-1] = temp
	return data
}
