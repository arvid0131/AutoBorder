package AutoBorder

import "fmt"

type QR_String struct {
	amt     float64
	com     string
	itm     string
	bitmask uint8
}

func generate_msg(qr QR_String) string {
	return fmt.Sprintf("C1235700653;%.2f;%s-%s;%b", qr.amt, qr.com, qr.itm, qr.bitmask)
}
