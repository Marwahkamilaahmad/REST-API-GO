package models

type Product struct {
	Id int64 `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"type:varchar(300)" json:"nama_product"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
}

// type Pasien struct{
// 	Id int64 `gorm:"type:bigint;primaryKey" json:"id"`
// 	NamaPasien string `gorm:"type:varchar(300)" json:"nama_pasien"`
// 	Usia int64 `gorm:"type:varchar(300)" json:"usia"`
// 	JenisKelamin string `gorm:"type:varchar(300)" json:"jenis_kelamin"`
// 	Alamat string `gorm:"type:text" json:"alamat"`
// 	RiwayatPenyakit string `gorm:"type:text" json:"riwayat_penyakit"`
// 	NomorKamar string `gorm:"type:varchar(300)" json:"nomor_kamar"`
// 	RumahSakitID  int64  `gorm:"type:bigint;foreignKey:RumahSakit;references:Id" json:"rumah_sakit_id"`
// }

// type RumahSakit struct{
// 	Id int64 `gorm:"type:bigint;primaryKey" json:"id"`
// 	Nama string `gorm:"type:varchar(300)" json:"nama"`
// 	Alamat string `gorm:"type:text" json:"alamat"`
// 	KodePos string `gorm:"type:varchar(300)" json:"kode_pos"`
// 	Kota string `gorm:"type:varchar(300)" json:"kota"`

// }

type Pasien struct {
	ID           int64  `gorm:"type:bigint;primaryKey" json:"id"`
	NamaPasien   string `gorm:"type:varchar(255)" json:"nama_pasien"`  // Adjusted length to 255
	Usia         int    `gorm:"type:int" json:"usia"`               // Changed type to int
	JenisKelamin string `gorm:"type:varchar(10)" json:"jenis_kelamin"` // Adjusted length to 10 for common values
	Alamat      string `gorm:"type:text" json:"alamat"`
	RiwayatPenyakit string `gorm:"type:text" json:"riwayat_penyakit"`
	NomorKamar   string `gorm:"type:varchar(20)" json:"nomor_kamar"`   // Adjusted length to 20
	RumahSakitID int64  `gorm:"type:bigint;foreignKey:RumahSakitID;references:ID" json:"rumah_sakit_id"` // Renamed for consistency
  }
  
  type RumahSakit struct {
	ID       int64  `gorm:"type:bigint;primaryKey" json:"id"`
	Nama     string `gorm:"type:varchar(255)" json:"nama"`     // Adjusted length to 255
	Alamat   string `gorm:"type:text" json:"alamat"`
	KodePos  string `gorm:"type:varchar(10)" json:"kode_pos"` // Adjusted length to 10 for common postal code formats
	Kota     string `gorm:"type:varchar(255)" json:"kota"`    // Adjusted length to 255
  }
  