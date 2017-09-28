package helper

import "fmt"

//删除重复的url
func RemoveDuplicates(urls []string) []string {
	ret := urls
	for k, v := range urls {
		for ck, cv := range urls{
			fmt.Print("比", k, v)
			fmt.Println("\n")
			fmt.Print("对",ck, cv)
			fmt.Println("\n")
			if(v == cv && k != ck){
				fmt.Println(ret[k:])
				fmt.Println(ret[k+1:])
				ret = append(ret[:k], ret[k+1:]...)
			}

		}
	}

	return ret
}
