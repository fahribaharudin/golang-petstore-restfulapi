package feature_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStoreAPet(t *testing.T) {

	// Feature	: Store a pet to the database from request data
	// Scenario	: The request data is valid

	Convey("Given a request to store a pet", t, func() {
		Convey("When Client send some data to the request", func() {
			// todo..

			Convey("And the request is valid", func() {
				// todo..
			})

			Convey("Then API store the request data to the database", func() {
				// todo..
			})

			Convey("Then Client should get 201 created http status from the API", func() {
				// todo..

				Convey("And Client should get success message indicating data was stored in the db", func() {
					// todo..
				})
			})
		})
	})
}
