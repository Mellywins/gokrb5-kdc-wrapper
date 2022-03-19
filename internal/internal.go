package internal

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
)

func __CHECK_HOST__() error {
	if runtime.GOOS != "linux" {
		return fmt.Errorf("os platform is not supported! only linux hosts are allowed")
	}
	return nil
}
func EnsureServiceIsRunning() (string, bool) {
	err := __CHECK_HOST__()
	if err != nil {
		panic(err)
	}
	output, err := exec.Command("systemctl", "status", "krb5-kdc"+".service").Output()
	if err == nil {
		if matched, err := regexp.MatchString("Active: active", string(output)); err == nil && matched {
			reg := regexp.MustCompile("Main PID: ([0-9]+)")
			data := reg.FindStringSubmatch(string(output))
			if len(data) > 1 {
				return "Service (pid  " + data[1] + ") is running...", true
			}
			return "Service is running...", true
		}
	}

	return "Service is stopped", false
}
