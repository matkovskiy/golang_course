package vehicle

import "fmt"

type Auto struct {
	Price, Speed int
	Name         string
}

func (a Auto) Move() {
	fmt.Println("The auto is moving")
}
func (a Auto) Stop() {
	fmt.Println("The auto has been stopped")
}
func (a Auto) ChangeSpeed() {
	fmt.Println("The auto changes speed")
}

func (a Auto) OnBoarding(p Passenger) {
	fmt.Printf("welcome on board, dear: %s, to our Car\n", p.Name)
}

func (a Auto) Disembarking(p Passenger) {
	fmt.Printf("Good Bye: %s.\nWe are waiting for you again on our bus our Auto\n", p.Name)
}
func (a Auto) ShowName() string {
	return a.Name
}

func (a Auto) AddToRoute(r Route) {
	r.Vehicles = append(r.Vehicles, a)
}
