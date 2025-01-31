// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jaegerremotesampling

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/jaegertracing/jaeger/proto-gen/api_v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configtls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestNewExtension(t *testing.T) {
	// test
	cfg := testConfig()
	cfg.Source.File = filepath.Join("testdata", "strategy.json")
	e := newExtension(cfg, componenttest.NewNopTelemetrySettings())

	// verify
	assert.NotNil(t, e)
}

func TestStartAndShutdownLocalFile(t *testing.T) {
	// prepare
	cfg := testConfig()
	cfg.Source.File = filepath.Join("testdata", "strategy.json")

	e := newExtension(cfg, componenttest.NewNopTelemetrySettings())
	require.NotNil(t, e)
	require.NoError(t, e.Start(context.Background(), componenttest.NewNopHost()))

	// test and verify
	assert.NoError(t, e.Shutdown(context.Background()))
}

func TestStartAndCallAndShutdownRemote(t *testing.T) {
	for _, tc := range []struct {
		name                     string
		remoteClientHeaderConfig map[string]configopaque.String
	}{
		{
			name: "no configured header additions",
		},
		{
			name: "configured header additions",
			remoteClientHeaderConfig: map[string]configopaque.String{
				"testheadername":    "testheadervalue",
				"anotherheadername": "anotherheadervalue",
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {

			// prepare the socket the mock server will listen at
			lis, err := net.Listen("tcp", "127.0.0.1:0")
			require.NoError(t, err)

			// create the mock server
			server := grpc.NewServer()

			// register the service
			mockServer := &samplingServer{}
			api_v2.RegisterSamplingManagerServer(server, mockServer)

			go func() {
				err = server.Serve(lis)
				require.NoError(t, err)
			}()

			// create the config, pointing to the mock server
			cfg := testConfig()
			cfg.GRPCServerSettings.NetAddr.Endpoint = "127.0.0.1:0"
			cfg.Source.Remote = &configgrpc.GRPCClientSettings{
				Endpoint: fmt.Sprintf("127.0.0.1:%d", lis.Addr().(*net.TCPAddr).Port),
				TLSSetting: configtls.TLSClientSetting{
					Insecure: true, // test only
				},
				WaitForReady: true,
				Headers:      tc.remoteClientHeaderConfig,
			}

			// create the extension
			e := newExtension(cfg, componenttest.NewNopTelemetrySettings())
			require.NotNil(t, e)

			// start the server
			assert.NoError(t, e.Start(context.Background(), componenttest.NewNopHost()))

			// make a call
			resp, err := http.Get("http://127.0.0.1:5778/sampling?service=foo")
			assert.NoError(t, err)
			assert.Equal(t, 200, resp.StatusCode)

			// shut down the server
			assert.NoError(t, e.Shutdown(context.Background()))

			// verify observed calls
			assert.Len(t, mockServer.observedCalls, 1)
			singleCall := mockServer.observedCalls[0]
			assert.Equal(t, &api_v2.SamplingStrategyParameters{
				ServiceName: "foo",
			}, singleCall.params)
			md, ok := metadata.FromIncomingContext(singleCall.ctx)
			assert.True(t, ok)
			for expectedHeaderName, expectedHeaderValue := range tc.remoteClientHeaderConfig {
				assert.Equal(t, []string{string(expectedHeaderValue)}, md.Get(expectedHeaderName))
			}
		})
	}
}

type samplingServer struct {
	api_v2.UnimplementedSamplingManagerServer
	observedCalls []observedCall
}

type observedCall struct {
	ctx    context.Context
	params *api_v2.SamplingStrategyParameters
}

func (s *samplingServer) GetSamplingStrategy(ctx context.Context, params *api_v2.SamplingStrategyParameters) (*api_v2.SamplingStrategyResponse, error) {
	s.observedCalls = append(s.observedCalls, observedCall{
		ctx:    ctx,
		params: params,
	})
	return &api_v2.SamplingStrategyResponse{
		StrategyType: api_v2.SamplingStrategyType_PROBABILISTIC,
	}, nil
}

func testConfig() *Config {
	cfg := createDefaultConfig().(*Config)
	cfg.HTTPServerSettings.Endpoint = "127.0.0.1:5778"
	cfg.GRPCServerSettings.NetAddr.Endpoint = "127.0.0.1:14250"
	return cfg
}
