CREATE DATABASE booking_room;

//post,get,getbyid,update
CREATE TABLE transaksi(
    id UUID,
    karyawanId string,
    Kamar int, Foreign key Room struct
    cek_in time.time
    cek_out time.time
    keterangan string,
    status string,
)

//post,getbyid,update,delete
CREATE TABLE user(
    id UUID,
    name string,
    divisi string,
    jabatan string, (admin, karyawan, GA, manajer ,direktur)
    email string,
    password string,
);

//post,getbyid,getall,update
CREATE TABLE Room(
    id UUID,
    tipe_kamar string,
    addons string,

);