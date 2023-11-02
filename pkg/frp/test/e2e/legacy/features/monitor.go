package features

import (
	"fmt"
	"strings"
	"time"

	"github.com/onsi/ginkgo/v2"

	"cyj/pkg/frp/pkg/util/log"
	"cyj/pkg/frp/test/e2e/framework"
	"cyj/pkg/frp/test/e2e/framework/consts"
	"cyj/pkg/frp/test/e2e/pkg/request"
)

var _ = ginkgo.Describe("[Feature: Monitor]", func() {
	f := framework.NewDefaultFramework()

	ginkgo.It("Prometheus metrics", func() {
		dashboardPort := f.AllocPort()
		serverConf := consts.LegacyDefaultServerConfig + fmt.Sprintf(`
		enable_prometheus = true
		dashboard_addr = 0.0.0.0
		dashboard_port = %d
		`, dashboardPort)

		clientConf := consts.LegacyDefaultClientConfig
		remotePort := f.AllocPort()
		clientConf += fmt.Sprintf(`
		[tcp]
		type = tcp
		local_port = {{ .%s }}
		remote_port = %d
		`, framework.TCPEchoServerPort, remotePort)

		f.RunProcesses([]string{serverConf}, []string{clientConf})

		framework.NewRequestExpect(f).Port(remotePort).Ensure()
		time.Sleep(500 * time.Millisecond)

		framework.NewRequestExpect(f).RequestModify(func(r *request.Request) {
			r.HTTP().Port(dashboardPort).HTTPPath("/metrics")
		}).Ensure(func(resp *request.Response) bool {
			log.Trace("prometheus metrics response: \n%s", resp.Content)
			if resp.Code != 200 {
				return false
			}
			if !strings.Contains(string(resp.Content), "traffic_in") {
				return false
			}
			return true
		})
	})
})
