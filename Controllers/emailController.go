package Controllers

import (
	"github.com/FaisalMashuri/emailServices/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmailController struct {
	emailService domain.EmailDomain
}

func NewEmailController(r *gin.RouterGroup, emailService domain.EmailDomain) {
	handler := &EmailController{
		emailService: emailService,
	}
	r.GET("/email", handler.SendEmail)
}

func (u *EmailController) SendEmail(c *gin.Context) {
	err := u.emailService.SendEmail("faisalmashuri16@gmail.com", "asdsad")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent"})
	return
}
