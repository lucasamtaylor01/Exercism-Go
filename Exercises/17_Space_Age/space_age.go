package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
    // Covert seconds to Earth years
	earth_years := seconds/31557600

    //Filter by planet name and covert
    switch planet {
        case "Mercury":
            return earth_years/0.2408467
        case "Venus":
            return earth_years/0.61519726
    	case "Earth":
        	return earth_years
        case "Mars":
            return earth_years/1.8808158
        case "Jupiter":
            return earth_years/11.862615
        case "Saturn":
            return earth_years/29.447498
        case "Uranus":
            return earth_years/84.016846
        case "Neptune":
            return earth_years/165.79132
        
        //If isn't a listed planet
        default:
    		return -1
    }
}

