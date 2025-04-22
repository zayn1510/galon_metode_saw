package services

import (
	"github.com/zayn1510/goarchi/app/resources"
	"github.com/zayn1510/goarchi/config"
	"github.com/zayn1510/goarchi/core/tools"
	"gorm.io/gorm"
	"sort"
	"strings"
)

type AlternatifService struct {
	db *gorm.DB
}

func NewAlternatifService() *AlternatifService {
	return &AlternatifService{
		db: config.GetDB(),
	}
}

func (s *AlternatifService) ShowAlternatif(userid uint) (*resources.DepotData, error) {

	depotModel, err := NewDepotService().FindAll(0, 10, "")
	if err != nil {
		return nil, err
	}
	userDetail, err := NewUserLocationService().IsUserExist(userid)
	if err != nil {
		return nil, err
	}
	depot := resources.GetDepotResource(depotModel)
	// Data Alternatif
	responseAlternatif := make([]*resources.DepotAlternatif, len(depot))
	for i, value := range depot {
		userLat := userDetail.Latitude
		userLon := userDetail.Longitude
		depotLat := depot[i].Latitude
		depotLon := depot[i].Longitude
		jarak := tools.Haversine(userLat, userLon, depotLat, depotLon)
		responseAlternatif[i] = &resources.DepotAlternatif{
			ID:             value.ID,
			NamaDepot:      value.NamaDepot,
			Alamat:         value.Alamat,
			Harga:          value.Harga,
			Jarak:          tools.FormatJarak(jarak),
			Diskon:         value.Diskon,
			Rating:         value.Rating,
			Distance:       jarak,
			Latitude:       value.Latitude,
			Longitude:      value.Longitude,
			NomorHandphone: value.NomorHandphone,
			Foto:           value.Foto,
			UpdatedAt:      value.UpdatedAt,
		}

	}
	sort.Slice(responseAlternatif, func(i, j int) bool {
		return responseAlternatif[i].Distance < responseAlternatif[j].Distance
	})

	// Normalisasi Data
	minHarga := tools.GetMinHarga(responseAlternatif)
	minJarak := tools.GetMinJarak(responseAlternatif)
	maxDiskon := tools.GetMaxDiskon(responseAlternatif)
	maxRating := tools.GetMaxRating(responseAlternatif)
	responseNormalisasi := make([]*resources.NormalisasiResource, len(depot))

	kriteria, err := NewKriteriaService().FindAll(0, 10, "")
	if err != nil {
		return nil, err
	}
	bobotKriteria := make(map[string]float64)
	tipeKriteria := make(map[string]int)
	for _, k := range kriteria {
		bobotKriteria[strings.ToLower(k.Keterangan)] = k.Bobot
		tipeKriteria[k.Keterangan] = k.Tipe
	}
	for index, alternatif := range responseAlternatif {

		hargaNorm := float64(minHarga) / float64(alternatif.Harga)
		jarakNorm := minJarak / alternatif.Distance
		diskonNorm := float64(alternatif.Diskon) / float64(maxDiskon)
		ratingNorm := alternatif.Rating / maxRating
		responseNormalisasi[index] = &resources.NormalisasiResource{
			IDDepot:  alternatif.ID,
			Depot:    alternatif.NamaDepot,
			Harga:    hargaNorm,
			Jarak:    tools.RoundToDecimal(jarakNorm, 6),
			Rating:   ratingNorm,
			Diskon:   diskonNorm,
			Distance: alternatif.Distance,
		}
	}

	// Hasil Ranking Metode saw
	responseHasil := make([]*resources.HasilSawResource, len(responseNormalisasi))
	for index, normalisasi := range responseNormalisasi {
		preferensi := (normalisasi.Harga * bobotKriteria["harga"]) +
			(normalisasi.Jarak * bobotKriteria["jarak"]) +
			(normalisasi.Diskon * bobotKriteria["diskon"]) +
			(normalisasi.Rating * bobotKriteria["rating"]) // Metode saw
		responseHasil[index] = &resources.HasilSawResource{
			IDDepot:        normalisasi.IDDepot,
			Depot:          normalisasi.Depot,
			Harga:          responseAlternatif[index].Harga,
			Jarak:          responseAlternatif[index].Jarak,
			Diskon:         responseAlternatif[index].Diskon,
			Rating:         responseAlternatif[index].Rating,
			Nilai:          tools.RoundToDecimal(preferensi, 4),
			Image:          responseAlternatif[index].Foto,
			Alamat:         responseAlternatif[index].Alamat,
			Latitude:       responseAlternatif[index].Latitude,
			Longitude:      responseAlternatif[index].Longitude,
			NomorHandphone: responseAlternatif[index].NomorHandphone,
			UpdatedAt:      responseAlternatif[index].UpdatedAt,
		}
	}
	tools.SortingHasil(responseHasil, "desc")
	return &resources.DepotData{
		Alternatif:          responseAlternatif,
		NormalisasiResource: responseNormalisasi,
		Hasil:               responseHasil,
	}, nil
}
