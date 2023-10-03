package vehicle

import "fmt"

type Airplane struct {
	Price, Speed int
	Name         string
}

func (a Airplane) Move() {
	fmt.Println("The Airplane is moving")
}
func (a Airplane) Stop() {
	fmt.Println("The Airplane has been stopped")
}
func (a Airplane) ChangeSpeed() {
	fmt.Println("The Airplane changes speed")
}

func (a Airplane) OnBoarding(p Passenger) {
	fmt.Printf("welcome on board, dear: %s, to our Airplane\n", p.Name)
}

func (a Airplane) Disembarking(p Passenger) {
	fmt.Printf("Good Bye: %s.\nWe are waiting for you again on our bus our Airplane\n", p.Name)
}

func (a Airplane) ShowName() string {
	return a.Name
}

func (a Airplane) AddToRoute(r Route) {
	r.Vehicles = append(r.Vehicles, a)
}
