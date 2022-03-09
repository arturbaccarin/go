package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	//print()
	//mars()
	//_fmt()
	//lightspeed()
	//_rand()
	//_contains()
	//elif()
	//countdown()
	//infinity()
	//loop_scope()
	//short_loop()
	//short_if()
	//pi()
	//third()
	//raw()
	//raw_lines()
	//casting()

	//kelvin := 294.0
	//celsius := kelvinToCelsius(kelvin)
	//fmt.Print(kelvin, "º K is ", celsius, "º C")

	//celsius()

	//array_()
	//array_loop()
	//array_range()
	//slicing_array()
	//append_array()

}

func print() {
	fmt.Println("Hello world")
}

func mars() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(149.0 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" years old.")
}

func _fmt() {
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 149.0*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 41*365/687)
}

func lightspeed() {
	const lightSpeed = 299792 // km/s
	var distance = 56000000   // km

	//var (
	//	distance = 56000000
	//	speed = 100800
	//)

	fmt.Println(distance/lightSpeed, "seconds")
	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")
}

func _rand() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)
	num = rand.Intn(10) + 1
	fmt.Println(num)
}

func _contains() {
	fmt.Println("You find yourself in a dimly lit cavern.")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit)
}

func compare() {
	fmt.Println("There is a sign near the entrance that reads 'No Minors'.")
	var age = 41
	var minor = age < 18
	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)
}

func _if() {
	var command = "go east"
	if command == "go east" {
		fmt.Println("You head further up the mountain.")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of your life.")
	} else {
		fmt.Println("Didn't quite get that.")
	}
}

func elif() {
	var room = "cave"
	if room == "cave" {
		fmt.Println("You find yourself in a dimly lit cavern.")
	} else if room == "entrance" {
		fmt.Println("There is a cavern entrance here and a path to the east.")
	} else if room == "mountain" {
		fmt.Println("There is a cliff here. A path leads west down the mountain.")
	} else {
		fmt.Println("Everything is white.")
	}
}

func switch_() {
	var room = "lake"

	switch {
	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case room == "underwater":
		fmt.Println("The water is freezing cold.")
	}
}

func countdown() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Liftoff")
}

func infinity() {
	var degrees = 0
	for {
		fmt.Println(degrees)
		degrees++
		if degrees >= 360 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}

func loop_scope() {
	count := 0

	for count = 10; count > 0; count-- {
		fmt.Println(count)
	}
	fmt.Println(count)
}

func short_loop() {
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}
}

func short_if() {
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}
}

func pi() {
	var pi64 = math.Pi
	var pi32 float32 = math.Pi

	fmt.Println(pi32)
	fmt.Println(pi64)
}

func third() {
	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%3f\n", third)
	fmt.Printf("%05.2f\n", third)
}

func inspect() {
	year := 2018
	fmt.Printf("Type %T for %v\n", year, year)
}

func inspect2() {
	days := 365.2546
	fmt.Printf("Type %T for %[1]v\n", days)
}

func raw() {
	peace := "peace"
	//var peace = "peace"
	//var peace string = "peace"
	fmt.Println(peace)
	fmt.Println("peace be upon you\nupon you be peace")
	fmt.Println(` strings can span multiple lines with the \n escape sequence `)
}

func raw_lines() {

	fmt.Println(`
    peace be upon you
    upon you be peace `)

}

func casting() {
	age := 41
	marsAge := float64(age)

	marsDays := 687.0
	earthDays := 365.2425

	marsAge = marsAge * earthDays / marsDays
	fmt.Println("I am", int(marsAge), "years old on Mars.")

}

// Function declaration - func Intn(n int) int
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsius() {
	type celsius float64
	var temperature celsius = 20
	fmt.Println(temperature)
}

func array_() {
	var planets [8]string
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2]
	fmt.Println(earth)
	fmt.Println(len(planets))
	fmt.Println(planets[3] == "")
	fmt.Println(dwarfs[0])
}

func array_loop() {
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(dwarf)
	}
}

func array_range() {
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}
}

func slicing_array() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]

	fmt.Println(terrestrial, gasGiants, iceGiants)

	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfSlice := dwarfArray
	fmt.Println(dwarfSlice)
}

func append_array() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs)

	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
	fmt.Println(dwarfs)
}

func map_dict() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %vº C.\n", temp)
	fmt.Println(temperature)
}
