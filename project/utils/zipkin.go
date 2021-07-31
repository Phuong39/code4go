package utils

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"log"
)

const Host = "192.168.140.128:9411"

func NewTracer(serviceName, ip string) opentracing.Tracer {
	// set up a span reporter
	reporter := zipkinhttp.NewReporter(fmt.Sprintf("http://%s/api/v2/spans", Host))
	defer reporter.Close()

	// create our local service endpoint
	endpoint, err := zipkin.NewEndpoint(serviceName, ip)
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}

	// initialize our tracer
	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
	}

	// use zipkin-go-opentracing to wrap our tracer
	tracer := zipkinot.Wrap(nativeTracer)

	// optionally set as Global OpenTracing tracer instance
	opentracing.SetGlobalTracer(tracer)

	return tracer
}
