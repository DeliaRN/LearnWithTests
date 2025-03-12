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

func SumAll(numbersToSum ...[]int) []int {

	//Two ways of writing it:
	/*FOR LOOP WITH A RANGE
	lengthOfArray := len(numbersToSum)
	sums := make([]int, lengthOfArray) //Create a slice of arrays of ints

		for i, numbers := range numbersToSum { //For each array...
			sums[i] = Sum(numbers) //Use of Sum function that takes an array
		}
	*/

	//USING APPEND FUNCTION
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	//This starts with an empty slice sums and append to it the result
	//of calling the Sum function for every vararg

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:] // from the first
		sums = append(sums, Sum(tail))
	}
	return sums
}
