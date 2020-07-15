package subtitles

// filterMerge combines all separate captions with same time spawn into one caption
func (subtitle *Subtitle) filterMerge() *Subtitle {

	idxToRemove := []int{}
	for i, cap := range subtitle.Captions {
		for j, pCap := range subtitle.Captions[0:i] {
			if pCap.Start == cap.Start && pCap.End == cap.End {
				subtitle.Captions[j].Text = append(pCap.Text, cap.Text...)
				idxToRemove = append(idxToRemove, i)
				break
			}
		}
	}

	newCaptions := []Caption{}
	for i, cap := range subtitle.Captions {
		if !contains(idxToRemove, i) {
			newCaptions = append(newCaptions, cap)
		}
	}

	subtitle.Captions = newCaptions

	return subtitle
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
