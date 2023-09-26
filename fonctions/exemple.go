package main

func main() {
	MyFunction()
	variable := MyFunction2(10)
	print(variable)
}

func MyFunction() {
	println("Trop fort")
}

func MyFunction2(nombre int) bool {
	if nombre > 5 {
		return true
	} else {
		return false
	}
}
