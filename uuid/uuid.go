package uuid

import (
	"fmt"
	"math/big"

	"github.com/google/uuid"
)

func StringToUUID(str string) (*uuid.UUID, error) {
	// Convert string to a big.Int to get corresponding bytes
	bigInt := new(big.Int)
	bigInt.SetString(str, 10)
	bytes := bigInt.Bytes()

	// Pad bytes array because uuid package wants 16 bytes
	l := len(bytes)
	padded := make([]byte, 16)
	copy(padded[16-l:], bytes)

	uuid, err := uuid.FromBytes(padded)
	if err != nil {
		return nil, fmt.Errorf("Cannot convert %s to a UUID.\nReason: %s\n", str, err.Error())
	}

	return &uuid, nil
}
