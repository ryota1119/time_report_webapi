package entities

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestSignedToken_String(t *testing.T) {
	tokenStr := "signed-token"
	signedToken := SignedToken(tokenStr)

	if signedToken.String() != tokenStr {
		t.Errorf("signedToken.String() = %s, want %s", signedToken.String(), tokenStr)
	}
}

func TestJti_String(t *testing.T) {
	tokenStr := "jti-token"
	jtiToken := Jti(tokenStr)

	if jtiToken.String() != tokenStr {
		t.Errorf("signedToken.String() = %s, want %s", jtiToken.String(), tokenStr)
	}
}

func TestNewAuthToken(test *testing.T) {
	accessToken := "accessToken"
	refreshToken := "refreshToken"
	expiredAt := time.Now().Add(15 * time.Minute)

	want := &AuthToken{
		AccessToken:  SignedToken(accessToken),
		RefreshToken: SignedToken(refreshToken),
		ExpiresAt:    expiredAt,
	}

	got := NewAuthToken(accessToken, refreshToken)

	if !cmp.Equal(want, got) {
		test.Errorf("NewAuthToken() = %+v, want %+v", got, want)
	}
}
