package classfile

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (info *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	info.referenceKind = reader.readUInt8()
	info.referenceIndex = reader.readUInt16()
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (info *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	info.descriptorIndex = reader.readUInt16()
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (info *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	info.bootstrapMethodAttrIndex = reader.readUInt16()
	info.nameAndTypeIndex = reader.readUInt16()
}
