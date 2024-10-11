package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/auth"
	"kingdom/model"
	"net/http"
)

type DomainDatabase interface {
	GetDomainByID(id uint) (*model.Domain, error)
	CreateDomain(domain *model.Domain) error
	GetDomains() ([]*model.Domain, error)
	UpdateDomain(domain *model.Domain) error
	DeleteDomain(id uint) error
	GetUserByID(id uint) (*model.User, error)
}

type DomainApi struct {
	DB DomainDatabase
}

func (a *DomainApi) CreateDomain(ctx *gin.Context) {
	user, _ := a.DB.GetUserByID(auth.GetUserID(ctx))
	if !user.Admin {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You can't access for this API"})
	}

	domain := &model.CreateDomain{}
	if err := ctx.ShouldBindJSON(domain); err == nil {
		internal := &model.Domain{
			Name:        domain.Name,
			Description: domain.Description,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateDomain(internal)); !success {
			return
		}
	}
	ctx.JSON(http.StatusCreated, domain)
}

func (a *DomainApi) GetDomainByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		domain, err := a.DB.GetDomainByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(http.StatusNotFound, err)
		}
		if domain != nil {
			ctx.JSON(http.StatusOK, domain)
		}
	})
}

func (a *DomainApi) GetDomains(ctx *gin.Context) {
	domains, err := a.DB.GetDomains()
	if success := SuccessOrAbort(ctx, 500, err); !success {
		ctx.JSON(http.StatusNotFound, err)
	}
	var resp []*model.Domain
	for _, domain := range domains {
		resp = append(resp, domain)
	}
	ctx.JSON(http.StatusOK, resp)
}

func (a *DomainApi) UpdateDomain(ctx *gin.Context) {
	user, _ := a.DB.GetUserByID(auth.GetUserID(ctx))
	if !user.Admin {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You can't access for this API"})
	}

	withID(ctx, "id", func(id uint) {
		var domain *model.UpdateDomain
		if err := ctx.ShouldBindJSON(domain); err == nil {
			oldDomain, err := a.DB.GetDomainByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				return
			}
			if oldDomain != nil {
				internal := &model.Domain{
					Name:        domain.Name,
					Description: domain.Description,
				}
				if success := SuccessOrAbort(ctx, 500, a.DB.UpdateDomain(internal)); !success {
					return
				}
				ctx.JSON(http.StatusOK, domain)
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Domain doesn't exist"})
		}
	})
}

func (a *DomainApi) DeleteDomain(ctx *gin.Context) {
	user, _ := a.DB.GetUserByID(auth.GetUserID(ctx))
	if !user.Admin {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You can't access for this API"})
	}

	withID(ctx, "id", func(id uint) {
		domain, err := a.DB.GetDomainByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if domain != nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteDomain(id)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"error": "Character was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Character doesn't exist"})
		}
	})
}
