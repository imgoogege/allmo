// Show all the packages in your project (including sub-packages' sub-packages)
package allmo

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

//show all modules.
func Mo(st string) []string {
	result := []string{}
	RestMap = make(map[string]int)
	syncNumber := make(chan struct{},50)
	tt(st,syncNumber)
	for { // 此处一直阻塞，直到 start和end在一秒的时间区域内都是相等的时候然后退出。
		// Blocked here until exit and end are equal in the time zone of one second and then exit.
		start := len(RestMap)
		time.Sleep(time.Second /10000000)
		end := len(RestMap)
		if start == end {
			break
		}
	}
	for k, _ := range RestMap {
		result = append(result, k)
	}
	fmt.Println(len(result))
	return result

}
func findGo() ([]byte, error) {
	return exec.Command("which", "go").Output()
}
func tt(s string,number chan struct{}) {
	number <- struct {}{}
	re := new(Result)
	dd, err := findGo()
	if err != nil {
		fmt.Println(err)
	}
	cmd := exec.Command(string(dd[:len(dd)-1]), "list", "-e", "-json", s)// 去除换行符
	data, err := cmd.Output()
	if err != nil {
		fmt.Print(err)
	}
	json.Unmarshal(data, re)
	<- number
	for _, v := range re.Imports {
		sy.Lock()
		RestMap[v]++
		sy.Unlock() // 如果在lock前调用unlock那么会发生error错误If you call unlock before lock, an error will occur.
		go tt(v,number)
	}

}
