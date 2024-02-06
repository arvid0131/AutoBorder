package AutoBorder

func parseCommittee(com string) {
	switch com {
	case "Café":
		qr.com = "c"
	case "AktU":
		qr.com = "a"
	case "D-shopen":
		qr.com = "d"
	case "D-chip":
		qr.com = "dc"
	case "Sexet - Alkoholfritt":
		qr.com = "s0"
	case "Sexet - Öl":
		qr.com = "s01"
	case "Sexet - Cider":
		qr.com = "s02"
	case "Sexet - Vin":
		qr.com = "s03"
	case "Sexet - Drink":
		qr.com = "s04"
	case "Sexet - Mat":
		qr.com = "s05"
	case "Sexet - Biljett":
		qr.com = "s06"
	case "UtEDischot":
		qr.com = "u"
	}
}
