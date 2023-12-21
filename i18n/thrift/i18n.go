// Code generated by Thrift Compiler (0.19.0). DO NOT EDIT.

package thrift

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	thrift "github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

// (needed by validator.)
var _ = strings.Contains
var _ = regexp.MatchString

// Attributes:
//   - Languages
type ListLanguagesRequest struct {
	Languages []string `thrift:"languages,1" db:"languages" json:"languages"`
}

func NewListLanguagesRequest() *ListLanguagesRequest {
	return &ListLanguagesRequest{}
}

func (p *ListLanguagesRequest) GetLanguages() []string {
	return p.Languages
}
func (p *ListLanguagesRequest) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.LIST {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ListLanguagesRequest) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin(ctx)
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]string, 0, size)
	p.Languages = tSlice
	for i := 0; i < size; i++ {
		var _elem0 string
		if v, err := iprot.ReadString(ctx); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem0 = v
		}
		p.Languages = append(p.Languages, _elem0)
	}
	if err := iprot.ReadListEnd(ctx); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *ListLanguagesRequest) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "ListLanguagesRequest"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ListLanguagesRequest) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "languages", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:languages: ", p), err)
	}
	if err := oprot.WriteListBegin(ctx, thrift.STRING, len(p.Languages)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Languages {
		if err := oprot.WriteString(ctx, string(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(ctx); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:languages: ", p), err)
	}
	return err
}

func (p *ListLanguagesRequest) Equals(other *ListLanguagesRequest) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if len(p.Languages) != len(other.Languages) {
		return false
	}
	for i, _tgt := range p.Languages {
		_src1 := other.Languages[i]
		if _tgt != _src1 {
			return false
		}
	}
	return true
}

func (p *ListLanguagesRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ListLanguagesRequest(%+v)", *p)
}

func (p *ListLanguagesRequest) Validate() error {
	return nil
}

// Attributes:
//   - Entries
//   - Timestamp
type ListLanguagesResponse struct {
	Entries   map[string]*LanguageEntry `thrift:"entries,1" db:"entries" json:"entries"`
	Timestamp int64                     `thrift:"timestamp,2" db:"timestamp" json:"timestamp"`
}

func NewListLanguagesResponse() *ListLanguagesResponse {
	return &ListLanguagesResponse{}
}

func (p *ListLanguagesResponse) GetEntries() map[string]*LanguageEntry {
	return p.Entries
}

func (p *ListLanguagesResponse) GetTimestamp() int64 {
	return p.Timestamp
}
func (p *ListLanguagesResponse) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.MAP {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ListLanguagesResponse) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin(ctx)
	if err != nil {
		return thrift.PrependError("error reading map begin: ", err)
	}
	tMap := make(map[string]*LanguageEntry, size)
	p.Entries = tMap
	for i := 0; i < size; i++ {
		var _key2 string
		if v, err := iprot.ReadString(ctx); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_key2 = v
		}
		_val3 := &LanguageEntry{}
		if err := _val3.Read(ctx, iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _val3), err)
		}
		p.Entries[_key2] = _val3
	}
	if err := iprot.ReadMapEnd(ctx); err != nil {
		return thrift.PrependError("error reading map end: ", err)
	}
	return nil
}

func (p *ListLanguagesResponse) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Timestamp = v
	}
	return nil
}

func (p *ListLanguagesResponse) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "ListLanguagesResponse"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ListLanguagesResponse) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "entries", thrift.MAP, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:entries: ", p), err)
	}
	if err := oprot.WriteMapBegin(ctx, thrift.STRING, thrift.STRUCT, len(p.Entries)); err != nil {
		return thrift.PrependError("error writing map begin: ", err)
	}
	for k, v := range p.Entries {
		if err := oprot.WriteString(ctx, string(k)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
		if err := v.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteMapEnd(ctx); err != nil {
		return thrift.PrependError("error writing map end: ", err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:entries: ", p), err)
	}
	return err
}

func (p *ListLanguagesResponse) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "timestamp", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:timestamp: ", p), err)
	}
	if err := oprot.WriteI64(ctx, int64(p.Timestamp)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.timestamp (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:timestamp: ", p), err)
	}
	return err
}

func (p *ListLanguagesResponse) Equals(other *ListLanguagesResponse) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if len(p.Entries) != len(other.Entries) {
		return false
	}
	for k, _tgt := range p.Entries {
		_src4 := other.Entries[k]
		if !_tgt.Equals(_src4) {
			return false
		}
	}
	if p.Timestamp != other.Timestamp {
		return false
	}
	return true
}

func (p *ListLanguagesResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ListLanguagesResponse(%+v)", *p)
}

func (p *ListLanguagesResponse) Validate() error {
	return nil
}

