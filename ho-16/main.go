package main

import "fmt"

var planets int = 12

func main() {

	if planets == 9 {
		fmt.Printf("Looks like you are counting Pluto. In that case, there are %d planets in the solar system\n", planets)
	} else if planets >= 9 {
		fmt.Printf("We've discovered new planets. There are now %d planets in the Solar System. The empire is expanding!\n", planets)
	} else if planets == 8 {
		fmt.Printf("So by saying there are %d planets in the Solar System, you do not count Pluto\n", planets)
	} else {
		fmt.Printf("OH NO!! WE HAVE LOST MORE PLANETS!! THERE ARE NOW ONLY %d PLANETS LEFT IN THE SOLAR SYSTEM! THIS IS AN EMERGENCY!! CALL THE PRESIDENT!!\n", planets)
	}

}

// I created an exe with <go build main.go>
// I ran the exe with <./main>
