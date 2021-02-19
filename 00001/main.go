package main

import "fmt"

func main() {
	for i := 0; i < 52; i++ {
		var color = i / 13
		var carNum = i % 13
		fmt.Print(i, " ")

		switch color {
		case 0:
			fmt.Print("红桃")
		case 1:
			fmt.Print("方块")
		case 2:
			fmt.Print("梅花")
		case 3:
			fmt.Print("黑桃")
		}

		switch carNum {
		case 0:
			fmt.Println("2")
		case 1:
			fmt.Println("3")
		case 2:
			fmt.Println("4")
		case 3:
			fmt.Println("5")
		case 4:
			fmt.Println("6")
		case 5:
			fmt.Println("7")
		case 6:
			fmt.Println("8")
		case 7:
			fmt.Println("9")
		case 8:
			fmt.Println("10")
		case 9:
			fmt.Println("J")
		case 10:
			fmt.Println("Q")
		case 11:
			fmt.Println("K")
		case 12:
			fmt.Println("A")
		}
	}
	/*
		1、方块 6、红心 7、黑桃  7、梅花 8、梅花 10
		2、黑桃 9、黑桃 2、红心 10、红心 5、方块 Q
		3、红桃 J、方块 K、黑桃  K、红心 3、黑桃 5
		4、梅花 3、方块 A、黑桃  A、梅花 K、梅花 5
		5、红心 A、红心 9、黑桃  J、黑桃 4、红心 K

		1 6     0 7     3  7     2 8     2 10
		3 9     3 2     0 10     0 5     1  Q
		0 J     1 K     3  K     0 3     3  5
		2 3     1 A     3  A     2 K     2  5
		0 A     0 9     3  J     3 4     0  K

	*/
	// 17  5 44 32 34
	// 46 39  8  3 23
	//  9 24 50  1 42
	// 27 25 51 37 29
	// 12  7 48 41 11
	/*
		17  5 44 32
		34 46 39  8
		 3 23  9 24
		50  1 42 27
		25 51 37 29
		12  7 48 41
		11
	*/

	/*
		var arr [52]int
		var n = 17
		for i := 0; i < 25; i++ {
			if n < 0 {
				n += 51
			}

			arr[n] = 1
			fmt.Print(arr, "  ")
			fmt.Println(n)
			n -= 12
		}
	*/

	var arr1 = [52]int{
		17, 5, 44, 32,
		34, 46, 39, 8,
		3, 23, 9, 24,
		50, 1, 42, 27,
		25, 51, 37, 29,
		12, 7, 48, 41,
		11,
	}

	var gap = ".."
	var arr [52]string
	for i, _ := range arr {
		arr[i] = gap
	}

	for l, v := range arr1 {
		for i, _ := range arr {
			arr[i] = gap
		}
		if l <= 24 {
			arr[v] = fmt.Sprintf("%2d", v)
		}
		fmt.Println(fmt.Sprintf("%2d", l), arr)
	}
}
