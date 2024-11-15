package jwt

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type UserAuthInfo struct {
	UserID        int64     `json:"UserID,string"`
	UserType      int       `json:"UserType"`
	Email         string    `json:"Email"`
	NickName      string    `json:"NickName"`
	LicenseNumber string    `json:"LicenseNumber"`
	CreatedAt     time.Time `json:"CreatedAt"`
	LastLoginAt   time.Time `json:"LastLoginAt"`
	Addr          string    `json:"Addr"`
	AddrName      string    `json:"AddrName"`
}

//
//func UserToAuthInfo(u model.Users) (UserAuthInfo, error) {
//	var ua UserAuthInfo
//	// Marshal the struct to JSON
//	jsonData, err := jsoniter.Marshal(u)
//	if err != nil {
//		return UserAuthInfo{}, err
//	}
//	// Unmarshal the JSON to map[string]interface{}
//	err = jsoniter.Unmarshal(jsonData, &ua)
//	if err != nil {
//		return UserAuthInfo{}, err
//	}
//	return ua, nil
//}

func ToUserAuthInfo(user interface{}) (UserAuthInfo, error) {
	var u UserAuthInfo
	var j []byte
	var err error

	switch v := user.(type) {
	case map[string]interface{}:
		j, err = jsoniter.Marshal(v)
	case map[string]string:
		j, err = jsoniter.Marshal(v)
	default:
		return u, errors.New("unsupported type")
	}

	if err != nil {
		return u, errors.New(err.Error())
	}

	err = jsoniter.Unmarshal(j, &u)
	if err != nil {
		return u, errors.New(err.Error())
	}

	return u, nil
}
