package twiliootp

import (
	"errors"
	"fmt"

	"github.com/abdotop/tools/countrycode"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type TwilioClient struct {
	account_sid string
	auth_token  string
	service_sid string
	client      *twilio.RestClient
}

func New(account_sid, auth_token, service_sid string) *TwilioClient {
	return &TwilioClient{
		account_sid: account_sid,
		auth_token:  auth_token,
		service_sid: service_sid,
		client: twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: account_sid,
			Password: auth_token,
		}),
	}
}

func (tc *TwilioClient) SendOtp(indicatif countrycode.Indicatif, number string) (*openapi.VerifyV2Verification, error) {

	to := fmt.Sprintf("%s%s", indicatif, number)
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	return tc.client.VerifyV2.CreateVerification(tc.service_sid, params)
}

func (tc *TwilioClient) VerifyOtp(indicatif countrycode.Indicatif, number, code string) (*openapi.VerifyV2VerificationCheck, error) {
	to := fmt.Sprintf("%s%s", indicatif, number)
	if code == "" {
		return nil, errors.New("code cannot be empty")
	}
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(to)
	params.SetCode(code)

	return tc.client.VerifyV2.CreateVerificationCheck(tc.service_sid, params)
}
