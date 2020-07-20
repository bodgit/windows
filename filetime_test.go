package windows

import (
	"encoding/hex"
	"fmt"
)

func ExampleFiletime_MarshalBinary() {
	ft := NsecToFiletime(1595265310000000000) // time.Date(2020, 7, 20, 17, 15, 10, 0, time.UTC).UnixNano()

	b, err := ft.MarshalBinary()
	if err != nil {
		panic(err)
	}

	fmt.Print(hex.Dump(b))
	// Output: 00000000  00 a3 7e 52 b9 5e d6 01                           |..~R.^..|
}

func ExampleFiletime_UnmarshalBinary() {
	ft := new(Filetime)

	if err := ft.UnmarshalBinary([]byte{0x00, 0xa3, 0x7e, 0x52, 0xb9, 0x5e, 0xd6, 0x01}); err != nil {
		panic(err)
	}

	fmt.Println(ft.Nanoseconds())
	// Output: 1595265310000000000
}
