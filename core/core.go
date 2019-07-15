package core

import (
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
}

type Container struct {
	Name	  string
	Replicas  int
	Address	  string
	Ports	  []operations.Ports
}

func NewCore(ctx context.Context) (c Core, err error) {
	c.eventsDocker = make(map[string]chan docker.Events)
	c.watchEvents(ctx)

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

func (c *Core) CreateDB(cont Container) (cs []Container, err error) {
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

					fmt.Println(event)
					wg.Done()
				}
			}
		}
	}()
	wg.Wait()

	return
}
