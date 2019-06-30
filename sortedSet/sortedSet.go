package sortedSet

import (
	"fmt"
	"sortedset"
	"time"
)

func main() {
	var ss = sortedset.New()
	var n = 1000
	//f,_ := os.Create("mm.prof")
	//pprof.StartCPUProfile(f)
	//
	//
	//
	//pprof.StopCPUProfile()

	t := time.Now()
	//num := t.Unix()
	for i:=0;i<n;i++{
		ss.AddOrUpdate(i,int64(i),nil)
	}

	d := ss.GetByRankRange(1,20,false)
	var ids = make([]int, len(d))
	for idx,node := range d{
		ids[idx] = node.Key()
	}


	fmt.Println(time.Since(t))
	fmt.Println(ids)
}

func testEcho()  {
	
}