// Attributes:
//   - Path
//   - Language
//   - Valid
//   - Payload
type LanguageEntry struct {
	Path     string `thrift:"path,1" db:"path" json:"path"`
	Language string `thrift:"language,2" db:"language" json:"language"`
	Valid    bool   `thrift:"valid,3" db:"valid" json:"valid"`
	// unused fields # 4 to 19
	Payload []byte `thrift:"payload,20" db:"payload" json:"payload"`
}

func NewLanguageEntry() *LanguageEntry {
	return &LanguageEntry{}
}

func (p *LanguageEntry) GetPath() string {
	return p.Path
}

func (p *LanguageEntry) GetLanguage() string {
	return p.Language
}

func (p *LanguageEntry) GetValid() bool {
	return p.Valid
}

func (p *LanguageEntry) GetPayload() []byte {
	return p.Payload
}
func (p *LanguageEntry) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 3:
			if fieldTypeId == thrift.BOOL {
				if err := p.ReadField3(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 20:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField20(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *LanguageEntry) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Path = v
	}
	return nil
}

func (p *LanguageEntry) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Language = v
	}
	return nil
}

func (p *LanguageEntry) ReadField3(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(ctx); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.Valid = v
	}
	return nil
}

func (p *LanguageEntry) ReadField20(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBinary(ctx); err != nil {
		return thrift.PrependError("error reading field 20: ", err)
	} else {
		p.Payload = v
	}
	return nil
}

func (p *LanguageEntry) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "LanguageEntry"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField3(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField20(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *LanguageEntry) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "path", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:path: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Path)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.path (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:path: ", p), err)
	}
	return err
}

func (p *LanguageEntry) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "language", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:language: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Language)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.language (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:language: ", p), err)
	}
	return err
}

func (p *LanguageEntry) writeField3(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "valid", thrift.BOOL, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:valid: ", p), err)
	}
	if err := oprot.WriteBool(ctx, bool(p.Valid)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.valid (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:valid: ", p), err)
	}
	return err
}

func (p *LanguageEntry) writeField20(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "payload", thrift.STRING, 20); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 20:payload: ", p), err)
	}
	if err := oprot.WriteBinary(ctx, p.Payload); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.payload (20) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 20:payload: ", p), err)
	}
	return err
}

func (p *LanguageEntry) Equals(other *LanguageEntry) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.Path != other.Path {
		return false
	}
	if p.Language != other.Language {
		return false
	}
	if p.Valid != other.Valid {
		return false
	}
	if bytes.Compare(p.Payload, other.Payload) != 0 {
		return false
	}
	return true
}

func (p *LanguageEntry) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LanguageEntry(%+v)", *p)
}

func (p *LanguageEntry) Validate() error {
	return nil
}

type I18N interface {
	// Parameters:
	//  - Req
	ListLanguages(ctx context.Context, req *ListLanguagesRequest) (_r *ListLanguagesResponse, _err error)
}

type I18NClient struct {
	c    thrift.TClient
	meta thrift.ResponseMeta
}

