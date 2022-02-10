package logsdk

import "strings"

type LogFilter func(str string)bool


func BlackWordFilter(words ...string)LogFilter{
	return func(str string) bool {
		for _,word:=range words{
			if strings.Contains(str,word){
				return true
			}
		}
		return false
	}
}