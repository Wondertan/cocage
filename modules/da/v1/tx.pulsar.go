package v1

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

var _ protoreflect.List = (*_MsgAttestDataCommitment_1_list)(nil)

type _MsgAttestDataCommitment_1_list struct {
	list *[][]byte
}

func (x *_MsgAttestDataCommitment_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgAttestDataCommitment_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfBytes((*x.list)[i])
}

func (x *_MsgAttestDataCommitment_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Bytes()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_MsgAttestDataCommitment_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Bytes()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgAttestDataCommitment_1_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message MsgAttestDataCommitment at list field DataCommitments as it is not of Message kind"))
}

func (x *_MsgAttestDataCommitment_1_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_MsgAttestDataCommitment_1_list) NewElement() protoreflect.Value {
	var v []byte
	return protoreflect.ValueOfBytes(v)
}

func (x *_MsgAttestDataCommitment_1_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_MsgAttestDataCommitment_3_list)(nil)

type _MsgAttestDataCommitment_3_list struct {
	list *[]*Attestation
}

func (x *_MsgAttestDataCommitment_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgAttestDataCommitment_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_MsgAttestDataCommitment_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Attestation)
	(*x.list)[i] = concreteValue
}

func (x *_MsgAttestDataCommitment_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Attestation)
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgAttestDataCommitment_3_list) AppendMutable() protoreflect.Value {
	v := new(Attestation)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgAttestDataCommitment_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_MsgAttestDataCommitment_3_list) NewElement() protoreflect.Value {
	v := new(Attestation)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_MsgAttestDataCommitment_3_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgAttestDataCommitment                  protoreflect.MessageDescriptor
	fd_MsgAttestDataCommitment_data_commitments protoreflect.FieldDescriptor
	fd_MsgAttestDataCommitment_end_height       protoreflect.FieldDescriptor
	fd_MsgAttestDataCommitment_attestations     protoreflect.FieldDescriptor
)

func init() {
	file_da_v1_tx_proto_init()
	md_MsgAttestDataCommitment = File_da_v1_tx_proto.Messages().ByName("MsgAttestDataCommitment")
	fd_MsgAttestDataCommitment_data_commitments = md_MsgAttestDataCommitment.Fields().ByName("data_commitments")
	fd_MsgAttestDataCommitment_end_height = md_MsgAttestDataCommitment.Fields().ByName("end_height")
	fd_MsgAttestDataCommitment_attestations = md_MsgAttestDataCommitment.Fields().ByName("attestations")
}

var _ protoreflect.Message = (*fastReflection_MsgAttestDataCommitment)(nil)

type fastReflection_MsgAttestDataCommitment MsgAttestDataCommitment

func (x *MsgAttestDataCommitment) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgAttestDataCommitment)(x)
}

