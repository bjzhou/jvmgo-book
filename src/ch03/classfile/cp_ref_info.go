package classfile

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (info *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	info.classIndex = reader.readUInt16()
	info.nameAndTypeIndex = reader.readUInt16()
}

func (info *ConstantMemberrefInfo) ClassName() string {
	return info.cp.getClassName(info.classIndex)
}

func (info *ConstantMemberrefInfo) NameAndType() (string, string) {
	return info.cp.getNameAndType(info.nameAndTypeIndex)
}

type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }
type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }