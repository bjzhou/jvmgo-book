package classfile

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (attr *CodeAttribute) readInfo(reader *ClassReader) {
	attr.maxStack = reader.readUInt16()
	attr.maxLocals = reader.readUInt16()
	length := reader.readUInt32()
	attr.code = reader.readBytes(length)
	attr.exceptionTable = readExceptionTable(reader)
	attr.attributes = readAttributes(reader, attr.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	length := reader.readUInt16()
	table := make([]*ExceptionTableEntry, length)
	for i := range table {
		table[i] = &ExceptionTableEntry{
			startPc:   reader.readUInt16(),
			endPc:     reader.readUInt16(),
			handlerPc: reader.readUInt16(),
			catchType: reader.readUInt16(),
		}
	}
	return table
}
