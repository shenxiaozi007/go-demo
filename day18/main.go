package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}

	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("B: ", i)
			wg.Done()
		}()
	}
	wg.Wait()
}


10 2 * * * www php /data1/wwwroot/system/system/updateInfo/updateSystem.php
0 2 * * 0 www php /data1/wwwroot/system/system/updateInfo/deleteNotice.php
*/5 8-20 * * * www bash /data1/wwwroot/system/system/StockInfo.sh
*/1 8-23 * * * www bash /data1/wwwroot/system/system/AutoKf.sh
*/2 * * * * www bash /data1/wwwroot/system/system/VideoRecord.sh
50 1,11,17 * * * www bash /data1/wwwroot/system/system/BakSql.sh
03 8,11,13,16,20 * * * www bash /data1/wwwroot/system/system/SyncOrder.sh
*/1 * * * * www bash /data1/wwwroot/system/system/StockNotice.sh
*/5 8-20 * * * www bash /data1/wwwroot/system/system/ProductStockInfo.sh
0 10,12,14,16,2 * * * www bash /data1/wwwroot/system/system/TotalMoneyAdviser.sh
