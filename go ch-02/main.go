package main

import "fmt"

func main() {
	var x int;
	var a,b,c = 45.5, "hello world", 6
	y := 40;
	x= 56;

	z := x + y;
	fmt.Println("x: ", x);
	fmt.Printf("y %T\n", y);
	fmt.Println("z: ", z);
	fmt.Printf("a: %T\n", a);
	fmt.Printf("b: %T\n", b);
	fmt.Printf("c: %T\n", c);


	if x>y {
		fmt.Printf("Hello i am if x %d\n", x)
	}else{
		fmt.Printf("Hello i am else y %d\n", y)
	}
	
}