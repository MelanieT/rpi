package pwm

// http://www.jumpnowtek.com/rpi/Using-the-Raspberry-Pi-Hardware-PWM-timers.html

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

const (
	SysFS = "/sys/class/pwm/pwmchip0"
)

type Driver interface {
	Export(ch int) error
	Unexport(ch int) error
	DutyCycle(ch, duty int) error
	Frequency(ch, freq int) error
	Enable(ch int) error
	Disable(ch int) error
}

func New() Driver {
	return &driver{
		sysfs:     SysFS,
		writeFile: ioutil.WriteFile,
	}
}

type driver struct {
	writeFile func(file string, data []byte, perm os.FileMode) error
	sysfs     string
}

func (d *driver) Export(ch int) error {
	file := filepath.Join(d.sysfs, "export")
	return d.writeFile(file, []byte(strconv.Itoa(ch)), 0600)
}

func (d *driver) Unexport(ch int) error {
	file := filepath.Join(d.sysfs, "uexport")
	return d.writeFile(file, []byte(strconv.Itoa(ch)), 0600)
}

func (d *driver) DutyCycle(ch, duty int) error {
	file := filepath.Join(d.sysfs, fmt.Sprintf("pwm%d", ch), "duty_cycle")
	return d.writeFile(file, []byte(strconv.Itoa(ch)), 0600)
}

func (d *driver) Frequency(ch, freq int) error {
	file := filepath.Join(d.sysfs, fmt.Sprintf("pwm%d", ch), "period")
	return d.writeFile(file, []byte(strconv.Itoa(ch)), 0600)
}

func (d *driver) Enable(ch int) error {
	file := filepath.Join(d.sysfs, fmt.Sprintf("pwm%d", ch), "enable")
	return d.writeFile(file, []byte("1"), 0600)
}
func (d *driver) Disable(ch int) error {
	file := filepath.Join(d.sysfs, fmt.Sprintf("pwm%d", ch), "enable")
	return d.writeFile(file, []byte("0"), 0600)
}
