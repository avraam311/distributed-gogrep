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
		end := start + baseSize
		if i == parts-1 {
			end += remainder
		}
		if end > total {
			end = total
		}
		result = append(result, s.strs[start:end])
		start = end
		if start >= total {
			break
		}
	}
	return result
}
