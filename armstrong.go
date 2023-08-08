package main

func main() {
	// count := 0 // <- initail count
	num := 1634
	temp := num
	sum := 0

	for num > 0 {
		// num /= 10
		// count = count + 1

		rem := num % 10
		// rem = count + 1
		sum = sum + (rem * rem * rem * rem)
		num = num / 10
	}

	if temp == sum {
		println("true")
	} else {
		println("false")
	}

	//println("Outside", count)
}
