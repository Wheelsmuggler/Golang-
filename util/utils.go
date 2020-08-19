package util

import (
	"github.com/huichen/sego"
)

var segmenter sego.Segmenter

func init()  {
	segmenter.LoadDictionary("/Users/hjzhou/Gocode/src/github.com/huichen/sego/data/dictionary.txt")
}

func RemoveRepeatElement(arr []string) (newArr []string) {
	newArr = make([]string,0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr,arr[i])
		}
	}
	return
}

func ExtractKeyWords(ping string) (keyWords [3][]string){
	text := []byte(ping)
	segments := segmenter.Segment(text)
	//fmt.Println(sego.SegmentsToString(segments, false))
	nums := sego.SegmentsToString(segments,false)
	nums = RemoveRepeatElement(nums)
	for i :=0;i<3 ; i++ {
		for j := 0;j< len(tags[i]); j++ {
			for k := 0; k< len(nums); k++ {
				if tags[i][j]==nums[k] {
					keyWords[i] = append(keyWords[i],nums[k])
				}
			}
		}
	}
	return
}