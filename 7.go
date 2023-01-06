package main

import "fmt"

type hero struct {
	id                    string
	cost, attack, defense int
}

type arrHero [123]hero

func main() {
	var A arrHero
	var m int

	isiData(&A, &m)
	sorting(&A, m)
	print(A, m)
}

func isiData(T *arrHero, n *int) {
	/* I.S. sejumlah data hero (ID hero, cost, attack, dan defence) telah siap pada piranti masukan
	   F.S. array T berisi sejumlah n data hero, masukan berakhir apabila ID hero adalah "0000"*/
	var ID string
	*n = 0
	fmt.Scan(&ID)
	for ID != "0000" && *n < 123 {
		T[*n].id = ID
		fmt.Scan(&T[*n].cost, &T[*n].attack, &T[*n].defense)
		*n++
		fmt.Scan(&ID)
	}
}

func sorting(T *arrHero, n int) {
	/* I.S. terdefinisi array T yang terdiri dari n data hero
	   F.S. array T terurut mengecil berdasarkan total attack dan defence menggunakan algoritma selection atau insertion sort */
	var i, pass int
	var temp hero

	for pass = 1; pass <= n-1; pass++ {
		i = pass
		temp = T[pass]
		for i > 0 && (temp.attack+temp.defense) > (T[i-1].attack+T[i-1].defense) {
			T[i] = T[i-1]
			i = i - 1
		}
		T[i] = temp
	}
}

func print(T arrHero, n int) {
	/* I.S. terdefinisi array T yang terdiri dari n data hero
	   F.S. menampilkan paling banyak 5 data array T yaitu id, total attack dan defence*/
	var i int
	i = 0
	for i < 5 {
		fmt.Println(T[i].id, (T[i].attack + T[i].defense))
		i++
	}
}
