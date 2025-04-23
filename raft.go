// Core Raft implementation - Consensus Module.

const DebugCM = 1

type LogEntry struct {
	Command any
	Term    int
}

type CMState int

const (
	Follower CMState = iota
	Candidate
	Leader
	Dead
)

func (s CMState) String() string {
	switch s {
	case Follower:
		return "Follower"
	case Candidate:
		return "Candidate"
	case Leader:
		return "Leader"
	case Dead:
		return "Dead"
	default:
		panic("unreachable")
	}
}