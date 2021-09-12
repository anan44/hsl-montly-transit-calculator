package main

import (
	"fmt"
	"hsl-transit/transit-calc/hsl"
	"sync"
)

func main() {
	home := hsl.Location{Address: "Leonkatu"}
	work := hsl.Location{Address: "Eteläesplanadi"}
	pint := hsl.Location{Address: "One Pint Pub"}
	wg := sync.WaitGroup{}
	routes := []hsl.Route{
		hsl.NewRoute("To Work", home, work),
		hsl.NewRoute("To Pint", home, pint),
	}
	for i := range routes {
		wg.Add(1)
		routes[i].Calculate(&wg)
	}
	wg.Wait()
	for _, route := range routes {
		fmt.Println(route)
	}
}
