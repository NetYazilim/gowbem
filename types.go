package wbem

import "encoding/xml"

type CIM struct {
	XMLName     xml.Name        `xml:"CIM"`
	CimVersion  string          `xml:"CIMVERSION,attr"`
	DtdVersion  string          `xml:"DTDVERSION,attr"`
	Message     *CimMessage     `xml:"MESSAGE,omitempty"`
	Declaration *CimDeclaration `xml:"DECLARATION,omitempty"`
}

//     <!-- Section: Declaration Elements -->
//     <xs:element name="DECLARATION">
//         <xs:annotation>
//             <xs:documentation>A set of CIM schema element declarations, consisting of logical declaration groups.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice maxOccurs="unbounded">
//                 <xs:element ref="DECLGROUP"/>
//                 <xs:element ref="DECLGROUP.WITHNAME"/>
//                 <xs:element ref="DECLGROUP.WITHPATH"/>
//             </xs:choice>
//         </xs:complexType>
//     </xs:element>
type CimDeclaration struct {
	XMLName    xml.Name          `xml:"DECLARATION"`
	DeclGroups []CimAnyDeclGroup `xml:",any,omitempty"`
}

type CimAnyDeclGroup struct {
	DeclGroup         *CimDeclGroup         `xml:"DECLGROUP,omitempty"`
	DeclGroupWithName *CimDeclGroupWithName `xml:"DECLGROUP.WITHNAME,omitempty"`
	DeclGroupWithPath *CimDeclGroupWithPath `xml:"DECLGROUP.WITHPATH,omitempty"`
}

func (self *CimAnyDeclGroup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if nil != self.DeclGroup {
		return e.Encode(self.DeclGroup)
	}

	if nil != self.DeclGroupWithName {
		return e.Encode(self.DeclGroupWithName)
	}

	if nil != self.DeclGroupWithPath {
		return e.Encode(self.DeclGroupWithPath)
	}

	return nil
}

func (self *CimAnyDeclGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if "DECLGROUP" == start.Name.Local {
		self.DeclGroup = &CimDeclGroup{}
		return d.DecodeElement(self.DeclGroup, &start)
	}
	if "DECLGROUP.WITHNAME" == start.Name.Local {
		self.DeclGroupWithName = &CimDeclGroupWithName{}
		return d.DecodeElement(self.DeclGroupWithName, &start)
	}

	if "DECLGROUP.WITHPATH" == start.Name.Local {
		self.DeclGroupWithPath = &CimDeclGroupWithPath{}
		return d.DecodeElement(self.DeclGroupWithPath, &start)
	}

	return nil
}

//     <xs:element name="DECLGROUP">
//         <xs:annotation>
//             <xs:documentation>A logical group of CIM schema element declarations (classes, instances, and
// qualifier types/declarations) without any path information.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:choice minOccurs="0">
//                     <xs:element ref="LOCALNAMESPACEPATH"/>
//                     <xs:element ref="NAMESPACEPATH"/>
//                 </xs:choice>
//                 <xs:element ref="QUALIFIER.DECLARATION" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="VALUE.OBJECT" minOccurs="0" maxOccurs="unbounded"/>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimDeclGroup struct {
	XMLName               xml.Name                  `xml:"DECLGROUP"`
	LocalNamespacePath    *CimLocalNamespacePath    `xml:"LOCALNAMESPACEPATH,omitempty"`
	NamespacePath         *CimNamespacePath         `xml:"NAMESPACEPATH,omitempty"`
	QualifierDeclarations []CimQualifierDeclaration `xml:"QUALIFIER.DECLARATION,omitempty"`
	ValueObjects          []CimValueObject          `xml:"VALUE.OBJECT,omitempty"`
}

//     <xs:element name="DECLGROUP.WITHNAME">
//         <xs:annotation>
//             <xs:documentation>A logical group of CIM schema element declarations (classes, instances, and
// qualifier types/declarations) with name (local path) information.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:choice minOccurs="0">
//                     <xs:element ref="LOCALNAMESPACEPATH"/>
//                     <xs:element ref="NAMESPACEPATH"/>
//                 </xs:choice>
//                 <xs:element ref="QUALIFIER.DECLARATION" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="VALUE.NAMEDOBJECT" minOccurs="0" maxOccurs="unbounded"/>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimDeclGroupWithName struct {
	XMLName               xml.Name                  `xml:"DECLGROUP.WITHNAME"`
	LocalNamespacePath    *CimLocalNamespacePath    `xml:"LOCALNAMESPACEPATH,omitempty"`
	NamespacePath         *CimNamespacePath         `xml:"NAMESPACEPATH,omitempty"`
	QualifierDeclarations []CimQualifierDeclaration `xml:"QUALIFIER.DECLARATION,omitempty"`
	ValueNamedObjects     []CimValueNamedObject     `xml:"VALUE.NAMEDOBJECT,omitempty"`
}

//     <xs:element name="DECLGROUP.WITHPATH">
//         <xs:annotation>
//             <xs:documentation>A logical group of CIM schema element declarations (classes, instances, and
// qualifier types/declarations) with path information.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice minOccurs="0" maxOccurs="unbounded">
//                 <xs:element ref="VALUE.OBJECTWITHPATH"/>
//                 <xs:element ref="VALUE.OBJECTWITHLOCALPATH"/>
//             </xs:choice>
//         </xs:complexType>
//     </xs:element>
type CimDeclGroupWithPath struct {
	XMLName                   xml.Name                      `xml:"DECLGROUP.WITHPATH"`
	ValueObjectWithPaths      []CimValueObjectWithPath      `xml:"VALUE.OBJECTWITHPATH,omitempty"`
	ValueObjectWithLocalPaths []CimValueObjectWithLocalPath `xml:"VALUE.OBJECTWITHLOCALPATH,omitempty"`
}

//     <xs:element name="QUALIFIER.DECLARATION">
//         <xs:annotation>
//             <xs:documentation>Defines a single CIM qualifier type/declaration.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="SCOPE" minOccurs="0"/>
//                 <xs:choice minOccurs="0">
//                     <xs:element ref="VALUE"/>
//                     <xs:element ref="VALUE.ARRAY"/>
//                 </xs:choice>
//             </xs:sequence>
//             <xs:attribute ref="NAME" use="required"/>
//             <xs:attribute ref="TYPE" use="required"/>
//             <xs:attribute name="ISARRAY" type="xs:boolean" default="false"/>
//             <xs:attribute ref="ARRAYSIZE"/>
//             <xs:attributeGroup ref="QualifierFlavor"/>
//         </xs:complexType>
//     </xs:element>
type CimQualifierDeclaration struct {
	CimQualifierFlavor

	XMLName   xml.Name `xml:"QUALIFIER.DECLARATION"`
	Name      string   `xml:"NAME,attr"`
	Type      string   `xml:"TYPE,attr"`
	IsArray   bool     `xml:"ISARRAY,attr,omitempty"`
	ArraySize int      `xml:"ARRAYSIZE,attr,omitempty"`

	Scope      *CimScope      `xml:"SCOPE,omitempty"`
	Value      *CimValue      `xml:"VALUE,omitempty"`
	ValueArray *CimValueArray `xml:"VALUE.ARRAY,omitempty"`
}

