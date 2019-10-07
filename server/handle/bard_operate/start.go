package bard_operate

import (
	"bard-gui/server/operate"
	"fmt"
	"net/http"
	"time"
)

//@uri	/bard/start
func Start(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var rsp string = "[1]未知错误"
	//wg := new(sync.WaitGroup)
	//wg.Add(1)
	go func() {
		Do(operate.Open)
		fmt.Println("after open")
		//wg.Done()
	}()
	time.Sleep(500*time.Millisecond)
	rsp = Do(operate.Status)

	fmt.Println(rsp)
	fmt.Fprint(w, rsp)


	// open结束后停止，实际上open在bard运行期间不会有返回 windows如此
	//wg.Wait()
}

