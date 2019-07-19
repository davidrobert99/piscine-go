package main

import "github.com/01-edu/z01"

type Door struct {
	state string
}

func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.state = "CLOSE"

}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = "OPEN"
}

func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?")
	return Door.state == "OPEN"
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state == "CLOSE"
}

func main() {
	door := &Door{state: "CLOSED"}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == "OPEN" {
		CloseDoor(door)
	}

}

/*
const (
	OPEN = true
	CLOSED = false
)*/
