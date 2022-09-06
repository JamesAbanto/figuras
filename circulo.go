package figuras

import (
	"fmt"
	"math"
)

type Circulo struct {
	Radio float32
}

func (r *Circulo) areaPer() string {
	Area := math.Pi * r.Radio * r.Radio
	Perimetro := 2 * math.Pi * r.Radio
	return fmt.Sprintf("Area: %v\nPerimetro: %v", Area, Perimetro)
}
