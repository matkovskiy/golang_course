package main

type login string
type password string

type attempt struct {
	Login    login    `json:"login"`
	Password password `json:"password"`
	Role     string   `json:"role"`
}

type Student struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Username  string `json:"Username"`
	Password  string `json:"Password"`
	ID        string `json:"ID"`
}

type Students struct {
	Student []Student `json:"students"`
}

type Task struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type Tasks struct {
	Task []Task `json:"tasks"`
}

type User struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Users struct {
	User []User `json:"users"`
}
