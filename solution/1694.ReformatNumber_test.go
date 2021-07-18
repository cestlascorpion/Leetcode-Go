package solution

import "testing"

func TestReformatNumber(t *testing.T) {
	number := "1-23-45 6"
	exp := "123-456"
	if ReformatNumber(number) != exp {
		t.Fail()
	}
}

func TestReformatNumber2(t *testing.T) {
	number := "1-23-45 6"
	exp := "123-456"
	if ReformatNumber(number) != exp {
		t.Fail()
	}
}

func TestReformatNumber3(t *testing.T) {
	number := "123 4-567"
	exp := "123-45-67"
	if ReformatNumber(number) != exp {
		t.Fail()
	}
}

func TestReformatNumber4(t *testing.T) {
	number := "123 4-5678"
	exp := "123-456-78"
	if ReformatNumber(number) != exp {
		t.Fail()
	}
}

func TestReformatNumber5(t *testing.T) {
	number := "12"
	exp := "12"
	if ReformatNumber(number) != exp {
		t.Fail()
	}
}

func TestReformatNumber6(t *testing.T) {
	number := "--17-5 229 35-39475 "
	exp := "175-229-353-94-75"
	if ReformatNumber(number) != exp {
		t.Fail()
	}
}
