package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Test_verifyToken(t *testing.T) {
	type args struct {
		tokenString string
		signingKey  []byte
	}
	tests := []struct {
		name    string
		args    args
		want    jwt.MapClaims
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := verifyToken(tt.args.tokenString, tt.args.signingKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("verifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verifyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createToken(t *testing.T) {
	type args struct {
		signingKey []byte
		username   string
		expiration time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createToken(tt.args.signingKey, tt.args.username, tt.args.expiration)
			if (err != nil) != tt.wantErr {
				t.Errorf("createToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("createToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