//     <xs:element name="SCOPE">
//         <xs:annotation>
//             <xs:documentation>Defines the scope of a qualifier type/declaration.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:attribute name="CLASS" type="xs:boolean" default="false"/>
//             <xs:attribute name="ASSOCIATION" type="xs:boolean" default="false"/>
//             <xs:attribute name="REFERENCE" type="xs:boolean" default="false"/>
//             <xs:attribute name="PROPERTY" type="xs:boolean" default="false"/>
//             <xs:attribute name="METHOD" type="xs:boolean" default="false"/>
//             <xs:attribute name="PARAMETER" type="xs:boolean" default="false"/>
//             <xs:attribute name="INDICATION" type="xs:boolean" default="false"/>
//         </xs:complexType>
//     </xs:element>
type CimScope struct {
	XMLName     xml.Name `xml:"SCOPE"`
	Class       bool     `xml:"CLASS,attr,omitempty"`
	Association bool     `xml:"ASSOCIATION,attr,omitempty"`
	Reference   bool     `xml:"REFERENCE,attr,omitempty"`
	Property    bool     `xml:"PROPERTY,attr,omitempty"`
	Method      bool     `xml:"METHOD,attr,omitempty"`
	Parameter   bool     `xml:"PARAMETER,attr,omitempty"`
	Indication  bool     `xml:"INDICATION,attr,omitempty"`
}

//     <!-- Section: Value Elements -->
//     <xs:element name="VALUE">
//         <xs:annotation>
//             <xs:documentation>Defines a non-reference, non-NULL scalar value.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:simpleType>
//             <xs:union memberTypes="StringValue_Type CharacterValue_Type RealValue_Type BooleanValue_Type IntegerValue_Type DateTimeValue_Type"/>
//         </xs:simpleType>
//     </xs:element>
type CimValue struct {
	XMLName xml.Name `xml:"VALUE"`
	Value   string   `xml:",chardata"`
}

//     <xs:element name="VALUE.ARRAY">
//         <xs:annotation>
//             <xs:documentation>Defines an non-reference array value.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice minOccurs="0" maxOccurs="unbounded">
//                 <xs:element ref="VALUE"/>
//                 <xs:element ref="VALUE.NULL"/>
//             </xs:choice>
//         </xs:complexType>
//     </xs:element>
// type CimValueArray struct {
// 	XMLName xml.Name         `xml:"VALUE.ARRAY"`
// 	Values  []CimValueOrNull //`xml:,omitempty`
// }
type CimValueArray []CimValueOrNull

type CimValueOrNull struct {
	Value *CimValue     `xml:"VALUE,omitempty"`
	Null  *CimValueNull `xml:"VALUE.NULL,omitempty"`
}

//     <xs:element name="VALUE.REFERENCE">
//         <xs:annotation>
//             <xs:documentation>Defines a non-NULL reference scalar value.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice>
//                 <xs:element ref="CLASSPATH"/>
//                 <xs:element ref="LOCALCLASSPATH"/>
//                 <xs:element ref="CLASSNAME"/>
//                 <xs:element ref="INSTANCEPATH"/>
//                 <xs:element ref="LOCALINSTANCEPATH"/>
//                 <xs:element ref="INSTANCENAME"/>
//             </xs:choice>
//         </xs:complexType>
//     </xs:element>
type CimValueReference struct {
	XMLName           xml.Name              `xml:"VALUE.REFERENCE"`
	ClassPath         *CimClassPath         `xml:"CLASSPATH,omitempty"`
	LocalClassPath    *CimLocalClassPath    `xml:"LOCALCLASSPATH,omitempty"`
	ClassName         *CimClassName         `xml:"CLASSNAME,omitempty"`
	InstancePath      *CimInstancePath      `xml:"INSTANCEPATH,omitempty"`
	LocalInstancePath *CimLocalInstancePath `xml:"LOCALINSTANCEPATH,omitempty"`
	InstanceName      *CimInstanceName      `xml:"INSTANCENAME,omitempty"`
}

//     <xs:element name="VALUE.REFARRAY">
//         <xs:annotation>
//             <xs:documentation>Defines a reference array value.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice minOccurs="0" maxOccurs="unbounded">
//                 <xs:element ref="VALUE.REFERENCE"/>
//                 <xs:element ref="VALUE.NULL"/>
//             </xs:choice>
//         </xs:complexType>
//     </xs:element>
// type CimValueRefArray struct {
// 	XMLName xml.Name                  `xml:"VALUE.REFARRAY"`
// 	Values  []CimValueReferenceOrNull //`xml:,omitempty`
// }
type CimValueRefArray []CimValueReferenceOrNull

type CimValueReferenceOrNull struct {
	Value *CimValueReference
	Null  *CimValueNull
}

//     <xs:element name="VALUE.OBJECT">
//         <xs:annotation>
//             <xs:documentation>Defines a value that comprises a CIM object (class or instance) definition.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice>
//                 <xs:element ref="CLASS"/>
//                 <xs:element ref="INSTANCE"/>
//             </xs:choice>
//         </xs:complexType>
//     </xs:element>
type CimValueObject struct {
	XMLName  xml.Name     `xml:"VALUE.OBJECT"`
	Instance *CimInstance `xml:"INSTANCE,omitempty"`
	Class    *CimClass    `xml:"CLASS,omitempty"`
}

//     <xs:element name="VALUE.NAMEDINSTANCE">
//         <xs:annotation>
//             <xs:documentation>Defines a value that comprises a named CIM instance definition.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="INSTANCENAME"/>
//                 <xs:element ref="INSTANCE"/>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimValueNamedInstance struct {
	XMLName  xml.Name        `xml:"VALUE.NAMEDINSTANCE"`
	Class    CimInstanceName `xml:"INSTANCENAME"`
	Instance CimInstance     `xml:"INSTANCE"`
}

//     <xs:element name="VALUE.NAMEDOBJECT">
//         <xs:annotation>
//             <xs:documentation>Defines a value that comprises a named CIM object (class or instance) definition.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice>
//                 <xs:element ref="CLASS"/>
//                 <xs:sequence>
//                     <xs:element ref="INSTANCENAME"/>
//                     <xs:element ref="INSTANCE"/>
//                 </xs:sequence>
//             </xs:choice>
//         </xs:complexType>
//     </xs:element>
type CimValueNamedObject struct {
	XMLName      xml.Name         `xml:"VALUE.NAMEDOBJECT"`
	Class        *CimClass        `xml:"CLASS,omitempty"`
	InstanceName *CimInstanceName `xml:"INSTANCENAME,omitempty"`
	Instance     *CimInstance     `xml:"INSTANCE,omitempty"`
}

//     <xs:element name="VALUE.OBJECTWITHPATH">
//         <xs:annotation>
//             <xs:documentation>Defines a value that comprises a CIM object (class or instance) definition with additional
// information that defines the absolute path to that object.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice>
//                 <xs:sequence>
//                     <xs:element ref="CLASSPATH"/>
//                     <xs:element ref="CLASS"/>
//                 </xs:sequence>
//                 <xs:sequence>
//                     <xs:element ref="INSTANCEPATH"/>
//                     <xs:element ref="INSTANCE"/>
//                 </xs:sequence>
//             </xs:choice>
//         </xs:complexType>
//     </xs:element>
type CimValueObjectWithPath struct {
	XMLName      xml.Name         `xml:"VALUE.OBJECTWITHPATH"`
	ClassPath    *CimClassPath    `xml:"CLASSPATH,omitempty"`
	Class        *CimClass        `xml:"CLASS,omitempty"`
	InstancePath *CimInstancePath `xml:"INSTANCEPATH,omitempty"`
	Instance     *CimInstance     `xml:"INSTANCE,omitempty"`
}

