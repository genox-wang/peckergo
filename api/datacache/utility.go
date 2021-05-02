package datacache

import "time"

func ymd() string {
	loc, err := time.LoadLocation("Asia/Chongqing")
	now := time.Now()
	if err == nil {
		now = now.In(loc)
	}
	return now.Format("060102")
}

// 获取缓存的时候为了在0点获取前一天的 pv，uv 时间向前调整1小时
func ymdGet() string {
	loc, err := time.LoadLocation("Asia/Chongqing")
	now := time.Now()
	if err == nil {
		now = now.In(loc)
	}
	now = now.Add(-time.Hour)
	return now.Format("060102")
}
