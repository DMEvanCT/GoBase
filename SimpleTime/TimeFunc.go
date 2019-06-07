package SimpleTime

import (
	"fmt"
	"time"
)

func TimeOut(years, months, days int) string {
	t := time.Now()
	expire := t.AddDate(years, months,days)
	edate := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		expire.Year(), expire.Month(), expire.Day(),
		expire.Hour(), expire.Minute(), expire.Second())

	return edate
}

