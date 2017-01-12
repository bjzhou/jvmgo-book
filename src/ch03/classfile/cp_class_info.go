package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (info *ConstantClassInfo) readInfo(reader *ClassReader) {
	info.nameIndex = reader.readUInt16()
}

func (info *ConstantClassInfo) String() string {
	return info.cp.getUtf8(info.nameIndex)
}