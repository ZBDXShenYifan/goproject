package array

type Array struct {
	data []int
	size int
}

func NewArray(capacity int) *Array {
	return &Array {
		data: make([]int, capacity)
	}
}

func NewArray() *Array{
	return &Array {
		data: make([]int, capacity)
	}
}

func (a *Array) GetCapacity(){
	return cap(a.data)
}

func (a *Array) Get