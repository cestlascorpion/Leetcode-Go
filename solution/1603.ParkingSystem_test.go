package solution

import "testing"

func TestParkingSystem_AddCar(t *testing.T) {
	p := ConstructorParkingSystem(1, 1, 0)
	if !p.AddCar(1) {
		t.Fail()
	}
	if !p.AddCar(2) {
		t.Fail()
	}
	if p.AddCar(3) {
		t.Fail()
	}
}
