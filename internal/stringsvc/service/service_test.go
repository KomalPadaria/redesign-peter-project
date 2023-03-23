package service

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_service_Count(t *testing.T) {
	svc := New()
	Convey("Given some string value", t, func() {
		x := "hello how are you?"
		Convey("Call the Count function", func() {
			c := svc.Count(x)
			Convey("The value should be equal to the expected value", func() {
				So(c, ShouldEqual, 18)
			})
		})
	})
}

func Test_service_Uppercase(t *testing.T) {
	svc := New()
	Convey("Given some string value", t, func() {
		x := "hello how are you?"
		Convey("Call the Uppercase function", func() {
			u := svc.Uppercase(x)
			Convey("The value should be equal to the expected value", func() {
				So(u, ShouldEqual, "HELLO HOW ARE YOU?")
			})
		})
	})
}
