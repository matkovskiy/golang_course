package observer

import "fmt"

type Observer interface {
	Update(string)
}
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify(message string)
}
type Player struct {
	Name string
}

func (p *Player) Update(message string) {
	fmt.Printf("[%s] Got a message: %s\n", p.Name, message)
}

type Game struct {
	Observers map[Observer]struct{}
}

func (g *Game) Register(observer Observer) {
	g.Observers[observer] = struct{}{}
}

func (g *Game) Unregister(observer Observer) {
	delete(g.Observers, observer)
}

func (g *Game) Notify(message string) {
	for observer := range g.Observers {
		observer.Update(message)
	}
}
