package main

import "fmt"

func main() {
	fmt.Println("Введите значение двух слагаемых для расчета суммы:")

	var item1, item2 int

	fmt.Scan(&item1)
	fmt.Scan(&item2)
	sum := item1
	for i := 0; i < item2; i++ {
		sum++
	}
	fmt.Println(sum)
}
