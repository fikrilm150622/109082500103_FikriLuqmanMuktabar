package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	for {
		fmt.Scanf("%c", &ch)
		if ch == '.' || *n >= NMAX {
			break
		}
		if ch != '\n' && ch != ' ' {
			t[*n] = ch
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	var temp rune
	for i := 0; i < n/2; i++ {
		temp = t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	var rev tabel
	for i := 0; i < n; i++ {
		rev[i] = t[i]
	}
	balikanArray(&rev, n)
	for i := 0; i < n; i++ {
		if t[i] != rev[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int
	fmt.Print("Teks         : ")
	isiArray(&tab, &m)

	var asli tabel
	for i := 0; i < m; i++ {
		asli[i] = tab[i]
	}	
	balikanArray(&tab, m)
	fmt.Print("Reverse teks : ")
	cetakArray(tab, m)
	fmt.Print("\nTeks         : ")
	cetakArray(asli, m)
	fmt.Print("Palindrom    ? ")
	if palindrom(asli, m) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}