package color

import (
	"strconv"
	"strings"

	"github.com/ethanbaker/colors/css"
	"github.com/ethanbaker/colors/sol"
)

/*
For future development with special escape code sequences
	ClearScreen = "\033[2J\033[H"
	ClearLine   = "\033[2K\033[G"
	CursorOff   = "\033[?25l"
	CursorOn    = "\033[?25h"
	StrikeOut   = "\033[9m"
*/

//Ascii Colors
var colorReplacerAsciiSol *strings.Replacer
var colorReplacerAsciiCss *strings.Replacer

func init() {
	colorReplacerAsciiSol = strings.NewReplacer(
		sol.AsciiIndex["Base03"], "",
		sol.AsciiIndex["Base02"], "",
		sol.AsciiIndex["Base01"], "",
		sol.AsciiIndex["Base00"], "",
		sol.AsciiIndex["Base0"], "",
		sol.AsciiIndex["Base1"], "",
		sol.AsciiIndex["Base2"], "",
		sol.AsciiIndex["Base3"], "",
		sol.AsciiIndex["Yellow"], "",
		sol.AsciiIndex["Orange"], "",
		sol.AsciiIndex["Red"], "",
		sol.AsciiIndex["Magenta"], "",
		sol.AsciiIndex["Violet"], "",
		sol.AsciiIndex["Blue"], "",
		sol.AsciiIndex["Cyan"], "",
		sol.AsciiIndex["Green"], "",
		sol.AsciiIndex["Reset"], "",
	)

	colorReplacerAsciiCss = strings.NewReplacer(
		css.AsciiIndex["AliceBlue"], "",
		css.AsciiIndex["AntiqueWhite"], "",
		css.AsciiIndex["Aqua"], "",
		css.AsciiIndex["AquaMarine"], "",
		css.AsciiIndex["Azure"], "",
		css.AsciiIndex["Beige"], "",
		css.AsciiIndex["Bisque"], "",
		css.AsciiIndex["Black"], "",
		css.AsciiIndex["BlanchedAlmond"], "",
		css.AsciiIndex["Blue"], "",
		css.AsciiIndex["BlueViolet"], "",
		css.AsciiIndex["Brown"], "",
		css.AsciiIndex["BurlyWood"], "",
		css.AsciiIndex["CadetBlue"], "",
		css.AsciiIndex["Chartreuse"], "",
		css.AsciiIndex["Chocolate"], "",
		css.AsciiIndex["Coral"], "",
		css.AsciiIndex["CornFlowerBlue"], "",
		css.AsciiIndex["CornSilk"], "",
		css.AsciiIndex["Crimson"], "",
		css.AsciiIndex["Cyan"], "",
		css.AsciiIndex["DarkBlue"], "",
		css.AsciiIndex["DarkCyan"], "",
		css.AsciiIndex["DarkGoldenRod"], "",
		css.AsciiIndex["DarkGray"], "",
		css.AsciiIndex["DarkGreen"], "",
		css.AsciiIndex["DarkKhaki"], "",
		css.AsciiIndex["DarkMagenta"], "",
		css.AsciiIndex["DarkOliveGreen"], "",
		css.AsciiIndex["DarkOrange"], "",
		css.AsciiIndex["DarkOrchid"], "",
		css.AsciiIndex["DarkRed"], "",
		css.AsciiIndex["DarkSalmon"], "",
		css.AsciiIndex["DarkSeaGreen"], "",
		css.AsciiIndex["DarkSlateBlue"], "",
		css.AsciiIndex["DarkSlateGray"], "",
		css.AsciiIndex["DarkTurquoise"], "",
		css.AsciiIndex["DarkViolet"], "",
		css.AsciiIndex["DeepPink"], "",
		css.AsciiIndex["DeepSkyBlue"], "",
		css.AsciiIndex["DimGray"], "",
		css.AsciiIndex["DodgerBlue"], "",
		css.AsciiIndex["Firebrick"], "",
		css.AsciiIndex["FloralWhite"], "",
		css.AsciiIndex["ForestGreen"], "",
		css.AsciiIndex["Fuchsia"], "",
		css.AsciiIndex["Gainsboro"], "",
		css.AsciiIndex["GhostWhite"], "",
		css.AsciiIndex["Gold"], "",
		css.AsciiIndex["GoldenRod"], "",
		css.AsciiIndex["Gray"], "",
		css.AsciiIndex["Green"], "",
		css.AsciiIndex["GreenYellow"], "",
		css.AsciiIndex["Honeydew"], "",
		css.AsciiIndex["HotPink"], "",
		css.AsciiIndex["IndianRed"], "",
		css.AsciiIndex["Indigo"], "",
		css.AsciiIndex["Ivory"], "",
		css.AsciiIndex["Khaki"], "",
		css.AsciiIndex["Lavender"], "",
		css.AsciiIndex["LavenderBlush"], "",
		css.AsciiIndex["LawnGreen"], "",
		css.AsciiIndex["LemonChiffon"], "",
		css.AsciiIndex["LightBlue"], "",
		css.AsciiIndex["LightCoral"], "",
		css.AsciiIndex["LightCyan"], "",
		css.AsciiIndex["LightGoldenRodYellow"], "",
		css.AsciiIndex["LightGreen"], "",
		css.AsciiIndex["LightGrey"], "",
		css.AsciiIndex["LightPink"], "",
		css.AsciiIndex["LightSalmon"], "",
		css.AsciiIndex["LightSeaGreen"], "",
		css.AsciiIndex["LightSkyBlue"], "",
		css.AsciiIndex["LightSlateGray"], "",
		css.AsciiIndex["LightSteelBlue"], "",
		css.AsciiIndex["LightYellow"], "",
		css.AsciiIndex["Lime"], "",
		css.AsciiIndex["LimeGreen"], "",
		css.AsciiIndex["Linen"], "",
		css.AsciiIndex["Magenta"], "",
		css.AsciiIndex["Maroon"], "",
		css.AsciiIndex["MediumAquaMarine"], "",
		css.AsciiIndex["MediumBlue"], "",
		css.AsciiIndex["MediumOrchid"], "",
		css.AsciiIndex["MediumPurple"], "",
		css.AsciiIndex["MediumSeaGreen"], "",
		css.AsciiIndex["MediumSlateBlue"], "",
		css.AsciiIndex["MediumSpringGreen"], "",
		css.AsciiIndex["MediumTurquoise"], "",
		css.AsciiIndex["MediumVioletRed"], "",
		css.AsciiIndex["MidnightBlue"], "",
		css.AsciiIndex["MintCream"], "",
		css.AsciiIndex["MistyRose"], "",
		css.AsciiIndex["Moccasin"], "",
		css.AsciiIndex["NavajoWhite"], "",
		css.AsciiIndex["Navy"], "",
		css.AsciiIndex["NavyBlue"], "",
		css.AsciiIndex["OldLace"], "",
		css.AsciiIndex["Olive"], "",
		css.AsciiIndex["OliveDrab"], "",
		css.AsciiIndex["Orange"], "",
		css.AsciiIndex["OrangeRed"], "",
		css.AsciiIndex["Orchid"], "",
		css.AsciiIndex["PaleGoldenRod"], "",
		css.AsciiIndex["PaleGreen"], "",
		css.AsciiIndex["PaleTurquoise"], "",
		css.AsciiIndex["PaleVioletRed"], "",
		css.AsciiIndex["PapayaWhip"], "",
		css.AsciiIndex["PeachPuff"], "",
		css.AsciiIndex["Peru"], "",
		css.AsciiIndex["Pink"], "",
		css.AsciiIndex["Plum"], "",
		css.AsciiIndex["PowderBlue"], "",
		css.AsciiIndex["Purple"], "",
		css.AsciiIndex["Red"], "",
		css.AsciiIndex["RosyBrown"], "",
		css.AsciiIndex["RoyalBlue"], "",
		css.AsciiIndex["SaddleBrown"], "",
		css.AsciiIndex["Salmon"], "",
		css.AsciiIndex["SandyBrown"], "",
		css.AsciiIndex["SeaGreen"], "",
		css.AsciiIndex["SeaShell"], "",
		css.AsciiIndex["Sienna"], "",
		css.AsciiIndex["Silver"], "",
		css.AsciiIndex["SkyBlue"], "",
		css.AsciiIndex["SlateBlue"], "",
		css.AsciiIndex["SlateGray"], "",
		css.AsciiIndex["Snow"], "",
		css.AsciiIndex["SpringGreen"], "",
		css.AsciiIndex["SteelBlue"], "",
		css.AsciiIndex["Tan"], "",
		css.AsciiIndex["Teal"], "",
		css.AsciiIndex["Thistle"], "",
		css.AsciiIndex["Tomato"], "",
		css.AsciiIndex["Turquoise"], "",
		css.AsciiIndex["Violet"], "",
		css.AsciiIndex["Wheat"], "",
		css.AsciiIndex["White"], "",
		css.AsciiIndex["WhiteSmoke"], "",
		css.AsciiIndex["Yellow"], "",
		css.AsciiIndex["YellowGreen"], "",
	)
}

