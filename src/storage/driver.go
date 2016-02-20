package storage

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/otsimo/api/apipb"
	"gopkg.in/mgo.v2/bson"
)

type Driver interface {
	Name() string
	GetById(bson.ObjectId) (*apipb.Content, error)
	GetBySlug(string) (*apipb.Content, error)
	List(apipb.ContentListRequest) ([]*apipb.Content, error)
	Put(*apipb.Content) error
	Update(*apipb.Content) error
	ChangeStatus(string, apipb.Content_Status) error
}

type RegisteredDriver struct {
	New   func(*cli.Context) (Driver, error)
	Flags []cli.Flag
}

var drivers map[string]*RegisteredDriver

func init() {
	drivers = make(map[string]*RegisteredDriver)
}

func Register(name string, rd *RegisteredDriver) error {
	if _, ext := drivers[name]; ext {
		return fmt.Errorf("Name already registered %s", name)
	}
	drivers[name] = rd
	return nil
}

func GetDriverNames() []string {
	drives := make([]string, 0)

	for name, _ := range drivers {
		drives = append(drives, name)
	}
	return drives
}

func GetDriver(name string) *RegisteredDriver {
	return drivers[name]
}
