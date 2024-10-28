package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type DomainDatabase interface {
	GetDomainByID(id uint) (*model.Domain, error)
	CreateDomain(domain *model.Domain) error
	GetDomains() ([]*model.Domain, error)
	UpdateDomain(domain *model.Domain) error
	DeleteDomain(id uint) error
}

type DomainApi struct {
	DB DomainDatabase
}

// CreateDomain godoc
//
// @Summary Create and returns domain or nil
// @Description Permissions for Admin
// @Tags Domain
// @Accept json
// @Produce json
// @Param domain body model.CreateDomain true "Domain data"
// @Success 201 {object} model.DomainExternal "Domain details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /domain [post]
func (a *DomainApi) CreateDomain(ctx *gin.Context) {
	domain := &model.CreateDomain{}
	if err := ctx.ShouldBindJSON(domain); err == nil {
		internal := &model.Domain{
			Name:        domain.Name,
			Description: domain.Description,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateDomain(internal)); !success {
			return
		}
		ctx.JSON(http.StatusCreated, ToDomainExternal(internal))
	}
}

// GetDomainByID godoc
//
// @Summary Returns Domain by id
// @Description Retrieve Domain details using its ID
// @Tags Domain
// @Accept json
// @Produce json
// @Param id path int true "domain id"
// @Success 200 {object} model.DomainExternal "domain details"
// @Failure 404 {string} string "Domain not found"
// @Router /domain/{id} [get]
func (a *DomainApi) GetDomainByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		domain, err := a.DB.GetDomainByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		if domain != nil {
			ctx.JSON(http.StatusOK, ToDomainExternal(domain))
		}
	})
}

// GetDomains godoc
//
// @Summary Returns all domains
// @Description Return all domains
// @Tags Domain
// @Accept json
// @Produce json
// @Success 200 {object} model.DomainExternal "Domain details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /domain [get]
func (a *DomainApi) GetDomains(ctx *gin.Context) {
	domains, err := a.DB.GetDomains()
	if success := SuccessOrAbort(ctx, 500, err); !success {
		ctx.JSON(http.StatusNotFound, err)
	}
	var resp []*model.DomainExternal
	for _, domain := range domains {
		resp = append(resp, ToDomainExternal(domain))
	}
	ctx.JSON(http.StatusOK, resp)
}

// UpdateDomain Updates Domain by ID
//
// @Summary Updates Domain by ID or nil
// @Description Permissions for Admin
// @Tags Domain
// @Accept json
// @Produce json
// @Param id path int true "Domain id"
// @Param character body model.UpdateDomain true "Domain data"
// @Success 200 {object} model.DomainExternal "Domain details"
// @Failure 403 {string} string "You can't access for this API"
// @Failure 404 {string} string "Domain doesn't exist"
// @Router /domain/{id} [patch]
func (a *DomainApi) UpdateDomain(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var domain *model.UpdateDomain
		if err := ctx.Bind(&domain); err == nil {
			oldDomain, err := a.DB.GetDomainByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				return
			}
			if oldDomain != nil {
				internal := &model.Domain{
					ID:          oldDomain.ID,
					Name:        domain.Name,
					Description: domain.Description,
				}
				if success := SuccessOrAbort(ctx, 500, a.DB.UpdateDomain(internal)); !success {
					return
				}
				ctx.JSON(http.StatusOK, ToDomainExternal(internal))
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Domain doesn't exist"})
		}
	})
}

// DeleteDomain Deletes Domain by ID
//
// @Summary Deletes Domain by ID or returns nil
// @Description Permissions for Admin
// @Tags Domain
// @Accept json
// @Produce json
// @Param id path int true "Domain id"
// @Success 204
// @Failure 404 {string} string "Domain doesn't exist"
// @Failure 403 {string} string "You can't access for this API"
// @Router /domain/{id} [delete]
func (a *DomainApi) DeleteDomain(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		domain, err := a.DB.GetDomainByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if domain != nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteDomain(id)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"error": "Domain was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Domain doesn't exist"})
		}
	})
}

func ToDomainExternal(domain *model.Domain) *model.DomainExternal {
	return &model.DomainExternal{
		ID:          domain.ID,
		Name:        domain.Name,
		Description: domain.Description,
	}
}
