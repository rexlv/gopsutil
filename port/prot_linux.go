package port

import (
	//"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func PortInfo(num string) (s *PortInfoStat, err error) {
	args := `netstat -a | grep :` + num + ` | awk '$1 == "tcp" && $NF == "LISTEN" {print $0}'`
	fmt.Println(args)
	//args := "ping 127.0.0.1"
	out, err := exec.Command("/bin/sh", "-c", args).Output()
	if err != nil {
		return
	}

	vals := strings.Fields(string(out))

	pis := &PortInfoStat{
		Id:    vals[3],
		Proto: vals[0],
		State: vals[5],
	}

	s.RecvQ, err = strconv.ParseUint(vals[1], 10, 0)
	if err != nil {
		return
	}

	s.SendQ, err = strconv.ParseUint(vals[2], 10, 0)
	if err != nil {
		return
	}

	s = pis
	return
}
