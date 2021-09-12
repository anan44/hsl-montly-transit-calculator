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
	toWork := hsl.NewRoute("To Work", home, work)
	toPint := hsl.NewRoute("To Pint", home, pint)
	routes := []*hsl.Route{&toWork, &toPint}
	for _, route := range routes {
		wg.Add(1)
		route.Calculate(&wg)
	}
	wg.Wait()
	for _, route := range routes {
		fmt.Println(*route)
	}
}
