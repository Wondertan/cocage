package module

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_MsgAttestDataRoot_3_list)(nil)

type _MsgAttestDataRoot_3_list struct {
	list *[]*Signature
}

func (x *_MsgAttestDataRoot_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgAttestDataRoot_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgAttestDataRoot_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Signature)
	(*x.list)[i] = concreteValue
}

func (x *_MsgAttestDataRoot_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Signature)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgAttestDataRoot_3_list) AppendMutable() protoreflect.Value {
	v := new(Signature)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgAttestDataRoot_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgAttestDataRoot_3_list) NewElement() protoreflect.Value {
	v := new(Signature)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgAttestDataRoot_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgAttestDataRoot            protoreflect.MessageDescriptor
	fd_MsgAttestDataRoot_data_root  protoreflect.FieldDescriptor
	fd_MsgAttestDataRoot_height     protoreflect.FieldDescriptor
	fd_MsgAttestDataRoot_signatures protoreflect.FieldDescriptor
)

func init() {
	file_da_v1_tx_proto_init()
	md_MsgAttestDataRoot = File_da_v1_tx_proto.Messages().ByName("MsgAttestDataRoot")
	fd_MsgAttestDataRoot_data_root = md_MsgAttestDataRoot.Fields().ByName("data_root")
	fd_MsgAttestDataRoot_height = md_MsgAttestDataRoot.Fields().ByName("height")
	fd_MsgAttestDataRoot_signatures = md_MsgAttestDataRoot.Fields().ByName("signatures")
}

var _ protoreflect.Message = (*fastReflection_MsgAttestDataRoot)(nil)

type fastReflection_MsgAttestDataRoot MsgAttestDataRoot

func (x *MsgAttestDataRoot) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgAttestDataRoot)(x)
}

