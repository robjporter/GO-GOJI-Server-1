package routes

import (
	"net/http"

	"../controllers"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func Routes(m *web.Mux) {
	goji.Handle("/api/v1/spark/*", m)
	goji.Get("/api/v1/spark", http.RedirectHandler("/api/v1/spark/", 301))

	m.Get("/api/v1/spark/", controllers.SparkHome)
	m.Get("/api/v1/spark/display", controllers.SparkDisplay)

	// ROOMS
	m.Get("/api/v1/spark/rooms/list/all", controllers.SparkRoomsListAll)
	m.Get("/api/v1/spark/rooms/name/:name", controllers.SparkRoomsCheckRoomExists)
	m.Get("/api/v1/spark/rooms/id/:id", controllers.SparkRoomsGetRoomNameFromID)
	m.Get("/api/v1/spark/rooms/create/:name", controllers.SparkRoomsCreateRoomByName)
	m.Get("/api/v1/spark/rooms/change/:name/:name2", controllers.SparkRoomsChangeName)
	m.Get("/api/v1/spark/rooms/delete/:name", controllers.SparkRoomsDeleteRoom)

	// PEOPLE
	m.Get("/api/v1/spark/people/find/me", controllers.SparkPeopleFindMe)
	m.Get("/api/v1/spark/people/find/email/:email", controllers.SparkPeopleFindEmail)
	m.Get("/api/v1/spark/people/find/name/:name", controllers.SparkPeopleFindName)
	m.Get("/api/v1/spark/people/find/email/:email/id", controllers.SparkPeopleFindEmailReturnID)
	m.Get("/api/v1/spark/people/find/name/:name/id", controllers.SparkPeopleFindNameReturnID)
	m.Get("/api/v1/spark/people/find/email/:email/emails", controllers.SparkPeopleFindEmailReturnEmails)
	m.Get("/api/v1/spark/people/find/name/:name/emails", controllers.SparkPeopleFindNameReturnEmails)
	m.Get("/api/v1/spark/people/find/email/:email/name", controllers.SparkPeopleFindEmailReturnName)
	m.Get("/api/v1/spark/people/find/name/:name/name", controllers.SparkPeopleFindNameReturnName)
	m.Get("/api/v1/spark/people/find/email/:email/avatar", controllers.SparkPeopleFindEmailReturnAvatar)
	m.Get("/api/v1/spark/people/find/name/:name/avatar", controllers.SparkPeopleFindNameReturnAvatar)
	m.Get("/api/v1/spark/people/find/email/:email/created", controllers.SparkPeopleFindEmailReturnCreated)
	m.Get("/api/v1/spark/people/find/name/:name/created", controllers.SparkPeopleFindNameReturnCreated)

	// MEMBERSHIP
	m.Get("/api/v1/spark/membership/find/room/:id", controllers.SparkMembershipByRoomID)
	m.Get("/api/v1/spark/membership/find/person/:id", controllers.SparkMembershipByPersonID)
	m.Get("/api/v1/spark/membership/find/email/:id", controllers.SparkMembershipByEmail)
}
