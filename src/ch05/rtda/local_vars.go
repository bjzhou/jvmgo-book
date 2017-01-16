package rtda

import "math"

type Slot struct {
	num int32
	ref *Object
}

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (localVars LocalVars) SetInt(index uint, val int32) {
	localVars[index].num = val
}

func (localVars LocalVars) GetInt(index uint) int32 {
	return localVars[index].num
}

func (localVars LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	localVars[index].num = int32(bits)
}

func (localVars LocalVars) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(localVars[index].num))
}

func (localVars LocalVars) SetLong(index uint, val int64) {
	localVars[index].num = int32(val)
	localVars[index+1].num = int32(val >> 32)
}

func (localVars LocalVars) GetLong(index uint) int64 {
	low := uint32(localVars[index].num)
	high := uint32(localVars[index+1].num)
	return int64(high) << 32 | int64(low)
}

func (localVars LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	localVars.SetLong(index, int64(bits))
}

func (localVars LocalVars) GetDouble(index uint) float64 {
	bits := uint64(localVars.GetLong(index))
	return math.Float64frombits(bits)
}

func (localVars LocalVars) SetRef(index uint, val *Object) {
	localVars[index].ref = val
}

func (localVars LocalVars) GetRef(index uint) *Object {
	return localVars[index].ref
}