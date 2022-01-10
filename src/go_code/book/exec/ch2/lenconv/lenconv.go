//长度换算
package lenconv

import "fmt"

type Inch float64
type Meter float64
const (
	SingleInToM = 0.0254
	SingleMToIn = 39.3700787
)

func (in Inch) String() string { return fmt.Sprintf("%gin", in)}
func (m Meter) String() string { return fmt.Sprintf("%gm", m)}
//InToM:英寸转化为米
func InToM(in Inch) Meter { return Meter(in * SingleInToM)}
//MToIn:米转化为英寸
func MToIn(m Meter) Inch { return Inch(m * SingleMToIn)} 

