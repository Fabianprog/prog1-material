package listprops

import "fmt"

func ExampleContains() {
	l := []string{"Hallo", "Welt"}
	fmt.Println(Contains(l, "Welt"))
	fmt.Println(Contains(l, "Haus"))

	// Output:
	// true
	// false
}
func ExampleContainsOnly() {
	l := []string{"Hallo", "Welt"}
	z := []string{"Hallo", "Hallo"}
	fmt.Println(ContainsOnly(l, "Welt"))
	fmt.Println(ContainsOnly(l, "Haus"))
	fmt.Println(ContainsOnly(z, "Hallo"))

	// Output:
	// false
	// false
	// true
}
func ExampleContainsN() {
	a := []string{"Hallo", "Hallo", "Hallo", "Tschau"}
	b := []string{"Hallo", "Hallo", "Tschau", "Tschau"}
	c := []string{"Hallo", "Hallo", "Hallo", "Hallo", "Tschau"}
	fmt.Println(ContainsN(a, "Hallo", 3))
	fmt.Println(ContainsN(b, "Hallo", 3))
	fmt.Println(ContainsN(c, "Hallo", 3))
	// Output:
	// true
	// false
	// true
}
func ExampleContainsNRow() {
	a := []string{"Hallo", "Hallo", "Hallo", "Tschau"}
	b := []string{"Hallo", "Hallo", "Tschau", "Tschau"}
	c := []string{"Hallo", "Hallo", "Hallo", "Hallo", "Tschau"}
	fmt.Println(ContainsNRow(a, "Hallo", 3))
	fmt.Println(ContainsNRow(b, "Tschau", 3))
	fmt.Println(ContainsNRow(c, "Hallo", 4))

	// Output:
	// true
	// false
	// true
}
