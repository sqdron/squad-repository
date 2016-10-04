package api

import (
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
	"testing"
)

func Test_Gtihub_Auth_URL(t *testing.T) {
	g := GithubAPI("qwerty", "")
	Convey("Test Authorization URL", t, func() {
		urlString := g.GetAuthUrl()
		r, _ := url.Parse(urlString)
		q := r.Query()
		So(urlString, ShouldNotBeNil)
		Convey("Check scope", func() {
			So(q.Get("scope"), ShouldEqual, "email repo")
		})
		Convey("Check client ID", func() {
			So(q.Get("client_id"), ShouldEqual, "qwerty")
		})
		Convey("Check state", func() {
			So(len(q.Get("state")), ShouldEqual, 10)
		})
	})
}
