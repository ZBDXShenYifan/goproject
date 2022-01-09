package tempconv
import "fmt"
//Celsius Fahrenheit不能使用算术表达式进行比较和合并
type Celsius float64
type Fahrenheit float64
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)
//Fahrenheit()显示类型转换，非函数
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32)}
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * / 9) }