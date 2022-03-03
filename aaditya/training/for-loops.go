package chessboard

//Concepts - Range iterations and custom type definitions

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool
// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	r,found := cb[rank]
    if !found {
        return 0
    }
	count:=0
    for _,v := range r{
        if v {
          count++
        }
    }
	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file <= 0 || file >= 9 {
        return 0
    }
	count:=0        

	for _,v := range cb {
        if v[file-1] {
            count++
        }
    }
	return count
}




// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	count:=0
	for _,v := range cb {
        for i:=0;i<8;i++ {
            if v[i] {
                count++
            }
        }
    }
	return count
}