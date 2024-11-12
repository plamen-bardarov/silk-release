package config

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"time"

	"code.cloudfoundry.org/lager/v3"
	"github.com/containernetworking/cni/pkg/skel"
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/containernetworking/plugins/pkg/ns"
)

const ipv6gateway = "fe80::1"

//go:generate counterfeiter -o fakes/hardwareAddressGenerator.go --fake-name HardwareAddressGenerator . hardwareAddressGenerator
type hardwareAddressGenerator interface {
	GenerateForContainer(containerIP net.IP) (net.HardwareAddr, error)
	GenerateForHost(containerIP net.IP) (net.HardwareAddr, error)
}

//go:generate counterfeiter -o fakes/deviceNameGenerator.go --fake-name DeviceNameGenerator . deviceNameGenerator
type deviceNameGenerator interface {
	GenerateForHost(containerIP net.IP) (string, error)
	GenerateTemporaryForContainer(containerIP net.IP) (string, error)
}

//go:generate counterfeiter -o fakes/namespaceAdapter.go --fake-name NamespaceAdapter . namespaceAdapter
type namespaceAdapter interface {
	GetNS(string) (ns.NetNS, error)
	GetCurrentNS() (ns.NetNS, error)
}

type ConfigCreator struct {
	HardwareAddressGenerator hardwareAddressGenerator
	DeviceNameGenerator      deviceNameGenerator
	NamespaceAdapter         namespaceAdapter
	Logger                   lager.Logger
}

func (c *ConfigCreator) Create(
	hostNS netNS,
	addCmdArgs *skel.CmdArgs,
	ipamResult *current.Result,
	mtu int,
	enableIPv6 bool,
) (*Config, error) {
	var conf Config
	var err error

	c.Logger.Debug("start")
	defer c.Logger.Debug("done")

	if addCmdArgs.IfName == "" {
		return nil, errors.New("IfName cannot be empty")
	}
	if len(addCmdArgs.IfName) > 15 {
		return nil, errors.New("IfName cannot be longer than 15 characters")
	}

	conf.Container.DeviceName = addCmdArgs.IfName
	conf.Container.Namespace, err = c.NamespaceAdapter.GetNS(addCmdArgs.Netns)
	if err != nil {
		return nil, fmt.Errorf("getting container namespace: %s", err)
	}

	if len(ipamResult.IPs) == 0 {
		return nil, errors.New("no IP address in IPAM result")
	}

	conf.Container.Address.IP = ipamResult.IPs[0].Address.IP

	conf.Container.TemporaryDeviceName, err = c.DeviceNameGenerator.GenerateTemporaryForContainer(conf.Container.Address.IP)
	if err != nil {
		return nil, fmt.Errorf("generating temporary container device name: %s", err)
	}

	conf.Container.Address.Hardware, err = c.HardwareAddressGenerator.GenerateForContainer(conf.Container.Address.IP)
	if err != nil {
		return nil, fmt.Errorf("generating container veth hardware address: %s", err)
	}

	conf.Container.MTU = mtu
	conf.Host.DeviceName, err = c.DeviceNameGenerator.GenerateForHost(conf.Container.Address.IP)
	if err != nil {
		return nil, fmt.Errorf("generating host device name: %s", err)
	}

	conf.Host.Namespace = hostNS
	conf.Host.Address.IP = net.IP{169, 254, 0, 1}

	conf.Host.Address.Hardware, err = c.HardwareAddressGenerator.GenerateForHost(conf.Container.Address.IP)
	if err != nil {
		return nil, fmt.Errorf("generating host veth hardware address: %s", err)
	}

	conf.Container.Routes = []*types.Route{
		{
			Dst: net.IPNet{
				IP:   net.IPv4zero,
				Mask: net.CIDRMask(0, 32),
			},
			GW: []byte{169, 254, 0, 1},
		},
	}

	if enableIPv6 {
		conf.Host.AddressIPv6.IP = net.ParseIP(ipv6gateway)
		conf.Host.AddressIPv6.Hardware = conf.Host.Address.Hardware

		conf.Container.AddressIPv6.IP = net.ParseIP(generateIPv6())
		conf.Container.AddressIPv6.Hardware = conf.Container.Address.Hardware

		conf.Container.RoutesIPv6 = []*types.Route{{GW: conf.Host.AddressIPv6.IP}}
	}

	return &conf, nil
}

// TODO Remove me. IPv6 PoC
const (
	ipv6Prefix = "2600:1f18:27b3:881e:9148::"
)

func generateIPv6() string {
	rand.Seed(time.Now().UnixNano())
	suffix := fmt.Sprintf("%x:%x", rand.Intn(0xffff), rand.Intn(0xffff))

	newIPv6 := ipv6Prefix + suffix

	return newIPv6
}
