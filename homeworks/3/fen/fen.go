package fen

import (
	"strconv"
	"strings"
)

type Distribution struct {
	whitePawns,
	whiteKnights,
	whiteBishops,
	whiteRooks,
	whiteQueens,
	whiteKing,
	blackPawns,
	blackKnights,
	blackBishops,
	blackRooks,
	blackQueens,
	blackKing uint64
}

func (d *Distribution) BBD() [12]uint64 {
	return [12]uint64{
		d.whitePawns,
		d.whiteKnights,
		d.whiteBishops,
		d.whiteRooks,
		d.whiteQueens,
		d.whiteKing,
		d.blackPawns,
		d.blackKnights,
		d.blackBishops,
		d.blackRooks,
		d.blackQueens,
		d.blackKing,
	}
}

func (d *Distribution) String() string {
	bbd := d.BBD()
	parts := make([]string, 12)

	for i, v := range bbd {
		parts[i] = strconv.FormatUint(v, 10)
	}

	return strings.Join(parts, "\n")
}

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
				distribution.blackKing |= mask
			case 'q':
				distribution.blackQueens |= mask
			case 'r':
				distribution.blackRooks |= mask
			case 'b':
				distribution.blackBishops |= mask
			case 'n':
				distribution.blackKnights |= mask
			case 'p':
				distribution.blackPawns |= mask
			case 'K':
				distribution.whiteKing |= mask
			case 'Q':
				distribution.whiteQueens |= mask
			case 'R':
				distribution.whiteRooks |= mask
			case 'B':
				distribution.whiteBishops |= mask
			case 'N':
				distribution.whiteKnights |= mask
			case 'P':
				distribution.whitePawns |= mask
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
