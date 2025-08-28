package systemctl

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/1Panel-dev/1Panel/agent/global"
	"github.com/pkg/errors"
)

func RunSystemCtl(args ...string) (string, error) {
	cmd := exec.Command("systemctl", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("failed to run command: %w", err)
	}
	return string(output), nil
}

func isSnapService(serviceName string) bool {
	cmd := exec.Command("snap", "services")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return false
	}
	return strings.Contains(out.String(), serviceName)
}

func isSnapServiceActive(serviceName string) bool {
	cmd := exec.Command("snap", "services")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return false
	}
	lines := strings.Split(out.String(), "\n")
	for _, line := range lines {
		if strings.Contains(line, serviceName) && strings.Contains(line, "active") {
			return true
		}
	}
	return false
}

func IsActive(serviceName string) (bool, error) {
	// [gsx]
	_, err := isDockerRunningMac()
	if err != nil {
		return false, err
	}
	return true, nil

	// out, err := RunSystemCtl("is-active", serviceName)
	// if err == nil {
	// 	return strings.TrimSpace(out) == "active", nil
	// }

	// if isSnapServiceActive(serviceName) {
	// 	return true, nil
	// }
	// return false, fmt.Errorf("service %s is not active: %v", serviceName, err)
}

func IsEnable(serviceName string) (bool, error) {
	out, err := RunSystemCtl("is-enabled", serviceName)
	if err == nil {
		return strings.TrimSpace(out) == "enabled", nil
	}

	if isSnapServiceActive(serviceName) {
		return true, nil
	}
	return false, fmt.Errorf("failed to determine if service %s is enabled: %v", serviceName, err)
}

func IsExist(serviceName string) (bool, error) {
	out, err := RunSystemCtl("is-enabled", serviceName)
	if err == nil || strings.Contains(out, "disabled") {
		return true, nil
	}
	if isSnapService(serviceName) {
		return true, nil
	}
	return false, nil
}

func handlerErr(out string, err error) error {
	if err != nil {
		if out != "" {
			return errors.New(out)
		}
		return err
	}
	return nil
}

func Restart(serviceName string) error {
	out, err := RunSystemCtl("restart", serviceName)
	if err == nil {
		return nil
	}
	if isSnapService(serviceName) {
		cmd := exec.Command("snap", "restart", serviceName)
		output, snapErr := cmd.CombinedOutput()
		return handlerErr(string(output), snapErr)
	}
	return handlerErr(out, err)
}

func Operate(operate, serviceName string) error {
	out, err := RunSystemCtl(operate, serviceName)
	if err == nil {
		return nil
	}

	if isSnapService(serviceName) && (operate == "start" || operate == "stop" || operate == "restart") {
		cmd := exec.Command("snap", operate, serviceName)
		output, snapErr := cmd.CombinedOutput()
		return handlerErr(string(output), snapErr)
	}
	return handlerErr(out, err)
}

// 检查 macOS 系统上 Docker 是否运行
func isDockerRunningMac() (bool, error) {
	// 方法1: 检查 Docker 应用进程是否存在
	cmd := exec.Command("pgrep", "docker")
	if err := cmd.Run(); err != nil {
		global.LOG.Errorf("docker not found, err: %v", err)
		return false, err
	}

	// 方法2: 检查 Docker 守护进程连接
	cmd = exec.Command("docker", "info")
	if err := cmd.Run(); err != nil {
		global.LOG.Errorf("docker info error, err: %v", err)
		return false, err
	}

	return true, nil
}
