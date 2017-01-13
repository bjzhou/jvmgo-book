package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberEntry
}

type LineNumberEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (attr *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	length := reader.readUInt16()
	attr.lineNumberTable = make([]*LineNumberEntry, length)
	for i := range attr.lineNumberTable {
		attr.lineNumberTable[i] = &LineNumberEntry{
			startPc:    reader.readUInt16(),
			lineNumber: reader.readUInt16(),
		}
	}
}
