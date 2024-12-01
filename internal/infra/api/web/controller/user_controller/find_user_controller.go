package user_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/opentelemetry"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/rest_err"
	"github.com/tiago-g-sales/leilao-goexpert/internal/usecase/user_usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

const (		
	REQUESTNAMEOTEL = "UserController"
)

type UserController struct {
	userUseCase user_usecase.UserUseCaseInterface
	TemplateData *opentelemetry.TemplateData
	 
}


func NewUserController(userUseCase user_usecase.UserUseCaseInterface, templateData *opentelemetry.TemplateData) *UserController {
	return &UserController{
		userUseCase: userUseCase,
		TemplateData: templateData,
	}
}

func (u *UserController) FindUserById(c *gin.Context) {
	
	carrier := propagation.HeaderCarrier(c.Request.Header)
	ctx := c.Request.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	ctx, spanInicial := u.TemplateData.OTELTracer.Start(ctx, REQUESTNAMEOTEL +" SPAN_INICIAL ")
	spanInicial.End()

	ctx, span := u.TemplateData.OTELTracer.Start(ctx, REQUESTNAMEOTEL + " Initial request FindUserById")
	defer span.End()


	userId := c.Param("userId")
	if err :=  uuid.Validate(userId); err != nil {
		 errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field: "userId",	
			Message: "Invalid UUID value",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	userData, err :=  u.userUseCase.FindUserById(context.Background(), userId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	ctx, spanEnd := u.TemplateData.OTELTracer.Start(ctx, REQUESTNAMEOTEL + "Finish request FindUserById" )
	defer spanEnd.End()

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(c.Request.Header))

	c.JSON(http.StatusOK, userData)
}