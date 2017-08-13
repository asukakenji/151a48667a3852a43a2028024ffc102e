package bitstring

type BitString64 uint64

const (
	Zero = BitString64(0)
)

func (bs BitString64) LeftShift(amount uint) BitString64 {
	return bs << amount
}

func (bs BitString64) RightShift(amount uint) BitString64 {
	return bs >> amount
}

// index starts with 0.
func (bs BitString64) IsSet(index uint) bool {
	return bs&(1<<index) != 0
}

func (bs BitString64) Set(index uint) BitString64 {
	return bs | (1 << index)
}

func (bs BitString64) Reset(index uint) BitString64 {
	return bs &^ (1 << index)
}

func (bs BitString64) UnsetRightmostOne() BitString64 {
	return bs & (bs - 1)
}

func (bs BitString64) SetRightmostZero() BitString64 {
	return bs | (bs + 1)
}

// Hacker's Delight: Figure 2-1
func (bs BitString64) NextValueWithSameBitCount() BitString64 {
	smallest := bs & -bs
	ripple := bs + smallest
	ones := bs ^ ripple
	ones = (ones >> 2) / smallest
	return ripple | ones
}

func (bs BitString64) InsertZero(index uint) BitString64 {
	maskR := (BitString64(1) << index) - 1
	maskL := ^maskR
	return ((bs & maskL) << 1) | (bs & maskR)
}
