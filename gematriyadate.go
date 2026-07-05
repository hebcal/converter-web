package main

import (
	"github.com/hebcal/gematriya"
	"github.com/hebcal/hdate"
)

// Hebrew month names with the ב prefix, indexed by hdate.HMonth
// (Nisan=1 .. Adar2=13). Ported from hebcal-web src/gematriyaDate.js.
var gematriyaMonthNames = []string{
	"",
	"בְּנִיסָן",
	"בְּאִיָיר",
	"בְּסִיוָן",
	"בְּתַמּוּז",
	"בְּאָב",
	"בֶּאֱלוּל",
	"בְּתִשְׁרֵי",
	"בְּחֶשְׁוָן",
	"בְּכִסְלֵו",
	"בְּטֵבֵת",
	"בִּשְׁבָט",
	"בַּאֲדָר",
	"בַּאֲדָר ב׳",
}

const gematriyaAdarI = "בַּאֲדָר א׳"

// gematriyaDate renders a Hebrew date in Hebrew letters with nikud,
// e.g. "כ׳ בְּתַמּוּז תשפ״ו".
func gematriyaDate(hd hdate.HDate) string {
	mm := hd.Month()
	var monthName string
	if mm == hdate.Adar1 && hdate.IsLeapYear(hd.Year()) {
		monthName = gematriyaAdarI
	} else {
		monthName = gematriyaMonthNames[mm]
	}
	return gematriya.Gematriya(hd.Day()) + " " + monthName + " " + gematriya.Gematriya(hd.Year())
}
