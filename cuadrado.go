package figuras

import "fmt"

type Cuadrado struct {
	Lado float32
}

func (c *Cuadrado) areaPer() string {
	Area := c.Lado * c.Lado
	Perimetro := 4 * c.Lado
	return fmt.Sprintf("Area: %v\nPerimetro: %v", Area, Perimetro)
}
