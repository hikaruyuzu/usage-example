package utility

func Avarage(value [8]int) float64 {
	temp := 0
	for _, v := range value {
		temp += v
	}
	length := len(value)
	result := float64(temp) / float64(length)
	return result
}

func AddHightScore(value [8]int) []int {
	var result []int

	for _, val := range value {
		if val >= 90 {
			result = append(result, val)
		}
	}
	return result
}
