package electionday

import "strconv"

//Concepts - Pointers in golang
 
// NewVoteCounter returns a new vote counter with
// a given number of inital votes.

func NewVoteCounter(initialVotes int) *int {         
	var newVoteCounter *int
    newVoteCounter = &initialVotes
    return newVoteCounter
}
        

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter== nil {
        return 0
    }else{
   		return *counter
    }
}

// IncrementVoteCount increments the value in a vote counter        

func IncrementVoteCount(counter *int, increment int) {
	*counter = *counter + increment
}

type ElectionResult struct {
	Name 	string
	Votes 	int
}

// NewElectionResult creates a new election result
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var electionResult ElectionResult
    electionResult = ElectionResult{
        Name : candidateName,
        Votes : votes,
    }
	var er *ElectionResult
    er = &electionResult
    return er
}

// DisplayResult creates a message with the result to be displayed

func DisplayResult(result *ElectionResult) string {
	winner:= result.Name + " (" + strconv.Itoa(result.Votes) + ")"
    return winner
}


// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map
func DecrementVotesOfCandidate(results map[string]int, candidate string) {

	v := results[candidate]
    results[candidate] = v-1
}