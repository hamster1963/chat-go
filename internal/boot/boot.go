package boot

import (
	"context"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/glog"
)

func Boot() (err error) {
	_, err = gcron.AddOnce(context.TODO(), "@every 2s", func(ctx context.Context) {
		glog.Debug(context.Background(), "定时任务启动中...")
		if err := Method(); err != nil {
			glog.Fatal(context.Background(), "定时任务启动失败: ", err)
		}
		glog.Debug(context.Background(), "定时任务启动成功")
	}, "开始启动定时任务")
	if err != nil {
		return err
	}
	return nil
}

func Method() (err error) {
	return err
}
