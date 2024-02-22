package raindrops
import "strconv"

func Convert(number int) string {
	if number % 3 == 0{
        if number % 5 == 0 && number % 7 == 0{
        	return "PlingPlangPlong"
        }else if number % 7 == 0{
        	return "PlingPlong"
        }else if number % 5 == 0{
            return "PlingPlang"
        }else{
        	return "Pling"
        }
    }else if number % 5 == 0{
    	if number % 7 == 0{
          return "PlangPlong"  
        }else{
        	return "Plang"
        }
        
    }else if number % 7 == 0{
    	return "Plong"
    }else{
    	return strconv.Itoa(number)
    }
}

