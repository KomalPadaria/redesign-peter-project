package jwt

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateToken(t *testing.T) {
	Convey("Given username and email", t, func() {
		username := "testuser"
		email := "test@email.com"
		Convey("Call the CreateToken function", func() {
			token, err := CreateToken(username, email)
			Convey("Error should be nil and the token value should not be empty", func() {
				So(err, ShouldBeNil)
				So(token, ShouldNotBeEmpty)
			})
		})
	})
}

func TestVerifyToken(t *testing.T) {
	Convey("Given username and email", t, func() {
		username := "testuser"
		email := "test@email.com"
		Convey("Call the CreateToken function", func() {
			token, _ := CreateToken(username, email)
			Convey("Call the VerifyToken function", func() {
				claims, err := VerifyToken(token)
				Convey("Error should be nil", func() {
					So(err, ShouldBeNil)
					Convey("Claims should contain expected values", func() {
						So(claims.Username, ShouldEqual, username)
						So(claims.Email, ShouldEqual, email)
					})
				})

			})
		})
	})
}
