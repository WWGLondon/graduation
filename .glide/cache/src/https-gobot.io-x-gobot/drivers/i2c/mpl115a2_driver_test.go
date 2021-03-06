package i2c

import (
	"strings"
	"testing"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/gobottest"
)

var _ gobot.Driver = (*MPL115A2Driver)(nil)

// --------- HELPERS
func initTestMPL115A2Driver() (driver *MPL115A2Driver) {
	driver, _ = initTestMPL115A2DriverWithStubbedAdaptor()
	return
}

func initTestMPL115A2DriverWithStubbedAdaptor() (*MPL115A2Driver, *i2cTestAdaptor) {
	adaptor := newI2cTestAdaptor()
	return NewMPL115A2Driver(adaptor), adaptor
}

// --------- TESTS

func TestNewMPL115A2Driver(t *testing.T) {
	// Does it return a pointer to an instance of MPL115A2Driver?
	var mpl interface{} = NewMPL115A2Driver(newI2cTestAdaptor())
	_, ok := mpl.(*MPL115A2Driver)
	if !ok {
		t.Errorf("NewMPL115A2Driver() should have returned a *MPL115A2Driver")
	}
}

// Methods
func TestMPL115A2Driver(t *testing.T) {
	mpl := initTestMPL115A2Driver()

	gobottest.Refute(t, mpl.Connection(), nil)
	gobottest.Assert(t, strings.HasPrefix(mpl.Name(), "MPL115A2"), true)
}

func TestMPL115A2DriverStart(t *testing.T) {
	mpl, adaptor := initTestMPL115A2DriverWithStubbedAdaptor()

	adaptor.i2cReadImpl = func(b []byte) (int, error) {
		copy(b, []byte{0x00, 0x01, 0x02, 0x04})
		return 4, nil
	}
	gobottest.Assert(t, mpl.Start(), nil)
	time.Sleep(100 * time.Millisecond)
	gobottest.Assert(t, mpl.Pressure, float32(50.007942))
	gobottest.Assert(t, mpl.Temperature, float32(116.58878))
}

func TestMPL115A2DriverHalt(t *testing.T) {
	mpl := initTestMPL115A2Driver()

	gobottest.Assert(t, mpl.Halt(), nil)
}
