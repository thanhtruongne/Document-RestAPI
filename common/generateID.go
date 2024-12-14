package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcutil/base58"
)

type UID struct {
	localID uint32
	object  int
	shareID uint32
}

func NewUID(localID uint32, object int, shareID uint32) UID {
	return UID{localID: localID, object: object, shareID: shareID}
}

func (uid UID) String() string {
	val := uint64(uid.localID)<<28 | uint64(uid.object)<<18 | uint64(uid.shareID)<<0
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}

func (uid UID) GetLocalID() uint32 {
	return uid.localID
}

func (uid UID) GetObject() int {
	return uid.object
}

func (uid UID) GetShareID() uint32 {
	return uid.shareID
}

func DecodeComposeID(s string) (UID, error) {
	uid, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return UID{}, err
	}

	if (1 << 18) > uid {
		return UID{}, errors.New("UID was wrong or invalid")

	}

	u := UID{
		localID: uint32(uid >> 18),
		object:  int(uid >> 18 & 0x3FF),
		shareID: uint32(uid >> 0 & 0x3FFFF),
	}
	return u, nil

}

func FromBase58(s string) (UID, error) {
	return DecodeComposeID(string(base58.Decode(s)))
}

func (uid UID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", uid.String())), nil
}

func (uid *UID) UnmarshalJSON(data []byte) error {
	decodeUID, err := FromBase58(strings.Replace(string(data), "\"", "", -1))
	if err != nil {
		return err
	}

	uid.localID = decodeUID.localID
	uid.object = decodeUID.object
	uid.shareID = decodeUID.shareID
	return nil
}

func (uid *UID) Value() (driver.Value, error) {
	if uid == nil {
		return nil, nil
	}
	return int64(uid.localID), nil
}

func (uid *UID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	var i uint32

	switch t := value.(type) {
	case int:
		i = uint32(t)
	case int8:
		i = uint32(t)
	case int16:
		i = uint32(t)
	case int32:
		i = uint32(t)
	case int64:
		i = uint32(t)
	case uint8:
		i = uint32(t)
	case uint16:
		i = uint32(t)
	case uint32:
		i = t
	case uint64:
		i = uint32(t)
	case []byte:
		a, err := strconv.Atoi(string(t))
		if err != nil {
			return nil
		}
		i = uint32(a)
	default:
		return errors.New("Invalid Scan Data UID")
	}

	*uid = NewUID(i, 0, 1)
	return nil
}
