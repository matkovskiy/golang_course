package vehicle

type Vehicle interface {
	ShowName() string
	Move()
	Stop()
	ChangeSpeed()
	OnBoarding(Passenger)
	Disembarking(Passenger)
}
