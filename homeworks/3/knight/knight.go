package knight

func Walks(x int) (int, uint64) {
	k := uint64(1) << x

	// knight that walks left long
	kll := k & 0xfcfcfcfcfcfcfcfc

	// knight that walks right long
	krl := k & 0x3f3f3f3f3f3f3f3f

	// knight that walks left short
	kls := k & 0xfefefefefefefefe

	// knight that walks right short
	krs := k & 0x7f7f7f7f7f7f7f7f

	// mask of all knight walks
	var mask uint64 = kls<<15 | kll<<6 | kll>>10 | kls>>17 | krs>>15 | krl>>6 | krl<<10 | krs<<17

	count := 0
	for m := mask; m > 0; m &= m - 1 {
		count++
	}

	return count, mask
}
