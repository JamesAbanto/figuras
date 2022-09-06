package figuras

import "fmt"

type Forma interface {
	areaPer() string
}

func AreaForma(forma Forma) {
	fmt.Println(forma.areaPer())
}
