package main

type Pilot struct {
	Name     string
	Life     float32
	Age      int
	Aircraft string
}

const AIRCRAFT1 = ""

func  main(){
	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}