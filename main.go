package main

import (
	"fmt"
	"hsl-transit/transit-calc/hsl"
	"sync"
)

func main() {
	home := hsl.Location{Address: "Leonkatu"}
	wg := sync.WaitGroup{}
	commutes := hsl.MonthlyCommutes{Routes: []hsl.Route{
		hsl.NewRoute("To Work", home, hsl.Location{Address: "Eteläesplanadi"}, 5),
		hsl.NewRoute("To Pint", home, hsl.Location{Address: "One Pint Pub"}, 3),
		hsl.NewRoute("To Station", home, hsl.Location{Address: "Rautatientori"}, 5),
	}}
	for i := range commutes.Routes {
		wg.Add(1)
		commutes.Routes[i].Estimate(&wg)
	}
	wg.Wait()
	fmt.Println(commutes.TotalDuration())
}
