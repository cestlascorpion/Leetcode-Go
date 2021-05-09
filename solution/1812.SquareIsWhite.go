package solution

func SquareIsWhite(coordinates string) bool {
	if len(coordinates) != 2 {
		panic("coordinates size != 2")
	}
	return (coordinates[0]+coordinates[1])%2 != 0
}
