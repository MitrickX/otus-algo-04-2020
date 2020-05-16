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

// Distribution of chess figures with extra masks - positions of all whites and of all blacks figures.
type Distribution struct {
	fen.Distribution
	whites uint64
	blacks uint64
}

// Convert from FEN to Distribution struct.
func Convert(code string) Distribution {
	d := fen.Convert(code)

	return Distribution{
		Distribution: d,
		whites:       d.WhiteQueens | d.WhiteBishops | d.WhiteRooks | d.WhiteKing | d.WhiteKnights | d.WhitePawns,
		blacks:       d.BlackQueens | d.BlackBishops | d.BlackRooks | d.BlackKing | d.BlackKnights | d.BlackPawns,
	}
}

// Position of figure that helps calculate mask of possible moves of figures.
type Position struct {
	pos    uint64 // position of white figure of predefined kind (only 3 kind is supported in this program)
	kind   int    // kind of figures (only 3 kind is supported in this program)
	whites uint64 // all whites positions mask except current figure
	blacks uint64 // all backs positions
}

// CalcMoves calculates mask of possible moves of 3 figures - white rook, white bishop, white queen.
func CalcMoves(code string) [3]uint64 {
	distribution := Convert(code)

	result := [3]uint64{}

	result[rook] = calcMoves(Position{
		pos:    distribution.WhiteRooks,
		kind:   rook,
		whites: distribution.whites ^ distribution.WhiteRooks,
		blacks: distribution.blacks,
	})

	result[bishop] = calcMoves(Position{
		pos:    distribution.WhiteBishops,
		kind:   bishop,
		whites: distribution.whites ^ distribution.WhiteBishops,
		blacks: distribution.blacks,
	})

	result[queen] = calcMoves(Position{
		pos:    distribution.WhiteQueens,
		kind:   queen,
		whites: distribution.whites ^ distribution.WhiteQueens,
		blacks: distribution.blacks,
	})

	return result
}

// calcMoves calculates all moves of current white figure of current kind.
func calcMoves(position Position) uint64 {
	mask := uint64(0)

	// up
	mask = calcMovesUp(position, mask)

	// left-up
	mask = calcMovesLeftUp(position, mask)

	// left
	mask = calcMovesLeft(position, mask)

	// left-down
	mask = calcMovesLeftDown(position, mask)

	// down
	mask = calcMoveDown(position, mask)

	// right-down
	mask = calcMovesRightDown(position, mask)

	// right
	mask = calcMovesRight(position, mask)

	// right-up
	mask = calcMoveRightUp(position, mask)

	return mask
}

// calcMovesOneDirection calculate all moves toward one direction.
// Method is abstract and parameterized with 2 functions:
// 	- couldMoveOneDirection - that is predicate can we go further from current position toward this direction
//	- move that is this actual move action toward this direction.
func calcMovesOneDirection(pos Position, mask uint64, couldMoveOneDirection func(p uint64, kind int) bool,
	move func(uint64) uint64) uint64 {
	//
	// start position - current position on figure in board
	p := pos.pos

	// until figure could be moved from current position we move and add position to mask
	for couldMoveOneDirection(p, pos.kind) {
		// move figure one cell toward direction
		p = move(p)

		// intersect with white figure - stops (can't go further) and not count this position (not add into mask)
		if p&pos.whites != 0 {
			break
		}

		// add this position into mask
		mask |= p

		// but if this position intersect with black figure we stops - can't go further
		if p&pos.blacks != 0 {
			break
		}
	}

	return mask
}

// calcMovesUp calculate all moves toward up direction.
func calcMovesUp(pos Position, mask uint64) uint64 {
	return calcMovesOneDirection(pos, mask, couldMoveUp, moveUp)
}

// calcMovesLeftUp calculate all moves toward left-up direction.
func calcMovesLeftUp(pos Position, mask uint64) uint64 {
	return calcMovesOneDirection(pos, mask, couldMoveLeftUp, moveLeftUp)
}

