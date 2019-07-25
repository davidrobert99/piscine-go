package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

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
	pr, pw := io.Pipe()
	defer pw.Close()

	// tell the command to write to our pipe
	cmd := exec.Command("cat", "fruit.txt")
	cmd.Stdout = pw

	cmd = exec.Command("ls", "-lah")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

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
	*/

}
