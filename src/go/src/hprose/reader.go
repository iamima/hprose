/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.net/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * hprose/reader.go                                       *
 *                                                        *
 * hprose Reader for Go.                                  *
 *                                                        *
 * LastModified: Jan 30, 2014                             *
 * Author: Ma Bingyao <andot@hprfc.com>                   *
 *                                                        *
\**********************************************************/

package hprose

import (
	"bufio"
	"bytes"
	"container/list"
	"io"
	"math/big"
	"reflect"
	"time"
	"uuid"
)

type Reader interface {
	Stream() *bufio.Reader
	CheckTag(byte) error
	CheckTags([]byte) (byte, error)
	Unserialize(interface{}) error
	ReadValue(reflect.Value) error
	ReadInt(byte) (int, error)
	ReadUint(byte) (uint, error)
	ReadInt64() (int64, error)
	ReadUint64() (uint64, error)
	ReadBigInt() (*big.Int, error)
	ReadFloat32() (float32, error)
	ReadFloat64() (float64, error)
	ReadBool() (bool, error)
	ReadDateTime() (time.Time, error)
	ReadDateWithoutTag() (time.Time, error)
	ReadTimeWithoutTag() (time.Time, error)
	ReadString() (string, error)
	ReadStringWithoutTag() (string, error)
	ReadBytes() (*[]byte, error)
	ReadBytesWithoutTag() (*[]byte, error)
	ReadUUID() (*uuid.UUID, error)
	ReadUUIDWithoutTag() (*uuid.UUID, error)
	ReadList() (*list.List, error)
	ReadListWithoutTag() (*list.List, error)
	ReadArray([]reflect.Value) error
	ReadSlice(interface{}) error
	ReadSliceWithoutTag(interface{}) error
	ReadMap(interface{}) error
	ReadMapWithoutTag(interface{}) error
	ReadObject(interface{}) error
	ReadObjectWithoutTag(interface{}) error
	ReadRaw() ([]byte, error)
	ReadRawTo(*bytes.Buffer) error
	Reset()
}

type reader struct {
	*simpleReader
	ref []interface{}
}

func NewReader(stream io.Reader) Reader {
	r := &reader{NewSimpleReader(stream).(*simpleReader), make([]interface{}, 0, 32)}
	r.setRef = func(p interface{}) {
		r.ref = append(r.ref, p)
	}
	r.readRef = func() (interface{}, error) {
		i, err := r.ReadInt(TagSemicolon)
		if err == nil {
			return r.ref[i], nil
		}
		return nil, err
	}
	return r
}

func (r *reader) Reset() {
	r.simpleReader.Reset()
	r.ref = r.ref[:0]
}
