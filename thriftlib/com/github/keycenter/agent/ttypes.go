// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package agent

import (
	"bytes"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type ErrorType int64

const (
	ErrorType_CLIENT_ERROR ErrorType = 0
	ErrorType_SERVER_ERROR ErrorType = 1
)

func (p ErrorType) String() string {
	switch p {
	case ErrorType_CLIENT_ERROR:
		return "CLIENT_ERROR"
	case ErrorType_SERVER_ERROR:
		return "SERVER_ERROR"
	}
	return "<UNSET>"
}

func ErrorTypeFromString(s string) (ErrorType, error) {
	switch s {
	case "CLIENT_ERROR":
		return ErrorType_CLIENT_ERROR, nil
	case "SERVER_ERROR":
		return ErrorType_SERVER_ERROR, nil
	}
	return ErrorType(0), fmt.Errorf("not a valid ErrorType string")
}

func ErrorTypePtr(v ErrorType) *ErrorType { return &v }

func (p ErrorType) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *ErrorType) UnmarshalText(text []byte) error {
	q, err := ErrorTypeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

type CompressionType int64

const (
	CompressionType_NONE   CompressionType = 0
	CompressionType_SNAPPY CompressionType = 1
)

func (p CompressionType) String() string {
	switch p {
	case CompressionType_NONE:
		return "NONE"
	case CompressionType_SNAPPY:
		return "SNAPPY"
	}
	return "<UNSET>"
}

func CompressionTypeFromString(s string) (CompressionType, error) {
	switch s {
	case "NONE":
		return CompressionType_NONE, nil
	case "SNAPPY":
		return CompressionType_SNAPPY, nil
	}
	return CompressionType(0), fmt.Errorf("not a valid CompressionType string")
}

func CompressionTypePtr(v CompressionType) *CompressionType { return &v }

func (p CompressionType) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

func (p *CompressionType) UnmarshalText(text []byte) error {
	q, err := CompressionTypeFromString(string(text))
	if err != nil {
		return err
	}
	*p = q
	return nil
}

// Attributes:
//  - Message
//  - CauseStacktrace
//  - ErrorType
type KeycenterException struct {
	Message         string     `thrift:"message,1,required" json:"message"`
	CauseStacktrace string     `thrift:"causeStacktrace,2,required" json:"causeStacktrace"`
	ErrorType       *ErrorType `thrift:"errorType,3" json:"errorType,omitempty"`
}

func NewKeycenterException() *KeycenterException {
	return &KeycenterException{}
}

func (p *KeycenterException) GetMessage() string {
	return p.Message
}

func (p *KeycenterException) GetCauseStacktrace() string {
	return p.CauseStacktrace
}

var KeycenterException_ErrorType_DEFAULT ErrorType

func (p *KeycenterException) GetErrorType() ErrorType {
	if !p.IsSetErrorType() {
		return KeycenterException_ErrorType_DEFAULT
	}
	return *p.ErrorType
}
func (p *KeycenterException) IsSetErrorType() bool {
	return p.ErrorType != nil
}

func (p *KeycenterException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetMessage bool = false
	var issetCauseStacktrace bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
			issetMessage = true
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
			issetCauseStacktrace = true
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetMessage {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Message is not set"))
	}
	if !issetCauseStacktrace {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field CauseStacktrace is not set"))
	}
	return nil
}

func (p *KeycenterException) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = v
	}
	return nil
}

func (p *KeycenterException) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.CauseStacktrace = v
	}
	return nil
}

func (p *KeycenterException) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		temp := ErrorType(v)
		p.ErrorType = &temp
	}
	return nil
}

func (p *KeycenterException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("KeycenterException"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *KeycenterException) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err)
	}
	if err := oprot.WriteString(string(p.Message)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err)
	}
	return err
}

func (p *KeycenterException) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("causeStacktrace", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:causeStacktrace: ", p), err)
	}
	if err := oprot.WriteString(string(p.CauseStacktrace)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.causeStacktrace (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:causeStacktrace: ", p), err)
	}
	return err
}

func (p *KeycenterException) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetErrorType() {
		if err := oprot.WriteFieldBegin("errorType", thrift.I32, 3); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:errorType: ", p), err)
		}
		if err := oprot.WriteI32(int32(*p.ErrorType)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.errorType (3) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 3:errorType: ", p), err)
		}
	}
	return err
}

func (p *KeycenterException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("KeycenterException(%+v)", *p)
}

func (p *KeycenterException) Error() string {
	return p.String()
}