// calcMovesLeft calculate all moves toward left direction.
func calcMovesLeft(pos Position, mask uint64) uint64 {
	return calcMovesOneDirection(pos, mask, couldMoveLeft, moveLeft)
}

// calcMovesLeftDown calculate all moves toward left-down direction.
func calcMovesLeftDown(pos Position, mask uint64) uint64 {
	return calcMovesOneDirection(pos, mask, couldMoveLeftDown, moveLeftDown)
}

// calcMoveDown calculate all moves toward down direction.
func calcMoveDown(pos Position, mask uint64) uint64 {
	return calcMovesOneDirection(pos, mask, couldMoveDown, moveDown)
}

// calcMovesRightDown calculate all moves toward right-down direction.
func calcMovesRightDown(pos Position, mask uint64) uint64 {
	return calcMovesOneDirection(pos, mask, couldMoveRightDown, moveRightDown)
}

// calcMovesRight calculate all moves toward right direction.
func calcMovesRight(pos Position, mask uint64) uint64 {
	return calcMovesOneDirection(pos, mask, couldMoveRight, moveRight)
}

// calcMoveRightUp calculate all moves toward right-up direction.
func calcMoveRightUp(pos Position, mask uint64) uint64 {
	return calcMovesOneDirection(pos, mask, couldMoveRightUp, moveRightUp)
}

// moveUp moves position one cell up.
func moveUp(pos uint64) uint64 {
	return pos << 8
}

// moveLeftUp moves position one cell left-up.
func moveLeftUp(pos uint64) uint64 {
	return pos << 7
}

// moveLeft moves position one cell left.
func moveLeft(pos uint64) uint64 {
	return pos >> 1
}

// moveLeftDown moves position one cell left-down.
func moveLeftDown(pos uint64) uint64 {
	return pos >> 9
}

// moveDown moves position one cell down.
func moveDown(pos uint64) uint64 {
	return pos >> 8
}

// moveRightDown moves position one cell right-down.
func moveRightDown(pos uint64) uint64 {
	return pos >> 7
}

// moveRight moves position one cell right.
func moveRight(pos uint64) uint64 {
	return pos << 1
}

// moveRightUp moves position one cell right-up.
func moveRightUp(pos uint64) uint64 {
	return pos << 9
}

// couldMoveUp checks could figure move from current position one cell up in clean board.
func couldMoveUp(pos uint64, kind int) bool {
	return kind != bishop && moveUp(pos) != 0
}

// couldMoveLeftUp checks could figure move from current position one cell left-up in clean board.
func couldMoveLeftUp(pos uint64, kind int) bool {
	return kind != rook && pos&0xfefefefefefefefe != 0 && moveUp(pos) != 0
}

// couldMoveLeft checks could figure move from current position one cell left in clean board.
func couldMoveLeft(pos uint64, kind int) bool {
	return kind != bishop && pos&0xfefefefefefefefe != 0
}

// couldMoveLeftDown checks could figure move from current position one cell left-down in clean board.
func couldMoveLeftDown(pos uint64, kind int) bool {
	return kind != rook && pos&0xfefefefefefefefe != 0 && moveDown(pos) != 0
}

// couldMoveDown checks could figure move from current position one cell down in clean board.
func couldMoveDown(pos uint64, kind int) bool {
	return kind != bishop && moveDown(pos) != 0
}

// couldMoveRightDown checks could figure move from current position one cell right-down in clean board.
func couldMoveRightDown(pos uint64, kind int) bool {
	return kind != rook && pos&0x7f7f7f7f7f7f7f7f != 0 && moveDown(pos) != 0
}

// couldMoveRight checks could figure move from current position one cell right in clean board.
func couldMoveRight(pos uint64, kind int) bool {
	return kind != bishop && pos&0x7f7f7f7f7f7f7f7f != 0
}

// couldMoveRightUp checks could figure move from current position one cell right-up in clean board.
func couldMoveRightUp(pos uint64, kind int) bool {
	return kind != rook && pos&0x7f7f7f7f7f7f7f7f != 0 && moveUp(pos) != 0
}
