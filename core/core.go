package core

import (
	"os"
	"fmt"
	"sync"
	"context"
	"encoding/json"
	"github.com/lucasmbaia/tdbas/docker"
	"github.com/lucasmbaia/tdbas/operations"
)

type Core struct {
	sync.RWMutex

	eventsDocker  map[string]chan docker.Events
	path	      string
}

type CoreConfig struct {
	PathVolume  string
}

type Container struct {
	Organization  string
	Name	      string
	Replicas      int
	Address	      string
	Ports	      []operations.Ports
}

func NewCore(ctx context.Context, cfg CoreConfig) (c Core, err error) {
	c.eventsDocker = make(map[string]chan docker.Events)
	c.watchEvents(ctx)
	c.path = cfg.PathVolume

	return
}

func (c *Core) watchEvents(ctx context.Context) {
	var (
		err	error
		errc	= make(chan error, 1)
		event	= make(chan []byte)
		ok	bool
	)

	go func() {
		errc <- docker.DockerEvents(ctx, event)
	}()

	go func() {
		for {
			select {
			case msg := <-event:
				var ev docker.Events

				if err = json.Unmarshal(msg, &ev); err != nil {
					break
				}

				c.Lock()
				if _, ok = c.eventsDocker[ev.Actor.Attributes.Name]; ok {
					c.eventsDocker[ev.Actor.Attributes.Name] <- ev
				}
				c.Unlock()
			case _ = <-errc:
				c.watchEvents(ctx)
			}
		}
	}()
}

func (c *Core) CreateDB(cont Container) (cs Container, err error) {
	var (
		ev    = make(chan docker.Events, 1)
		path  string
		wg    sync.WaitGroup
	)

	path = fmt.Sprintf("%s%s/%s", c.path, cont.Organization, cont.Name)
	if _, err = os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			return
		}
	}

	c.Lock()
	c.eventsDocker[cont.Name] = ev
	c.Unlock()

	go func() {
		operations.CreateContainer(cont.Name, path)
	}()

	wg.Add(1)
	go func() {
		for {
			select {
			case event := <-ev:
				switch event.Action{
				case "start":
					cs.Name = event.Actor.Attributes.Name
					cs.Address, err = operations.AddressContainer(cs.Name)
					cs.Ports, err = operations.PortsContainer(cs.Name)

					wg.Done()
				}
			}
		}
	}()
	wg.Wait()

	return
}

/*func (c *Core) CreateDB(cont Container) (cs []Container, err error) {
	var (
		ev    = make(chan docker.Events, 1)
		wg    sync.WaitGroup
	)

	c.Lock()
	for i := 1; i <= cont.Replicas; i++ {
		c.eventsDocker[fmt.Sprintf("%s-%d", cont.Name, i)] = ev
	}
	c.Unlock()

	wg.Add(cont.Replicas)
	go func() {
		for i := 1; i <= cont.Replicas; i++ {
			operations.CreateContainer(fmt.Sprintf("%s-%d", cont.Name, i))
		}
	}()

	go func() {
		for {
			select {
			case event := <-ev:
				switch event.Action{
				case "start":
					var ct = Container{Name: event.Actor.Attributes.Name}

					ct.Address, err = operations.AddressContainer(ct.Name)
					ct.Ports, err = operations.PortsContainer(ct.Name)
					cs = append(cs, ct)

					wg.Done()
				}
			}
		}
	}()
	wg.Wait()

	return
}*/
