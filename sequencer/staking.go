package sequencer

func (s *Sequencer) checkStacking() error {
	if s.blockNumber%uint64(s.stakingInterval) == 0 {
		// TODO check stacking status
		return nil
	}
	return nil
}
