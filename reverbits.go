package piscine

import (
	"strconv"
)

func ReverseBits(octet byte) byte {
	aux := int64(octet)
	str := strconv.FormatInt(aux, 2)
	for len(str) != 8 {
		str = Concat("0", str)
	}
	str = StrRev(str)
	strBaseDec := ConvertBase(str, "01", "0123456789")
	valor := Atoi(strBaseDec)
	return byte(valor)
}
