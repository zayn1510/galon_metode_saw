package resources

import "github.com/zayn1510/goarchi/app/models"

type DepotResource struct {
	ID             uint    `json:"id"`
	NamaDepot      string  `json:"nama_depot"`
	Alamat         string  `json:"alamat"`
	KecamatanId    uint64  `json:"kecamatan_id"`
	KecamatanName  string  `json:"kecamatan_name"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	NomorHandphone string  `json:"nomor_handphone"`
	Harga          int     `json:"harga"`
	Diskon         int     `json:"diskon"`
	Foto           string  `json:"foto,omitempty"`
	Rating         float64 `json:"rating"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	DeletedAt      string  `json:"deleted_at"`
}

type DepotAlternatif struct {
	ID             uint    `json:"id"`
	NamaDepot      string  `json:"nama_depot"`
	Alamat         string  `json:"alamat"`
	Harga          int     `json:"harga"`
	Jarak          string  `json:"jarak"`
	Diskon         int     `json:"diskon"`
	Rating         float64 `json:"rating"`
	Latitude       float64 `json:"latitude,omitempty"`
	Longitude      float64 `json:"longitude,omitempty"`
	UserLat        float64 `json:"user_lat,omitempty"`
	UserLong       float64 `json:"user_long,omitempty"`
	KecamatanId    uint64  `json:"kecamatan_id,omitempty"`
	KecamatanName  string  `json:"kecamatan_name,omitempty"`
	Distance       float64 `json:"distance"`
	Foto           string  `json:"foto,omitempty"`
	NomorHandphone string  `json:"nomor_handphone"`
	UpdatedAt      string  `json:"updated_at"`
}

func NewDepotResource(m models.Depot) *DepotResource {
	var deletedAt string
	if m.DeletedAt.Valid {
		deletedAt = m.DeletedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &DepotResource{ // Mengembalikan pointer agar lebih ringan
		ID:             m.ID,
		NamaDepot:      m.NamaDepot,
		Alamat:         m.Alamat,
		KecamatanId:    m.KecamatanID,
		KecamatanName:  m.Kecamatan.NamaKecamatan,
		Latitude:       m.Latitude,
		Longitude:      m.Longitude,
		NomorHandphone: m.NomorHandphone,
		Harga:          m.Harga,
		Diskon:         m.Diskon,
		Rating:         m.Rating,
		Foto:           m.Foto,
		CreatedAt:      m.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:      m.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt:      deletedAt,
	}
}

func GetDepotResource(data []models.Depot) []*DepotResource {
	resources := make([]*DepotResource, len(data))
	for i, v := range data {
		resources[i] = NewDepotResource(v)
	}
	return resources
}
