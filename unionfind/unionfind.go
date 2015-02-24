package main

import "fmt"

type UF struct {
	arrayN []int
}

func (u *UF) init(size int) {
	for i := 0; i < size; i++ {
		u.arrayN = append(u.arrayN, i)
	}
	fmt.Println(u.arrayN)
}

func (u *UF) Union(a int, b int) {
	aid := u.arrayN[a]
	bid := u.arrayN[b]
	for i, _ := range u.arrayN {
		if u.arrayN[i] == aid {
			u.arrayN[i] = bid
		}
	}
}

func (u *UF) Connected(a int, b int) bool {
	return u.arrayN[a] == u.arrayN[b]
}

func main() {
	myUF := UF{}
	myUF.init(20)
	myUF.Union(2, 3)
	myUF.Union(1, 3)
	myUF.Union(16, 3)
	myUF.Union(3, 17)
	fmt.Println(myUF.Connected(1, 17))
	fmt.Println(myUF.Connected(1, 9))

	fmt.Println(myUF.arrayN)
}
