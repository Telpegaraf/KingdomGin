package api

import (
	"github.com/gin-gonic/gin"
	"kingdom/model"
	"net/http"
)

type ActionDatabase interface {
	GetActionByID(id uint) (*model.Action, error)
	CreateAction(Action *model.Action) error
	GetActions() ([]*model.Action, error)
	UpdateAction(Action *model.Action) error
	DeleteAction(id uint) error
}

type ActionApi struct {
	DB ActionDatabase
}

// CreateAction godoc
//
// @Summary Create and returns Action or nil
// @Description Permissions for Admin
// @Tags Action
// @Accept json
// @Produce json
// @Param action body model.CreateAction true "Action data"
// @Success 201 {object} model.ActionExternal "Action details"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "You can't access for this API"
// @Router /action [post]
func (a *ActionApi) CreateAction(ctx *gin.Context) {
	action := &model.CreateAction{}
	if err := ctx.ShouldBindJSON(action); err == nil {
		internal := &model.Action{
			Name: action.Name,
		}
		if success := SuccessOrAbort(ctx, 500, a.DB.CreateAction(internal)); !success {
			return
		}
		ctx.JSON(http.StatusCreated, ToActionExternal(internal))
	}
}

// GetActionByID godoc
//
// @Summary Returns Action by id
// @Description Retrieve Action details using its ID
// @Tags Action
// @Accept json
// @Produce json
// @Param id path int true "Action id"
// @Success 200 {object} model.Action "Action details"
// @Failure 404 {string} string "Action not found"
// @Router /action/{id} [get]
func (a *ActionApi) GetActionByID(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		action, err := a.DB.GetActionByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			ctx.JSON(http.StatusNotFound, err)
			return
		}
		if action != nil {
			ctx.JSON(http.StatusOK, ToActionExternal(action))
		}
	})
}

// GetActions godoc
//
// @Summary Returns all Actions
// @Description Return all Actions
// @Tags Action
// @Accept json
// @Produce json
// @Success 200 {object} model.Action "Action details"
// @Failure 401 {string} string ""Unauthorized"
// @Router /action [get]
func (a *ActionApi) GetActions(ctx *gin.Context) {
	actions, err := a.DB.GetActions()
	if success := SuccessOrAbort(ctx, 500, err); !success {
		ctx.JSON(http.StatusNotFound, err)
	}
	var resp []*model.ActionExternal
	for _, action := range actions {
		resp = append(resp, ToActionExternal(action))
	}
	ctx.JSON(http.StatusOK, resp)
}

// UpdateAction Updates Action by ID
//
// @Summary Updates Action by ID or nil
// @Description Permissions for Admin
// @Tags Action
// @Accept json
// @Produce json
// @Param id path int true "Action id"
// @Param action body model.UpdateAction true "Action data"
// @Success 200 {object} model.ActionExternal "Action details"
// @Failure 403 {string} string "You can't access for this API"
// @Failure 404 {string} string "Action doesn't exist"
// @Router /action/{id} [patch]
func (a *ActionApi) UpdateAction(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		var action *model.UpdateAction
		if err := ctx.Bind(&action); err == nil {
			oldAction, err := a.DB.GetActionByID(id)
			if success := SuccessOrAbort(ctx, 500, err); !success {
				return
			}
			if oldAction != nil {
				internal := &model.Action{
					ID:   oldAction.ID,
					Name: action.Name,
				}
				if success := SuccessOrAbort(ctx, 500, a.DB.UpdateAction(internal)); !success {
					return
				}
				ctx.JSON(http.StatusOK, ToActionExternal(internal))
			}
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Action doesn't exist"})
		}
	})
}

// DeleteAction Deletes Action by ID
//
// @Summary Deletes Action by ID or returns nil
// @Description Permissions for Admin
// @Tags Action
// @Accept json
// @Produce json
// @Param id path int true "Action id"
// @Success 204
// @Failure 404 {string} string "Action doesn't exist"
// @Failure 403 {string} string "You can't access for this API"
// @Router /action/{id} [delete]
func (a *ActionApi) DeleteAction(ctx *gin.Context) {
	withID(ctx, "id", func(id uint) {
		action, err := a.DB.GetActionByID(id)
		if success := SuccessOrAbort(ctx, 500, err); !success {
			return
		}
		if action != nil {
			if success := SuccessOrAbort(ctx, 500, a.DB.DeleteAction(id)); !success {
				return
			}
			ctx.JSON(http.StatusNoContent, gin.H{"error": "Action was deleted"})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Action doesn't exist"})
		}
	})
}

func ToActionExternal(action *model.Action) *model.ActionExternal {
	return &model.ActionExternal{
		ID:   action.ID,
		Name: action.Name,
	}
}
