package models

// Stat struct
type Stat struct {
	CountMutantDna int     `json:"count_mutant_dna"`
	CountHumanDna  int     `json:"count_human_dna"`
	Ratio          float64 `json:"ratio"`
}

func (s *Stat) calculateRatio() {
	if s.CountHumanDna != 0 {
		s.Ratio = float64(s.CountMutantDna) / float64(s.CountHumanDna)
	}
}

func (s *Stat) countDna(mutant []Mutant) {
	for _, m := range mutant {
		if m.IsMutant {
			s.CountMutantDna++
		} else {
			s.CountHumanDna++
		}
	}
}

// GetStats calculate stats
func (s *Stat) GetStats() error {
	var mutant Mutant
	data, err := mutant.FindAll()

	if err != nil {
		return err
	}

	s.countDna(data)
	s.calculateRatio()
	return nil
}
