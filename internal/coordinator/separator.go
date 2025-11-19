package coordinator

type Separator struct {
	strs []string
}

func NewSeparator(strs []string) *Separator {
	return &Separator{
		strs: strs,
	}
}

func (s *Separator) Separate() [][]string {
	var result [][]string
	total := len(s.strs)
	parts := 5

	baseSize := total / parts
	remainder := total % parts

	start := 0
	for i := 0; i < parts; i++ {
		size := baseSize
		if i < remainder {
			size++
		}
		end := start + size
		if end > total {
			end = total
		}
		result = append(result, s.strs[start:end])
		start = end
	}
	return result
}
