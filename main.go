package main

import ("fmt"
		"icp/unittesting/count"
)


func main() {
	fmt.Println("Hello, 世界");	
	result := count.Tambah(1,2);
	fmt.Println("Hasil: ", result);
}