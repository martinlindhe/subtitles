package subtitles

import "os"

// SubFinder represents a video being queried for subtitles
type SubFinder struct {
	FileName  string
	Language  string
	VideoFile *os.File
	Quiet     bool
}

// NewSubFinder creates a SubFilePair object used to download subs for a video
func NewSubFinder(video *os.File, fileName string, language string) *SubFinder {

	return &SubFinder{
		FileName:  fileName,
		Language:  language,
		VideoFile: video,
	}
}
