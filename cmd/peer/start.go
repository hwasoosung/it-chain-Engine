package peer

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/it-chain/it-chain-Engine/gateway"
	"github.com/syossan27/tebata"
	"github.com/urfave/cli"
)

func StartCmd() cli.Command {

	return cli.Command{
		Name:    "start",
		Aliases: []string{"s"},
		Usage:   "start peer as background",
		Flags: []cli.Flag{
			cli.BoolTFlag{
				Name:  "damon, d",
				Usage: "",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("peer is starting...")
			start(c.Bool("damon"))
			return nil
		},
	}
}

//start peer
func start(damon bool) {

	if damon {
		args := os.Args[1:]
		args = append(args, "-d=false")

		cmd := exec.Command(os.Args[0], args...)
		err := cmd.Start()

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("[PID]", cmd.Process.Pid)

		os.Exit(0)
	}

	t := tebata.New(syscall.SIGINT, syscall.SIGTERM)
	err := t.Reserve(stopGateway)

	pidValue, err := Create("my.pid")

	fmt.Println(pidValue)

	if err != nil {
		log.Fatalf("fail to write pid [%s]", err.Error())
		os.Exit(1)
	}

	err = gateway.Start()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

var (
	ErrPidExists = errors.New("pid file exists.")
	Debug        bool
)

func Create(pidfile string) (int, error) {

	if _, err := os.Stat(pidfile); !os.IsNotExist(err) {
		// file exists
		if pid, _ := GetValue(pidfile); pid > 0 {
			if ok, _ := IsActive(pid); ok {
				return pid, ErrPidExists
			}
		}
	}

	if pf, err := os.OpenFile(pidfile, os.O_RDWR|os.O_CREATE, 0600); err != nil {

		return 0, errors.New(fmt.Sprintf("Error when create pid file: %s\n", err.Error()))

	} else {
		pid := os.Getpid()
		pf.Write([]byte(strconv.Itoa(pid)))
		return pid, nil
	}
}

func IsActive(pid int) (bool, error) {
	if pid <= 0 {
		return false, errors.New("process id error.")
	}
	p, err := os.FindProcess(pid)
	if err != nil {
		if Debug {
			fmt.Printf("find process: %s\n", err.Error())
		}
		return false, err
	}

	if err := p.Signal(os.Signal(syscall.Signal(0))); err != nil {
		if Debug {
			fmt.Printf("send signal [0]: %s\n", err.Error())
		}
		return false, err
	}

	return true, nil
}

func GetValue(pidfile string) (int, error) {
	value, err := ioutil.ReadFile(pidfile)
	if err != nil {
		if Debug {
			fmt.Printf("read pid file: %s\n", err.Error())
		}
		return 0, err
	}
	pid, err := strconv.ParseInt(string(value), 10, 32)
	if err != nil {
		if Debug {
			fmt.Printf("trans pid to int: %s\n", err.Error())
		}
		return 0, err
	}
	return int(pid), nil
}

func stopGateway() {
	gateway.Stop()
	log.Println("stopped by signal")
}