//     <xs:element name="VALUE.OBJECTWITHLOCALPATH">
//         <xs:annotation>
//             <xs:documentation>Defines a value that comprises a CIM object (class or instance) definition with additional
// information that defines the local path to that object.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice>
//                 <xs:sequence>
//                     <xs:element ref="LOCALCLASSPATH"/>
//                     <xs:element ref="CLASS"/>
//                 </xs:sequence>
//                 <xs:sequence>
//                     <xs:element ref="LOCALINSTANCEPATH"/>
//                     <xs:element ref="INSTANCE"/>
//                 </xs:sequence>
//             </xs:choice>
//         </xs:complexType>
//     </xs:element>
type CimValueObjectWithLocalPath struct {
	XMLName      xml.Name              `xml:"VALUE.OBJECTWITHLOCALPATH"`
	ClassPath    *CimLocalClassPath    `xml:"LOCALCLASSPATH,omitempty"`
	Class        *CimClass             `xml:"CLASS,omitempty"`
	InstancePath *CimLocalInstancePath `xml:"LOCALINSTANCEPATH,omitempty"`
	Instance     *CimInstance          `xml:"INSTANCE,omitempty"`
}

//     <xs:element name="VALUE.NULL">
//         <xs:annotation>
//             <xs:documentation>Defines the NULL value.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType/>
//     </xs:element>
type CimValueNull struct {
	XMLName xml.Name `xml:"VALUE.NULL"`
}

//     <xs:element name="VALUE.INSTANCEWITHPATH">
//         <xs:annotation>
//             <xs:documentation>Defines a value that comprises a CIM instance definition with additional information that
// defines the absolute path to that object.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="INSTANCEPATH"/>
//                 <xs:element ref="INSTANCE"/>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimValueInstanceWithPath struct {
	XMLName      xml.Name        `xml:"VALUE.INSTANCEWITHPATH"`
	InstancePath CimInstancePath `xml:"INSTANCEPATH"`
	Instance     CimInstance     `xml:"INSTANCE"`
}

//     <!-- Section: Naming and Location Elements -->
//     <xs:element name="NAMESPACEPATH">
//         <xs:annotation>
//             <xs:documentation>Defines an (absolute) namespace path.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="HOST"/>
//                 <xs:element ref="LOCALNAMESPACEPATH"/>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimNamespacePath struct {
	XMLName            xml.Name              `xml:"NAMESPACEPATH"`
	Host               CimHost               `xml:"HOST"`
	LocalNamespacePath CimLocalNamespacePath `xml:"LOCALNAMESPACEPATH"`
}

//     <xs:element name="LOCALNAMESPACEPATH">
//         <xs:annotation>
//             <xs:documentation>Defines a local namespace path (one without a host component).
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="NAMESPACE" maxOccurs="unbounded"/>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimLocalNamespacePath struct {
	XMLName   xml.Name       `xml:"LOCALNAMESPACEPATH"`
	Namespace []CimNamespace `xml:"NAMESPACE"`
}

//     <xs:element name="HOST">
//         <xs:annotation>
//             <xs:documentation>Defines a host name and optionally a port number.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:simpleType>
//             <xs:restriction base="xs:string"/>
//         </xs:simpleType>
//     </xs:element>
type CimHost struct {
	XMLName xml.Name `xml:"HOST"`
	Value   string   `xml:",chardata"`
}

//     <xs:element name="NAMESPACE">
//         <xs:annotation>
//             <xs:documentation>Defines a single namespace within the namespace component of a namespace path.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:attribute ref="NAME" use="required"/>
//         </xs:complexType>
//     </xs:element>
type CimNamespace struct {
	XMLName xml.Name `xml:"NAMESPACE"`
	Name    string   `xml:"NAME,attr"`
}

//     <xs:element name="CLASSPATH">
//         <xs:annotation>
//             <xs:documentation>Defines the absolute path to a CIM class.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="NAMESPACEPATH"/>
//                 <xs:element ref="CLASSNAME"/>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimClassPath struct {
	XMLName       xml.Name         `xml:"CLASSPATH"`
	NamespacePath CimNamespacePath `xml:"NAMESPACEPATH"`
	ClassName     CimClassName     `xml:"CLASSNAME"`
}

//     <xs:element name="LOCALCLASSPATH">
//         <xs:annotation>
//             <xs:documentation>Defines the local path to a CIM class.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="LOCALNAMESPACEPATH"/>
//                 <xs:element ref="CLASSNAME"/>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimLocalClassPath struct {
	XMLName       xml.Name              `xml:"LOCALCLASSPATH"`
	NamespacePath CimLocalNamespacePath `xml:"LOCALNAMESPACEPATH"`
	ClassName     CimClassName          `xml:"CLASSNAME"`
}

//     <xs:element name="CLASSNAME">
//         <xs:annotation>
//             <xs:documentation>Defines the name of a CIM class.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:attribute name="NAME" type="CIMClassName_Type" use="required"/>
//         </xs:complexType>
//     </xs:element>
type CimClassName struct {
	XMLName xml.Name `xml:"CLASSNAME"`
	Name    string   `xml:"NAME,attr"`
}

//     <xs:element name="INSTANCEPATH">
//         <xs:annotation>
//             <xs:documentation>Defines the absolute path to a CIM instance.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="NAMESPACEPATH"/>
//                 <xs:element ref="INSTANCENAME"/>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimInstancePath struct {
	XMLName       xml.Name         `xml:"INSTANCEPATH"`
	NamespacePath CimNamespacePath `xml:"NAMESPACEPATH"`
	InstanceName  CimInstanceName  `xml:"INSTANCENAME"`
}

//     <xs:element name="LOCALINSTANCEPATH">
//         <xs:annotation>
//             <xs:documentation>Defines the local path to a CIM instance.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="LOCALNAMESPACEPATH"/>
//                 <xs:element ref="INSTANCENAME"/>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimLocalInstancePath struct {
	XMLName            xml.Name              `xml:"LOCALINSTANCEPATH"`
	LocalNamespacePath CimLocalNamespacePath `xml:"LOCALNAMESPACEPATH"`
	InstanceName       CimInstanceName       `xml:"INSTANCENAME"`
}

//     <xs:element name="INSTANCENAME">
//         <xs:annotation>
//             <xs:documentation>Defines the location of a CIM instance within a namespace (it is referred to in DSP0004
// as a model path). For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice>
//                 <xs:element ref="KEYBINDING" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="KEYVALUE" minOccurs="0"/>
//                 <xs:element ref="VALUE.REFERENCE" minOccurs="0"/>
//             </xs:choice>
//             <xs:attribute ref="CLASSNAME" use="required"/>
//         </xs:complexType>
//     </xs:element>
type CimInstanceName struct {
	XMLName        xml.Name           `xml:"INSTANCENAME"`
	ClassName      string             `xml:"CLASSNAME,attr"`
	KeyBindings    []CimKeyBinding    `xml:"KEYBINDING,omitempty"`
	KeyValue       *CimKeyValue       `xml:"KEYVALUE,omitempty"`
	ValueReference *CimValueReference `xml:"VALUE.REFERENCE,omitempty"`
}

//     <xs:element name="OBJECTPATH">
//         <xs:annotation>
//             <xs:documentation>Defines the full path to a single CIM object (class or instance).
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice>
//                 <xs:element ref="INSTANCEPATH"/>
//                 <xs:element ref="CLASSPATH"/>
//             </xs:choice>
//         </xs:complexType>
//     </xs:element>
type CimObjectPath struct {
	XMLName      xml.Name         `xml:"OBJECTPATH"`
	InstancePath *CimInstancePath `xml:"INSTANCEPATH,omitempty"`
	ClassPath    *CimClassPath    `xml:"CLASSPATH,omitempty"`
}

//     <xs:element name="KEYBINDING">
//         <xs:annotation>
//             <xs:documentation>Defines a key binding (= key property name and value used in an instance path).
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice>
//                 <xs:element ref="KEYVALUE"/>
//                 <xs:element ref="VALUE.REFERENCE"/>
//             </xs:choice>
//             <xs:attribute ref="NAME" use="required"/>
//         </xs:complexType>
//     </xs:element>
type CimKeyBinding struct {
	XMLName        xml.Name           `xml:"KEYBINDING"`
	Name           string             `xml:"NAME,attr"`
	KeyValue       *CimKeyValue       `xml:"KEYVALUE,omitempty"`
	ValueReference *CimValueReference `xml:"VALUE.REFERENCE,omitempty"`
}

