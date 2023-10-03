package vehicle

import "fmt"

type Trian struct {
	Price, Speed int
	Name         string
}

func (t Trian) Move() {
	fmt.Println("The Trian is moving")
}

func (t Trian) Stop() {
	fmt.Println("The Trian has been stopped")
}
func (t Trian) ChangeSpeed() {
	fmt.Println("The Trian changes speed")
}

func (t Trian) OnBoarding(p Passenger) {
	fmt.Printf("welcome on board, dear: %s, to our Trian\n", p.Name)
}

func (a Trian) Disembarking(p Passenger) {
	fmt.Printf("Good Bye: %s.\nWe are waiting for you again on our bus our Trian\n", p.Name)
}

func (t Trian) ShowName() string {
	return t.Name
}

func (t Trian) AddToRoute(r Route) {
	r.Vehicles = append(r.Vehicles, t)
}
