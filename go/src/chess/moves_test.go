package chess

import (
	"fmt"
	"testing"
)

func TestGetMoves(t *testing.T) {
	node := createChessNode(initial_board_json)
	legal_black_moves := getMoves("b", node.board)
	fmt.Println("legal_black_moves = ", legal_black_moves)
	fmt.Println("count = ", len(legal_black_moves))
	if len(legal_black_moves) != 20 {
		t.Error("Initial moves for black should be 20")
	}
}

func TestGetLegalMoves(t *testing.T) {
	node := createChessNode(initial_board_json)
	legal_black_moves := getLegalMoves("b", node.board)
	fmt.Println("legal_black_moves = ", legal_black_moves)
	fmt.Println("count = ", len(legal_black_moves))
	if len(legal_black_moves) != 20 {
		t.Error("Initial moves for black should be 20")
	}
}

func TestGetKingLocations(t *testing.T) {
	node := createChessNode(initial_board_json)
	whiteKingLoc := getKingLocation("w", node.board)
	fmt.Println("white king location = ", whiteKingLoc)
	blackKingLoc := getKingLocation("b", node.board)
	fmt.Println("black king location = ", blackKingLoc)
}

func TestF7PawnShouldNotMove(t *testing.T) {
	fenPlacement := "r1bqkbnr/1ppppppp/p7/n6Q/4P3/3P1N2/PPP2PPP/RNB1KB1R"
	fmt.Println("fen = ", fenPlacement)
	board := createBoardUsingFen(fenPlacement)
	node := ChessNode{board: board}
	node.legal_moves = getLegalMoves("b", node.board)
}

func TestKingShouldNotKillQueen(t *testing.T) {
	fenPlacement := "rnb1k2r/ppqppQ1p/7b/2p5/2B1P3/3P4/PPP2PPP/RN2K1NR"
	fmt.Println("fen = ", fenPlacement)
	board := createBoardUsingFen(fenPlacement)
	node := ChessNode{board: board}
	node.legal_moves = getLegalMoves("b", node.board)
	fmt.Println("# moves = ", len(node.legal_moves))
	fmt.Println(node.legal_moves)
}
