package main

type smallStruct struct {
	a,b int64
	c,d float64
}

//go:noinline
func smallAllocation() *smallStruct {
	return &smallStruct{}
}

func main() {
	smallAllocation()
}

type User struct {
	UserName string `json:"userName"`
}
