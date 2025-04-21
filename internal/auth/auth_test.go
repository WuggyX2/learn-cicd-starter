package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestValidHeader(t *testing.T) {
	testHeader := make(http.Header)
	token := "1234567"
	testHeader.Set("Authorization", "BorkBork "+token)
	testHeader.Set("Content-Type", "application/json")

	result, err := auth.GetAPIKey(testHeader)

	if err != nil {
		t.Errorf("Error retrieving token from auth header: %s", err)
		return
	}

	if result != token {
		t.Errorf("result does not match expected, was:%s expecetd: %s", result, token)
	}

}

func TestAuthorizationHeaderDoesNotExit(t *testing.T) {
	testHeader := make(http.Header)
	testHeader.Set("Content-Type", "application/json")

	_, err := auth.GetAPIKey(testHeader)

	if err == nil {
		t.Errorf(
			"There supposed to be an no auth header found error but function completed succesfully",
		)
	}
}

func TestAuthHeaderNotValidFormat(t *testing.T) {
	testHeader := make(http.Header)
	token := "1234567"
	testHeader.Set("Authorization", token)
	testHeader.Set("Content-Type", "application/json")

	_, err := auth.GetAPIKey(testHeader)

	if err == nil {
		t.Errorf("Error expected, but function completed succesfully.")
	}
}
