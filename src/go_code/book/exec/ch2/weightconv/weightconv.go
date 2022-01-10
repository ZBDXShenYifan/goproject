package weightconv

import "fmt"

type Kilogram float64
type Pound float64

func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg)}
func (p Pound) String() string { return fmt.Sprintf("%gp", p)}

func KToP(kg Kilogram) Pound { return Pound(kg * 2.2046226) }
func PToK(p Pound) Kilogram { return Kilogram(p * 0.4535924 ) }