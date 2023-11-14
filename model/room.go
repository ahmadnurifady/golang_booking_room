package model

import "time"

type Room struct {
	Id        string       `json:"id"`
	RoomType  string       `json:"roomType"`
	Capacity  string       `json:"capacity"`
	Facility  RoomFacility `json:"facility"`
	Status    string       `json:"status"` //untuk status hanya ada dua yaitu Available atau Booked
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
}

type RoomFacility struct {
	Id               string    `json:"id"`
	RoomDescription  string    `json:"description"`
	Fwifi            bool      `json:"wifi"`
	FsoundSystem     bool      `json:"soundSystem"`
	Fprojector       bool      `json:"projector"`
	FscreenProjector bool      `json:"screenProjector"`
	Fchairs          bool      `json:"chairs"`
	Ftables          bool      `json:"tables"`
	FsoundProof      bool      `json:"soundProof"`
	FsmonkingArea    bool      `json:"smokingArea"`
	Ftelevison       bool      `json:"television"`
	FAc              bool      `json:"ac"`
	Fbathroom        bool      `json:"bathroom"`
	FcoffeMaker      bool      `json:"coffe maker"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
