package consensus

import "github.com/Lagrange-Labs/lagrange-node/consensus/types"

// State handles the consensus process.
type State struct {
	types.RoundState

	storage storageInterface
}

// NewState returns a new State.
func NewState(cfg *Config, storage storageInterface) *State {
	return &State{
		storage: storage,
	}
}

// OnStart loads the first unverified block and starts the receive routine.
func (s *State) OnStart() error {
	return nil
}
