package rtda

import "math"

type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (stack *OperandStack) PushInt(val int32) {
	stack.slots[stack.size].num = val
	stack.size++
}

func (stack *OperandStack) PopInt() int32 {
	stack.size--
	return stack.slots[stack.size].num
}

func (stack *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	stack.PushInt(int32(bits))
}

func (stack *OperandStack) PopFloat() float32 {
	return math.Float32frombits(uint32(stack.PopInt()))
}

func (stack *OperandStack) PushLong(val int64) {
	stack.PushInt(int32(val))
	stack.PushInt(int32(val >> 32))
}

func (stack *OperandStack) PopLong() int64 {
	high := uint32(stack.PopInt())
	low := uint32(stack.PopInt())
	return int64(high) << 32 | int64(low)
}

func (stack *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	stack.PushLong(int64(bits))
}

func (stack *OperandStack) PopDouble() float64 {
	return math.Float64frombits(uint64(stack.PopLong()))
}

func (stack *OperandStack) PushRef(val *Object) {
	stack.slots[stack.size].ref = val
	stack.size++
}

func (stack *OperandStack) PopRef() *Object {
	stack.size--
	ref := stack.slots[stack.size].ref
	stack.slots[stack.size].ref = nil
	return ref
}