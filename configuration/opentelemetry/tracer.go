package opentelemetry

import (
	"go.opentelemetry.io/otel/trace"
)

type TemplateData struct {
	OTELTracer         trace.Tracer
}

func NewTemplateData( otelTracer trace.Tracer) *TemplateData {
	return &TemplateData{
		OTELTracer:         otelTracer,
	}
}

