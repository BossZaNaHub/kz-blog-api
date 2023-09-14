package rpservices

import (
	"github.com/gofiber/fiber/v2"
	rpmodels "github.com/kuroshibaz/app/role_permission/models"
)

func (svc *defaultService) CreatePermission(data *rpmodels.Permission) *fiber.Error {
	return svc.rp.CreatePermission(data)
}
