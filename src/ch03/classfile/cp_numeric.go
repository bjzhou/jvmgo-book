package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

type ConstantLongInfo struct {
	val int64
}

type ConstantFloatInfo struct {
	val float32
}

type ConstantDoubleInfo struct {
	val float64
}

func (info *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	info.val = int32(reader.readUInt32())
}

func (info *ConstantFloatInfo) readInfo(reader *ClassReader) {
	info.val = math.Float32frombits(reader.readUInt32())
}

func (info *ConstantLongInfo) readInfo(reader *ClassReader) {
	info.val = int64(reader.readUInt64())
}

func (info *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	info.val = math.Float64frombits(reader.readUInt64())
}
