package countrycode

import "regexp"

type Indicatif string
type CountryCode string

const (
	SN CountryCode = "SN"
)

type CountryInfo struct {
	Reg       *regexp.Regexp
	Indicatif Indicatif
}

var CountrysInfo = map[CountryCode]CountryInfo{
	SN: {regexp.MustCompile(`^7[0678][0-9]{7}$`), "+221"},
}

func (info *CountryInfo) VerifyPhoneNumber(code CountryCode, phone string) bool {
	return info.Reg.MatchString(phone)
}
