package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    steps := 0
    if n > 0{
       	for n != 1{
            if (n % 2 == 0){
            	n  = n/2    
            }else{
            	n = 3*n+1
            }
        	steps+=1
        }
    	return steps, nil 
    }else{
    	return 0, errors.New("input must be a positive number")
    }

}

