package number_of_students_doing_homework_at_a_given_time

func BusyStudent(startTime, endTime []int, queryTime int) int {
	// 简单枚举
	// const timeMax = 1000
	// timeline := make([]int, timeMax)
	// for i := range startTime {
	// 	for j := startTime[i] - 1; j < endTime[i]; j++ {
	// 		timeline[j]++
	// 	}
	// }
	// return timeline[queryTime-1]

	// 区间数组
	const timeMax = 1000
	timeline := make([]int, timeMax+1)
	for i := range startTime {
		timeline[startTime[i]-1]++
		timeline[endTime[i]]--
	}
	result := 0
	for i := 0; i < queryTime; i++ {
		result += timeline[i]
	}
	return result
}
