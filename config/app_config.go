package config

const (
	//auth
	UserSesion = "user"

	UserAdmin = "24d66cb1-73ae-4b37-a53b-e646297fa21a"

	AuthGroup        = "/auth"
	AuthRegister     = "/register"
	AuthLogin        = "/login"
	AuthRefreshToken = "/refresh-token"

	//user
	UserGroup  = "/users"
	UserPost   = "/"
	UserGet    = "/:id"
	UserDelete = "/:id"
	UserGetAll = "/"
	UserUpdate = "/"

	BookingGroup          = "/booking"
	BookingPost           = "/"
	BookingGet            = "/:id"
	BookingGetAll         = "/"
	BookingGetAllByStatus = "/status/:status"
	Approval              = "/approval"

	RoomGroup = "/rooms"
	RoomPost  = "/"
	RoomGet   = "/:id"
)
