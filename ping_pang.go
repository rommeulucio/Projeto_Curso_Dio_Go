package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PingPang") //SE FOR MULTIPLO DE 3 E 5 FALA PINGPANG
		} else if i%3 == 0 {
			fmt.Println("Ping") //SE FOR MULTIPLO DE 3 PING
		} else if i%5 == 0 {
			fmt.Println("Pang") //SE FOR MULTIPLO DE 5 FALA PANG
		} else {
			fmt.Println(i)

		}
	}
}
