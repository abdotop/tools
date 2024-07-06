package jwt

import (
	"encoding/base64"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the utils package for GetSecret
type MockUtils struct {
	mock.Mock
}

func (m *MockUtils) GetSecret(secretName string) (string, error) {
	args := m.Called(secretName)
	return args.String(0), args.Error(1)
}

func TestLoadPrivateKeyFromEnv(t *testing.T) {
	// Setup
	key := "TEST_PRIVATE_KEY"
	expectedKey := `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCgLUFMFgAXdUBd
/df2TiEMtpbEEsquQBX8AiOsI1DSq99PZ3UkJddmQCVFCrXV/AqHSonMEX3wuS0j
3p8eoo7nCk60OHJVpbLnnttJ4riFNG5QA7FOjm+L6LUsyrc5fiClKkCfRltxWMw6
aVSb+ykFOYcBlT6IrTJqK9OalhKiVbtCWSIjMFJDEraTgC9b9UraSsXzQHxil/94
qjZhnJOhTySq1hqLSTR0OZtEU0PVg9JzMWgORDUJ9OTeBShKICflGhp+bEVwR94s
sX2gH41cXPXiZLzoRj/CApNieicsGO+uTvl/Hbelf1KS0k7uTc+qKywd7WMXuZM6
VSr0OmPpAgMBAAECggEACGERrkWk29VLCJe1x1ljNndrR+bNQkQROlmQzOmciebK
k+xvNvTBUcSnhnb21/K7SMRMFNgx6MjBHSQhWhCKLfkusJQW3Bbi1uXLJjB1c2dw
oWAuE3RVTsH9z1Hr95aa5dUB5Cshr7pj1abqsY24IcE2iC1K5Aq+bUotY8P3xdTl
z9WvlQ7LP+hR7hXXDZFFdi0o6VDOzcrCg1LFN54HMI49MlBjeyEIqxeXfkZ6LK+r
sFSrk6AR2eLnzgk9LoBjEbAyIQ+aTbhIdR4Fhltt+J13S8x/w18IcBuyxqnCChP0
fv7Cwkcl828rJ72pEJ+Wn0nNUbfCxxJWW581HsDEMQKBgQDX8C4MTOr4wTtTznaT
5gxCQzbNI0ww5minhqJMC3e+09ieqLz5wn8UNL/BJWZmc6h2oAxGEsLz6nGslBgH
knOyRSH3xPtemkcXsMhdcdNlD5QAAH2W5cbq0eOmWvMKuzhwRTqK24n0DZmRZbfS
FIjSlarZVJMt3jDQSZgVoBqWeQKBgQC95LqXI7kmkQqAOaAYonlVPmgSpnoT3c/8
EOLdwOVKlvooqmGmG6Urbgq52Enq+DC8TxJvQQ8teCC5VXgQmT15eWHH2W32Wk3c
HjgGC7Wm/DXdh33gYSOtmWPAznoE7Ci6GwUpY1uYKoFE0YiHo1nWC7SAQnjVKete
pqj/V0mc8QKBgFY80FcPPOPdX9EnaknEnO61oyQnzZzV18PXy+csyUTHnAI5B0eD
unaaXl5Hjm6qEARYBK7TfVImNgGjHzme7l+qplcqLu1oFa5LZqmS9Maugv/BMmba
GyfabN0aA2gsvuxvaWLdLGnwKH1drxzSIJZhOVsUILc3vizgx0ETsKqZAoGARzh2
UyFu2+wiSB1U0mh9oV8aoYQW3USgMSYUMJ+cX/FoOcBNh4Nu228WTsW0U5BqlvYB
MADytFcNzvUjZvZSfDDuX4pJF9CeyyP6VyolI1AM091xdKp6Oy4An9vRH++DBuoD
cZZ2UAgBG9KKpnS5yPHy7PgBYdGrGYDgeYQ/c8ECgYEAg2P7aC99PwijTysJ9Ckc
9zkYs7/Tn7fcc0uuVk5k2ECtMVXqFX2EmDvQtznVJoEe77PBgJsuzuKK9gdXvQnc
pOhZMsRY8Zzqeo9dpnG2VaJqbi0IfiYmIeN/cXg6gjxmv9Ml0yM/FJU7et+l4gf7
ss3EEpCrzbIhrDZgb0+aMCk=
-----END PRIVATE KEY-----
`
	os.Setenv(key, base64.StdEncoding.EncodeToString([]byte(expectedKey)))
	defer os.Unsetenv(key)

	j := &jwt_tools{}

	// Test
	err := j.LoadPrivateKeyFromEnv(key)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, j.privateKey)
}

