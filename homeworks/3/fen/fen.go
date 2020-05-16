package fen

import (
	"strconv"
	"strings"
)

// Distribution of chess figures.
type Distribution struct {
	WhitePawns,
	WhiteKnights,
	WhiteBishops,
	WhiteRooks,
	WhiteQueens,
	WhiteKing,
	BlackPawns,
	BlackKnights,
	BlackBishops,
	BlackRooks,
	BlackQueens,
	BlackKing uint64
}

// BBD format.
func (d *Distribution) BBD() [12]uint64 {
	return [12]uint64{
		d.WhitePawns,
		d.WhiteKnights,
		d.WhiteBishops,
		d.WhiteRooks,
		d.WhiteQueens,
		d.WhiteKing,
		d.BlackPawns,
		d.BlackKnights,
		d.BlackBishops,
		d.BlackRooks,
		d.BlackQueens,
		d.BlackKing,
	}
}

// String format of BDD.
func (d *Distribution) String() string {
	bbd := d.BBD()
	parts := make([]string, 12)

	for i, v := range bbd {
		parts[i] = strconv.FormatUint(v, 10)
	}

	return strings.Join(parts, "\n")
}

// Convert from FEN format to Distribution structure.
func Convert(code string) Distribution {
	distribution := Distribution{}

	lines := strings.Split(code, "/")

	for row := 0; row < 8; row++ {
		line := lines[7-row]
		lineLen := len(line)

		col := 0

		for i := 0; i < lineLen; i++ {
			cell := 8*row + col
			mask := uint64(1 << cell)

			switch line[i] {
			case 'k':
				distribution.BlackKing |= mask
			case 'q':
				distribution.BlackQueens |= mask
			case 'r':
				distribution.BlackRooks |= mask
			case 'b':
				distribution.BlackBishops |= mask
			case 'n':
				distribution.BlackKnights |= mask
			case 'p':
				distribution.BlackPawns |= mask
			case 'K':
				distribution.WhiteKing |= mask
			case 'Q':
				distribution.WhiteQueens |= mask
			case 'R':
				distribution.WhiteRooks |= mask
			case 'B':
				distribution.WhiteBishops |= mask
			case 'N':
				distribution.WhiteKnights |= mask
			case 'P':
				distribution.WhitePawns |= mask
			default:
				shift := int(line[i] - '0')
				col += shift - 1 // shift column position on board (minus 1, cause of increment in the end of loop body
			}

			// increment column position on board
			col++
		}
	}

	return distribution
}
