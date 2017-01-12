package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (info *ConstantStringInfo) readInfo(reader *ClassReader) {
	info.stringIndex = reader.readUInt16()
}

func (info *ConstantStringInfo) String() string {
	return info.cp.getUtf8(info.stringIndex)
}
