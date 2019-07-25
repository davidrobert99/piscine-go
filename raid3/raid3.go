package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	linha := 0
	coluna := 0
	variavel := true
	for i := range b {
		if b[i] != 10 && variavel {
			coluna++

		} else {
			variavel = false
			if b[i] == 10 {
				linha++
			}
		}
	}
	if linha > 0 && coluna > 0 {
		cmd0 := exec.Command("./raid1a", strconv.Itoa(coluna), strconv.Itoa(linha))
		cmd1 := exec.Command("./raid1b", strconv.Itoa(coluna), strconv.Itoa(linha))
		cmd2 := exec.Command("./raid1c", strconv.Itoa(coluna), strconv.Itoa(linha))
		cmd3 := exec.Command("./raid1d", strconv.Itoa(coluna), strconv.Itoa(linha))
		cmd4 := exec.Command("./raid1e", strconv.Itoa(coluna), strconv.Itoa(linha))
		dados0, _ := cmd0.Output()
		dados1, _ := cmd1.Output()
		dados2, _ := cmd2.Output()
		dados3, _ := cmd3.Output()
		dados4, _ := cmd4.Output()
		barra := false
		if string(dados0) == string(b) {
			fmt.Printf("[raid1a] [%v] [%v]", coluna, linha)
			barra = true
		}
		if string(dados1) == string(b) {
			if barra {
				fmt.Print(" || ")
			}
			fmt.Printf("[raid1b] [%v] [%v]", coluna, linha)
			barra = true
		}
		if string(dados2) == string(b) {
			if barra {
				fmt.Print(" || ")
			}
			fmt.Printf("[raid1c] [%v] [%v]", coluna, linha)
			barra = true
		}
		if string(dados3) == string(b) {
			if barra {
				fmt.Print(" || ")
			}
			fmt.Printf("[raid1d] [%v] [%v]", coluna, linha)
			barra = true
		}
		if string(dados4) == string(b) {
			if barra {
				fmt.Print(" || ")
			}
			fmt.Printf("[raid1e] [%v] [%v]", coluna, linha)
			barra = true
		}
		if !barra {
			fmt.Print("Not a Raid function")
		}
		fmt.Print("\n")
	}
}

/*func ExampleCmd_Output() {
	out, err := exec.Command("./raid1a", "2", "2").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}

/*
func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	linha := 0
	coluna := 0
	variavel := true
	for i := range b {
		if b[i] != 10 && variavel {
			coluna++

		} else {
			variavel = false
			if b[i] == 10 {
				linha++
			}
		}
	}
	raid := ""
	if b[0] == 111 {
		raid = "raid1a"
		fmt.Printf("[%v] [%v] [%v] \n", raid, coluna, linha)
	}
	if b[0] == 47 {
		raid = "raid1b"
		fmt.Printf("[%v] [%v] [%v] \n", raid, coluna, linha)
	}
	if b[0] == 65 && b[coluna-1] == 65 {
		raid = "raid1c"
		fmt.Printf("[%v] [%v] [%v] \n", raid, coluna, linha)
	} else {
		if coluna > 2 && linha > 2 {
			if b[0] == 65 && b[coluna-1] == 67 && b[len(b)-2] == 67 {
				raid = "raid1d"
				fmt.Printf("[%v] [%v] [%v] \n", raid, coluna, linha)
			}
			if b[0] == 65 && b[coluna-1] == 67 && b[len(b)-2] == 65 {
				raid = "raid1e"
				fmt.Printf("[%v] [%v] [%v] \n", raid, coluna, linha)
			}
		}
	}
		//argumentos := os.Args
	/*
		pr, pw := io.Pipe()
		defer pw.Close()

		// tell the command to write to our pipe
		cmd := exec.Command("./raid1a", "4", "4")
		cmd.Stdout = pw

		go func() {
			defer pr.Close()
			// copy the data written to the PipeReader via the cmd to stdout
			if _, err := io.Copy(os.Stdout, pr); err != nil {
				log.Fatal(err)
			}

		}()

		// run the command, which writes all output to the PipeWriter
		// which then ends up in the PipeReader
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}


} */
