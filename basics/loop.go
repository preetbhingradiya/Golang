package basics

func Loop() {

	num := 10
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			println(i, "is even")
		} else {
			println(i, "is odd")
		}
	}

	//Iterate over a collection(slice)
	numbers := []int{23, 45, 67, 89, 12, 34}
	for index, value := range numbers {
		if index%2 == 0 {
			continue
		}
		println("Index:", index, "Value:", value)
	}

	//while loop equivalent ( not any while and do_while in golang use for with condition only)
	count := 5
	for count > 0 {
		println("Count is:", count)
		count--
	}

}
