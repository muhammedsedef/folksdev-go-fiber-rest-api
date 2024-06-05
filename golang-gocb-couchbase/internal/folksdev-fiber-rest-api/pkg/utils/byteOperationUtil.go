package utils

//go:generate mockery --name=IByteOperationUtil
type IByteOperationUtil interface {
	MbToBytes(mb uint) uint
}

type byteOperationUtil struct {
}

func NewByteOperationUtil() *byteOperationUtil {
	return &byteOperationUtil{}
}

func (util *byteOperationUtil) MbToBytes(mb uint) uint {
	return mb << (10 * 2)
}
