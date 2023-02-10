package guid

import (
	"crypto/sha1"
	"strings"
	"github.com/google/uuid"
)

func createGuid(bytes []byte, array []byte) string {
	hasher := sha1.New()
	hasher.Write(array)
	hasher.Write(bytes)
	hash := hasher.Sum(nil)
	array2 := make([]byte, 16)
	copy(array2, hash)
	array2[6] = byte((array2[6] & uint8(0xF)) | uint8(5 << 4))
	array2[8] = byte((array2[8] & uint8(0x3F)) | uint8(0x80))
	uuid, _ := uuid.FromBytes(array2)
	return uuid.String()
}

func GenerateGuid(input ...string) string {
	namespaceUUID, _ := uuid.Parse("11fb06fb-712d-4ddd-98c7-e71bbd588830")
	namespaceBytes := [16]byte(namespaceUUID)
	inputBytes := []byte(strings.Join(input, ""))
	return createGuid(inputBytes, namespaceBytes[:])
}
