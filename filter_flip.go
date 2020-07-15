package subtitles

// filterFlip reverts the line order of each caption
func (subtitle *Subtitle) filterFlip() *Subtitle {
	for i, cap := range subtitle.Captions {
		flipped := []string{}
		for _, line := range cap.Text {
			flipped = append([]string{line}, flipped...)
		}
		subtitle.Captions[i].Text = flipped
	}
	return subtitle
}