func NewI18NClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *I18NClient {
	return &I18NClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewI18NClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *I18NClient {
	return &I18NClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewI18NClient(c thrift.TClient) *I18NClient {
	return &I18NClient{
		c: c,
	}
}

func (p *I18NClient) Client_() thrift.TClient {
	return p.c
}

func (p *I18NClient) LastResponseMeta_() thrift.ResponseMeta {
	return p.meta
}

func (p *I18NClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
	p.meta = meta
}

// Parameters:
//   - Req
func (p *I18NClient) ListLanguages(ctx context.Context, req *ListLanguagesRequest) (_r *ListLanguagesResponse, _err error) {
	var _args5 I18NListLanguagesArgs
	_args5.Req = req
	var _result7 I18NListLanguagesResult
	var _meta6 thrift.ResponseMeta
	_meta6, _err = p.Client_().Call(ctx, "ListLanguages", &_args5, &_result7)
	p.SetLastResponseMeta_(_meta6)
	if _err != nil {
		return
	}
	if _ret8 := _result7.GetSuccess(); _ret8 != nil {
		return _ret8, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "ListLanguages failed: unknown result")
}

type I18NProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      I18N
}

func (p *I18NProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *I18NProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *I18NProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewI18NProcessor(handler I18N) *I18NProcessor {

	self9 := &I18NProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self9.processorMap["ListLanguages"] = &i18NProcessorListLanguages{handler: handler}
	return self9
}

func (p *I18NProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
	if err2 != nil {
		return false, thrift.WrapTException(err2)
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(ctx, thrift.STRUCT)
	iprot.ReadMessageEnd(ctx)
	x10 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
	x10.Write(ctx, oprot)
	oprot.WriteMessageEnd(ctx)
	oprot.Flush(ctx)
	return false, x10

}

type i18NProcessorListLanguages struct {
	handler I18N
}

func (p *i18NProcessorListLanguages) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	var _write_err11 error
	args := I18NListLanguagesArgs{}
	if err2 := args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "ListLanguages", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelCauseFunc
		ctx, cancel = context.WithCancelCause(ctx)
		defer cancel(nil)
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelCauseFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel(thrift.ErrAbandonRequest)
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := I18NListLanguagesResult{}
	if retval, err2 := p.handler.ListLanguages(ctx, args.Req); err2 != nil {
		tickerCancel()
		err = thrift.WrapTException(err2)
		if errors.Is(err2, thrift.ErrAbandonRequest) {
			return false, thrift.WrapTException(err2)
		}
		if errors.Is(err2, context.Canceled) {
			if err := context.Cause(ctx); errors.Is(err, thrift.ErrAbandonRequest) {
				return false, thrift.WrapTException(err)
			}
		}
		_exc12 := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing ListLanguages: "+err2.Error())
		if err2 := oprot.WriteMessageBegin(ctx, "ListLanguages", thrift.EXCEPTION, seqId); err2 != nil {
			_write_err11 = thrift.WrapTException(err2)
		}
		if err2 := _exc12.Write(ctx, oprot); _write_err11 == nil && err2 != nil {
			_write_err11 = thrift.WrapTException(err2)
		}
		if err2 := oprot.WriteMessageEnd(ctx); _write_err11 == nil && err2 != nil {
			_write_err11 = thrift.WrapTException(err2)
		}
		if err2 := oprot.Flush(ctx); _write_err11 == nil && err2 != nil {
			_write_err11 = thrift.WrapTException(err2)
		}
		if _write_err11 != nil {
			return false, thrift.WrapTException(_write_err11)
		}
		return true, err
	} else {
		result.Success = retval
	}
	tickerCancel()
	if err2 := oprot.WriteMessageBegin(ctx, "ListLanguages", thrift.REPLY, seqId); err2 != nil {
		_write_err11 = thrift.WrapTException(err2)
	}
	if err2 := result.Write(ctx, oprot); _write_err11 == nil && err2 != nil {
		_write_err11 = thrift.WrapTException(err2)
	}
	if err2 := oprot.WriteMessageEnd(ctx); _write_err11 == nil && err2 != nil {
		_write_err11 = thrift.WrapTException(err2)
	}
	if err2 := oprot.Flush(ctx); _write_err11 == nil && err2 != nil {
		_write_err11 = thrift.WrapTException(err2)
	}
	if _write_err11 != nil {
		return false, thrift.WrapTException(_write_err11)
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//   - Req
type I18NListLanguagesArgs struct {
	Req *ListLanguagesRequest `thrift:"req,1" db:"req" json:"req"`
}

func NewI18NListLanguagesArgs() *I18NListLanguagesArgs {
	return &I18NListLanguagesArgs{}
}

var I18NListLanguagesArgs_Req_DEFAULT *ListLanguagesRequest

func (p *I18NListLanguagesArgs) GetReq() *ListLanguagesRequest {
	if !p.IsSetReq() {
		return I18NListLanguagesArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *I18NListLanguagesArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *I18NListLanguagesArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *I18NListLanguagesArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Req = &ListLanguagesRequest{}
	if err := p.Req.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
	}
	return nil
}

func (p *I18NListLanguagesArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "ListLanguages_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *I18NListLanguagesArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "req", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:req: ", p), err)
	}
	if err := p.Req.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Req), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:req: ", p), err)
	}
	return err
}

func (p *I18NListLanguagesArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("I18NListLanguagesArgs(%+v)", *p)
}

// Attributes:
//   - Success
type I18NListLanguagesResult struct {
	Success *ListLanguagesResponse `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewI18NListLanguagesResult() *I18NListLanguagesResult {
	return &I18NListLanguagesResult{}
}

var I18NListLanguagesResult_Success_DEFAULT *ListLanguagesResponse

func (p *I18NListLanguagesResult) GetSuccess() *ListLanguagesResponse {
	if !p.IsSetSuccess() {
		return I18NListLanguagesResult_Success_DEFAULT
	}
	return p.Success
}
func (p *I18NListLanguagesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *I18NListLanguagesResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *I18NListLanguagesResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &ListLanguagesResponse{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *I18NListLanguagesResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "ListLanguages_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *I18NListLanguagesResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *I18NListLanguagesResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("I18NListLanguagesResult(%+v)", *p)
}