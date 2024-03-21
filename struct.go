// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	type Coordenada struct {
		latitude  float64
		longitude float64
	}

	coordenada := Coordenada{latitude: 37.42, longitude: -122.08}

	coordenada.latitude +=100

	fmt.Println(coordenada)
	fmt.Println(coordenada.latitude)
	fmt.Printf("%+v", coordenada)

}
