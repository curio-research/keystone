package server

func HashNumbers(num1, num2 int) int {
	// a prime number used in the hashing process
	prime := 31

	// combine the two numbers using bitwise operations
	hashValue := num1*prime + num2

	return hashValue
}
