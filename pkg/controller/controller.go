package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/14jasimmtp/Url-Shortner/pkg/db"
	"github.com/14jasimmtp/Url-Shortner/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UrlHandler struct {
	DB       *gorm.DB
	validate *validator.Validate
}

type URLRequest struct {
	URL string `json:"url" validate:"url"`
}

func NewController(db *gorm.DB) *UrlHandler {
	return &UrlHandler{DB: db,validate: validator.New()}
}

func (u *UrlHandler) ShortUrl(c echo.Context) error {
	var req URLRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error": err.Error()})
	}
	if err := u.validate.Struct(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	var Key string
	if u.DB.Raw("SELECT key FROM u_rls WHERE url = ?", req.URL).Scan(&Key).RowsAffected >0{
		return c.JSON(http.StatusOK, echo.Map{"Message": "success", "URL": fmt.Sprintf("http://localhost:3000/%s", Key)})
	}

	Key, err := utils.GenerateRandomString()
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error": err.Error()})
	}
	if err := u.DB.Create(&db.URls{Key: Key, URL: req.URL}).Error; err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"Message": "success", "URL": fmt.Sprintf("http://localhost:3000/%s", Key)})
}

func (u *UrlHandler) GetUrl(c echo.Context) error {
	var url string
	Key := c.Param("ShotKey")
	query := u.DB.Raw(`SELECT url FROM u_rls WHERE Key = ?`, Key).Scan(&url)
	if query.Error != nil {
		errs := errors.New("page not found")
		return c.JSON(http.StatusBadGateway, echo.Map{"error": errs.Error(), "serverError": query.Error.Error()})
	}
	if query.RowsAffected < 1{
		return c.JSON(http.StatusNotFound,echo.Map{"error" :"page not found"})
	}
	return c.Redirect(http.StatusPermanentRedirect, url)
}
