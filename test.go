package main

func main() {

	var password string = "ac"

	switch password {
	case "Ac":
		{
			print("success")
		}
	case "ac":
		{
			print("fail")
		}
	}

	print(add(1, 2))

	println(addAll(11, 2, 3, 4, 5))

}

func add(a int, b int) int {
	return 1
}

func addAll(li ...int) int {
	result := 0
	for idx, n := range li {
		println("hi: ", idx, n)
		result += n
	}

	return result
}
