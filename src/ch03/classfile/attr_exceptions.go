package classfile

type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (attr *ExceptionsAttribute) readInfo(reader *ClassReader) {
	attr.exceptionIndexTable = reader.readUInt16s()
}

func (attr *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return attr.exceptionIndexTable
}
