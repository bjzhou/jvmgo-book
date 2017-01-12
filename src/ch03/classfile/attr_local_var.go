package classfile

type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableEntry
}

type LocalVariableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (attr *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	length := reader.readUInt16()
	table := make([]*LocalVariableEntry, length)
	for i := range table {
		attr.localVariableTable[i] = &LocalVariableEntry{
			startPc:         reader.readUInt16(),
			length:          reader.readUInt16(),
			nameIndex:       reader.readUInt16(),
			descriptorIndex: reader.readUInt16(),
			index:           reader.readUInt16(),
		}
	}
}
