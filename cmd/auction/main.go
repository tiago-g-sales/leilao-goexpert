package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"time"

	"github.com/spf13/viper"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/database/mongodb"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/opentelemetry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

)


type TemplateData struct {
	Title              string
	ResponseTime       time.Duration
	ExternalCallMethod string
	ExternalCallURL    string
	Content            string
	RequestNameOTEL    string
	OTELTracer         trace.Tracer
}



func init() {
	viper.AutomaticEnv()
}

func main() {

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)


	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	otel_active := viper.GetString("OTEL")

	if otel_active != "false" {
		shutdown, err := opentelemetry.InitProvider(viper.GetString("OTEL_SERVICE_NAME"), viper.GetString("OTEL_EXPORTER_OTLP_ENDPOINT"))
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
		}()

		if err != nil {
			log.Fatal(err)
		}
	}

	tracer := otel.Tracer("microservice-tracer")

	templateData := &TemplateData{
		Title:              viper.GetString("TITLE"),
		ResponseTime:       time.Duration(viper.GetInt("RESPONSE_TIME")),
		ExternalCallURL:    viper.GetString("EXTERNAL_CALL_URL"),
		ExternalCallMethod: viper.GetString("EXTERNAL_CALL_METHOD"),
		RequestNameOTEL:    viper.GetString("REQUEST_NAME_OTEL"),
		OTELTracer:         tracer,
	}

	fmt.Println(templateData)

	_ , err :=  mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}


}