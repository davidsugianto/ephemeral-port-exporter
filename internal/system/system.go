package system

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetEphemeralPortRange() (int, int, error) {
	f, err := os.Open("/proc/sys/net/ipv4/ip_local_port_range")
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		return 0, 0, errors.New("failed to read port range line")
	}

	fields := strings.Fields(scanner.Text())
	if len(fields) != 2 {
		return 0, 0, errors.New("invalid port range format")
	}

	start, err1 := strconv.Atoi(fields[0])
	end, err2 := strconv.Atoi(fields[1])
	if err1 != nil || err2 != nil {
		return 0, 0, errors.New("failed to parse port range")
	}

	return start, end, nil
}

func CountUsedEphemeralPorts(start, end int) (int, error) {
	cmd := fmt.Sprintf("ss -Htan | awk '{print $4}' | awk -F: '{print $NF}' | awk -v s=%d -v e=%d '$1 >= s && $1 <= e' | sort -n | uniq | wc -l", start, end)

	usedOutput, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return 0, fmt.Errorf("failed to get used port count: %v", err)
	}
	used, _ := strconv.Atoi(strings.TrimSpace(string(usedOutput)))

	return used, nil
}