func (x *MsgAttestDataCommitment) slowProtoReflect() protoreflect.Message {
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

var _fastReflection_MsgAttestDataCommitment_messageType fastReflection_MsgAttestDataCommitment_messageType
var _ protoreflect.MessageType = fastReflection_MsgAttestDataCommitment_messageType{}

type fastReflection_MsgAttestDataCommitment_messageType struct{}

func (x fastReflection_MsgAttestDataCommitment_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgAttestDataCommitment)(nil)
}
func (x fastReflection_MsgAttestDataCommitment_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgAttestDataCommitment)
}
func (x fastReflection_MsgAttestDataCommitment_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAttestDataCommitment
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgAttestDataCommitment) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAttestDataCommitment
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgAttestDataCommitment) Type() protoreflect.MessageType {
	return _fastReflection_MsgAttestDataCommitment_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgAttestDataCommitment) New() protoreflect.Message {
	return new(fastReflection_MsgAttestDataCommitment)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgAttestDataCommitment) Interface() protoreflect.ProtoMessage {
	return (*MsgAttestDataCommitment)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgAttestDataCommitment) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.DataCommitments) != 0 {
		value := protoreflect.ValueOfList(&_MsgAttestDataCommitment_1_list{list: &x.DataCommitments})
		if !f(fd_MsgAttestDataCommitment_data_commitments, value) {
			return
		}
	}
	if x.EndHeight != uint64(0) {
		value := protoreflect.ValueOfUint64(x.EndHeight)
		if !f(fd_MsgAttestDataCommitment_end_height, value) {
			return
		}
	}
	if len(x.Attestations) != 0 {
		value := protoreflect.ValueOfList(&_MsgAttestDataCommitment_3_list{list: &x.Attestations})
		if !f(fd_MsgAttestDataCommitment_attestations, value) {
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
func (x *fastReflection_MsgAttestDataCommitment) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "da.v1.MsgAttestDataCommitment.data_commitments":
		return len(x.DataCommitments) != 0
	case "da.v1.MsgAttestDataCommitment.end_height":
		return x.EndHeight != uint64(0)
	case "da.v1.MsgAttestDataCommitment.attestations":
		return len(x.Attestations) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataCommitment"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataCommitment does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAttestDataCommitment) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "da.v1.MsgAttestDataCommitment.data_commitments":
		x.DataCommitments = nil
	case "da.v1.MsgAttestDataCommitment.end_height":
		x.EndHeight = uint64(0)
	case "da.v1.MsgAttestDataCommitment.attestations":
		x.Attestations = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataCommitment"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataCommitment does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgAttestDataCommitment) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "da.v1.MsgAttestDataCommitment.data_commitments":
		if len(x.DataCommitments) == 0 {
			return protoreflect.ValueOfList(&_MsgAttestDataCommitment_1_list{})
		}
		listValue := &_MsgAttestDataCommitment_1_list{list: &x.DataCommitments}
		return protoreflect.ValueOfList(listValue)
	case "da.v1.MsgAttestDataCommitment.end_height":
		value := x.EndHeight
		return protoreflect.ValueOfUint64(value)
	case "da.v1.MsgAttestDataCommitment.attestations":
		if len(x.Attestations) == 0 {
			return protoreflect.ValueOfList(&_MsgAttestDataCommitment_3_list{})
		}
		listValue := &_MsgAttestDataCommitment_3_list{list: &x.Attestations}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataCommitment"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataCommitment does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgAttestDataCommitment) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "da.v1.MsgAttestDataCommitment.data_commitments":
		lv := value.List()
		clv := lv.(*_MsgAttestDataCommitment_1_list)
		x.DataCommitments = *clv.list
	case "da.v1.MsgAttestDataCommitment.end_height":
		x.EndHeight = value.Uint()
	case "da.v1.MsgAttestDataCommitment.attestations":
		lv := value.List()
		clv := lv.(*_MsgAttestDataCommitment_3_list)
		x.Attestations = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataCommitment"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataCommitment does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgAttestDataCommitment) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "da.v1.MsgAttestDataCommitment.data_commitments":
		if x.DataCommitments == nil {
			x.DataCommitments = [][]byte{}
		}
		value := &_MsgAttestDataCommitment_1_list{list: &x.DataCommitments}
		return protoreflect.ValueOfList(value)
	case "da.v1.MsgAttestDataCommitment.attestations":
		if x.Attestations == nil {
			x.Attestations = []*Attestation{}
		}
		value := &_MsgAttestDataCommitment_3_list{list: &x.Attestations}
		return protoreflect.ValueOfList(value)
	case "da.v1.MsgAttestDataCommitment.end_height":
		panic(fmt.Errorf("field end_height of message da.v1.MsgAttestDataCommitment is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataCommitment"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataCommitment does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgAttestDataCommitment) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "da.v1.MsgAttestDataCommitment.data_commitments":
		list := [][]byte{}
		return protoreflect.ValueOfList(&_MsgAttestDataCommitment_1_list{list: &list})
	case "da.v1.MsgAttestDataCommitment.end_height":
		return protoreflect.ValueOfUint64(uint64(0))
	case "da.v1.MsgAttestDataCommitment.attestations":
		list := []*Attestation{}
		return protoreflect.ValueOfList(&_MsgAttestDataCommitment_3_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataCommitment"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataCommitment does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgAttestDataCommitment) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in da.v1.MsgAttestDataCommitment", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgAttestDataCommitment) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAttestDataCommitment) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgAttestDataCommitment) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgAttestDataCommitment) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgAttestDataCommitment)
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
		if len(x.DataCommitments) > 0 {
			for _, b := range x.DataCommitments {
				l = len(b)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.EndHeight != 0 {
			n += 1 + runtime.Sov(uint64(x.EndHeight))
		}
		if len(x.Attestations) > 0 {
			for _, e := range x.Attestations {
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
		x := input.Message.Interface().(*MsgAttestDataCommitment)
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
		if len(x.Attestations) > 0 {
			for iNdEx := len(x.Attestations) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Attestations[iNdEx])
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
		if x.EndHeight != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.EndHeight))
			i--
			dAtA[i] = 0x10
		}
		if len(x.DataCommitments) > 0 {
			for iNdEx := len(x.DataCommitments) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.DataCommitments[iNdEx])
				copy(dAtA[i:], x.DataCommitments[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DataCommitments[iNdEx])))
				i--
				dAtA[i] = 0xa
			}
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
		x := input.Message.Interface().(*MsgAttestDataCommitment)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAttestDataCommitment: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAttestDataCommitment: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DataCommitments", wireType)
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
				x.DataCommitments = append(x.DataCommitments, make([]byte, postIndex-iNdEx))
				copy(x.DataCommitments[len(x.DataCommitments)-1], dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EndHeight", wireType)
				}
				x.EndHeight = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.EndHeight |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Attestations", wireType)
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
				x.Attestations = append(x.Attestations, &Attestation{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Attestations[len(x.Attestations)-1]); err != nil {
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
	md_MsgAttestDataCommitmentResponse protoreflect.MessageDescriptor
)

func init() {
	file_da_v1_tx_proto_init()
	md_MsgAttestDataCommitmentResponse = File_da_v1_tx_proto.Messages().ByName("MsgAttestDataCommitmentResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgAttestDataCommitmentResponse)(nil)

type fastReflection_MsgAttestDataCommitmentResponse MsgAttestDataCommitmentResponse

func (x *MsgAttestDataCommitmentResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgAttestDataCommitmentResponse)(x)
}

func (x *MsgAttestDataCommitmentResponse) slowProtoReflect() protoreflect.Message {
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

var _fastReflection_MsgAttestDataCommitmentResponse_messageType fastReflection_MsgAttestDataCommitmentResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgAttestDataCommitmentResponse_messageType{}

type fastReflection_MsgAttestDataCommitmentResponse_messageType struct{}

func (x fastReflection_MsgAttestDataCommitmentResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgAttestDataCommitmentResponse)(nil)
}
func (x fastReflection_MsgAttestDataCommitmentResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgAttestDataCommitmentResponse)
}
func (x fastReflection_MsgAttestDataCommitmentResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAttestDataCommitmentResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgAttestDataCommitmentResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAttestDataCommitmentResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgAttestDataCommitmentResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgAttestDataCommitmentResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgAttestDataCommitmentResponse) New() protoreflect.Message {
	return new(fastReflection_MsgAttestDataCommitmentResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgAttestDataCommitmentResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgAttestDataCommitmentResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgAttestDataCommitmentResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgAttestDataCommitmentResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataCommitmentResponse"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataCommitmentResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAttestDataCommitmentResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataCommitmentResponse"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataCommitmentResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgAttestDataCommitmentResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataCommitmentResponse"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataCommitmentResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgAttestDataCommitmentResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataCommitmentResponse"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataCommitmentResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgAttestDataCommitmentResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataCommitmentResponse"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataCommitmentResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgAttestDataCommitmentResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: da.v1.MsgAttestDataCommitmentResponse"))
		}
		panic(fmt.Errorf("message da.v1.MsgAttestDataCommitmentResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgAttestDataCommitmentResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in da.v1.MsgAttestDataCommitmentResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgAttestDataCommitmentResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAttestDataCommitmentResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgAttestDataCommitmentResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgAttestDataCommitmentResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgAttestDataCommitmentResponse)
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
		x := input.Message.Interface().(*MsgAttestDataCommitmentResponse)
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
		x := input.Message.Interface().(*MsgAttestDataCommitmentResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAttestDataCommitmentResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAttestDataCommitmentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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

type MsgAttestDataCommitment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataCommitments [][]byte       `protobuf:"bytes,1,rep,name=data_commitments,json=dataCommitments,proto3" json:"data_commitments,omitempty"`
	EndHeight       uint64         `protobuf:"varint,2,opt,name=end_height,json=endHeight,proto3" json:"end_height,omitempty"`
	Attestations    []*Attestation `protobuf:"bytes,3,rep,name=attestations,proto3" json:"attestations,omitempty"`
}

func (x *MsgAttestDataCommitment) Reset() {
	*x = MsgAttestDataCommitment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_da_v1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAttestDataCommitment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAttestDataCommitment) ProtoMessage() {}

// Deprecated: Use MsgAttestDataCommitment.ProtoReflect.Descriptor instead.
func (*MsgAttestDataCommitment) Descriptor() ([]byte, []int) {
	return file_da_v1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgAttestDataCommitment) GetDataCommitments() [][]byte {
	if x != nil {
		return x.DataCommitments
	}
	return nil
}

func (x *MsgAttestDataCommitment) GetEndHeight() uint64 {
	if x != nil {
		return x.EndHeight
	}
	return 0
}

func (x *MsgAttestDataCommitment) GetAttestations() []*Attestation {
	if x != nil {
		return x.Attestations
	}
	return nil
}

type MsgAttestDataCommitmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgAttestDataCommitmentResponse) Reset() {
	*x = MsgAttestDataCommitmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_da_v1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAttestDataCommitmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAttestDataCommitmentResponse) ProtoMessage() {}

// Deprecated: Use MsgAttestDataCommitmentResponse.ProtoReflect.Descriptor instead.
func (*MsgAttestDataCommitmentResponse) Descriptor() ([]byte, []int) {
	return file_da_v1_tx_proto_rawDescGZIP(), []int{1}
}

var File_da_v1_tx_proto protoreflect.FileDescriptor

var file_da_v1_tx_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x11, 0x64, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x17, 0x4d,
	0x73, 0x67, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c,
	0x52, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x36, 0x0a, 0x0c, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x21, 0x0a, 0x1f, 0x4d, 0x73, 0x67, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x5f, 0x0a, 0x03, 0x4d,
	0x73, 0x67, 0x12, 0x58, 0x0a, 0x0e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1e, 0x2e, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x73, 0x67,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x26, 0x2e, 0x64, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x73, 0x67,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x57, 0x6f, 0x6e, 0x64, 0x65,
	0x72, 0x74, 0x61, 0x6e, 0x2f, 0x64, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f,
	0x64, 0x61, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_da_v1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_da_v1_tx_proto_goTypes = []interface{}{
	(*MsgAttestDataCommitment)(nil),         // 0: da.v1.MsgAttestDataCommitment
	(*MsgAttestDataCommitmentResponse)(nil), // 1: da.v1.MsgAttestDataCommitmentResponse
	(*Attestation)(nil),                     // 2: da.v1.Attestation
}
var file_da_v1_tx_proto_depIdxs = []int32{
	2, // 0: da.v1.MsgAttestDataCommitment.attestations:type_name -> da.v1.Attestation
	0, // 1: da.v1.Msg.AttestDataRoot:input_type -> da.v1.MsgAttestDataCommitment
	1, // 2: da.v1.Msg.AttestDataRoot:output_type -> da.v1.MsgAttestDataCommitmentResponse
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
	file_da_v1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_da_v1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAttestDataCommitment); i {
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
			switch v := v.(*MsgAttestDataCommitmentResponse); i {
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
			NumMessages:   2,
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
