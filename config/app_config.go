package config

const (
	//auth
	UserSesion = "user"

	DownloadReport = "/download"
	UserAdmin      = "08f0cea4-2d3e-40e6-b2ee-18c7ec5507d8"

	AuthGroup        = "/auth"
	AuthRegister     = "/register"
	AuthLogin        = "/login"
	AuthRefreshToken = "/refresh-token"

	//User
	UserGroup  = "/users"
	UserPost   = "/"
	UserGet    = "/:id"
	UserDelete = "/:id"
	UserGetAll = "/"
	UserUpdate = "/"

	//booking
	BookingGroup          = "/booking"
	BookingPost           = "/"
	BookingGet            = "/:id"
	BookingGetAll         = "/"
	BookingGetAllByStatus = "/status/:status"
	Approval              = "/approval"

	//room
	RoomGroup         = "/rooms"
	RoomPost          = "/create"
	RoomGetByroomType = "/" //query
	RoomGetAll        = "/get"
	RoomGetById       = "/:id"
	RoomGetByStatus   = "/status"
	RoomDelete        = "/:id"
	RoomUpdate        = "/:id"
	RoomUpdateStatus  = "/status/:id"
)
