package classfile

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (attr *SourceFileAttribute) readInfo(reader *ClassReader) {
	attr.sourceFileIndex = reader.readUInt16()
}

func (attr *SourceFileAttribute) FileName() string {
	return attr.cp.getUtf8(attr.sourceFileIndex)
}
