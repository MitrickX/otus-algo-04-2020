package moves

import (
	"github.com/MitrickX/otus-algo-04-2020/homeworks/3/fen"
)

// figure kinds enum.
const (
	rook = iota
	bishop
	queen
)

func CalcMoves(code string) [3]uint64 {
	distribution := fen.Convert(code)

	result := [3]uint64{}

	result[rook] = calcMoves(distribution.WhiteRooks, rook)
	result[bishop] = calcMoves(distribution.WhiteBishops, bishop)
	result[queen] = calcMoves(distribution.WhiteQueens, queen)

	return result
}

func calcMoves(pos uint64, kind int) uint64 {
	return calcMovesMask(pos, kind, uint64(0))
}

func calcMovesMask(pos uint64, kind int, mask uint64) uint64 {
	// up
	for p := pos; canMoveUp(p, kind); p <<= 8 {
		mask |= p
	}

	// left-up
	for p := pos; canMoveLeftUp(p, kind); p <<= 7 {
		mask |= p
	}

	// left
	for p := pos; canMoveLeft(p, kind); p >>= 1 {
		mask |= p
	}

	// left-down
	for p := pos; canMoveLeftDown(p, kind); p >>= 9 {
		mask |= p
	}

	// down
	for p := pos; canMoveDown(p, kind); p >>= 8 {
		mask |= p
	}

	// right-down
	for p := pos; canMoveRightDown(p, kind); p >>= 7 {
		mask |= p
	}

	// right
	for p := pos; canMoveRight(p, kind); p <<= 1 {
		mask |= p
	}

	// right-up
	for p := pos; canMoveRightUp(p, kind); p <<= 9 {
		mask |= p
	}

	return mask ^ pos
}

func canMoveUp(pos uint64, kind int) bool {
	return kind != bishop && pos != 0
}

func canMoveLeftUp(pos uint64, kind int) bool {
	return kind != rook && pos != 0 && pos&0x7f7f7f7f7f7f7f7f != 0
}

func canMoveLeft(pos uint64, kind int) bool {
	return kind != bishop && pos&0x7f7f7f7f7f7f7f7f != 0
}

func canMoveLeftDown(pos uint64, kind int) bool {
	return kind != rook && pos != 0 && pos&0x7f7f7f7f7f7f7f7f != 0
}

func canMoveDown(pos uint64, kind int) bool {
	return kind != bishop && pos != 0
}

func canMoveRightDown(pos uint64, kind int) bool {
	return kind != rook && pos != 0 && pos&0xfefefefefefefefe != 0
}

func canMoveRight(pos uint64, kind int) bool {
	return kind != bishop && pos&0xfefefefefefefefe != 0
}

func canMoveRightUp(pos uint64, kind int) bool {
	return kind != rook && pos != 0 && pos&0xfefefefefefefefe != 0
}
