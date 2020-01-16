package cache

import (
	"errors"
	"fmt"
	"golang.org/x/exp/utf8string"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
)

func TestMemoryCache_GetSize(t *testing.T) {
	//t1 := make(map[string]string)
	//
	//fmt.Println(unsafe.Sizeof(t1))
	//
	//t1["xxx"] = "xxxxx"
	//t1["xxx1"] = "xxxxx"
	//t1["xxx2"] = "xxxxx"
	//t1["xxx3"] = "xxxxx"
	//
	//f := unsafe.Sizeof(t1)
	//fmt.Println(f)
	//
	//fmt.Println(t1)

	name := "我是胡八一"

	//fmt.Println("name[:4] = ", name[:4])
	//fmt.Println("name[:1]", substr(name, 0, 1))
	//fmt.Println("substr:")
	//fmt.Println(substrRange(name, -1, 1))
	start := time.Now()
	for i:=0; i< 100000; i++ {
		substr(name, 0, 1)
	}
	middle := time.Now()
	for i:=0; i< 100000; i++ {
		substrRange(name, 0, 1)
	}
	end := time.Now()

	fmt.Printf("utf8string used %v\n", middle.Sub(start))
	fmt.Printf("range used %v\n", end.Sub(middle))
	//var re string
	////var j = 0
	//for i,c := range name {
	//	re += string(c)
	//	fmt.Println(i)
	//	if i == 4 {
	//		break
	//	}
	//}
	//fmt.Println(len(name))
	//fmt.Println(re)
}

func substrRange(str string, start, length int) (string, error) {
	l := utf8.RuneCountInString(str)
	if l == 0 {
		return "", errors.New("Str is empty! ")
	}

	if start < 0 {
		start = l + start
	}

	if (start + length) > l {
		return str, errors.New("Str length is less than start+length! ")
	}
	sb := new(strings.Builder)
	var pos = 0
	for _, c := range str {
		if pos >= start {
			if pos >= start+length {
				break
			}
			sb.WriteRune(c)

		}
		pos++
	}

	return sb.String(), nil
}

func substr(str string, start, length int) (string, error) {
	uStr := utf8string.NewString(str)
	l := uStr.RuneCount()
	if start < 0 {
		start = l + start
	}

	if (start + length) > l {
		return str, errors.New("Str length is less than start+length! ")
	}
	sb := new(strings.Builder)
	for i := start; i < length; i++ {
		sb.WriteRune(uStr.At(i))
	}
	return sb.String(),nil
}