func (x *MsgAttestDataRoot) slowProtoReflect() protoreflect.Message {
	mi := &file_da_v1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgAttestDataRoot_messageType fastReflection_MsgAttestDataRoot_messageType
var _ protoreflect.MessageType = fastReflection_MsgAttestDataRoot_messageType{}

type fastReflection_MsgAttestDataRoot_messageType struct{}

func (x fastReflection_MsgAttestDataRoot_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgAttestDataRoot)(nil)
}
func (x fastReflection_MsgAttestDataRoot_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgAttestDataRoot)
}
func (x fastReflection_MsgAttestDataRoot_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAttestDataRoot
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgAttestDataRoot) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAttestDataRoot
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgAttestDataRoot) Type() protoreflect.MessageType {
	return _fastReflection_MsgAttestDataRoot_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgAttestDataRoot) New() protoreflect.Message {
	return new(fastReflection_MsgAttestDataRoot)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgAttestDataRoot) Interface() protoreflect.ProtoMessage {
	return (*MsgAttestDataRoot)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgAttestDataRoot) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.DataRoot) != 0 {
		value := protoreflect.ValueOfBytes(x.DataRoot)
		if !f(fd_MsgAttestDataRoot_data_root, value) {
			return
		}
	}
	if x.Height != int64(0) {
		value := protoreflect.ValueOfInt64(x.Height)
		if !f(fd_MsgAttestDataRoot_height, value) {
			return
		}
	}
	if len(x.Signatures) != 0 {
		value := protoreflect.ValueOfList(&_MsgAttestDataRoot_3_list{list: &x.Signatures})
		if !f(fd_MsgAttestDataRoot_signatures, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgAttestDataRoot) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "da.v1.MsgAttestDataRoot.data_root":
		return len(x.DataRoot) != 0
	case "da.v1.MsgAttestDataRoot.height":
		return x.Height != int64(0)
	case "da.v1.MsgAttestDataRoot.signatures":
		return len(x.Signatures) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataRoot"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataRoot does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAttestDataRoot) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "da.v1.MsgAttestDataRoot.data_root":
		x.DataRoot = nil
	case "da.v1.MsgAttestDataRoot.height":
		x.Height = int64(0)
	case "da.v1.MsgAttestDataRoot.signatures":
		x.Signatures = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataRoot"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataRoot does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgAttestDataRoot) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "da.v1.MsgAttestDataRoot.data_root":
		value := x.DataRoot
		return protoreflect.ValueOfBytes(value)
	case "da.v1.MsgAttestDataRoot.height":
		value := x.Height
		return protoreflect.ValueOfInt64(value)
	case "da.v1.MsgAttestDataRoot.signatures":
		if len(x.Signatures) == 0 {
			return protoreflect.ValueOfList(&_MsgAttestDataRoot_3_list{})
		}
		listValue := &_MsgAttestDataRoot_3_list{list: &x.Signatures}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataRoot"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataRoot does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAttestDataRoot) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "da.v1.MsgAttestDataRoot.data_root":
		x.DataRoot = value.Bytes()
	case "da.v1.MsgAttestDataRoot.height":
		x.Height = value.Int()
	case "da.v1.MsgAttestDataRoot.signatures":
		lv := value.List()
		clv := lv.(*_MsgAttestDataRoot_3_list)
		x.Signatures = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataRoot"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataRoot does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAttestDataRoot) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "da.v1.MsgAttestDataRoot.signatures":
		if x.Signatures == nil {
			x.Signatures = []*Signature{}
		}
		value := &_MsgAttestDataRoot_3_list{list: &x.Signatures}
		return protoreflect.ValueOfList(value)
	case "da.v1.MsgAttestDataRoot.data_root":
		panic(fmt.Errorf("field data_root of message da.v1.MsgAttestDataRoot is not mutable"))
	case "da.v1.MsgAttestDataRoot.height":
		panic(fmt.Errorf("field height of message da.v1.MsgAttestDataRoot is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataRoot"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataRoot does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgAttestDataRoot) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "da.v1.MsgAttestDataRoot.data_root":
		return protoreflect.ValueOfBytes(nil)
	case "da.v1.MsgAttestDataRoot.height":
		return protoreflect.ValueOfInt64(int64(0))
	case "da.v1.MsgAttestDataRoot.signatures":
		list := []*Signature{}
		return protoreflect.ValueOfList(&_MsgAttestDataRoot_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataRoot"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataRoot does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgAttestDataRoot) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in da.v1.MsgAttestDataRoot", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgAttestDataRoot) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAttestDataRoot) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgAttestDataRoot) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgAttestDataRoot) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgAttestDataRoot)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.DataRoot)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Height != 0 {
			n += 1 + runtime.Sov(uint64(x.Height))
		}
		if len(x.Signatures) > 0 {
			for _, e := range x.Signatures {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgAttestDataRoot)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Signatures) > 0 {
			for iNdEx := len(x.Signatures) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Signatures[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0x1a
			}
		}
		if x.Height != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Height))
			i--
			dAtA[i] = 0x10
		}
		if len(x.DataRoot) > 0 {
			i -= len(x.DataRoot)
			copy(dAtA[i:], x.DataRoot)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DataRoot)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgAttestDataRoot)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAttestDataRoot: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAttestDataRoot: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DataRoot", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.DataRoot = append(x.DataRoot[:0], dAtA[iNdEx:postIndex]...)
				if x.DataRoot == nil {
					x.DataRoot = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
				}
				x.Height = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Height |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Signatures", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Signatures = append(x.Signatures, &Signature{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Signatures[len(x.Signatures)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_Signature                   protoreflect.MessageDescriptor
	fd_Signature_validator_address protoreflect.FieldDescriptor
	fd_Signature_signature         protoreflect.FieldDescriptor
)

func init() {
	file_da_v1_tx_proto_init()
	md_Signature = File_da_v1_tx_proto.Messages().ByName("Signature")
	fd_Signature_validator_address = md_Signature.Fields().ByName("validator_address")
	fd_Signature_signature = md_Signature.Fields().ByName("signature")
}

var _ protoreflect.Message = (*fastReflection_Signature)(nil)

type fastReflection_Signature Signature

func (x *Signature) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Signature)(x)
}

