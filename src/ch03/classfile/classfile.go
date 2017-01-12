package classfile

type ClassFile struct {
	//magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	/*	defer func() {
			if r := recover(); r != nil {
				var ok bool
				err, ok = r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
			}
		}()*/
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

func (cf *ClassFile) SuperClassName() string {
	return cf.constantPool.getClassName(cf.superClass)
}

func (cf *ClassFile) InterfaceNames() []string {
	names := make([]string, len(cf.interfaces))
	for i := range names {
		names[i] = cf.constantPool.getClassName(cf.interfaces[i])
	}
	return names
}

func (cf *ClassFile) Fields() []string {
	names := make([]string, len(cf.fields))
	for i := range names {
		names[i] = cf.fields[i].Name()
	}
	return names
}

func (cf *ClassFile) Methods() []string {
	names := make([]string, len(cf.methods))
	for i := range names {
		names[i] = cf.methods[i].Name()
	}
	return names
}

func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	cf.constantPool = cf.readConstantPool(reader)
	cf.accessFlags = reader.readUInt16()
	cf.thisClass = reader.readUInt16()
	cf.superClass = reader.readUInt16()
	cf.interfaces = reader.readUInt16s()
	cf.fields = readMembers(reader, cf.constantPool)
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool)
}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUInt32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUInt16()
	cf.majorVersion = reader.readUInt16()
	if cf.majorVersion == 45 {
		return
	} else if cf.majorVersion > 45 && cf.majorVersion <= 52 && cf.minorVersion == 0 {
		return
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (cf *ClassFile) readConstantPool(reader *ClassReader) ConstantPool {
	count := int(reader.readUInt16())
	cp := make([]ConstantInfo, count)
	for i := 1; i < count; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}