//     <xs:element name="KEYVALUE">
//         <xs:annotation>
//             <xs:documentation>Defines the value of a non-reference (and scalar) key binding.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType mixed="true">
//             <xs:attribute name="VALUETYPE" default="string">
//                 <xs:simpleType>
//                     <xs:restriction base="xs:NMTOKEN">
//                         <xs:enumeration value="string"/>
//                         <xs:enumeration value="boolean"/>
//                         <xs:enumeration value="numeric"/>
//                     </xs:restriction>
//                 </xs:simpleType>
//             </xs:attribute>
//             <xs:attribute ref="TYPE" use="required"/>
//         </xs:complexType>
//     </xs:element>
type CimKeyValue struct {
	XMLName   xml.Name `xml:"KEYVALUE"`
	ValueType string   `xml:"VALUETYPE,attr,omitempty"`
	Type      string   `xml:"TYPE,attr,omitempty"`
	Value     string   `xml:",chardata"`
}

//     <!-- Section: Object Definition Elements -->
//     <xs:element name="CLASS">
//         <xs:annotation>
//             <xs:documentation>Defines a CIM class.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="QUALIFIER" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:choice minOccurs="0" maxOccurs="unbounded">
//                     <xs:element ref="PROPERTY"/>
//                     <xs:element ref="PROPERTY.ARRAY"/>
//                     <xs:element ref="PROPERTY.REFERENCE"/>
//                 </xs:choice>
//                 <xs:element ref="METHOD" minOccurs="0" maxOccurs="unbounded"/>
//             </xs:sequence>
//             <xs:attribute ref="NAME" use="required"/>
//             <xs:attribute ref="SUPERCLASS"/>
//         </xs:complexType>
//     </xs:element>
type CimClass struct {
	XMLName    xml.Name         `xml:"CLASS"`
	Name       string           `xml:"NAME,attr"`
	SuperClass string           `xml:"SUPERCLASS,attr,omitempty"`
	Qualifiers []CimQualifier   `xml:"QUALIFIER,omitempty"`
	Properties []CimAnyProperty `xml:",any,omitempty"`
	Methods    []CimMethod      `xml:"METHOD,omitempty"`
}

type CimAnyProperty struct {
	Property          *CimProperty          `xml:"PROPERTY,omitempty"`
	PropertyArray     *CimPropertyArray     `xml:"PROPERTY.ARRAY,omitempty"`
	PropertyReference *CimPropertyReference `xml:"PROPERTY.REFERENCE,omitempty"`
}

func (self *CimAnyProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if nil != self.Property {
		return e.Encode(self.Property)
	}

	if nil != self.PropertyArray {
		return e.Encode(self.PropertyArray)
	}

	if nil != self.PropertyReference {
		return e.Encode(self.PropertyReference)
	}
	return nil
}

func (self *CimAnyProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if "PROPERTY" == start.Name.Local {
		self.Property = &CimProperty{}
		return d.DecodeElement(self.Property, &start)
	}

	if "PROPERTY.ARRAY" == start.Name.Local {
		self.PropertyArray = &CimPropertyArray{}
		return d.DecodeElement(self.PropertyArray, &start)
	}

	if "PROPERTY.REFERENCE" == start.Name.Local {
		self.PropertyReference = &CimPropertyReference{}
		return d.DecodeElement(self.PropertyReference, &start)
	}
	return nil
}

//     <xs:element name="INSTANCE">
//         <xs:annotation>
//             <xs:documentation>Defines a CIM instance.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="QUALIFIER" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:choice minOccurs="0" maxOccurs="unbounded">
//                     <xs:element ref="PROPERTY"/>
//                     <xs:element ref="PROPERTY.ARRAY"/>
//                     <xs:element ref="PROPERTY.REFERENCE"/>
//                 </xs:choice>
//             </xs:sequence>
//             <xs:attribute ref="CLASSNAME" use="required"/>
//             <xs:attribute ref="xml:lang"/>
//         </xs:complexType>
//     </xs:element>
type CimInstance struct {
	XMLName    xml.Name         `xml:"INSTANCE"`
	ClassName  string           `xml:"CLASSNAME,attr"`
	Lang       string           `xml:"xml lang,attr,omitempty"`
	Qualifiers []CimQualifier   `xml:"QUALIFIER,omitempty"`
	Properties []CimAnyProperty `xml:",any,omitempty"`
}

//     <xs:element name="QUALIFIER">
//         <xs:annotation>
//             <xs:documentation>Defines a CIM qualifier value.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:choice minOccurs="0">
//                     <xs:element ref="VALUE"/>
//                     <xs:element ref="VALUE.ARRAY"/>
//                 </xs:choice>
//             </xs:sequence>
//             <xs:attribute ref="NAME" use="required"/>
//             <xs:attribute ref="TYPE" use="required"/>
//             <xs:attribute ref="PROPAGATED"/>
//             <xs:attributeGroup ref="QualifierFlavor"/>
//             <xs:attribute ref="xml:lang"/>
//         </xs:complexType>
//     </xs:element>
type CimQualifier struct {
	CimQualifierFlavor

	XMLName    xml.Name      `xml:"QUALIFIER"`
	Name       string        `xml:"NAME,attr"`
	Type       string        `xml:"TYPE,attr"`
	Propagated bool          `xml:"PROPAGATED,attr,omitempty"`
	Lang       string        `xml:"xml lang,attr,omitempty"`
	Value      *CimValue     `xml:"VALUE,omitempty"`
	ValueArray CimValueArray `xml:"VALUE.ARRAY,omitempty"`
}

//     <xs:attributeGroup name="QualifierFlavor">
//         <xs:annotation>
//             <xs:documentation>Defines the flavor settings for a CIM qualifier declaration;
// this attribute group corresponds to the %QualifierFlavor entity in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:attribute name="OVERRIDABLE" type="xs:boolean" default="true"/>
//         <xs:attribute name="TOSUBCLASS" type="xs:boolean" default="true"/>
//         <xs:attribute name="TOINSTANCE" type="xs:boolean" default="false"/>
//         <xs:attribute name="TRANSLATABLE" type="xs:boolean" default="false"/>
//     </xs:attributeGroup>
type CimQualifierFlavor struct {
	Overridable  bool `xml:"OVERRIDABLE,attr,omitempty"`
	ToSubclass   bool `xml:"TOSUBCLASS,attr,omitempty"`
	ToInstance   bool `xml:"TOINSTANCE,attr,omitempty"`
	Translatable bool `xml:"TRANSLATABLE,attr,omitempty"`
}

//     <xs:element name="PROPERTY">
//         <xs:annotation>
//             <xs:documentation>Defines a non-reference scalar property, that is used as a property value in a CIM instance
// or as a property declaration in a CIM class.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="QUALIFIER" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="VALUE" minOccurs="0"/>
//             </xs:sequence>
//             <xs:attribute ref="NAME" use="required"/>
//             <xs:attribute ref="TYPE" use="required"/>
//             <xs:attribute ref="CLASSORIGIN"/>
//             <xs:attribute ref="PROPAGATED"/>
//             <xs:attribute ref="EmbeddedObject"/>
//             <xs:attribute ref="xml:lang"/>
//         </xs:complexType>
//     </xs:element>
type CimProperty struct {
	XMLName        xml.Name       `xml:"PROPERTY"`
	Name           string         `xml:"NAME,attr"`
	Type           string         `xml:"TYPE,attr"`
	ClassOrigin    string         `xml:"CLASSORIGIN,attr,omitempty"`
	Propagated     bool           `xml:"PROPAGATED,attr,omitempty"`
	EmbeddedObject string         `xml:"EmbeddedObject,attr,omitempty"`
	Lang           string         `xml:"xml lang,attr,omitempty"`
	Qualifiers     []CimQualifier `xml:"QUALIFIER",omitempty`
	Value          *CimValue      `xml:"VALUE,omitempty"`
}

