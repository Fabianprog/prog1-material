package duration

import "fmt"

func ExampleDuration() {

	fmt.Println(FromSekunden(1).Sekunden())
	fmt.Println(FromMinuten(1).Minuten())
	fmt.Println(FromStunden(1).Sekunden())
	fmt.Println(FromTage(1).Stunden())
	fmt.Println(FromWochen(1).Tage())
	fmt.Println(FromJahre(1).Tage())

	//Output:

}
