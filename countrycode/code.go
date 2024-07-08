package countrycode

import "regexp"

type Indicatif string
type CountryCode string

const (
	SN CountryCode = "SN"
)

type CountryInfo struct {
	reg       *regexp.Regexp
	indicatif Indicatif
}

var CountrysInfo = map[CountryCode]CountryInfo{
	SN: {regexp.MustCompile(`^7[0678][0-9]{7}$`), "+221"},
}

func VerifyPhoneNumber(code CountryCode, phone string) (Indicatif, bool) {
	info, ok := CountrysInfo[code]
	if !ok {
		return "", false
	}
	return info.indicatif, info.reg.MatchString(phone)
}