//     <xs:element name="PROPERTY.ARRAY">
//         <xs:annotation>
//             <xs:documentation>Defines a non-reference array property, that is used as a property value in a CIM instance
// or as a property declaration in a CIM class.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="QUALIFIER" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="VALUE.ARRAY" minOccurs="0"/>
//             </xs:sequence>
//             <xs:attribute ref="NAME" use="required"/>
//             <xs:attribute ref="TYPE" use="required"/>
//             <xs:attribute ref="ARRAYSIZE"/>
//             <xs:attribute ref="CLASSORIGIN"/>
//             <xs:attribute ref="PROPAGATED"/>
//             <xs:attribute ref="EmbeddedObject"/>
//             <xs:attribute ref="xml:lang"/>
//         </xs:complexType>
//     </xs:element>
type CimPropertyArray struct {
	XMLName        xml.Name       `xml:"PROPERTY.ARRAY"`
	Name           string         `xml:"NAME,attr"`
	Type           string         `xml:"TYPE,attr"`
	ArraySize      int            `xml:"ARRAYSIZE,attr,omitempty"`
	ClassOrigin    string         `xml:"CLASSORIGIN,attr,omitempty"`
	Propagated     bool           `xml:"PROPAGATED,attr,omitempty"`
	EmbeddedObject string         `xml:"EmbeddedObject,attr,omitempty"`
	Lang           string         `xml:"xml lang,attr,omitempty"`
	Qualifiers     []CimQualifier `xml:"QUALIFIER,omitempty"`
	ValueArray     CimValueArray  `xml:"VALUE.ARRAY,omitempty"`
}

//     <xs:element name="PROPERTY.REFERENCE">
//         <xs:annotation>
//             <xs:documentation>Defines a scalar reference property, that is used as a property value in a CIM instance
// or as a property declaration in a CIM class.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="QUALIFIER" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="VALUE.REFERENCE" minOccurs="0"/>
//             </xs:sequence>
//             <xs:attribute ref="NAME" use="required"/>
//             <xs:attribute ref="REFERENCECLASS"/>
//             <xs:attribute ref="CLASSORIGIN"/>
//             <xs:attribute ref="PROPAGATED"/>
//         </xs:complexType>
//     </xs:element>
type CimPropertyReference struct {
	XMLName        xml.Name           `xml:"PROPERTY.REFERENCE"`
	Name           string             `xml:"NAME,attr"`
	ReferenceClass string             `xml:"REFERENCECLASS,attr,omitempty"`
	ClassOrigin    string             `xml:"CLASSORIGIN,attr,omitempty"`
	Propagated     bool               `xml:"PROPAGATED,attr,omitempty"`
	Qualifiers     []CimQualifier     `xml:"QUALIFIER,omitempty"`
	ValueReference *CimValueReference `xml:"VALUE.REFERENCE,omitempty"`
}

//     <xs:element name="METHOD">
//         <xs:annotation>
//             <xs:documentation>Defines a CIM method within a class declaration.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="QUALIFIER" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:choice minOccurs="0" maxOccurs="unbounded">
//                     <xs:element ref="PARAMETER"/>
//                     <xs:element ref="PARAMETER.REFERENCE"/>
//                     <xs:element ref="PARAMETER.ARRAY"/>
//                     <xs:element ref="PARAMETER.REFARRAY"/>
//                 </xs:choice>
//             </xs:sequence>
//             <xs:attribute ref="NAME" use="required"/>
//             <xs:attribute ref="TYPE"/>
//             <xs:attribute ref="CLASSORIGIN"/>
//             <xs:attribute ref="PROPAGATED"/>
//         </xs:complexType>
//     </xs:element>
type CimMethod struct {
	XMLName     xml.Name          `xml:"METHOD"`
	Name        string            `xml:"NAME,attr"`
	Type        string            `xml:"TYPE,attr,omitempty"`
	ClassOrigin string            `xml:"CLASSORIGIN,attr,omitempty"`
	Propagated  bool              `xml:"PROPAGATED,attr,omitempty"`
	Qualifiers  []CimQualifier    `xml:"QUALIFIER,omitempty"`
	Parameters  []CimAnyParameter `xml:",any,omitempty"`
}

type CimAnyParameter struct {
	Parameter          *CimParameter          `xml:"PARAMETER,omitempty"`
	ParameterReference *CimParameterReference `xml:"PARAMETER.REFERENCE,omitempty"`
	ParameterArray     *CimParameterArray     `xml:"PARAMETER.ARRAY,omitempty"`
	ParameterRefArray  *CimParameterRefArray  `xml:"PARAMETER.REFARRAY,omitempty"`
}

func (self *CimAnyParameter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if nil != self.Parameter {
		return e.Encode(self.Parameter)
	}

	if nil != self.ParameterReference {
		return e.Encode(self.ParameterReference)
	}

	if nil != self.ParameterArray {
		return e.Encode(self.ParameterArray)
	}

	if nil != self.ParameterRefArray {
		return e.Encode(self.ParameterRefArray)
	}
	return nil
}

func (self *CimAnyParameter) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if "PARAMETER" == start.Name.Local {
		self.Parameter = &CimParameter{}
		return d.DecodeElement(self.Parameter, &start)
	}
	if "PARAMETER.REFERENCE" == start.Name.Local {
		self.ParameterReference = &CimParameterReference{}
		return d.DecodeElement(self.ParameterReference, &start)
	}

	if "PARAMETER.ARRAY" == start.Name.Local {
		self.ParameterArray = &CimParameterArray{}
		return d.DecodeElement(self.ParameterArray, &start)
	}

	if "PARAMETER.REFARRAY" == start.Name.Local {
		self.ParameterRefArray = &CimParameterRefArray{}
		return d.DecodeElement(self.ParameterRefArray, &start)
	}

	return nil
}

//     <xs:element name="PARAMETER">
//         <xs:annotation>
//             <xs:documentation>Defines a non-reference scalar CIM parameter within a method in a class declaration.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="QUALIFIER" minOccurs="0" maxOccurs="unbounded"/>
//             </xs:sequence>
//             <xs:attribute ref="NAME" use="required"/>
//             <xs:attribute ref="TYPE" use="required"/>
//         </xs:complexType>
//     </xs:element>
type CimParameter struct {
	XMLName    xml.Name       `xml:"PARAMETER"`
	Name       string         `xml:"NAME,attr"`
	Type       string         `xml:"TYPE,attr"`
	Qualifiers []CimQualifier `xml:"QUALIFIER,omitempty"`
}

//     <xs:element name="PARAMETER.REFERENCE">
//         <xs:annotation>
//             <xs:documentation>Defines a reference-typed scalar CIM parameter within a method in a class declaration.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="QUALIFIER" minOccurs="0" maxOccurs="unbounded"/>
//             </xs:sequence>
//             <xs:attribute ref="NAME" use="required"/>
//             <xs:attribute ref="REFERENCECLASS"/>
//         </xs:complexType>
//     </xs:element>
type CimParameterReference struct {
	XMLName        xml.Name       `xml:"PARAMETER.REFERENCE"`
	Name           string         `xml:"NAME,attr"`
	ReferenceClass string         `xml:"REFERENCECLASS,attr,omitempty"`
	Qualifiers     []CimQualifier `xml:"QUALIFIER,omitempty"`
}

