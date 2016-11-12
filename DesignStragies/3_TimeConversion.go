package HackerRank

import (
	"fmt"
	"strings"
	"strconv"
)

func Main3(){
	var timeStr string
	var timeStrArr []string
	fmt.Scanf("%s",&timeStr)
	timeStrArr = strings.Split(timeStr,":")
	var hour int
	hour,_ = strconv.Atoi(timeStrArr[0])
	if strings.Contains(timeStrArr[2], "PM"){
		if hour != 12{
			hour += 12
		}

	}else{
		if hour == 12{
			hour = 0
		}
	}
	fmt.Printf("%02d:%s:%s",hour,timeStrArr[1],timeStrArr[2][0:2])
}
