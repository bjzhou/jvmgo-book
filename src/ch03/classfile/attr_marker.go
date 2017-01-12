package classfile

type MarkerAttribute struct{}

func (attr *MarkerAttribute) readInfo(reader *ClassReader) {}

type DeprecatedAttribute struct{ MarkerAttribute }
type SyntheticAttribute struct{ MarkerAttribute }
