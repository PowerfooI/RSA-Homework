package server

import (
	"encoding/hex"
	"errors"
	"golangv/rsa"
	"math/big"
)

type Key struct {
	NLength int    `json:"n_len,omitempty"`
	NValue  string `json:"n_val,omitempty"`
	EValue  string `json:"e_val,omitempty"`
	DValue  string `json:"d_val,omitempty"`
	PValue  string `json:"p_val,omitempty"`
	QValue  string `json:"q_val,omitempty"`
}

func genKeyTypeFromStr(k *Key) (rsa.Key, error) {
	if k.EValue != "" && (k.PValue != "" || k.QValue != "" || k.DValue != "") {
		return nil, errors.New("illegal pattern for keys")
	}
	if k.NLength == 0 {
		return nil, errors.New("0 length of key is not permitted")
	}
	if k.NValue == "" {
		return nil, errors.New("empty n_val")
	}
	nVal, ok := new(big.Int).SetString(k.NValue, 16)
	if !ok {
		return nil, errors.New("illegal hex pattern of key")
	}

	if k.EValue != "" {
		eVal, ok := new(big.Int).SetString(k.EValue, 16)
		if !ok {
			return nil, errors.New("illegal hex pattern of key")
		}
		return &rsa.PublicKey{
			NLength: k.NLength,
			NValue:  nVal,
			EValue:  eVal,
		}, nil
	} else {
		if k.DValue == "" || k.PValue == "" || k.QValue == "" {
			return nil, errors.New("d, p, q are all required for private key")
		}
		dVal, ok := new(big.Int).SetString(k.DValue, 16)
		if !ok {
			return nil, errors.New("illegal hex pattern of key")
		}
		pVal, ok := new(big.Int).SetString(k.PValue, 16)
		if !ok {
			return nil, errors.New("illegal hex pattern of key")
		}
		qVal, ok := new(big.Int).SetString(k.QValue, 16)
		if !ok {
			return nil, errors.New("illegal hex pattern of key")
		}
		return &rsa.PrivateKey{
			NLength: k.NLength,
			NValue:  nVal,
			DValue:  dVal,
			PValue:  pVal,
			QValue:  qVal,
		}, nil
	}
}

func genStrFromKeyType(key rsa.Key) *Key {
	if key.IsPublic() {
		pKey := key.(*rsa.PublicKey)
		return &Key{
			NLength: pKey.NLength,
			EValue:  hex.EncodeToString(pKey.EValue.Bytes()),
			NValue:  hex.EncodeToString(pKey.NValue.Bytes()),
		}
	} else {
		pKey := key.(*rsa.PrivateKey)
		return &Key{
			NLength: pKey.NLength,
			NValue:  hex.EncodeToString(pKey.NValue.Bytes()),
			DValue:  hex.EncodeToString(pKey.DValue.Bytes()),
			PValue:  hex.EncodeToString(pKey.PValue.Bytes()),
			QValue:  hex.EncodeToString(pKey.QValue.Bytes()),
		}
	}
}
