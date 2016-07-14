package kernal

func InArray(elem string, array []string)(result bool) {
	result = false
	for _,value:=range array{
		if(value == elem){
			result = true
		}
	}
	return
}