// Decolor removes any of the ANSI color escapes known to this package.
func DecolorAsciiSol(s string) string {
	return colorReplacerAsciiSol.Replace(s)
}

func DecolorAsciiCss(s string) string {
	return colorReplacerAsciiCss.Replace(s)
}

func SearchAsciiIndex(s string) string {
	if val, ok := AsciiIndex[strings.ToLower(s)]; ok {
		return val + strings.ToLower(s)
	} else {
		return s + " is not a valid color. To see what colors are valid, run the sample? command."
	}
}

func PrintAsciiIndex() string {
	s := ""
	for k, v := range AsciiIndex {
		s = s + v + k + "\n"
	}
	return s
}

func AsciiRgb(r string, g string, b string) string {
	s := "\033[38;2;" + r + ";" + g + ";" + b + "m"
	return s
}

func AsciiRgbBackground(r string, g string, b string) string {
	s := "\033[48;2;" + r + ";" + g + ";" + b + "m"
	return s
}

var AsciiIndex = map[string]string{
	"sol.yellow":               HexValtoAscii(sol.Yellow),
	"sol.orange":               HexValtoAscii(sol.Orange),
	"sol.red":                  HexValtoAscii(sol.Red),
	"sol.magenta":              HexValtoAscii(sol.Magenta),
	"sol.violet":               HexValtoAscii(sol.Violet),
	"sol.blue":                 HexValtoAscii(sol.Blue),
	"sol.cyan":                 HexValtoAscii(sol.Cyan),
	"sol.green":                HexValtoAscii(sol.Green),
	"css.aliceblue":            HexValtoAscii(css.AliceBlue),
	"css.antiquewhite":         HexValtoAscii(css.AntiqueWhite),
	"css.aqua":                 HexValtoAscii(css.Aqua),
	"css.aquamarine":           HexValtoAscii(css.AquaMarine),
	"css.azure":                HexValtoAscii(css.Azure),
	"css.beige":                HexValtoAscii(css.Beige),
	"css.bisque":               HexValtoAscii(css.Bisque),
	"css.black":                HexValtoAscii(css.Black),
	"css.blanchedalmond":       HexValtoAscii(css.BlanchedAlmond),
	"css.blue":                 HexValtoAscii(css.Blue),
	"css.blueviolet":           HexValtoAscii(css.BlueViolet),
	"css.brown":                HexValtoAscii(css.Brown),
	"css.burlywood":            HexValtoAscii(css.BurlyWood),
	"css.cadetblue":            HexValtoAscii(css.CadetBlue),
	"css.chartreuse":           HexValtoAscii(css.Chartreuse),
	"css.chocolate":            HexValtoAscii(css.Chocolate),
	"css.coral":                HexValtoAscii(css.Coral),
	"css.cornflowerblue":       HexValtoAscii(css.CornFlowerBlue),
	"css.cornsilk":             HexValtoAscii(css.CornSilk),
	"css.crimson":              HexValtoAscii(css.Crimson),
	"css.cyan":                 HexValtoAscii(css.Cyan),
	"css.darkblue":             HexValtoAscii(css.DarkBlue),
	"css.darkcyan":             HexValtoAscii(css.DarkCyan),
	"css.darkgoldenrod":        HexValtoAscii(css.DarkGoldenRod),
	"css.darkgray":             HexValtoAscii(css.DarkGray),
	"css.darkgreen":            HexValtoAscii(css.DarkGreen),
	"css.darkkhaki":            HexValtoAscii(css.DarkKhaki),
	"css.darkmagenta":          HexValtoAscii(css.DarkMagenta),
	"css.darkolivegreen":       HexValtoAscii(css.DarkOliveGreen),
	"css.darkorange":           HexValtoAscii(css.DarkOrange),
	"css.darkorchid":           HexValtoAscii(css.DarkOrchid),
	"css.darkred":              HexValtoAscii(css.DarkRed),
	"css.darksalmon":           HexValtoAscii(css.DarkSalmon),
	"css.darkseagreen":         HexValtoAscii(css.DarkSeaGreen),
	"css.darkslateblue":        HexValtoAscii(css.DarkSlateBlue),
	"css.darkslategray":        HexValtoAscii(css.DarkSlateGray),
	"css.darkturquoise":        HexValtoAscii(css.DarkTurquoise),
	"css.darkviolet":           HexValtoAscii(css.DarkViolet),
	"css.deeppink":             HexValtoAscii(css.DeepPink),
	"css.deepskyblue":          HexValtoAscii(css.DeepSkyBlue),
	"css.dimgray":              HexValtoAscii(css.DimGray),
	"css.dodgerblue":           HexValtoAscii(css.DodgerBlue),
	"css.firebrick":            HexValtoAscii(css.Firebrick),
	"css.floralwhite":          HexValtoAscii(css.FloralWhite),
	"css.forestgreen":          HexValtoAscii(css.ForestGreen),
	"css.fuchsia":              HexValtoAscii(css.Fuchsia),
	"css.gainsboro":            HexValtoAscii(css.Gainsboro),
	"css.ghostwhite":           HexValtoAscii(css.GhostWhite),
	"css.gold":                 HexValtoAscii(css.Gold),
	"css.goldenrod":            HexValtoAscii(css.GoldenRod),
	"css.gray":                 HexValtoAscii(css.Gray),
	"css.green":                HexValtoAscii(css.Green),
	"css.greenyellow":          HexValtoAscii(css.GreenYellow),
	"css.honeydew":             HexValtoAscii(css.Honeydew),
	"css.hotpink":              HexValtoAscii(css.HotPink),
	"css.indianred":            HexValtoAscii(css.IndianRed),
	"css.indigo":               HexValtoAscii(css.Indigo),
	"css.ivory":                HexValtoAscii(css.Ivory),
	"css.khaki":                HexValtoAscii(css.Khaki),
	"css.lavender":             HexValtoAscii(css.Lavender),
	"css.lavenderblush":        HexValtoAscii(css.LavenderBlush),
	"css.lawngreen":            HexValtoAscii(css.LawnGreen),
	"css.lemonchiffon":         HexValtoAscii(css.LemonChiffon),
	"css.lightblue":            HexValtoAscii(css.LightBlue),
	"css.lightcoral":           HexValtoAscii(css.LightCoral),
	"css.lightcyan":            HexValtoAscii(css.LightCyan),
	"css.lightgoldenrodyellow": HexValtoAscii(css.LightGoldenRodYellow),
	"css.lightgreen":           HexValtoAscii(css.LightGreen),
	"css.lightgrey":            HexValtoAscii(css.LightGrey),
	"css.lightpink":            HexValtoAscii(css.LightPink),
	"css.lightsalmon":          HexValtoAscii(css.LightSalmon),
	"css.lightseagreen":        HexValtoAscii(css.LightSeaGreen),
	"css.lightskyblue":         HexValtoAscii(css.LightSkyBlue),
	"css.lightslategray":       HexValtoAscii(css.LightSlateGray),
	"css.lightsteelblue":       HexValtoAscii(css.LightSteelBlue),
	"css.lightyellow":          HexValtoAscii(css.LightYellow),
	"css.lime":                 HexValtoAscii(css.Lime),
	"css.limegreen":            HexValtoAscii(css.LimeGreen),
	"css.linen":                HexValtoAscii(css.Linen),
	"css.magenta":              HexValtoAscii(css.Magenta),
	"css.maroon":               HexValtoAscii(css.Maroon),
	"css.mediumaquamarine":     HexValtoAscii(css.MediumAquaMarine),
	"css.mediumblue":           HexValtoAscii(css.MediumBlue),
	"css.mediumorchid":         HexValtoAscii(css.MediumOrchid),
	"css.mediumpurple":         HexValtoAscii(css.MediumPurple),
	"css.mediumseagreen":       HexValtoAscii(css.MediumSeaGreen),
	"css.mediumslateblue":      HexValtoAscii(css.MediumSlateBlue),
	"css.mediumspringgreen":    HexValtoAscii(css.MediumSpringGreen),
	"css.mediumturquoise":      HexValtoAscii(css.MediumTurquoise),
	"css.mediumvioletred":      HexValtoAscii(css.MediumVioletRed),
	"css.midnightblue":         HexValtoAscii(css.MidnightBlue),
	"css.mintcream":            HexValtoAscii(css.MintCream),
	"css.mistyrose":            HexValtoAscii(css.MistyRose),
	"css.moccasin":             HexValtoAscii(css.Moccasin),
	"css.navajowhite":          HexValtoAscii(css.NavajoWhite),
	"css.navy":                 HexValtoAscii(css.Navy),
	"css.navyblue":             HexValtoAscii(css.NavyBlue),
	"css.oldlace":              HexValtoAscii(css.OldLace),
	"css.olive":                HexValtoAscii(css.Olive),
	"css.olivedrab":            HexValtoAscii(css.OliveDrab),
	"css.orange":               HexValtoAscii(css.Orange),
	"css.orangered":            HexValtoAscii(css.OrangeRed),
	"css.orchid":               HexValtoAscii(css.Orchid),
	"css.palegoldenrod":        HexValtoAscii(css.PaleGoldenRod),
	"css.palegreen":            HexValtoAscii(css.PaleGreen),
	"css.paleturquoise":        HexValtoAscii(css.PaleTurquoise),
	"css.palevioletred":        HexValtoAscii(css.PaleVioletRed),
	"css.papayawhip":           HexValtoAscii(css.PapayaWhip),
	"css.peachpuff":            HexValtoAscii(css.PeachPuff),
	"css.peru":                 HexValtoAscii(css.Peru),
	"css.pink":                 HexValtoAscii(css.Pink),
	"css.plum":                 HexValtoAscii(css.Plum),
	"css.powderblue":           HexValtoAscii(css.PowderBlue),
	"css.purple":               HexValtoAscii(css.Purple),
	"css.red":                  HexValtoAscii(css.RebeccaPurple),
	"css.rebeccapurple":        HexValtoAscii(css.Red),
	"css.rosybrown":            HexValtoAscii(css.RosyBrown),
	"css.royalblue":            HexValtoAscii(css.RoyalBlue),
	"css.saddlebrown":          HexValtoAscii(css.SaddleBrown),
	"css.salmon":               HexValtoAscii(css.Salmon),
	"css.sandybrown":           HexValtoAscii(css.SandyBrown),
	"css.seagreen":             HexValtoAscii(css.SeaGreen),
	"css.seashell":             HexValtoAscii(css.SeaShell),
	"css.sienna":               HexValtoAscii(css.Sienna),
	"css.silver":               HexValtoAscii(css.Silver),
	"css.skyblue":              HexValtoAscii(css.SkyBlue),
	"css.slateblue":            HexValtoAscii(css.SlateBlue),
	"css.slategray":            HexValtoAscii(css.SlateGray),
	"css.snow":                 HexValtoAscii(css.Snow),
	"css.springgreen":          HexValtoAscii(css.SpringGreen),
	"css.steelblue":            HexValtoAscii(css.SteelBlue),
	"css.tan":                  HexValtoAscii(css.Tan),
	"css.teal":                 HexValtoAscii(css.Teal),
	"css.thistle":              HexValtoAscii(css.Thistle),
	"css.tomato":               HexValtoAscii(css.Tomato),
	"css.turquoise":            HexValtoAscii(css.Turquoise),
	"css.violet":               HexValtoAscii(css.Violet),
	"css.wheat":                HexValtoAscii(css.Wheat),
	"css.white":                HexValtoAscii(css.White),
	"css.whitesmoke":           HexValtoAscii(css.WhiteSmoke),
	"css.yellow":               HexValtoAscii(css.Yellow),
	"css.yellowgreen":          HexValtoAscii(css.YellowGreen),
}

