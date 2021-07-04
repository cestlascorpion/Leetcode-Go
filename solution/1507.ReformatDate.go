package solution

import "strings"

func getDay(str string) string {
	day := string([]byte(str)[0 : len(str)-2])
	if len(day) == 1 {
		day = "0" + day
	}
	return day
}

func getMonth(str string) string {
	switch str {
	case "Jan":
		return "01"
	case "Feb":
		return "02"
	case "Mar":
		return "03"
	case "Apr":
		return "04"
	case "May":
		return "05"
	case "Jun":
		return "06"
	case "Jul":
		return "07"
	case "Aug":
		return "08"
	case "Sep":
		return "09"
	case "Oct":
		return "10"
	case "Nov":
		return "11"
	case "Dec":
		return "12"
	}
	return "??"
}

func ReformatDate(date string) string {
	elements := strings.Split(date, " ")
	return elements[2] + "-" + getMonth(elements[1]) + "-" + getDay(elements[0])
}
