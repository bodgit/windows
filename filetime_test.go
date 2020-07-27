package windows

import (
	"fmt"
	"time"
)

func ExampleNsecToFiletime() {
	ft := NsecToFiletime(time.Date(2020, 7, 20, 17, 15, 10, 0, time.UTC).UnixNano())

	fmt.Println(ft)
	// Output: {1384030976 30826169}
}

func ExampleFiletime_Nanoseconds() {
	ft := Filetime{
		1384030976,
		30826169,
	}

	fmt.Println(time.Unix(0, ft.Nanoseconds()).UTC())
	// Output: 2020-07-20 17:15:10 +0000 UTC
}
