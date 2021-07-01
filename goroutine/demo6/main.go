package main

import "fmt"

type User struct {
	Name string
	Age int
}

func main() {

	users := []User{{"jerry",20},{"zhangsna",25},{"tom",30}}

	m := make(map[string]*User)

	//for i := 0; i < len(users); i++ {
	//	m[users[i].Name] = &users[i]
	//}

	//for i, _ := range users{
	//	m[users[i].Name] = &users[i]
	//}

	for _, v := range users {
		u := v
		m[v.Name] = &u
	}

	for k,v := range m {
		fmt.Println(k,v)
	}
}
