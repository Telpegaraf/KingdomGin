package api

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"io"
	"kingdom/model"
	"log"
	"net/http"
	"os"
)

type TraditionDatabase interface {
	GetTraditionByID(id uint) (*model.Tradition, error)
	GetTraditionByName(name string) (*model.Tradition, error)
	CreateTradition(Tradition *model.Tradition) error
	GetTraditions() ([]*model.Tradition, error)
	UpdateTradition(Tradition *model.Tradition) error
	DeleteTradition(id uint) error
}

type TraditionApi struct {
	DB TraditionDatabase
}

// CreateTradition godoc
//
// @Summary Create and returns Tradition or nil
// @Description Permissions for Admin
// @Tags Tradition
// @Accept json
// @Produce json
// @Param tradition body model.CreateTradition true "Tradition data"
// @Success 201 {object} model.TraditionExternal "Tradition details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /tradition [post]
func (a *TraditionApi) CreateTradition(ctx *gin.Context) {
	tradition := &model.CreateTradition{}
	if err := ctx.ShouldBindJSON(tradition); err == nil {
		internal := &model.Tradition{
			Name:        tradition.Name,
			Description: tradition.Description,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateTradition(internal)); !success {
			return
		}
		ctx.JSON(http.StatusCreated, ToTraditionExternal(internal))
	}
}

// LoadTradition godoc
//
// @Summary Create Tradition from csv file on server or nil
// @Description Permissions for Admin
// @Tags Tradition
// @Accept json
// @Produce json
// @Success 201 {object} model.TraditionExternal "Tradition details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /tradition/load [post]
func (a *TraditionApi) LoadTradition(ctx *gin.Context) {
	file, err := os.Open("./csv/Tradition.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	reader := csv.NewReader(file)
	reader.Comma = ';'
	var traditions []model.Tradition

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		tradition := model.Tradition{
			Name:        record[0],
			Description: record[1],
		}
		traditions = append(traditions, tradition)
		if existTradition, err := a.DB.GetTraditionByName(tradition.Name); err == nil && existTradition != nil {
			continue
		}
		err = a.DB.CreateTradition(&tradition)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
}

// GetTraditionByID godoc
//
// @Summary Returns Tradition by id
// @Description Retrieve Tradition details using its ID
// @Tags Tradition
// @Accept json
// @Produce json
// @Param id path int true "Tradition id"
// @Success 200 {object} model.Tradition "Tradition details"
// @Failure 404 {string} string "Tradition not found"
// @Router /tradition/{id} [get]
func (a *TraditionApi) GetTraditionByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		tradition, err := a.DB.GetTraditionByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		if tradition != nil {
			ctx.JSON(http.StatusOK, ToTraditionExternal(tradition))
		}
	})
}

// GetTraditions godoc
//
// @Summary Returns all Traditions
// @Description Return all Traditions
// @Tags Tradition
// @Accept json
// @Produce json
// @Success 200 {object} model.Tradition "Tradition details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /tradition [get]
func (a *TraditionApi) GetTraditions(ctx *gin.Context) {
	traditions, err := a.DB.GetTraditions()
	if success := SuccessOrAbort(ctx, 500, err); !success {
		ctx.JSON(http.StatusNotFound, err)
	}
	var resp []*model.TraditionExternal
	for _, tradition := range traditions {
		resp = append(resp, ToTraditionExternal(tradition))
	}
	ctx.JSON(http.StatusOK, resp)
}

// UpdateTradition Updates Tradition by ID
//
// @Summary Updates Tradition by ID or nil
// @Description Permissions for Admin
// @Tags Tradition
// @Accept json
// @Produce json
// @Param id path int true "Tradition id"
// @Param character body model.UpdateTradition true "Tradition data"
// @Success 200 {object} model.TraditionExternal "Tradition details"
// @Failure 403 {string} string "You can't access for this API"
// @Failure 404 {string} string "Tradition doesn't exist"
// @Router /tradition/{id} [patch]
func (a *TraditionApi) UpdateTradition(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var tradition *model.UpdateTradition
		if err := ctx.Bind(&tradition); err == nil {
			oldTradition, err := a.DB.GetTraditionByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				return
			}
			if oldTradition != nil {
				internal := &model.Tradition{
					ID:          oldTradition.ID,
					Name:        tradition.Name,
					Description: tradition.Description,
				}
				if success := SuccessOrAbort(ctx, 500, a.DB.UpdateTradition(internal)); !success {
					return
				}
				ctx.JSON(http.StatusOK, ToTraditionExternal(internal))
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Tradition doesn't exist"})
		}
	})
}

// DeleteTradition Deletes Tradition by ID
//
// @Summary Deletes Tradition by ID or returns nil
// @Description Permissions for Admin
// @Tags Tradition
// @Accept json
// @Produce json
// @Param id path int true "Tradition id"
// @Success 204
// @Failure 404 {string} string "Tradition doesn't exist"
// @Failure 403 {string} string "You can't access for this API"
// @Router /tradition/{id} [delete]
func (a *TraditionApi) DeleteTradition(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		tradition, err := a.DB.GetTraditionByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if tradition != nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteTradition(id)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"error": "Tradition was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Tradition doesn't exist"})
		}
	})
}

func ToTraditionExternal(tradition *model.Tradition) *model.TraditionExternal {
	return &model.TraditionExternal{
		ID:          tradition.ID,
		Name:        tradition.Name,
		Description: tradition.Description,
	}
}
