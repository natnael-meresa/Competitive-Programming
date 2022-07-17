package main

func isPowerOfThree(n int) bool {
	if n <= 1 {
		return n == 1
	}
	if n%3 != 0 {
		return false
	}
	return isPowerOfThree(n / 3)
}

func main() {

	isPowerOfThree(9)

}