func (x *Signature) slowProtoReflect() protoreflect.Message {
	mi := &file_da_v1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Signature_messageType fastReflection_Signature_messageType
var _ protoreflect.MessageType = fastReflection_Signature_messageType{}

type fastReflection_Signature_messageType struct{}

func (x fastReflection_Signature_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Signature)(nil)
}
func (x fastReflection_Signature_messageType) New() protoreflect.Message {
	return new(fastReflection_Signature)
}
func (x fastReflection_Signature_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Signature
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Signature) Descriptor() protoreflect.MessageDescriptor {
	return md_Signature
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Signature) Type() protoreflect.MessageType {
	return _fastReflection_Signature_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Signature) New() protoreflect.Message {
	return new(fastReflection_Signature)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Signature) Interface() protoreflect.ProtoMessage {
	return (*Signature)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Signature) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ValidatorAddress != "" {
		value := protoreflect.ValueOfString(x.ValidatorAddress)
		if !f(fd_Signature_validator_address, value) {
			return
		}
	}
	if len(x.Signature) != 0 {
		value := protoreflect.ValueOfBytes(x.Signature)
		if !f(fd_Signature_signature, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Signature) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "da.v1.Signature.validator_address":
		return x.ValidatorAddress != ""
	case "da.v1.Signature.signature":
		return len(x.Signature) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.Signature"))
		}
		panic(fmt.Errorf("message da.v1.Signature does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Signature) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "da.v1.Signature.validator_address":
		x.ValidatorAddress = ""
	case "da.v1.Signature.signature":
		x.Signature = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.Signature"))
		}
		panic(fmt.Errorf("message da.v1.Signature does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Signature) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "da.v1.Signature.validator_address":
		value := x.ValidatorAddress
		return protoreflect.ValueOfString(value)
	case "da.v1.Signature.signature":
		value := x.Signature
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.Signature"))
		}
		panic(fmt.Errorf("message da.v1.Signature does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Signature) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "da.v1.Signature.validator_address":
		x.ValidatorAddress = value.Interface().(string)
	case "da.v1.Signature.signature":
		x.Signature = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.Signature"))
		}
		panic(fmt.Errorf("message da.v1.Signature does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Signature) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "da.v1.Signature.validator_address":
		panic(fmt.Errorf("field validator_address of message da.v1.Signature is not mutable"))
	case "da.v1.Signature.signature":
		panic(fmt.Errorf("field signature of message da.v1.Signature is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.Signature"))
		}
		panic(fmt.Errorf("message da.v1.Signature does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Signature) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "da.v1.Signature.validator_address":
		return protoreflect.ValueOfString("")
	case "da.v1.Signature.signature":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.Signature"))
		}
		panic(fmt.Errorf("message da.v1.Signature does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Signature) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in da.v1.Signature", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Signature) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Signature) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Signature) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Signature) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Signature)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.ValidatorAddress)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Signature)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Signature)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Signature) > 0 {
			i -= len(x.Signature)
			copy(dAtA[i:], x.Signature)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Signature)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.ValidatorAddress) > 0 {
			i -= len(x.ValidatorAddress)
			copy(dAtA[i:], x.ValidatorAddress)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ValidatorAddress)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Signature)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Signature: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Signature: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ValidatorAddress = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Signature = append(x.Signature[:0], dAtA[iNdEx:postIndex]...)
				if x.Signature == nil {
					x.Signature = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_MsgAttestDataRootResponse protoreflect.MessageDescriptor
)

func init() {
	file_da_v1_tx_proto_init()
	md_MsgAttestDataRootResponse = File_da_v1_tx_proto.Messages().ByName("MsgAttestDataRootResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgAttestDataRootResponse)(nil)

type fastReflection_MsgAttestDataRootResponse MsgAttestDataRootResponse

func (x *MsgAttestDataRootResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgAttestDataRootResponse)(x)
}

func (x *MsgAttestDataRootResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_da_v1_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgAttestDataRootResponse_messageType fastReflection_MsgAttestDataRootResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgAttestDataRootResponse_messageType{}

type fastReflection_MsgAttestDataRootResponse_messageType struct{}

func (x fastReflection_MsgAttestDataRootResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgAttestDataRootResponse)(nil)
}
func (x fastReflection_MsgAttestDataRootResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgAttestDataRootResponse)
}
func (x fastReflection_MsgAttestDataRootResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAttestDataRootResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgAttestDataRootResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAttestDataRootResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgAttestDataRootResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgAttestDataRootResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgAttestDataRootResponse) New() protoreflect.Message {
	return new(fastReflection_MsgAttestDataRootResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgAttestDataRootResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgAttestDataRootResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgAttestDataRootResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_MsgAttestDataRootResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataRootResponse"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataRootResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAttestDataRootResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataRootResponse"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataRootResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgAttestDataRootResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataRootResponse"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataRootResponse does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAttestDataRootResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataRootResponse"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataRootResponse does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAttestDataRootResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataRootResponse"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataRootResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgAttestDataRootResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataRootResponse"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataRootResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgAttestDataRootResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in da.v1.MsgAttestDataRootResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgAttestDataRootResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAttestDataRootResponse) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_MsgAttestDataRootResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgAttestDataRootResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgAttestDataRootResponse)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgAttestDataRootResponse)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*MsgAttestDataRootResponse)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAttestDataRootResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAttestDataRootResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: da/v1/tx.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MsgAttestDataRoot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataRoot   []byte       `protobuf:"bytes,1,opt,name=data_root,json=dataRoot,proto3" json:"data_root,omitempty"`
	Height     int64        `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Signatures []*Signature `protobuf:"bytes,3,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *MsgAttestDataRoot) Reset() {
	*x = MsgAttestDataRoot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_da_v1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAttestDataRoot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAttestDataRoot) ProtoMessage() {}

