package classfile

type MemberInfo struct {
	constantPool    ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, constantPool ConstantPool) []*MemberInfo {
	n := reader.readUInt16()
	members := make([]*MemberInfo, n)
	for i := range members {
		members[i] = readMember(reader, constantPool)
	}
	return members
}

func readMember(reader *ClassReader, constantPool ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    constantPool,
		accessFlags:     reader.readUInt16(),
		nameIndex:       reader.readUInt16(),
		descriptorIndex: reader.readUInt16(),
		attributes:      readAttributes(reader, constantPool),
	}
}

func (info *MemberInfo) Name() string {
	return info.constantPool.getUtf8(info.nameIndex)
}
