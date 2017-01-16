package classfile

type ConstantPool []ConstantInfo

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if info := cp[index]; info != nil {
		return info
	}
	panic("Invalid constant pool index!")
}

func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	info := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := cp.getUtf8(info.nameIndex)
	_type := cp.getUtf8(info.descriptorIndex)
	return name, _type
}

func (cp ConstantPool) getClassName(index uint16) string {
	info := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(info.nameIndex)
}

func (cp ConstantPool) getUtf8(index uint16) string {
	return cp.getConstantInfo(index).(*ConstantUtf8Info).str
}
