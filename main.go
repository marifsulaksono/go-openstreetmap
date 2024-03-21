package main

func main() {
	// Get the venue location
	params := "?q=Borobudur&format=json" // query params search
	PlaceIdentifier(params)

	// Get the route from start to end location
	vehicle := "car"           // type of vehicle
	startLat := "-7.7605736"   // latitude starting point location (ex. Kraksaan, Probolinggo, Jawa Timur)
	startLong := "113.4164516" // longtitude starting point location (ex. Kraksaan, Probolinggo, Jawa Timur)
	endLat := "-8.0115593"     // latitude starting point location (ex. Malang, Jawa Timur)
	endLong := "112.6444829"   // longtitude starting point location (ex. Malang, Jawa Timur)
	GetMapDirectionUsingGraphhopperEngine(vehicle, startLat, startLong, endLat, endLong)
}
