package subber

import (
	"log"
	"strings"
)

// CleanupSub performs cleanup on .srt data, returning a string. caller is responsible for passing UTF8 string
func CleanupSub(utf8 string, filterName string, keepAds bool) (string, error) {

	captions := parseSrt(utf8)
	if !keepAds {
		captions = removeAds(captions)
	}

	captions = filterSubs(captions, filterName)

	out := renderSrt(captions)

	return out, nil
}

// RemoveAds removes advertisement from the subtitles
func removeAds(subs []caption) []caption {

	var res []caption

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
		"encoded by",
		"web-dl",
		"subscene",
		"seriessub",
		"addic7ed", "addicted.com", "vaioholics",
		"sdimedia", "sdi media",
		"allsubs.org", "hdbits.org", "bierdopje.com", "subcentral",
		"cssubs", "tvsub", "uksubtitles",
		"ragbear.com", "ydy.com", "yyets.net", "indivx.net", "sub-way.fr",
		"forom.com", "forom. com", "facebook.com", "hdvietnam.com",
		"napisy.org", "1000fr.com",
		"opensubtitles", "open subtitles", "s u b t i t l e",
		"sous-titres.eu", "300mbfilms.com", "put.io", "subtitulos.es", "osdb.link",
		"simail.si", "sf.net", "vitac.com",
		"yify-torrents", "yify torrents",
		"thepiratebay", "anoxmous", "verdikt", "la fisher team", "red bee media",
		"mkv player", "best watched using", "advertise your product", "remove all ads",
		"memoryonsmells", "1st-booking",

		// swedish:
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
		"mediatextgruppen", "visiontext", "scandinavian text service",
		"jhs International",

		// french:
		"relecture et corrections finales:",
	}

	seq := 1
	for orgSeq, sub := range subs {

		isAd := false

		for _, line := range sub.Text {
			x := strings.ToLower(line)
			for _, adLine := range ads {
				if !isAd && strings.Contains(x, adLine) {
					isAd = true
					log.Println("[ads]", orgSeq, sub.Text, "matched", adLine)
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