// Deprecated: Use MsgAttestDataRoot.ProtoReflect.Descriptor instead.
func (*MsgAttestDataRoot) Descriptor() ([]byte, []int) {
	return file_da_v1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgAttestDataRoot) GetDataRoot() []byte {
	if x != nil {
		return x.DataRoot
	}
	return nil
}

func (x *MsgAttestDataRoot) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *MsgAttestDataRoot) GetSignatures() []*Signature {
	if x != nil {
		return x.Signatures
	}
	return nil
}

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	Signature        []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_da_v1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_da_v1_tx_proto_rawDescGZIP(), []int{1}
}

func (x *Signature) GetValidatorAddress() string {
	if x != nil {
		return x.ValidatorAddress
	}
	return ""
}

func (x *Signature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type MsgAttestDataRootResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgAttestDataRootResponse) Reset() {
	*x = MsgAttestDataRootResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_da_v1_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAttestDataRootResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAttestDataRootResponse) ProtoMessage() {}

// Deprecated: Use MsgAttestDataRootResponse.ProtoReflect.Descriptor instead.
func (*MsgAttestDataRootResponse) Descriptor() ([]byte, []int) {
	return file_da_v1_tx_proto_rawDescGZIP(), []int{2}
}

var File_da_v1_tx_proto protoreflect.FileDescriptor

var file_da_v1_tx_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x22, 0x7a, 0x0a, 0x11, 0x4d, 0x73, 0x67, 0x41, 0x74,
	0x74, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x30, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x22, 0x56, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x2b, 0x0a, 0x11, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x1b, 0x0a, 0x19, 0x4d,
	0x73, 0x67, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x6f, 0x6f, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x53, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12,
	0x4c, 0x0a, 0x0e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x6f, 0x6f,
	0x74, 0x12, 0x18, 0x2e, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x6f, 0x6f, 0x74, 0x1a, 0x20, 0x2e, 0x64, 0x61,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x20, 0x5a,
	0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x57, 0x6f, 0x6e, 0x64,
	0x65, 0x72, 0x74, 0x61, 0x6e, 0x2f, 0x64, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_da_v1_tx_proto_rawDescOnce sync.Once
	file_da_v1_tx_proto_rawDescData = file_da_v1_tx_proto_rawDesc
)

func file_da_v1_tx_proto_rawDescGZIP() []byte {
	file_da_v1_tx_proto_rawDescOnce.Do(func() {
		file_da_v1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_da_v1_tx_proto_rawDescData)
	})
	return file_da_v1_tx_proto_rawDescData
}

var file_da_v1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_da_v1_tx_proto_goTypes = []interface{}{
	(*MsgAttestDataRoot)(nil),         // 0: da.v1.MsgAttestDataRoot
	(*Signature)(nil),                 // 1: da.v1.Signature
	(*MsgAttestDataRootResponse)(nil), // 2: da.v1.MsgAttestDataRootResponse
}
var file_da_v1_tx_proto_depIdxs = []int32{
	1, // 0: da.v1.MsgAttestDataRoot.signatures:type_name -> da.v1.Signature
	0, // 1: da.v1.Msg.AttestDataRoot:input_type -> da.v1.MsgAttestDataRoot
	2, // 2: da.v1.Msg.AttestDataRoot:output_type -> da.v1.MsgAttestDataRootResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_da_v1_tx_proto_init() }
func file_da_v1_tx_proto_init() {
	if File_da_v1_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_da_v1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAttestDataRoot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_da_v1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_da_v1_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAttestDataRootResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_da_v1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_da_v1_tx_proto_goTypes,
		DependencyIndexes: file_da_v1_tx_proto_depIdxs,
		MessageInfos:      file_da_v1_tx_proto_msgTypes,
	}.Build()
	File_da_v1_tx_proto = out.File
	file_da_v1_tx_proto_rawDesc = nil
	file_da_v1_tx_proto_goTypes = nil
	file_da_v1_tx_proto_depIdxs = nil
}
