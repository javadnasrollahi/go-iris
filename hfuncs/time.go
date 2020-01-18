package hfuncs

import (
	"fmt"
	"time"
)

func Timestamp() string {
	return time.Now().In(time.FixedZone("UTC-8", +(3*60*60 + 30*60))).Format(time.RFC3339)
}
func Get5Char() string {
	timeString := fmt.Sprintf("%d", time.Now().Round(time.Millisecond).UnixNano()/(int64(time.Millisecond)/int64(time.Nanosecond)))
	subTimeString := timeString[len(timeString)-5 : len(timeString)]
	var code string
	for i := 0; i < len(subTimeString); i++ {
		code += string(subTimeString[i])
	}
	return code
}
