package subtitles

import (
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

// CleanupSub parses .srt or .ssa, performs cleanup and renders to a .srt, returning a string. caller is responsible for passing UTF8 string
func CleanupSub(utf8 string, filterName string, keepAds bool, sync int) (string, error) {
	var subtitle Subtitle
	var err error

	if looksLikeSrt(utf8) {
		subtitle, err = NewFromSRT(utf8)
	} else {
		// falls back on .ssa decoding, for now
		subtitle, err = NewFromSSA(utf8)
	}
	if err != nil {
		return "", err
	}

	if !keepAds {
		subtitle.removeAds()
	}

	if sync != 0 {
		subtitle.resyncSubs(sync)
	}

	subtitle.filterSubs(filterName)
	out := subtitle.AsSRT()

	return out, nil
}

func (subtitle *Subtitle) resyncSubs(sync int) {
	// log.Printf("resyncing with %d\n", sync)
	for i := range subtitle.Captions {
		subtitle.Captions[i].Start = subtitle.Captions[i].Start.
			Add(time.Duration(sync) * time.Millisecond)
		subtitle.Captions[i].End = subtitle.Captions[i].End.
			Add(time.Duration(sync) * time.Millisecond)
	}
}

// RemoveAds removes advertisement from the subtitles
func (subtitle *Subtitle) removeAds() *Subtitle {
	ads := []string{
		// english:
		"captions paid for by",
		"english subtitles",
		"subtitles:", "subtitles by",
		"subtitles downloaded",
		"captioning by", "captions by",
		"transcript :", "transcript:", "transcript by",
		"sync, corrected", "synced and corrected",
		"sync and corrected", "sync & corrections",
		"sync and corrections",
		"traduction:", "transcript par",
		"corrections by",
		"sync by n17t01",
		"sync,", "synchro :", "synchro:", "synced by", "synchronized by",
		"synchronization by", "synchronisation:",
		"resynchronization:",
		"resync:", "resynchro", "resync by",
		"translation by",
		"encoded by",
		"web-dl",
		"subscene",
		"seriessub",
		"addic7ed", "addicted.com", "vaioholics",
		"sdimedia", "sdi media",
		"allsubs.org", "hdbits.org", "bierdopje.com", "subcentral",
		"cssubs", "tvsub", "uksubtitles",
		"ragbear.com", "ydy.com", "yyets.net", "indivx.net", "sub-way.fr", "blogspot",
		"forom.com", "forom. com", "facebook.com", "hdvietnam.com", "sapo.pt", "softhome.net",
		"@gmail.com", "@hotmail.com", "@hotmail.fr",
		"napisy.org", "1000fr.com",
		"opensubtitles", "open subtitles", "s u b t i t l e",
		"sous-titres.eu", "300mbfilms.com", "put.io", "subtitulos.es", "osdb.link", "300mbunited",
		"simail.si", "sf.net", "vitac.com", "rapidpremium", "psarips",
		"yify-torrents", "yify torrents",
		"thepiratebay", "anoxmous", "verdikt", "la fisher team", "red bee media",
		"mkv player", "best watched using", "advertise your product", "remove all ads",
		"memoryonsmells", "1st-booking",
		":[gwc]:", "ripped with subrip", "titra film",

		// swedish:
		"swedish subtitles", "svenska undertexter", "internationella undertexter",
		"svensktextning", "(c) sveriges televisionit", "sveriges television ab",
		"undertexter.se", "undertexter. se", "swesub.nu", "divxsweden",
		"undertext av", "översatt av", "översättning:", "översättning av", "rättad av",
		"synkad av", "synkat av", "synk:", "synkning:", "redigerad av", "textning:",
		"svensk text", "text:", "omsynk:", "omsynkad",
		"transkribering:", "piratpartiet.se",
		"korrektur:", "korrekturläst", "texter på nätet", "text hämtad från",
		"din filmsajt på nätet", "din största filmsajt på nätet",
		"alltid nya texter",
		"senaste undertexter på",
		"programtextning", "översättargrupp",
		"mediatextgruppen", "visiontext", "scandinavian text service",
		"jhs International", "svensk medietext",

		// french:
		"relecture et corrections finales:",
	}

	seq := 1
	res := []Caption{}
	for orgSeq, sub := range subtitle.Captions {

		isAd := false

		for _, line := range sub.Text {
			x := strings.ToLower(line)
			for _, adLine := range ads {
				if !isAd && strings.Contains(x, adLine) {
					isAd = true
					log.Println("[ads]", (orgSeq + 1), sub.Text, "matched", adLine)
					break
				}
			}
		}

		if !isAd {
			sub.Seq = seq
			res = append(res, sub)
			seq++
		}
	}
	subtitle.Captions = res
	return subtitle
}
