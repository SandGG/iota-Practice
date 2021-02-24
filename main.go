//iota -> enumerated constants
package main

import "fmt"

func main() {
	fmt.Println("Starts at value 0")
	const (
		c0 = iota //0
		c1        //1
		c2        //2
	)
	fmt.Printf("const1: %d, const2: %d, const3: %d\n", c0, c1, c2)

	fmt.Println("Starts at value 1")
	const (
		co1 = iota + 1 //1
		co2            //2
		co3            //3
	)
	fmt.Printf("const1: %d, const2: %d, const3: %d\n", co1, co2, co3)

	fmt.Println("Ignore one value")
	const (
		con1 = iota + 1 //1
		con2            //2
		_
		con4 //3
	)
	fmt.Printf("const1: %d, const2: %d, const3: %d\n", con1, con2, con4)

	fmt.Println("Permissions")
	const (
		x = 1 << iota //00000001
		w             //00000010
		r             //00000100
	)
	var bitActive = r | w //00000110
	fmt.Printf("x: %d, w: %d, r: %d, bitActive: %d\n", x, w, r, bitActive)

	fmt.Println("Constants of package OS")
	const (
		O_RDONLY = iota //0
		O_WRONLY        //1
		O_RDWR          //2

		O_CREAT = 1 << (iota + 3) //64
		O_EXCL                    //128
		_
		O_TRUNC  //512
		O_APPEND //1024

	)
	fmt.Printf("O_RDONLY: %d, O_WRONLY: %d, O_RDWR: %d, O_CREAT: %d, O_EXCL: %d, O_TRUNC: %d, O_APPEND: %d", O_RDONLY, O_WRONLY, O_RDWR, O_CREAT, O_EXCL, O_TRUNC, O_APPEND)
}