//     <xs:element name="PARAMETER.ARRAY">
//         <xs:annotation>
//             <xs:documentation>Defines a non-reference array CIM parameter within a method in a class declaration.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="QUALIFIER" minOccurs="0" maxOccurs="unbounded"/>
//             </xs:sequence>
//             <xs:attribute ref="NAME" use="required"/>
//             <xs:attribute ref="TYPE" use="required"/>
//             <xs:attribute ref="ARRAYSIZE"/>
//         </xs:complexType>
//     </xs:element>
type CimParameterArray struct {
	XMLName    xml.Name       `xml:"PARAMETER.ARRAY"`
	Name       string         `xml:"NAME,attr"`
	Type       string         `xml:"TYPE,attr"`
	ArraySize  int            `xml:"ARRAYSIZE,attr,omitempty"`
	Qualifiers []CimQualifier `xml:"QUALIFIER,omitempty"`
}

//     <xs:element name="PARAMETER.REFARRAY">
//         <xs:annotation>
//             <xs:documentation>Defines a reference-typed array CIM parameter within a method in a class declaration.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="QUALIFIER" minOccurs="0" maxOccurs="unbounded"/>
//             </xs:sequence>
//             <xs:attribute ref="NAME" use="required"/>
//             <xs:attribute ref="REFERENCECLASS"/>
//             <xs:attribute ref="ARRAYSIZE"/>
//         </xs:complexType>
//     </xs:element>
type CimParameterRefArray struct {
	XMLName        xml.Name       `xml:"PARAMETER.REFARRAY"`
	Name           string         `xml:"NAME,attr"`
	ReferenceClass string         `xml:"REFERENCECLASS,attr"`
	ArraySize      int            `xml:"ARRAYSIZE,attr,omitempty"`
	Qualifiers     []CimQualifier `xml:"QUALIFIER,omitempty"`
}

//     <!-- Section: Message Elements -->
//     <xs:element name="MESSAGE">
//         <xs:annotation>
//             <xs:documentation>Defines a CIM message in the CIM-XML protocol.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice>
//                 <xs:element ref="SIMPLEREQ"/>
//                 <xs:element ref="MULTIREQ"/>
//                 <xs:element ref="SIMPLERSP"/>
//                 <xs:element ref="MULTIRSP"/>
//                 <xs:element ref="SIMPLEEXPREQ"/>
//                 <xs:element ref="MULTIEXPREQ"/>
//                 <xs:element ref="SIMPLEEXPRSP"/>
//                 <xs:element ref="MULTIEXPRSP"/>
//             </xs:choice>
//             <xs:attribute name="ID" type="xs:string" use="required"/>
//             <xs:attribute name="PROTOCOLVERSION" type="VersionMN_Type" use="required"/>
//         </xs:complexType>
//     </xs:element>
type CimMessage struct {
	XMLName         xml.Name         `xml:"MESSAGE"`
	Id              string           `xml:"ID,attr"`
	ProtocolVersion string           `xml:"PROTOCOLVERSION,attr"`
	SimpleReq       *CimSimpleReq    `xml:"SIMPLEREQ,omitempty"`
	MultiReq        *CimMultiReq     `xml:"MULTIREQ,omitempty"`
	SimpleRsp       *CimSimpleRsp    `xml:"SIMPLERSP,omitempty"`
	MultiRsp        *CimMultiRsp     `xml:"MULTIRSP,omitempty"`
	SimpleExpReq    *CimSimpleExpReq `xml:"SIMPLEEXPREQ,omitempty"`
	MultiExpReq     *CimMultiExpReq  `xml:"MULTIEXPREQ,omitempty"`
	SimpleExpRsp    *CimSimpleExpRsp `xml:"SIMPLEEXPRSP,omitempty"`
	MultiExpRsp     *CimMultiExpRsp  `xml:"MULTIEXPRSP,omitempty"`
}

//     <xs:element name="MULTIREQ">
//         <xs:annotation>
//             <xs:documentation>Defines a multiple CIM operation request.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="SIMPLEREQ" minOccurs="2" maxOccurs="unbounded"/>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimMultiReq struct {
	XMLName    xml.Name       `xml:"MULTIREQ"`
	SimpleReqs []CimSimpleReq `xml:"SIMPLEREQ"`
}

//     <xs:element name="SIMPLEREQ">
//         <xs:annotation>
//             <xs:documentation>Defines a simple CIM operation request.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="CORRELATOR" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:choice>
//                     <xs:element ref="METHODCALL"/>
//                     <xs:element ref="IMETHODCALL"/>
//                 </xs:choice>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimSimpleReq struct {
	XMLName     xml.Name        `xml:"SIMPLEREQ"`
	Correlators []CimCorrelator `xml:"CORRELATOR,omitempty"`
	MethodCall  *CimMethodCall  `xml:"METHODCALL,omitempty"`
	IMethodCall *CimIMethodCall `xml:"IMETHODCALL,omitempty"`
}

//     <xs:element name="METHODCALL">
//         <xs:annotation>
//             <xs:documentation>Defines a single extrinsic (= class-defined) method invocation request.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:choice>
//                     <xs:element ref="LOCALCLASSPATH"/>
//                     <xs:element ref="LOCALINSTANCEPATH"/>
//                 </xs:choice>
//                 <xs:element ref="PARAMVALUE" minOccurs="0" maxOccurs="unbounded"/>
//             </xs:sequence>
//             <xs:attribute ref="NAME" use="required"/>
//         </xs:complexType>
//     </xs:element>
type CimMethodCall struct {
	XMLName           xml.Name              `xml:"METHODCALL"`
	Name              string                `xml:"NAME,attr"`
	LocalClassPath    *CimLocalClassPath    `xml:"LOCALCLASSPATH,omitempty"`
	LocalInstancePath *CimLocalInstancePath `xml:"LOCALINSTANCEPATH,omitempty"`
	ParamValues       []CimParamValue       `xml:"PARAMVALUE,omitempty"`
}

//     <xs:element name="PARAMVALUE">
//         <xs:annotation>
//             <xs:documentation>Defines a parameter value that is used in an extrinsic (= class defined) and - for
// historical reasons - also in an intrinsic method (= operation) request or response.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice minOccurs="0">
//                 <xs:element ref="VALUE"/>
//                 <xs:element ref="VALUE.REFERENCE"/>
//                 <xs:element ref="VALUE.ARRAY"/>
//                 <xs:element ref="VALUE.REFARRAY"/>
//                 <xs:element ref="CLASSNAME"/>
//                 <xs:element ref="INSTANCENAME"/>
//                 <xs:element ref="CLASS"/>
//                 <xs:element ref="INSTANCE"/>
//                 <xs:element ref="VALUE.NAMEDINSTANCE"/>
//             </xs:choice>
//             <xs:attribute ref="NAME" use="required"/>
//             <xs:attribute ref="PARAMTYPE"/>
//             <xs:attribute ref="EmbeddedObject"/>
//         </xs:complexType>
//     </xs:element>
type CimParamValue struct {
	XMLName            xml.Name               `xml:"PARAMVALUE"`
	Name               string                 `xml:"NAME,attr"`
	ParamType          string                 `xml:"PARAMTYPE,attr,omitempty"`
	EmbeddedObject     string                 `xml:"EmbeddedObject,attr,omitempty"`
	Value              *CimValue              `xml:"VALUE,omitempty"`
	ValueReference     *CimValueReference     `xml:"VALUE.REFERENCE,omitempty"`
	ValueArray         CimValueArray          `xml:"VALUE.ARRAY,omitempty"`
	ValueRefArray      CimValueRefArray       `xml:"VALUE.REFARRAY,omitempty"`
	ClassName          *CimClassName          `xml:"CLASSNAME,omitempty"`
	InstanceName       *CimInstanceName       `xml:"INSTANCENAME,omitempty"`
	Class              *CimClass              `xml:"CLASS,omitempty"`
	Instance           *CimInstance           `xml:"INSTANCE,omitempty"`
	ValueNamedInstance *CimValueNamedInstance `xml:"VALUE.NAMEDINSTANCE,omitempty"`
}

