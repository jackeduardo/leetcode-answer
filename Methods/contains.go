package Methods

import "reflect"

func Contains(res[][]int,temp []int) bool{
	if res==nil{
		return false
	}
	for _, v := range res {
		if reflect.DeepEqual(temp,v){
			return true
		}
	}
	return false
}