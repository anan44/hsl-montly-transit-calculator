package main

import (
	"fmt"
	"hsl-transit/transit-calc/hsl"
)

func main() {
	home := hsl.Location{Address: "Leonkatu"}
	commutes := hsl.NewMonthlyCommutes([]hsl.Route{
		hsl.NewRoute("To Work", home, hsl.Location{Address: "Eteläesplanadi"}, 5),
		hsl.NewRoute("To Pint", home, hsl.Location{Address: "One Pint Pub"}, 3),
		hsl.NewRoute("To Station", home, hsl.Location{Address: "Rautatientori"}, 5),
	})
	fmt.Println(commutes.TotalDuration())
}
