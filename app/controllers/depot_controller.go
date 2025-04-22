package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/zayn1510/goarchi/app/models"
	"github.com/zayn1510/goarchi/app/requests"
	"github.com/zayn1510/goarchi/app/resources"
	"github.com/zayn1510/goarchi/app/services"
	"github.com/zayn1510/goarchi/core/tools"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

const maxFileSize = 5 * 1024 * 1024

type DepotController struct {
	service *services.DepotService
}

func NewDepotController() *DepotController {
	return &DepotController{
		service: services.NewDepotService(),
	}
}

func (c *DepotController) Create(ctx *gin.Context) {
	var req requests.CreateDepotRequest
	if err := ctx.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		resources.BadRequest(ctx, err)
		return
	}

	if err, validation := requests.Validate(req); err != nil {
		resources.BadRequest(ctx, validation)
		return
	}
	filePath := ""
	fileNameNew := "default.png"
	pathName := fmt.Sprintf("./public/uploads")
	if req.Foto != nil {
		ext := strings.ToLower(filepath.Ext(req.Foto.Filename))
		if !allowedExtensions[ext] {
			resources.BadRequest(ctx, fmt.Errorf("unsupported file type: %s", ext))
			return
		}

		filePath = fmt.Sprintf("%s", pathName)
		fileName, err := tools.SaveResizedImage(req.Foto, filePath)
		if err != nil {
			resources.InternalError(ctx, err)
			return
		}
		fileNameNew = fileName
	}

	data := models.Depot{
		KecamatanID:    req.KecamatanId,
		NamaDepot:      req.NamaDepot,
		Alamat:         req.Alamat,
		NomorHandphone: req.NomorHandphone,
		Harga:          req.Harga,
		Diskon:         req.Diskon,
		Rating:         req.Rating,
		Latitude:       req.Latitude,
		Longitude:      req.Longitude,
		Foto:           fileNameNew,
	}
	if err := c.service.Create(&data); err != nil {
		resources.InternalError(ctx, err)
		return
	}
	resources.Success(ctx, "success", req)
}

func (c *DepotController) Show(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)

	limitStr := ctx.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * limit

	filterStr := ctx.DefaultQuery("filter", "")
	data, err := c.service.FindAll(offset, limit, filterStr)
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	response := resources.GetDepotResource(data)
	totaldata, err := c.service.CountAll()
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	total := int(math.Ceil(float64(totaldata) / float64(limit)))
	resources.SuccessWithPaginaition(ctx, "success", response, &total, &page, &limit)
}

func (c *DepotController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	var req requests.UpdateDepotRequest
	if err := ctx.ShouldBindWith(&req, binding.FormMultipart); err != nil {
		resources.BadRequest(ctx, err)
		return
	}

	var filePath string
	filePath = "default.png"
	pathName := fmt.Sprintf("./public/uploads")

	fileNameNew := req.FotoLama
	if req.Foto != nil {
		ext := strings.ToLower(filepath.Ext(req.Foto.Filename))
		if !allowedExtensions[ext] {
			resources.BadRequest(ctx, fmt.Errorf("unsupported file type: %s", ext))
			return
		}
		filePath = fmt.Sprintf("%s", pathName)
		fileName, err := tools.SaveResizedImage(req.Foto, filePath)
		if err != nil {
			resources.InternalError(ctx, err)
			return
		}
		fileNameNew = fileName
		filePathOld := filepath.Join("./public/uploads", req.FotoLama)
		if req.FotoLama != "" {
			if _, err := os.Stat(filePathOld); err == nil {
				if err := os.Remove(filePathOld); err != nil {
					resources.InternalError(ctx, err)
					return
				}
			}
		}

	}
	// buat map untuk field yang akan diupdate
	data := make(map[string]interface{})
	if req.NamaDepot != nil {
		data["nama_depot"] = *req.NamaDepot
	}
	if req.Alamat != nil {
		data["alamat"] = *req.Alamat
	}
	if req.Latitude != nil {
		data["latitude"] = *req.Latitude
	}
	if req.Longitude != nil {
		data["longitude"] = *req.Longitude
	}
	if req.NomorHandphone != nil {
		data["nomor_handphone"] = *req.NomorHandphone
	}
	if req.Harga != nil {
		data["harga"] = *req.Harga
	}
	if req.Diskon != nil {
		data["diskon"] = *req.Diskon
	}
	if req.Rating != nil {
		data["rating"] = *req.Rating
	}
	if req.KecamatanId != nil {
		data["kecamatan_id"] = *req.KecamatanId
	}
	if filePath != "" {
		data["foto"] = fileNameNew
	}

	if err := c.service.Update(data, uint(id)); err != nil {
		resources.InternalError(ctx, err)
		return
	}

	resources.Success(ctx, "success", req)
}
func (c *DepotController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	depot, err := c.service.Delete(uint(id))
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	filePath := filepath.Join("./public/uploads", depot.Foto)
	if depot.Foto != "default" {
		if _, err := os.Stat(filePath); err == nil {
			if err := os.Remove(filePath); err != nil {
				resources.InternalError(ctx, err)
				return
			}
		}
	}

	resources.Success(ctx, "success")
}
func (c *DepotController) PreviewFile(ctx *gin.Context) {
	filename := ctx.Param("filename")

	// Tentukan path file foto
	filePath := filepath.Join("./public/uploads", filename)

	// Cek apakah file ada
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		resources.NotFound(ctx, fmt.Errorf("file not found"))
		return
	}
	// Kirim file sebagai response
	ctx.File(filePath)
}
func (c *DepotController) DetailDepotById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	depot, err := c.service.FindById(uint(id))
	if err != nil {
		resources.InternalError(ctx, err)
		return
	}
	detail := resources.NewDepotResource(*depot)

	resources.Success(ctx, "success", detail)
}
