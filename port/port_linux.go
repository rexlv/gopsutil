// +build linux

package port

import (
	"os/exec"
	"strconv"
	"strings"
)

func PortInfo(num []string) (s []*PortInfoStat, err error) {
	args := `netstat -anp | grep -E '`
	for _, n := range num {
		args += ":" + n + "|"
	}
	args = args[:len(args)-1] + "'"
	out, err := exec.Command("/bin/sh", "-c", args).Output()
	if err != nil {
		return
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		vals := strings.Fields(line)

		pis := PortInfoStat{}
		if len(vals) > 5 {
			pis.Id = strings.Split(vals[3], ":")[1]
			pis.Proto = vals[0]
			pis.State = vals[5]
			pis.RecvQ, err = strconv.ParseUint(vals[1], 10, 0)
			if err != nil {
				pis.RecvQ = 0
			}

			pis.SendQ, err = strconv.ParseUint(vals[2], 10, 0)
			if err != nil {
				pis.SendQ = 0
			}
		}

		s = append(s, &pis)
	}

	for _, n := range num {
		ok := false
		for _, pi := range s {
			if n == pi.Id {
				ok = true
			}
		}
		if !ok {
			s = append(s, &PortInfoStat{
				Id:    n,
				Proto: "Unknown",
				State: "Unknown",
			})
		}
	}

	return
}
