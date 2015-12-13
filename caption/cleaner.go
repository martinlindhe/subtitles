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
		"sdimedia", "sdi media",
		"allsubs.org", "hdbits.org", "bierdopje.com", "subcentral", "cssubs", "tvsub",
		"ragbear.com", "ydy.com", "yyets.net", "indivx.net", "sub-way.fr",
		"forom.com", "forom. com", "facebook.com", "hdvietnam.com",
		"napisy.org", "1000fr.com",
		"opensubtitles.org", "open subtitles", "s u b t i t l e",
		"sous-titres.eu", "300mbfilms.com", "put.io", "subtitulos.es", "osdb.link",
		"simail.si", "sf.net", "yify-torrents", "vitac.com",
		"thepiratebay", "anoxmous", "verdikt", "la fisher team", "red bee media",
		"mkv player", "best watched using", "advertise your product", "remove all ads",
		"memoryonsmells", "1st-booking",

		// swe subs:
		"swedish subtitles", "svenska undertexter", "internationella undertexter",
		"undertexter.se", "undertexter. se", "swesub.nu", "divxsweden",
		"undertext av", "översatt av", "översättning:", "översättning av", "rättad av",
		"synkad av", "synkat av", "synk:", "synkning:", "redigerad av",
		"svensk text", "text av", "text:", "omsynk:", "omsynkad",
		"transkribering:", "piratpartiet.se",
		"korrektur:", "korrekturläst", "texter på nätet", "text hämtad från",
		"din filmsajt på nätet", "din största filmsajt på nätet",
		"alltid nya texter",
		"senaste undertexter på",
		"programtextning", "översättargrupp",
		"mediatextgruppen", "visiontext",
	}

	seq := 1
	for orgSeq, sub := range subs {

		isAd := false

		for _, line := range sub.Text {
			x := strings.ToLower(line)
			for _, adLine := range ads {
				if !isAd && strings.Contains(x, adLine) {
					isAd = true
					fmt.Println("[ads]", orgSeq, sub.Text, "matched", adLine)
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
