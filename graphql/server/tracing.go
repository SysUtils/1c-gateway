package server

import (
	"context"
	"fmt"
	"github.com/graph-gophers/graphql-go/errors"
	"github.com/graph-gophers/graphql-go/introspection"
	"github.com/graph-gophers/graphql-go/trace"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type Tracer struct {
	metrics *prometheus.HistogramVec
}

func NewTracer(metrics *prometheus.HistogramVec) *Tracer {
	return &Tracer{
		metrics: metrics,
	}
}

func (t *Tracer) TraceQuery(ctx context.Context, queryString string, operationName string, variables map[string]interface{}, varTypes map[string]*introspection.Type) (context.Context, trace.TraceQueryFinishFunc) {
	span, spanCtx := opentracing.StartSpanFromContext(ctx, "GraphQL request")
	span.SetTag("graphql.query", queryString)

	if operationName != "" {
		span.SetTag("graphql.operationName", operationName)
	}

	if len(variables) != 0 {
		span.LogFields(log.Object("graphql.variables", variables))
	}

	return spanCtx, func(errs []*errors.QueryError) {
		if len(errs) > 0 {
			msg := errs[0].Error()
			if len(errs) > 1 {
				msg += fmt.Sprintf(" (and %d more errors)", len(errs)-1)
			}
			ext.Error.Set(span, true)
			span.SetTag("graphql.error", msg)
		}
		span.Finish()
	}
}

func noop(*errors.QueryError) {}

func (t *Tracer) TraceField(ctx context.Context, label, typeName, fieldName string, trivial bool, args map[string]interface{}) (context.Context, trace.TraceFieldFinishFunc) {
	if trivial {
		return ctx, noop
	}

	start := time.Now()

	span, spanCtx := opentracing.StartSpanFromContext(ctx, label)
	span.SetTag("graphql.type", typeName)
	span.SetTag("graphql.field", fieldName)
	for name, value := range args {
		span.SetTag("graphql.args."+name, value)
	}

	return spanCtx, func(err *errors.QueryError) {
		duration := time.Now().Sub(start)
		code := "200"
		if err != nil {
			ext.Error.Set(span, true)
			span.SetTag("graphql.error", err.Error())
			code = "500"
		}
		t.metrics.With(prometheus.Labels{"type": typeName, "resolver": fieldName, "code": code}).Observe(duration.Seconds())
		span.Finish()
	}
}
