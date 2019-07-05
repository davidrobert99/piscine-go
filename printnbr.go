
package piscine

import "github.com/01-edu/z01"




func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune(rune(48))
	}else {
		conta  := 0
		aux := n
		for aux!= 0{
			conta++
			aux = aux/10
		}
		vet:=make([] int, conta)
		if n > 0{
			for n%10 != 0 || n/10 > 0{
				conta--
				vet[conta] = n%10
				n = n/10
			}
		}else{
			z01.PrintRune(rune(45))
			n = n*-1
			for n%10 != 0 || n/10 > 0{
				conta--
				vet[conta] = n%10
				n = n/10
			}
		}
		for i:=0 ; i<len(vet); i++{
			z01.PrintRune(rune(vet[i]+ 48))
		}
	}
}



