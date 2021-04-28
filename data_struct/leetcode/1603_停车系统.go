package leetcode

type ParkingSystem struct {
	val [3]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{val: [3]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.val[carType] ==0{
		this.val[carType]++
		return true
	}
	return false
}
