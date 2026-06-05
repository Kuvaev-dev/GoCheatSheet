package main

import (
	"cmp"
	"fmt"
)

func main() {
	// 1. Hello, World
	fmt.Println("Hello, World")

	// 2. Calc
	var a int = 10
	var b int = 10
	fmt.Println(a + b)

	// 3. Func Def
	c := 20
	fmt.Println(c)

	// 4. Enum
	const (
		One = iota
		Two
		Three
	)
	fmt.Println(One)

	// 5. Float
	var d float64 = 2.2
	fmt.Println(d)

	// 6. Rune - int for Unicode
	var e rune = 1
	fmt.Println(e)

	// 7. Array

	// Static
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr[0])

	// Dynamic
	darr := []int{}
	darr = append(darr, 1)
	darr = append(darr, 2)
	darr = append(darr, 3)
	darr = append(darr, 4)
	darr = append(darr, 5)
	// From first to pre-last element (5 is not included)
	darr2 := darr[0:5]
	fmt.Println(darr2)
	fmt.Println(len(darr))
	fmt.Println(cap(darr))

	// 8. Map
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	delete(m, "a")
	fmt.Println(m)

	// 9. Struct
	type Food struct {
		Calories int
		Weight   int
	}
	p := Food{Calories: 1000, Weight: 100}
	fmt.Println(p.Calories, p.Weight)

	// 10. Short Def Inside If
	if n := Sum(10, 10); n > 1 {
		fmt.Println("OK")
	}

	// 11. Loops

	// 11.1. CLassic
	for i := 0; i < 5; i++ {
		fmt.Println("OK")
	}

	// 11.2. While-like
	q := 5
	for q < 5 {
		q--
		fmt.Println("OK")
	}

	// 11.3. Infinity
	//for { fmt.Println("OK") }

	// 12. range
	for i, v := range []string{"a", "b"} {
		fmt.Println(i, v)
	}
	for i, v := range m {
		fmt.Println(i, v)
	}

	// 13.1. Value Return
	fmt.Println(Sum(1, 2))
	// 13.2. Several Values Return
	fmt.Println(SumWithMsg(1, 2))
	// 13.3. Named return values + naked return
	fmt.Println(Calc(10))
	// 13.4. Variadic-function
	fmt.Println(ArrSum(1, 2, 3))

	// 14. defer - executes in the end of the function
	//defer fmt.Println("end of func")

	// 15. panic - stops execution
	//panic("critical error")

	// 16. Pointers
	x := 10
	y := &x         // Var Address
	fmt.Println(*y) // 10
	*y = 20         // Changes Original x

	// 17
	person := Human{
		Name: "Vasya",
		Age:  10,
	}
	// 17.1. Value receiver
	fmt.Println(person.GetNameData())
	fmt.Println(person.GetFullData())
	// 17.2. Pointer receiver
	person.ChangeAge(11)
	fmt.Println(person.GetFullData())

	// 18. In-building - instead inheritance
	house := House{
		Human:  person,
		Floors: 10,
	}
	fmt.Println(house.Human.Name, house.Human.Age, house.Floors)

	// 19. Interfaces
	PrintStatus(house.Human)

	// 20. Type switch
	// switch v := person.(type) {
	// case Human:
	// 	fmt.Println("Human")
	// case House:
	// 	fmt.Println("House")
	// default:
	// 	fmt.Println("Undefined")
	// }

	// 21. Errors
	// return fmt.Errorf("Error opening file: %w", err)
	// if errors.Is(err, os.ErrNotExist) { ... }

	// 22. Gorutines - easy threads
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	val := <-ch
	fmt.Println(val)

	ch1 := make(chan int)
	ch2 := make(chan int)
	select {
	case tempv := <-ch1:
		fmt.Println(tempv)
	case ch2 <- val:
		fmt.Println(val)
	default: // non-blocking
	}
}

// 13. Functions

// 13.1. Value Return
func Sum(a int, b int) int {
	return a + b
}

// 13.2. Several Values Return
func SumWithMsg(a int, b int) (int, string) {
	if (a + b) > 10 {
		return a + b, "Geater"
	}
	return a + b, "Less"
}

// 13.3. Named return values + naked return
func Calc(x int) (sum, prod int) {
	sum = x + 1
	prod = x * 2
	return
}

// 13.4. Variadic-functions
func ArrSum(nums ...int) int {
	summary := 0
	for i := 0; i < len(nums); i++ {
		summary += nums[i]
	}
	return summary
}

// 13.5. Сlosure
func MakeAdder(x int) func(int) int {
	return func(y int) int { return x + y }
}

// 17. Struct
type Human struct {
	Name string
	Age  int
}

// 17.1. Value receiver
func (r Human) GetNameData() string {
	return r.Name
}
func (r Human) GetFullData() (string, int) {
	return r.Name, r.Age
}

// 17.2. Pointer receiver
func (r *Human) ChangeAge(newAge int) {
	r.Age = newAge
}

// 18. In-building - instead inheritance
type House struct {
	Human
	Floors int
}

// 19. Interfaces

// 19.1. Declaration
type Entity interface {
	IsAlive() bool
}

// 19.2. Implementation of method for Human
func (h Human) IsAlive() bool {
	return h.Age >= 0
}

// 19.3. Using
func PrintStatus(e Entity) {
	if e.IsAlive() {
		fmt.Println("Human is alive")
	}
}

// 23. Generics
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