//     <xs:element name="IMETHODCALL">
//         <xs:annotation>
//             <xs:documentation>Defines a single intrinsic method (=operation) invocation request.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="LOCALNAMESPACEPATH"/>
//                 <xs:element ref="IPARAMVALUE" minOccurs="0" maxOccurs="unbounded"/>
//             </xs:sequence>
//             <xs:attribute ref="NAME" use="required"/>
//         </xs:complexType>
//     </xs:element>
type CimIMethodCall struct {
	XMLName            xml.Name              `xml:"IMETHODCALL"`
	Name               string                `xml:"NAME,attr"`
	LocalNamespacePath CimLocalNamespacePath `xml:"LOCALNAMESPACEPATH"`
	ParamValues        []CimIParamValue      `xml:"IPARAMVALUE,omitempty"`
}

//     <xs:element name="IPARAMVALUE">
//         <xs:annotation>
//             <xs:documentation>Defines a parameter value that is used in a intrinsic method (= operation) request or response.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice minOccurs="0">
//                 <xs:element ref="VALUE"/>
//                 <xs:element ref="VALUE.ARRAY"/>
//                 <xs:element ref="VALUE.REFERENCE"/>
//                 <xs:element ref="CLASSNAME"/>
//                 <xs:element ref="INSTANCENAME"/>
//                 <xs:element ref="QUALIFIER.DECLARATION"/>
//                 <xs:element ref="CLASS"/>
//                 <xs:element ref="INSTANCE"/>
//                 <xs:element ref="VALUE.NAMEDINSTANCE"/>
//             </xs:choice>
//             <xs:attribute ref="NAME" use="required"/>
//         </xs:complexType>
//     </xs:element>
type CimIParamValue struct {
	XMLName              xml.Name                 `xml:"IPARAMVALUE"`
	Name                 string                   `xml:"NAME,attr"`
	Value                *CimValue                `xml:"VALUE,omitempty"`
	ValueReference       *CimValueReference       `xml:"VALUE.REFERENCE,omitempty"`
	ValueArray           CimValueArray            `xml:"VALUE.ARRAY,omitempty"`
	ClassName            *CimClassName            `xml:"CLASSNAME,omitempty"`
	InstanceName         *CimInstanceName         `xml:"INSTANCENAME,omitempty"`
	QualifierDeclaration *CimQualifierDeclaration `xml:""QUALIFIER.DECLARATION,omitempty"`
	Class                *CimClass                `xml:"CLASS,omitempty"`
	Instance             *CimInstance             `xml:"INSTANCE,omitempty"`
	ValueNamedInstance   *CimValueNamedInstance   `xml:"VALUE.NAMEDINSTANCE,omitempty"`
}

//     <xs:element name="MULTIRSP">
//         <xs:annotation>
//             <xs:documentation>Defines a multiple CIM operation response.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="SIMPLERSP" minOccurs="2" maxOccurs="unbounded"/>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimMultiRsp struct {
	XMLName    xml.Name       `xml:"MULTIRSP"`
	SimpleRsps []CimSimpleRsp `xml:"SIMPLEREQ"`
}

//     <xs:element name="SIMPLERSP">
//         <xs:annotation>
//             <xs:documentation>Defines a simple CIM operation response.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice>
//                 <xs:element ref="METHODRESPONSE"/>
//                 <xs:element ref="IMETHODRESPONSE"/>
//             </xs:choice>
//         </xs:complexType>
//     </xs:element>
type CimSimpleRsp struct {
	XMLName         xml.Name            `xml:"SIMPLERSP"`
	MethodResponse  *CimMethodResponse  `xml:"METHODRESPONSE,omitempty"`
	IMethodResponse *CimIMethodResponse `xml:"IMETHODRESPONSE,omitempty"`
}

//     <xs:element name="METHODRESPONSE">
//         <xs:annotation>
//             <xs:documentation>Defines a single extrinsic (= class-defined) method invocation response.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice>
//                 <xs:element ref="ERROR"/>
//                 <xs:sequence>
//                     <xs:element ref="RETURNVALUE" minOccurs="0"/>
//                     <xs:element ref="PARAMVALUE" minOccurs="0" maxOccurs="unbounded"/>
//                 </xs:sequence>
//             </xs:choice>
//             <xs:attribute ref="NAME" use="required"/>
//         </xs:complexType>
//     </xs:element>
type CimMethodResponse struct {
	XMLName     xml.Name        `xml:"METHODRESPONSE"`
	Name        string          `xml:"NAME,attr"`
	Error       *CimError       `xml:"ERROR,omitempty"`
	ReturnValue *CimReturnValue `xml:"RETURNVALUE,omitempty"`
	ParamValue  []CimParamValue `xml:"PARAMVALUE,omitempty"`
}

//     <xs:element name="IMETHODRESPONSE">
//         <xs:annotation>
//             <xs:documentation>Defines a single intrinsic method (=operation) invocation response.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice>
//                 <xs:element ref="ERROR"/>
//                 <xs:sequence>
//                     <xs:element ref="IRETURNVALUE" minOccurs="0"/>
//                     <xs:element ref="PARAMVALUE" minOccurs="0" maxOccurs="unbounded"/>
//                 </xs:sequence>
//             </xs:choice>
//             <xs:attribute ref="NAME" use="required"/>
//         </xs:complexType>
//     </xs:element>
type CimIMethodResponse struct {
	XMLName     xml.Name         `xml:"IMETHODRESPONSE"`
	Name        string           `xml:"NAME,attr"`
	Error       *CimError        `xml:"ERROR,omitempty"`
	ReturnValue *CimIReturnValue `xml:"IRETURNVALUE,omitempty"`
	ParamValue  []CimParamValue  `xml:"PARAMVALUE,omitempty"`
}

//     <xs:element name="ERROR">
//         <xs:annotation>
//             <xs:documentation>Defines a fundamental error that prevented a method from executing normally
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="INSTANCE" minOccurs="0" maxOccurs="unbounded"/>
//             </xs:sequence>
//             <xs:attribute name="CODE" type="CIMStatusCode_Type" use="required"/>
//             <xs:attribute name="DESCRIPTION" type="xs:string"/>
//         </xs:complexType>
//     </xs:element>
type CimError struct {
	XMLName     xml.Name      `xml:"ERROR"`
	Code        int           `xml:"NAME,attr"`
	Description string        `xml:"DESCRIPTION,attr,omitempty"`
	Instance    []CimInstance `xml:"INSTANCE,omitempty"`
}

//     <xs:element name="RETURNVALUE">
//         <xs:annotation>
//             <xs:documentation>Defines the return value of an extrinsic (= class defined) method within a response.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice minOccurs="0">
//                 <xs:element ref="VALUE"/>
//                 <xs:element ref="VALUE.REFERENCE"/>
//             </xs:choice>
//             <xs:attribute ref="PARAMTYPE"/>
//             <xs:attribute ref="EmbeddedObject"/>
//         </xs:complexType>
//     </xs:element>
type CimReturnValue struct {
	XMLName        xml.Name           `xml:"RETURNVALUE"`
	ParamType      string             `xml:"PARAMTYPE,attr"`
	EmbeddedObject string             `xml:"EmbeddedObject,attr"`
	Value          *CimValue          `xml:"VALUE,omitempty"`
	ValueReference *CimValueReference `xml:"VALUE.REFERENCE,omitempty"`
}

