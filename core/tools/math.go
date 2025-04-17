package tools

import (
	"fmt"
	"github.com/zayn1510/goarchi/app/resources"
	"math"
	"sort"
	"strings"
)

const EarthRadiusKM = 6371.0

func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	dLat := toRadians(lat2 - lat1)
	dLon := toRadians(lon2 - lon1)

	lat1Rad := toRadians(lat1)
	lat2Rad := toRadians(lat2)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return EarthRadiusKM * c
}

func toRadians(degree float64) float64 {
	return degree * math.Pi / 180
}
func FormatJarak(jarak float64) string {
	if jarak < 0.5 {
		return fmt.Sprintf("%.0f m - Terdekat", jarak*1000)
	} else if jarak < 2 {
		return fmt.Sprintf("%.2f Km - Lumayan", jarak)
	}
	return fmt.Sprintf("%.2f Km - Jauh", jarak)
}
func RoundToDecimal(value float64, places int) float64 {
	multiplier := math.Pow(10, float64(places))
	return math.Round(value*multiplier) / multiplier
}
func GetMinHarga(data []*resources.DepotAlternatif) int {
	min := data[0].Harga
	for _, d := range data {
		if d.Harga < min {
			min = d.Harga
		}
	}
	return min
}

func GetMinJarak(data []*resources.DepotAlternatif) float64 {
	min := data[0].Distance
	for _, d := range data {
		if d.Distance < min {
			min = d.Distance
		}
	}
	return min
}

func GetMaxJarak(data []*resources.DepotAlternatif) float64 {
	maxJarak := data[0].Distance
	for _, d := range data {
		if d.Distance > maxJarak {
			maxJarak = d.Distance
		}
	}
	return maxJarak
}
func GetMaxDiskon(data []*resources.DepotAlternatif) int {
	max := data[0].Diskon
	for _, d := range data {
		if d.Diskon > max {
			max = d.Diskon
		}
	}
	return max
}

func GetMaxRating(data []*resources.DepotAlternatif) float64 {
	max := data[0].Rating
	for _, d := range data {
		if d.Rating > max {
			max = d.Rating
		}
	}
	return max
}
func SortingHasil(data []*resources.HasilSawResource, order string) {
	sort.Slice(data, func(i, j int) bool {
		if strings.ToLower(order) == "desc" {
			return data[i].Nilai > data[j].Nilai
		}
		return data[i].Nilai < data[j].Nilai
	})
}
