package design_parking_system

type ParkingSystem struct {
	big, medium, small int
}

func Constructor(big, medium, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}

func (ps *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		ps.big--
		return ps.big >= 0
	case 2:
		ps.medium--
		return ps.medium >= 0
	case 3:
		ps.small--
		return ps.small >= 0
	default:
		panic("invalid car type")
	}
}
