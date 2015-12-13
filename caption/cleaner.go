package caption

import (
	"fmt"
	"strings"
)

// CleanSubs removes advertisement from the subtitles
func CleanSubs(subs []Caption) []Caption {

	var res []Caption

	ads := []string{

		// eng subs:
		"captions paid for by",
		"subtitles:", "subtitles by", "captioning by", "captions by",
		"transcript :", "transcript:", "transcript by", "sync and corrected",
		"sync, corrected", "traduction:", "transcript par",
		"sync by n17t01",
		"sync,", "synchro :", "synchro:", "synchronized by", "synchronization by",
		"synchronisation:", "relecture et corrections finales:", "resynchronization:",
		"resync:", "resynchro", "resync by",
		"encoded by",
		"subscene",
		"seriessub",
		"addic7ed", "addicted.com",
		"allsubs.org", "hdbits.org", "bierdopje.com", "subcentral", "cssubs", "tvsub",
		"ragbear.com", "ydy.com", "yyets.net", "indivx.net", "sub-way.fr",
		"forom.com", "forom. com", "facebook.com",
		"napisy.org", "1000fr.com", "opensubtitles.org", "o p e n s u b t i t l e s",
		"sous-titres.eu", "300mbfilms.com", "put.io", "subtitulos.es", "osdb.link",
		"simail.si", "sf.net", "yify-torrents", "vitac.com",
		"thepiratebay", "anoxmous", "verdikt", "la fisher team", "red bee media",
		"memoryonsmells", "mkv player",
		"1st-booking",

		// swe subs:
		"swedish subtitles",
		"undertexter.se", "undertexter. se", "swesub.nu", "divxsweden.net",
		"undertext av", "översatt av", "översättning av", "rättad av", "synkad av",
		"synkat av", "text av", "text:", "synk:", "synkning:", "transkribering:",
		"korrektur:", "mediatextgruppen", "texter på nätet",
	}

	seq := 1
	for orgSeq, sub := range subs {

		isAd := false

		for _, line := range sub.Text {
			x := strings.ToLower(line)
			for _, adLine := range ads {
				if !isAd && strings.Contains(x, adLine) {
					isAd = true
					fmt.Println("Removing caption", orgSeq, sub.Text)
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

	return res
}
