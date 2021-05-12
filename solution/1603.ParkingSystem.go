package solution

const (
	BIG = iota + 1
	MEDIUM
	SMALL
)

type ParkingSystem struct {
	cap [3]int
}

func ConstructorParkingSystem(big int, medium int, small int) ParkingSystem {
	p := ParkingSystem{
		cap: [3]int{big, medium, small},
	}
	return p
}

func (p *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case BIG:
		if p.cap[BIG-1] > 0 {
			p.cap[BIG-1]--
			return true
		}
	case MEDIUM:
		if p.cap[MEDIUM-1] > 0 {
			p.cap[MEDIUM-1]--
			return true
		}
	case SMALL:
		if p.cap[SMALL-1] > 0 {
			p.cap[SMALL-1]--
			return true
		}
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
