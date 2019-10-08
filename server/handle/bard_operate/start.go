package bard_operate

import (
	"bard-gui/server/operate"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

//@uri	/bard/start
func Start(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var rsp string = "[1]未知错误"
	//wg := new(sync.WaitGroup)
	//wg.Add(1)
	if runtime.GOOS == "windows" {
		go func() {
			Do(operate.Open)
			//wg.Done()
		}()
		time.Sleep(500*time.Millisecond)
		rsp = Do(operate.Status)
	}
	rsp = Do(operate.Open)

	fmt.Println(rsp)
	fmt.Fprint(w, rsp)

	// open结束后停止，实际上open在bard运行期间不会有返回 windows如此
	//wg.Wait()
}

