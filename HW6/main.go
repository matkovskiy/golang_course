package main

import (
	"github.com/matkovskiy/golang_course/HW6/pkg/post"
	"github.com/matkovskiy/golang_course/HW6/pkg/vehicle"
)

func main() {
	//Exercise 0001
	branch_0001 := post.PostOfficeBranch{City: "Odesa", Name: "Branch_0001"}
	branch_0001.Proccessing(post.Box{Sender: "Vasya", Destination_country: "Ukraine", Source_county: "Greece"})
	branch_0001.Proccessing(post.Envelope{Sender: "Vasya", Source_county: "Romania", Destination_country: "Poland"})

	// Exercise 0002
	println("========================================================")
	Vasya := vehicle.Passenger{Name: "Vasya", Sex: "male", Weight: 100}
	Petya := vehicle.Passenger{Name: "Petya", Sex: "male", Weight: 80}
	Linda := vehicle.Passenger{Name: "Linda", Sex: "female", Weight: 55}

	thecar := vehicle.Auto{Price: 1000, Speed: 100, Name: "BMW"}
	thetain := vehicle.Trian{Price: 100000, Speed: 50, Name: "Train1"}
	theairplaine := vehicle.Airplane{Price: 5000000, Speed: 800, Name: "Boing"}
	vehicle.Vehicle.Move(thecar)
	vehicle.Vehicle.Move(thetain)
	vehicle.Vehicle.Move(theairplaine)

	vehicle.Vehicle.OnBoarding(thecar, Vasya)
	vehicle.Vehicle.OnBoarding(thetain, Petya)
	vehicle.Vehicle.OnBoarding(theairplaine, Linda)

	vehicle.Vehicle.Stop(thecar)
	vehicle.Vehicle.Stop(thetain)
	vehicle.Vehicle.Stop(theairplaine)

	vehicle.Vehicle.ChangeSpeed(thecar)
	vehicle.Vehicle.ChangeSpeed(thetain)
	vehicle.Vehicle.ChangeSpeed(theairplaine)

	vehicle.Vehicle.Disembarking(thecar, Vasya)
	vehicle.Vehicle.Disembarking(thetain, Petya)
	vehicle.Vehicle.Disembarking(theairplaine, Linda)

	Odesa := vehicle.Route{Name: "Odesa"}
	Kyiv := vehicle.Route{Name: "Kyiv"}
	Odesa.AddToRoute(thecar)
	Odesa.AddToRoute(thetain)

	Kyiv.AddToRoute(theairplaine)
	Odesa.ShowVehicles()
	Kyiv.ShowVehicles()

}
