package king

func Walks(x int) (int, uint64) {
	k := uint64(1) << x

	kl := k & 0xfefefefefefefefe
	kr := k & 0x7f7f7f7f7f7f7f7f

	// mask of all king walks
	var mask uint64 = kl<<7 | k<<8 | kr<<9 | kl>>1 | kr<<1 | kl>>9 | k>>8 | kr>>7

	count := 0
	for m := mask; m > 0; m &= m - 1 {
		count++
	}

	return count, mask
}
