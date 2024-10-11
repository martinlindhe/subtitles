package subtitles

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

// ResyncSubs adjust text timing by `sync` milliseconds
func (subtitle *Subtitle) ResyncSubs(sync int) {
	// log.Printf("resyncing with %d\n", sync)
	for i := range subtitle.Captions {
		subtitle.Captions[i].Start = subtitle.Captions[i].Start.
			Add(time.Duration(sync) * time.Millisecond)
		subtitle.Captions[i].End = subtitle.Captions[i].End.
			Add(time.Duration(sync) * time.Millisecond)
	}
}

var (
	advertisements = []string{
		// english:
		"captions paid for by",
		"english subtitles",
		"subtitles -", "subtitles:", "subtitles by",
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
		"downloaded from",
		"web-dl",
		"subscene",
		"broadcasttext",
		"seriessub", "subtitlesource",
		"addic7ed", "addicted.com", "vaioholics",
		"sdimedia", "sdi media",
		"allsubs.org", "hdbits.org", "bierdopje.com", "subcentral", "mkvcage",
		"cssubs", "tvsub", "uksubtitles",
		"ragbear.com", "ydy.com", "yyets.net", "indivx.net", "sub-way.fr", "blogspot",
		"forom.com", "forom. com", "facebook.com", "hdvietnam.com", "sapo.pt", "softhome.net",
		"americascardroom.com",
		"@gmail.com", "@hotmail.com", "@hotmail.fr",
		"napisy.org", "1000fr.com", "yts.mx",
		"opensubtitles", "open subtitles", "s u b t i t l e",
		"sous-titres.eu", "300mbfilms.com", "put.io", "subtitulos.es", "osdb.link", "300mbunited",
		"simail.si", "sf-film.dk", "sf.net", "vitac.com", "rapidpremium", "psarips",
		"yify-torrents", "yify torrents", "yify movies",
		"thepiratebay", "anoxmous", "verdikt", "la fisher team", "red bee media",
		"mkv player", "best watched using", "advertise your product", "remove all ads",
		"memoryonsmells", "1st-booking",
		":[gwc]:", "ripped by", "ripped with subrip", "titra film",
		"hiqve",
		"trimark home video",

		// swedish:
		"swedish subtitles", "svenska undertexter", "internationella undertexter",
		"svensktextning", "(c) sveriges televisionit", "sveriges television ab",
		"undertexter.se", "undertexter. se", "swesub.nu", "divxsweden", "undertext.com", "undertext.se",
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

		// norweigan:
		"norsk tekst:",

		// danish:
		"oversættelse:", "tekster:",

		// finnish:
		"suomennos:",

		// french:
		"relecture et corrections finales:", "sous-titres par", "el uploador",
	}
)

// RemoveAds removes advertisement from the subtitles
func (subtitle *Subtitle) RemoveAds(outputPrefix string) *Subtitle {

	info := color.New(color.FgGreen, color.BgBlack).SprintFunc()
	orange := color.New(color.FgBlack, color.BgHiYellow).SprintFunc()

	seq := 1
	res := []Caption{}
	for orgSeq, sub := range subtitle.Captions {
		isAd := false
		for _, line := range sub.Text {
			x := strings.ToLower(line)
			for _, adLine := range advertisements {
				if !isAd && strings.Contains(x, adLine) {
					isAd = true

					fmt.Printf("%s [ads] %d %s matched (%s)\n", info(outputPrefix), (orgSeq + 1), orange(sub.Text), adLine)
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
