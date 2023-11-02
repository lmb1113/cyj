package basic

import (
	"bufio"
	"fmt"
	"net"
	"net/http"

	"github.com/onsi/ginkgo/v2"

	"cyj/pkg/frp/pkg/util/util"
	"cyj/pkg/frp/test/e2e/framework"
	"cyj/pkg/frp/test/e2e/framework/consts"
	"cyj/pkg/frp/test/e2e/mock/server/streamserver"
	"cyj/pkg/frp/test/e2e/pkg/request"
	"cyj/pkg/frp/test/e2e/pkg/rpc"
)

var _ = ginkgo.Describe("[Feature: TCPMUX httpconnect]", func() {
	f := framework.NewDefaultFramework()

	getDefaultServerConf := func(httpconnectPort int) string {
		conf := consts.DefaultServerConfig + `
		tcpmuxHTTPConnectPort = %d
		`
		return fmt.Sprintf(conf, httpconnectPort)
	}
	newServer := func(port int, respContent string) *streamserver.Server {
		return streamserver.New(
			streamserver.TCP,
			streamserver.WithBindPort(port),
			streamserver.WithRespContent([]byte(respContent)),
		)
	}

	proxyURLWithAuth := func(username, password string, port int) string {
		if username == "" {
			return fmt.Sprintf("http://127.0.0.1:%d", port)
		}
		return fmt.Sprintf("http://%s:%s@127.0.0.1:%d", username, password, port)
	}

	ginkgo.It("Route by HTTP user", func() {
		vhostPort := f.AllocPort()
		serverConf := getDefaultServerConf(vhostPort)

		fooPort := f.AllocPort()
		f.RunServer("", newServer(fooPort, "foo"))

		barPort := f.AllocPort()
		f.RunServer("", newServer(barPort, "bar"))

		otherPort := f.AllocPort()
		f.RunServer("", newServer(otherPort, "other"))

		clientConf := consts.DefaultClientConfig
		clientConf += fmt.Sprintf(`
			[[proxies]]
			name = "foo"
			type = "tcpmux"
			multiplexer = "httpconnect"
			localPort = %d
			customDomains = ["normal.example.com"]
			routeByHTTPUser = "user1"

			[[proxies]]
			name = "bar"
			type = "tcpmux"
			multiplexer = "httpconnect"
			localPort = %d
			customDomains = ["normal.example.com"]
			routeByHTTPUser = "user2"

			[[proxies]]
			name = "catchAll"
			type = "tcpmux"
			multiplexer = "httpconnect"
			localPort = %d
			customDomains = ["normal.example.com"]
			`, fooPort, barPort, otherPort)

		f.RunProcesses([]string{serverConf}, []string{clientConf})

		// user1
		framework.NewRequestExpect(f).Explain("user1").
			RequestModify(func(r *request.Request) {
				r.Addr("normal.example.com").Proxy(proxyURLWithAuth("user1", "", vhostPort))
			}).
			ExpectResp([]byte("foo")).
			Ensure()

		// user2
		framework.NewRequestExpect(f).Explain("user2").
			RequestModify(func(r *request.Request) {
				r.Addr("normal.example.com").Proxy(proxyURLWithAuth("user2", "", vhostPort))
			}).
			ExpectResp([]byte("bar")).
			Ensure()

		// other user
		framework.NewRequestExpect(f).Explain("other user").
			RequestModify(func(r *request.Request) {
				r.Addr("normal.example.com").Proxy(proxyURLWithAuth("user3", "", vhostPort))
			}).
			ExpectResp([]byte("other")).
			Ensure()
	})

	ginkgo.It("Proxy auth", func() {
		vhostPort := f.AllocPort()
		serverConf := getDefaultServerConf(vhostPort)

		fooPort := f.AllocPort()
		f.RunServer("", newServer(fooPort, "foo"))

		clientConf := consts.DefaultClientConfig
		clientConf += fmt.Sprintf(`
			[[proxies]]
			name = "test"
			type = "tcpmux"
			multiplexer = "httpconnect"
			localPort = %d
			customDomains = ["normal.example.com"]
			httpUser = "test"
			httpPassword = "test"
		`, fooPort)

		f.RunProcesses([]string{serverConf}, []string{clientConf})

		// not set auth header
		framework.NewRequestExpect(f).Explain("no auth").
			RequestModify(func(r *request.Request) {
				r.Addr("normal.example.com").Proxy(proxyURLWithAuth("", "", vhostPort))
			}).
			ExpectError(true).
			Ensure()

		// set incorrect auth header
		framework.NewRequestExpect(f).Explain("incorrect auth").
			RequestModify(func(r *request.Request) {
				r.Addr("normal.example.com").Proxy(proxyURLWithAuth("test", "invalid", vhostPort))
			}).
			ExpectError(true).
			Ensure()

		// set correct auth header
		framework.NewRequestExpect(f).Explain("correct auth").
			RequestModify(func(r *request.Request) {
				r.Addr("normal.example.com").Proxy(proxyURLWithAuth("test", "test", vhostPort))
			}).
			ExpectResp([]byte("foo")).
			Ensure()
	})

	ginkgo.It("TCPMux Passthrough", func() {
		vhostPort := f.AllocPort()
		serverConf := getDefaultServerConf(vhostPort)
		serverConf += `
		tcpmuxPassthrough = true
		`

		var (
			respErr            error
			connectRequestHost string
		)
		newServer := func(port int) *streamserver.Server {
			return streamserver.New(
				streamserver.TCP,
				streamserver.WithBindPort(port),
				streamserver.WithCustomHandler(func(conn net.Conn) {
					defer conn.Close()

					// read HTTP CONNECT request
					bufioReader := bufio.NewReader(conn)
					req, err := http.ReadRequest(bufioReader)
					if err != nil {
						respErr = err
						return
					}
					connectRequestHost = req.Host

					// return ok response
					res := util.OkResponse()
					if res.Body != nil {
						defer res.Body.Close()
					}
					_ = res.Write(conn)

					buf, err := rpc.ReadBytes(conn)
					if err != nil {
						respErr = err
						return
					}
					_, _ = rpc.WriteBytes(conn, buf)
				}),
			)
		}

		localPort := f.AllocPort()
		f.RunServer("", newServer(localPort))

		clientConf := consts.DefaultClientConfig
		clientConf += fmt.Sprintf(`
			[[proxies]]
			name = "test"
			type = "tcpmux"
			multiplexer = "httpconnect"
			localPort = %d
			customDomains = ["normal.example.com"]
			`, localPort)

		f.RunProcesses([]string{serverConf}, []string{clientConf})

		framework.NewRequestExpect(f).
			RequestModify(func(r *request.Request) {
				r.Addr("normal.example.com").Proxy(proxyURLWithAuth("", "", vhostPort)).Body([]byte("frp"))
			}).
			ExpectResp([]byte("frp")).
			Ensure()
		framework.ExpectNoError(respErr)
		framework.ExpectEqualValues(connectRequestHost, "normal.example.com")
	})
})