//     <xs:element name="IRETURNVALUE">
//         <xs:annotation>
//             <xs:documentation>Defines the return value of an intrinsic (= operation) method within a response.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice>
//                 <xs:element ref="CLASSNAME" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="INSTANCENAME" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="VALUE" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="VALUE.OBJECTWITHPATH" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="VALUE.OBJECTWITHLOCALPATH" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="VALUE.OBJECT" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="OBJECTPATH" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="QUALIFIER.DECLARATION" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="VALUE.ARRAY" minOccurs="0"/>
//                 <xs:element ref="VALUE.REFERENCE" minOccurs="0"/>
//                 <xs:element ref="CLASS" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="INSTANCE" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="INSTANCEPATH" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="VALUE.NAMEDINSTANCE" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="VALUE.INSTANCEWITHPATH" minOccurs="0" maxOccurs="unbounded"/>
//             </xs:choice>
//         </xs:complexType>
//     </xs:element>
type CimIReturnValue struct {
	XMLName                  xml.Name                      `xml:"IRETURNVALUE"`
	ClassName                []CimClassName                `xml:"CLASSNAME,omitempty"`
	InstanceName             []CimInstanceName             `xml:"INSTANCENAME,omitempty"`
	Value                    []CimValue                    `xml:"VALUE,omitempty"`
	ValueObjectWithPath      []CimValueObjectWithPath      `xml:"VALUE.OBJECTWITHPATH,omitempty"`
	ValueObjectWithLocalPath []CimValueObjectWithLocalPath `xml:"VALUE.OBJECTWITHLOCALPATH,omitempty"`
	ValueObject              []CimValueObject              `xml:"VALUE.OBJECT,omitempty"`
	ObjectPath               []CimObjectPath               `xml:"OBJECTPATH,omitempty"`
	QualifierDeclaration     []CimQualifierDeclaration     `xml:"QUALIFIER.DECLARATION,omitempty"`
	ValueArray               CimValueArray                 `xml:"VALUE.ARRAY,omitempty"`
	ValueReference           *CimValueReference            `xml:"VALUE.REFERENCE,omitempty"`
	Class                    []CimClass                    `xml:"CLASS,omitempty"`
	Instance                 []CimInstance                 `xml:"INSTANCE,omitempty"`
	InstancePath             []CimInstancePath             `xml:"INSTANCEPATH,omitempty"`
	ValueNamedInstance       []CimValueNamedInstance       `xml:"VALUE.NAMEDINSTANCE,omitempty"`
	ValueInstanceWithPath    []CimValueInstanceWithPath    `xml:"VALUE.INSTANCEWITHPATH,omitempty"`
}

//     <xs:element name="MULTIEXPREQ">
//         <xs:annotation>
//             <xs:documentation>Defines a multiple CIM export (= listener operation) request.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="SIMPLEEXPREQ" minOccurs="2" maxOccurs="unbounded"/>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimMultiExpReq struct {
	XMLName       xml.Name          `xml:"MULTIEXPREQ"`
	SimpleExpReqs []CimSimpleExpReq `xml:"SIMPLEEXPREQ"`
}

//     <xs:element name="SIMPLEEXPREQ">
//         <xs:annotation>
//             <xs:documentation>Defines a simple CIM export (= listener operation) request.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="CORRELATOR" minOccurs="0" maxOccurs="unbounded"/>
//                 <xs:element ref="EXPMETHODCALL"/>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimSimpleExpReq struct {
	XMLName       xml.Name         `xml:"SIMPLEEXPREQ"`
	Correlator    []CimCorrelator  `xml:"CORRELATOR,omitempty"`
	ExpMethodCall CimExpMethodCall `xml:"EXPMETHODCALL"`
}

//     <xs:element name="EXPMETHODCALL">
//         <xs:annotation>
//             <xs:documentation>Defines a single export method (= listener operation) invocation request.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="EXPPARAMVALUE" minOccurs="0" maxOccurs="unbounded"/>
//             </xs:sequence>
//             <xs:attribute ref="NAME" use="required"/>
//         </xs:complexType>
//     </xs:element>
//     </xs:element>
type CimExpMethodCall struct {
	XMLName       xml.Name           `xml:"EXPMETHODCALL"`
	Name          string             `xml:"NAME,attr"`
	ExpParamValue []CimExpParamValue // `xml:"EXPPARAMVALUE,omitempty"`
}

//     <xs:element name="MULTIEXPRSP">
//         <xs:annotation>
//             <xs:documentation>Defines a multiple CIM export response.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="SIMPLEEXPRSP" minOccurs="2" maxOccurs="unbounded"/>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimMultiExpRsp struct {
	XMLName       xml.Name          `xml:"MULTIEXPRSP"`
	SimpleExpRsps []CimSimpleExpRsp // `xml:"SIMPLEEXPRSP,omitempty"`
}

//     <xs:element name="SIMPLEEXPRSP">
//         <xs:annotation>
//             <xs:documentation>Defines a simple CIM export response.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="EXPMETHODRESPONSE"/>
//             </xs:sequence>
//         </xs:complexType>
//     </xs:element>
type CimSimpleExpRsp struct {
	XMLName           xml.Name             `xml:"SIMPLEEXPRSP"`
	ExpMethodResponse CimExpMethodResponse `xml:"EXPMETHODRESPONSE"`
}

//     <xs:element name="EXPMETHODRESPONSE">
//         <xs:annotation>
//             <xs:documentation>Defines a single export method (= listener operation) invocation response.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice>
//                 <xs:element ref="ERROR"/>
//                 <xs:element ref="IRETURNVALUE" minOccurs="0"/>
//             </xs:choice>
//             <xs:attribute ref="NAME" use="required"/>
//         </xs:complexType>
//     </xs:element>
type CimExpMethodResponse struct {
	XMLName     xml.Name         `xml:"EXPMETHODRESPONSE"`
	Name        string           `xml:"NAME,attr"`
	Error       *CimError        `xml:"ERROR,omitempty"`
	ReturnValue *CimIReturnValue `xml:"IRETURNVALUE,omitempty"`
}

//     <xs:element name="EXPPARAMVALUE">
//         <xs:annotation>
//             <xs:documentation>Defines a parameter value that is used in a request or response of an export method
// (= listener operation). For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:choice>
//                 <xs:element ref="INSTANCE" minOccurs="0"/>
//             </xs:choice>
//             <xs:attribute ref="NAME" use="required"/>
//         </xs:complexType>
//     </xs:element>
type CimExpParamValue struct {
	XMLName  xml.Name     `xml:"EXPPARAMVALUE"`
	Name     string       `xml:"NAME,attr"`
	Instance *CimInstance `xml:"IRETURNVALUE,omitempty"`
}

// <!--
// **************************************************
// CHANGE NOTE: The ENUMERATIONCONTEXT element was
// removed in version 2.4.0 of this document.
// **************************************************
// -->
//     <xs:element name="CORRELATOR">
//         <xs:annotation>
//             <xs:documentation>Defines an operation correlator.
// For details, see the corresponding element defined in DSP0201.</xs:documentation>
//         </xs:annotation>
//         <xs:complexType>
//             <xs:sequence>
//                 <xs:element ref="VALUE"/>
//             </xs:sequence>
//             <xs:attribute ref="NAME" use="required"/>
//             <xs:attribute ref="TYPE" use="required"/>
//         </xs:complexType>
//     </xs:element>
type CimCorrelator struct {
	XMLName xml.Name `xml:"CORRELATOR"`
	Name    string   `xml:"NAME,attr"`
	Type    string   `xml:"TYPE,attr"`
	Value   CimValue `xml:"VALUE"`
}