//Hex (int) colors
func SearchHexIndex(s string) int32 {
	if val, ok := HexIndex[strings.ToLower(s)]; ok {
		return val
	} else {
		return 0
	}
}

func PrintHexIndex() string {
	s := ""
	for k, v := range HexIndex {
		s = s + "#" + strconv.FormatInt(int64(v), 16) + k + "\n"
	}
	return s
}

var HexIndex = map[string]int32{
	"sol.yellow":               sol.Yellow,
	"sol.orange":               sol.Orange,
	"sol.red":                  sol.Red,
	"sol.magenta":              sol.Magenta,
	"sol.violet":               sol.Violet,
	"sol.blue":                 sol.Blue,
	"sol.cyan":                 sol.Cyan,
	"sol.green":                sol.Green,
	"css.aliceblue":            css.AliceBlue,
	"css.antiquewhite":         css.AntiqueWhite,
	"css.aqua":                 css.Aqua,
	"css.aquamarine":           css.AquaMarine,
	"css.azure":                css.Azure,
	"css.beige":                css.Beige,
	"css.bisque":               css.Bisque,
	"css.black":                css.Black,
	"css.blanchedalmond":       css.BlanchedAlmond,
	"css.blue":                 css.Blue,
	"css.blueviolet":           css.BlueViolet,
	"css.brown":                css.Brown,
	"css.burlywood":            css.BurlyWood,
	"css.cadetblue":            css.CadetBlue,
	"css.chartreuse":           css.Chartreuse,
	"css.chocolate":            css.Chocolate,
	"css.coral":                css.Coral,
	"css.cornflowerblue":       css.CornFlowerBlue,
	"css.cornsilk":             css.CornSilk,
	"css.crimson":              css.Crimson,
	"css.cyan":                 css.Cyan,
	"css.darkblue":             css.DarkBlue,
	"css.darkcyan":             css.DarkCyan,
	"css.darkgoldenrod":        css.DarkGoldenRod,
	"css.darkgray":             css.DarkGray,
	"css.darkgreen":            css.DarkGreen,
	"css.darkkhaki":            css.DarkKhaki,
	"css.darkmagenta":          css.DarkMagenta,
	"css.darkolivegreen":       css.DarkOliveGreen,
	"css.darkorange":           css.DarkOrange,
	"css.darkorchid":           css.DarkOrchid,
	"css.darkred":              css.DarkRed,
	"css.darksalmon":           css.DarkSalmon,
	"css.darkseagreen":         css.DarkSeaGreen,
	"css.darkslateblue":        css.DarkSlateBlue,
	"css.darkslategray":        css.DarkSlateGray,
	"css.darkturquoise":        css.DarkTurquoise,
	"css.darkviolet":           css.DarkViolet,
	"css.deeppink":             css.DeepPink,
	"css.deepskyblue":          css.DeepSkyBlue,
	"css.dimgray":              css.DimGray,
	"css.dodgerblue":           css.DodgerBlue,
	"css.firebrick":            css.Firebrick,
	"css.floralwhite":          css.FloralWhite,
	"css.forestgreen":          css.ForestGreen,
	"css.fuchsia":              css.Fuchsia,
	"css.gainsboro":            css.Gainsboro,
	"css.ghostwhite":           css.GhostWhite,
	"css.gold":                 css.Gold,
	"css.goldenrod":            css.GoldenRod,
	"css.gray":                 css.Gray,
	"css.green":                css.Green,
	"css.greenyellow":          css.GreenYellow,
	"css.honeydew":             css.Honeydew,
	"css.hotpink":              css.HotPink,
	"css.indianred":            css.IndianRed,
	"css.indigo":               css.Indigo,
	"css.ivory":                css.Ivory,
	"css.khaki":                css.Khaki,
	"css.lavender":             css.Lavender,
	"css.lavenderblush":        css.LavenderBlush,
	"css.lawngreen":            css.LawnGreen,
	"css.lemonchiffon":         css.LemonChiffon,
	"css.lightblue":            css.LightBlue,
	"css.lightcoral":           css.LightCoral,
	"css.lightcyan":            css.LightCyan,
	"css.lightgoldenrodyellow": css.LightGoldenRodYellow,
	"css.lightgreen":           css.LightGreen,
	"css.lightgrey":            css.LightGrey,
	"css.lightpink":            css.LightPink,
	"css.lightsalmon":          css.LightSalmon,
	"css.lightseagreen":        css.LightSeaGreen,
	"css.lightskyblue":         css.LightSkyBlue,
	"css.lightslategray":       css.LightSlateGray,
	"css.lightsteelblue":       css.LightSteelBlue,
	"css.lightyellow":          css.LightYellow,
	"css.lime":                 css.Lime,
	"css.limegreen":            css.LimeGreen,
	"css.linen":                css.Linen,
	"css.magenta":              css.Magenta,
	"css.maroon":               css.Maroon,
	"css.mediumaquamarine":     css.MediumAquaMarine,
	"css.mediumblue":           css.MediumBlue,
	"css.mediumorchid":         css.MediumOrchid,
	"css.mediumpurple":         css.MediumPurple,
	"css.mediumseagreen":       css.MediumSeaGreen,
	"css.mediumslateblue":      css.MediumSlateBlue,
	"css.mediumspringgreen":    css.MediumSpringGreen,
	"css.mediumturquoise":      css.MediumTurquoise,
	"css.mediumvioletred":      css.MediumVioletRed,
	"css.midnightblue":         css.MidnightBlue,
	"css.mintcream":            css.MintCream,
	"css.mistyrose":            css.MistyRose,
	"css.moccasin":             css.Moccasin,
	"css.navajowhite":          css.NavajoWhite,
	"css.navy":                 css.Navy,
	"css.navyblue":             css.NavyBlue,
	"css.oldlace":              css.OldLace,
	"css.olive":                css.Olive,
	"css.olivedrab":            css.OliveDrab,
	"css.orange":               css.Orange,
	"css.orangered":            css.OrangeRed,
	"css.orchid":               css.Orchid,
	"css.palegoldenrod":        css.PaleGoldenRod,
	"css.palegreen":            css.PaleGreen,
	"css.paleturquoise":        css.PaleTurquoise,
	"css.palevioletred":        css.PaleVioletRed,
	"css.papayawhip":           css.PapayaWhip,
	"css.peachpuff":            css.PeachPuff,
	"css.peru":                 css.Peru,
	"css.pink":                 css.Pink,
	"css.plum":                 css.Plum,
	"css.powderblue":           css.PowderBlue,
	"css.purple":               css.Purple,
	"css.red":                  css.RebeccaPurple,
	"css.rosybrown":            css.Red,
	"css.royalblue":            css.RosyBrown,
	"css.saddlebrown":          css.RoyalBlue,
	"css.salmon":               css.SaddleBrown,
	"css.sandybrown":           css.Salmon,
	"css.seagreen":             css.SandyBrown,
	"css.seashell":             css.SeaGreen,
	"css.sienna":               css.SeaShell,
	"css.silver":               css.Sienna,
	"css.skyblue":              css.Silver,
	"css.slateblue":            css.SkyBlue,
	"css.slategray":            css.SlateBlue,
	"css.snow":                 css.SlateGray,
	"css.springgreen":          css.Snow,
	"css.steelblue":            css.SpringGreen,
	"css.tan":                  css.SteelBlue,
	"css.teal":                 css.Tan,
	"css.thistle":              css.Teal,
	"css.tomato":               css.Thistle,
	"css.turquoise":            css.Tomato,
	"css.violet":               css.Turquoise,
	"css.wheat":                css.Violet,
	"css.white":                css.Wheat,
	"css.whitesmoke":           css.White,
	"css.yellow":               css.WhiteSmoke,
	"css.yellowgreen":          css.Yellow,
}
