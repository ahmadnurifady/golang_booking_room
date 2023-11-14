CREATE DATABASE booking_room;

//post,get,getbyid,update
CREATE TABLE transaksi(
    id          UUID,
    userId      string,     //Foreign key ke user id
    roomType    string,        //Foreign key ke room id
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

CREATE TABLE facility(
    Id               string,    
	RoomDescription  string,    
	Fwifi            string,      
	FsoundSystem     string,      
	Fprojector       string,      
	FscreenProjector string,      
	Fchairs          string,      
	Ftables          string,      
	FsoundProof      string,      
	FsmonkingArea    string,      
	Ftelevison       string,      
	FAc              string,
    Fbathroom        string,      
	FcoffeMaker      string,      
	CreatedAt        time.Time, 
	UpdatedAt        time.Time, 
);