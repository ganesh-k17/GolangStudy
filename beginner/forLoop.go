package main

import "fmt"

func main() {

	/* for traditional loop */
	count := 5
	for i := 1; i < count; i++ {
		fmt.Println("hello")
	}

	fmt.Println("")

	i := 1
	for i < count {
		fmt.Println("hello1")
		i++
	}

	fmt.Println("")

	for i := 1; ; i++ {
		fmt.Println("hello2")
		if i == 5 {
			break
		}
	}

	fmt.Println("")

	/* for range */
	for _, v := range []string{"Raja", "Vikram"} {
		fmt.Println(v)
	}

	fmt.Println("")

	/* for range printing index in first parameter*/
	for i, v := range []string{"Raja", "Vikram"} {
		fmt.Printf("%d %s", i, v)
		fmt.Println("")
	}

	fmt.Println("")

	/* Infinite loop / alternative for while */
	k := 1
	for {
		fmt.Println("hello")
		if k == 10 {
			break
		}
		k++
	}

	fmt.Println("")

	/* loop range over string */

	for range "hello" {
		fmt.Println("hello range")
	}

	fmt.Println("")
}
