package iteration

const repeatCount = 5

// Repeat exported for testing
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

// Some for examples

// import "fmt"

// func main() {
//	i := 1
//	for i <= 3 {
//		fmt.Println(i)
//		i = i + 1
//	}

//	for j := 7; j <= 9; j++ {
//		fmt.Println(j)
//	}

//	for {
//		fmt.Println("loop")
//		break
//	}

//	for n := 0; n <= 5; n++
//		if n%2 == 0 {
//			continue
//		}
//		fmt.Println(n)
//}
