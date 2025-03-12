package ArraysAndSlices

func Sum(numbers []int) int {
	sum := 0

	//Two ways of writing it

	/*SIMPLE FOR LOOP:
	for int i := 0; i<5; i++ {
		sum += numbers[i]
	}
	*/

	//FOR LOOP USING RANGE
	for _, number := range numbers {
		sum += number
	}
	return sum
}
