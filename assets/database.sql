CREATE DATABASE booking_room;

//post,get,getbyid,update
CREATE TABLE transaksi(
    id          UUID,
    userId      string,
    roomType    int,
    cek_in      time.time,
    cek_out     time.time,
    description string,
    status      string,     //pending, decline atau accept
)

//post,getbyid,update,delete
CREATE TABLE user(
    id          UUID,
    name        string,
    divisi      string,
    jabatan     string,         //manager, direktur itu di set di jabatan dan rolenya adalah admin
    email       string,
    password    string,
    role        string,         //admin, karyawan atau GA
);

//post,getbyid,getall,update
CREATE TABLE Room(
    id UUID,
	RoomType  string,       
	Capacity  string,      
	Facility  int,            //Foreign key ke Facility id
	Status    string,        //untuk status hanya ada dua yaitu Available atau Booked
	CreatedAt time.Time,    
	UpdatedAt time.Time,    
);

CREATE TABLE faciliti(
    Id               string,    
	RoomDescription  string,    
	Fwifi            bool,      
	FsoundSystem     bool,      
	Fprojector       bool,      
	FscreenProjector bool,      
	Fchairs          bool,      
	Ftables          bool,      
	FsoundProof      bool,      
	FsmonkingArea    bool,      
	Ftelevison       bool,      
	FAc              bool,
    Fbathroom        bool,      
	FcoffeMaker      bool,      
	CreatedAt        time.Time, 
	UpdatedAt        time.Time, 
);