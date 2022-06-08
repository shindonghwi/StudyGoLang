package main

func main() {

	for i := 0; i < 5; i++ {
		println("data: ", i)
	}

	var payType string = "card"

	if payType == "card" {
		println("this is card")
	} else if payType == "cash" {
		println("this is cash")
	} else {
		println("this is else")
	}

}
