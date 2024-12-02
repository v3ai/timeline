package main

import "fmt"


type person struct {
	firstName string
	lastName string
	born int
	died int
	bc_ad string
}


func main() {


	Tutankhamun := person{firstName: "Tutankhamun", lastName: "", born: 1341, died: 1323, bc_ad: "BC"}	
	// figure out siddartha here
	Socrates := person{firstName: "Socrates", lastName: "", born: 470, died: 399, bc_ad: "BC"}
	// figure out plato's exact birthdate, or put error bars
	Alexander := person{firstName: "Alexander", lastName: "the Great", born: 356, died: 323, bc_ad: "BC"}

	Archimedes := person{firstName: "Archimedes", lastName: "of Syracuse", born: 287, died: 212, bc_ad: "BC"}

	JuliusCaesar := person{firstName: "Julius", lastName: "Caesar", born: 100, died: 44, bc_ad: "BC"}

	GenghisKhan := person{firstName: "Genghis", lastName: "Khan", born: 1162, died: 1227, bc_ad: "AD"}

	LeonardodaVinci := person{firstName: "Leonardo", lastName: "da Vinci", born: 1452, died: 1519, bc_ad: "AD"}
	


	
	people := []person{Tutankhamun, Socrates, Alexander, Archimedes, JuliusCaesar, GenghisKhan, LeonardodaVinci}
	

	for i := 0; i < len(people); i++ {
		fmt.Println(people[i])
	}

}