func TestLoadPublicKeyFromEnv(t *testing.T) {
	// Setup
	key := "TEST_PUBLIC_KEY"
	expectedKey := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoC1BTBYAF3VAXf3X9k4h
DLaWxBLKrkAV/AIjrCNQ0qvfT2d1JCXXZkAlRQq11fwKh0qJzBF98LktI96fHqKO
5wpOtDhyVaWy557bSeK4hTRuUAOxTo5vi+i1LMq3OX4gpSpAn0ZbcVjMOmlUm/sp
BTmHAZU+iK0yaivTmpYSolW7QlkiIzBSQxK2k4AvW/VK2krF80B8Ypf/eKo2YZyT
oU8kqtYai0k0dDmbRFND1YPSczFoDkQ1CfTk3gUoSiAn5RoafmxFcEfeLLF9oB+N
XFz14mS86EY/wgKTYnonLBjvrk75fx23pX9SktJO7k3PqissHe1jF7mTOlUq9Dpj
6QIDAQAB
-----END PUBLIC KEY-----
`
	os.Setenv(key, base64.StdEncoding.EncodeToString([]byte(expectedKey)))
	defer os.Unsetenv(key)

	j := &jwt_tools{}

	// Test
	err := j.LoadPublicKeyFromEnv(key)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, j.publicKey)
}

func TestGenerateToken(t *testing.T) {
	// Setup
	key := "TEST_PRIVATE_KEY"
	expectedKey := `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCgLUFMFgAXdUBd
/df2TiEMtpbEEsquQBX8AiOsI1DSq99PZ3UkJddmQCVFCrXV/AqHSonMEX3wuS0j
3p8eoo7nCk60OHJVpbLnnttJ4riFNG5QA7FOjm+L6LUsyrc5fiClKkCfRltxWMw6
aVSb+ykFOYcBlT6IrTJqK9OalhKiVbtCWSIjMFJDEraTgC9b9UraSsXzQHxil/94
qjZhnJOhTySq1hqLSTR0OZtEU0PVg9JzMWgORDUJ9OTeBShKICflGhp+bEVwR94s
sX2gH41cXPXiZLzoRj/CApNieicsGO+uTvl/Hbelf1KS0k7uTc+qKywd7WMXuZM6
VSr0OmPpAgMBAAECggEACGERrkWk29VLCJe1x1ljNndrR+bNQkQROlmQzOmciebK
k+xvNvTBUcSnhnb21/K7SMRMFNgx6MjBHSQhWhCKLfkusJQW3Bbi1uXLJjB1c2dw
oWAuE3RVTsH9z1Hr95aa5dUB5Cshr7pj1abqsY24IcE2iC1K5Aq+bUotY8P3xdTl
z9WvlQ7LP+hR7hXXDZFFdi0o6VDOzcrCg1LFN54HMI49MlBjeyEIqxeXfkZ6LK+r
sFSrk6AR2eLnzgk9LoBjEbAyIQ+aTbhIdR4Fhltt+J13S8x/w18IcBuyxqnCChP0
fv7Cwkcl828rJ72pEJ+Wn0nNUbfCxxJWW581HsDEMQKBgQDX8C4MTOr4wTtTznaT
5gxCQzbNI0ww5minhqJMC3e+09ieqLz5wn8UNL/BJWZmc6h2oAxGEsLz6nGslBgH
knOyRSH3xPtemkcXsMhdcdNlD5QAAH2W5cbq0eOmWvMKuzhwRTqK24n0DZmRZbfS
FIjSlarZVJMt3jDQSZgVoBqWeQKBgQC95LqXI7kmkQqAOaAYonlVPmgSpnoT3c/8
EOLdwOVKlvooqmGmG6Urbgq52Enq+DC8TxJvQQ8teCC5VXgQmT15eWHH2W32Wk3c
HjgGC7Wm/DXdh33gYSOtmWPAznoE7Ci6GwUpY1uYKoFE0YiHo1nWC7SAQnjVKete
pqj/V0mc8QKBgFY80FcPPOPdX9EnaknEnO61oyQnzZzV18PXy+csyUTHnAI5B0eD
unaaXl5Hjm6qEARYBK7TfVImNgGjHzme7l+qplcqLu1oFa5LZqmS9Maugv/BMmba
GyfabN0aA2gsvuxvaWLdLGnwKH1drxzSIJZhOVsUILc3vizgx0ETsKqZAoGARzh2
UyFu2+wiSB1U0mh9oV8aoYQW3USgMSYUMJ+cX/FoOcBNh4Nu228WTsW0U5BqlvYB
MADytFcNzvUjZvZSfDDuX4pJF9CeyyP6VyolI1AM091xdKp6Oy4An9vRH++DBuoD
cZZ2UAgBG9KKpnS5yPHy7PgBYdGrGYDgeYQ/c8ECgYEAg2P7aC99PwijTysJ9Ckc
9zkYs7/Tn7fcc0uuVk5k2ECtMVXqFX2EmDvQtznVJoEe77PBgJsuzuKK9gdXvQnc
pOhZMsRY8Zzqeo9dpnG2VaJqbi0IfiYmIeN/cXg6gjxmv9Ml0yM/FJU7et+l4gf7
ss3EEpCrzbIhrDZgb0+aMCk=
-----END PRIVATE KEY-----
`
	os.Setenv(key, base64.StdEncoding.EncodeToString([]byte(expectedKey)))
	defer os.Unsetenv(key)

	j := New(1)

	// Test
	err := j.LoadPrivateKeyFromEnv(key)

	assert.NoError(t, err)
	// Test
	token, err := j.GenerateToken("testData")

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateToken(t *testing.T) {
	// Setup
	key := "TEST_PRIVATE_KEY"
	expectedKey := `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCgLUFMFgAXdUBd
/df2TiEMtpbEEsquQBX8AiOsI1DSq99PZ3UkJddmQCVFCrXV/AqHSonMEX3wuS0j
3p8eoo7nCk60OHJVpbLnnttJ4riFNG5QA7FOjm+L6LUsyrc5fiClKkCfRltxWMw6
aVSb+ykFOYcBlT6IrTJqK9OalhKiVbtCWSIjMFJDEraTgC9b9UraSsXzQHxil/94
qjZhnJOhTySq1hqLSTR0OZtEU0PVg9JzMWgORDUJ9OTeBShKICflGhp+bEVwR94s
sX2gH41cXPXiZLzoRj/CApNieicsGO+uTvl/Hbelf1KS0k7uTc+qKywd7WMXuZM6
VSr0OmPpAgMBAAECggEACGERrkWk29VLCJe1x1ljNndrR+bNQkQROlmQzOmciebK
k+xvNvTBUcSnhnb21/K7SMRMFNgx6MjBHSQhWhCKLfkusJQW3Bbi1uXLJjB1c2dw
oWAuE3RVTsH9z1Hr95aa5dUB5Cshr7pj1abqsY24IcE2iC1K5Aq+bUotY8P3xdTl
z9WvlQ7LP+hR7hXXDZFFdi0o6VDOzcrCg1LFN54HMI49MlBjeyEIqxeXfkZ6LK+r
sFSrk6AR2eLnzgk9LoBjEbAyIQ+aTbhIdR4Fhltt+J13S8x/w18IcBuyxqnCChP0
fv7Cwkcl828rJ72pEJ+Wn0nNUbfCxxJWW581HsDEMQKBgQDX8C4MTOr4wTtTznaT
5gxCQzbNI0ww5minhqJMC3e+09ieqLz5wn8UNL/BJWZmc6h2oAxGEsLz6nGslBgH
knOyRSH3xPtemkcXsMhdcdNlD5QAAH2W5cbq0eOmWvMKuzhwRTqK24n0DZmRZbfS
FIjSlarZVJMt3jDQSZgVoBqWeQKBgQC95LqXI7kmkQqAOaAYonlVPmgSpnoT3c/8
EOLdwOVKlvooqmGmG6Urbgq52Enq+DC8TxJvQQ8teCC5VXgQmT15eWHH2W32Wk3c
HjgGC7Wm/DXdh33gYSOtmWPAznoE7Ci6GwUpY1uYKoFE0YiHo1nWC7SAQnjVKete
pqj/V0mc8QKBgFY80FcPPOPdX9EnaknEnO61oyQnzZzV18PXy+csyUTHnAI5B0eD
unaaXl5Hjm6qEARYBK7TfVImNgGjHzme7l+qplcqLu1oFa5LZqmS9Maugv/BMmba
GyfabN0aA2gsvuxvaWLdLGnwKH1drxzSIJZhOVsUILc3vizgx0ETsKqZAoGARzh2
UyFu2+wiSB1U0mh9oV8aoYQW3USgMSYUMJ+cX/FoOcBNh4Nu228WTsW0U5BqlvYB
MADytFcNzvUjZvZSfDDuX4pJF9CeyyP6VyolI1AM091xdKp6Oy4An9vRH++DBuoD
cZZ2UAgBG9KKpnS5yPHy7PgBYdGrGYDgeYQ/c8ECgYEAg2P7aC99PwijTysJ9Ckc
9zkYs7/Tn7fcc0uuVk5k2ECtMVXqFX2EmDvQtznVJoEe77PBgJsuzuKK9gdXvQnc
pOhZMsRY8Zzqeo9dpnG2VaJqbi0IfiYmIeN/cXg6gjxmv9Ml0yM/FJU7et+l4gf7
ss3EEpCrzbIhrDZgb0+aMCk=
-----END PRIVATE KEY-----
`
	os.Setenv(key, base64.StdEncoding.EncodeToString([]byte(expectedKey)))
	defer os.Unsetenv(key)

	// Setup
	key2 := "TEST_PUBLIC_KEY"
	expectedKey2 := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoC1BTBYAF3VAXf3X9k4h
DLaWxBLKrkAV/AIjrCNQ0qvfT2d1JCXXZkAlRQq11fwKh0qJzBF98LktI96fHqKO
5wpOtDhyVaWy557bSeK4hTRuUAOxTo5vi+i1LMq3OX4gpSpAn0ZbcVjMOmlUm/sp
BTmHAZU+iK0yaivTmpYSolW7QlkiIzBSQxK2k4AvW/VK2krF80B8Ypf/eKo2YZyT
oU8kqtYai0k0dDmbRFND1YPSczFoDkQ1CfTk3gUoSiAn5RoafmxFcEfeLLF9oB+N
XFz14mS86EY/wgKTYnonLBjvrk75fx23pX9SktJO7k3PqissHe1jF7mTOlUq9Dpj
6QIDAQAB
-----END PUBLIC KEY-----
`
	os.Setenv(key2, base64.StdEncoding.EncodeToString([]byte(expectedKey2)))
	defer os.Unsetenv(key2)

	j := New(1)

	// Test
	err := j.LoadPrivateKeyFromEnv(key)

	assert.NoError(t, err)

	// Test
	err = j.LoadPublicKeyFromEnv(key2)

	// Assert
	assert.NoError(t, err)

	// Test
	token, err := j.GenerateToken("testData")

	// Assert
	assert.NoError(t, err)

	// Test
	claims, err := j.ValidateToken(token)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, claims)
}

func TestLoadPrivateKeyFromSecretsManager(t *testing.T) {
	// Setup
	mockUtils := new(MockUtils)
	mockUtils.On("GetSecret", "TEST_SECRET").Return("fakePrivateKey", nil)

	j := &jwt_tools{}

	// Test
	err := j.LoadPrivateKeyFromSecretsManager("TEST_SECRET")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, j.privateKey)
	mockUtils.AssertExpectations(t)
}

func TestLoadPublicKeyFromSecretsManager(t *testing.T) {
	// Setup
	mockUtils := new(MockUtils)
	mockUtils.On("GetSecret", "TEST_SECRET").Return("fakePublicKey", nil)

	j := &jwt_tools{}

	// Test
	err := j.LoadPublicKeyFromSecretsManager("TEST_SECRET")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, j.publicKey)
	mockUtils.AssertExpectations(t)
}
