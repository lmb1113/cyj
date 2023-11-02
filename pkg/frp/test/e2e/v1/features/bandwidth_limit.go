package features

import (
	"fmt"
	"strings"
	"time"

	"github.com/onsi/ginkgo/v2"

	plugin "qh-tool/pkg/frp/pkg/plugin/server"
	"qh-tool/pkg/frp/test/e2e/framework"
	"qh-tool/pkg/frp/test/e2e/framework/consts"
	"qh-tool/pkg/frp/test/e2e/mock/server/streamserver"
	pluginpkg "qh-tool/pkg/frp/test/e2e/pkg/plugin"
	"qh-tool/pkg/frp/test/e2e/pkg/request"
)

var _ = ginkgo.Describe("[Feature: Bandwidth Limit]", func() {
	f := framework.NewDefaultFramework()

	ginkgo.It("Proxy Bandwidth Limit by Client", func() {
		serverConf := consts.DefaultServerConfig
		clientConf := consts.DefaultClientConfig

		localPort := f.AllocPort()
		localServer := streamserver.New(streamserver.TCP, streamserver.WithBindPort(localPort))
		f.RunServer("", localServer)

		remotePort := f.AllocPort()
		clientConf += fmt.Sprintf(`
			[[proxies]]
			name = "tcp"
			type = "tcp"
			localPort = %d
			remotePort = %d
			transport.bandwidthLimit = "10KB"
			`, localPort, remotePort)

		f.RunProcesses([]string{serverConf}, []string{clientConf})

		content := strings.Repeat("a", 50*1024) // 5KB
		start := time.Now()
		framework.NewRequestExpect(f).Port(remotePort).RequestModify(func(r *request.Request) {
			r.Body([]byte(content)).Timeout(30 * time.Second)
		}).ExpectResp([]byte(content)).Ensure()

		duration := time.Since(start)
		framework.Logf("request duration: %s", duration.String())

		framework.ExpectTrue(duration.Seconds() > 8, "100Kb with 10KB limit, want > 8 seconds, but got %s", duration.String())
	})

	ginkgo.It("Proxy Bandwidth Limit by Server", func() {
		// new test plugin server
		newFunc := func() *plugin.Request {
			var r plugin.Request
			r.Content = &plugin.NewProxyContent{}
			return &r
		}
		pluginPort := f.AllocPort()
		handler := func(req *plugin.Request) *plugin.Response {
			var ret plugin.Response
			content := req.Content.(*plugin.NewProxyContent)
			content.BandwidthLimit = "10KB"
			content.BandwidthLimitMode = "server"
			ret.Content = content
			return &ret
		}
		pluginServer := pluginpkg.NewHTTPPluginServer(pluginPort, newFunc, handler, nil)

		f.RunServer("", pluginServer)

		serverConf := consts.DefaultServerConfig + fmt.Sprintf(`
		[[httpPlugins]]
		name = "test"
		addr = "127.0.0.1:%d"
		path = "/handler"
		ops = ["NewProxy"]
		`, pluginPort)
		clientConf := consts.DefaultClientConfig

		localPort := f.AllocPort()
		localServer := streamserver.New(streamserver.TCP, streamserver.WithBindPort(localPort))
		f.RunServer("", localServer)

		remotePort := f.AllocPort()
		clientConf += fmt.Sprintf(`
			[[proxies]]
			name = "tcp"
			type = "tcp"
			localPort = %d
			remotePort = %d
			`, localPort, remotePort)

		f.RunProcesses([]string{serverConf}, []string{clientConf})

		content := strings.Repeat("a", 50*1024) // 5KB
		start := time.Now()
		framework.NewRequestExpect(f).Port(remotePort).RequestModify(func(r *request.Request) {
			r.Body([]byte(content)).Timeout(30 * time.Second)
		}).ExpectResp([]byte(content)).Ensure()

		duration := time.Since(start)
		framework.Logf("request duration: %s", duration.String())

		framework.ExpectTrue(duration.Seconds() > 8, "100Kb with 10KB limit, want > 8 seconds, but got %s", duration.String())
	})
})
