package main

import "fmt"

func main() {
	fmt.Println("Methods in golang")
	//no inheritence in golang; No super or parent
	Surya := User{"Surya", "surya@go.dev", true, 23}
	fmt.Println(Surya)
	fmt.Printf("surya details are : %v\n", Surya)
	fmt.Printf("surya with extra details are : %+v\n", Surya)
	fmt.Printf("name is %v and email is %v\n", Surya.Name, Surya.Email)
	Surya.GetStatus()
	Surya.NewMail()
	fmt.Printf("name is %v and email is %v\n", Surya.Name, Surya.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active : ", u.Status)
}

func(u User)NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is : ",u.Email)
}
