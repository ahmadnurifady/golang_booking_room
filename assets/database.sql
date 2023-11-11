CREATE DATABASE booking_room;

CREATE TABLE admins(
    id UUID,
    username string UNIQUE,
    password string,
);

CREATE TABLE general_advisors(
    id UUID,
    username string UNIQUE,
    password string,
);

CREATE TABLE FORM(
    id UUID,
    karyawanId string,
    tipeKamar int,
    jadwal time.Time,
    keterangan string,
    status string,
)

CREATE TABLE karyawan(
    id UUID,
    name string,
    divisi string,
    jabatan string,
    contact string,
);

CREATE TABLE Room(
    id UUID,
    type string,
    addons string,
);