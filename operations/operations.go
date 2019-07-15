package operations

import (
	"strings"
	"github.com/lucasmbaia/tdbas/utils"
)

const (
	TIMEOUT_DEFAULT_COMMAND	= 60
)

type Ports struct {
	Source	      string
	Destinations  []string
}

func CreateContainer(name string) (id string, err error) {
	var (
		result	[]string
		args	[]string
	)

	args = []string{"run", "-e", "ACCEPT_EULA=Y", "-e", "SA_PASSWORD=totvs@123", "-P", "--expose=1433", "--name", name, "-d", "mcr.microsoft.com/mssql/server:2017-latest"}

	if result, err = utils.Cmd("docker", args, 60); err != nil {
		return
	}

	id = result[0]
	return
}

func AddressContainer(name string) (addr string, err error) {
	var (
		result	[]string
	)

	if result, err = utils.Cmd("docker", []string{"inspect", "-f", "{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}", name}, TIMEOUT_DEFAULT_COMMAND); err != nil {
		return
	}

	addr = result[0]
	return
}

func PortsContainer(name string) (p []Ports, err error) {
	var (
		result	[]string
		ports	[]string
		pc	= make(map[string][]string)
	)

	if result, err = utils.Cmd("docker", []string{"inspect", "--format={{range $p, $conf := .NetworkSettings.Ports}}{{$p}}:{{(index $conf 0).HostPort}}-{{end}}", name}, TIMEOUT_DEFAULT_COMMAND); err != nil {
		return
	}

	ports = strings.Split(result[0], "-")
	ports = ports[:len(ports)-1]

	for _, port := range ports {
		var p = strings.Split(port, "/tcp:")

		if _, ok := pc[p[0]]; ok {
			pc[p[0]] = append(pc[p[0]], p[1])
		} else {
			pc[p[0]] = []string{p[1]}
		}
	}

	for src, dst := range pc {
		p = append(p, Ports{Source: src, Destinations: dst})
	}

	return
}
