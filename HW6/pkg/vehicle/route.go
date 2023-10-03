package vehicle

import "fmt"

type Route struct {
	Name     string
	Vehicles []Vehicle
}

func (r *Route) AddToRoute(v Vehicle) {
	r.Vehicles = append(r.Vehicles, v)
}

func (r *Route) ShowVehicles() {
	for _, v := range r.Vehicles {

		fmt.Printf("Route name is:%s, Vehicle=%s\n", r.Name, v.ShowName())
	}
}
