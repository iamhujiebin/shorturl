package logic

import (
	"context"
	"errors"
	"github.com/tal-tech/go-zero/core/mr"
	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"
	"strconv"

	"github.com/tal-tech/go-zero/core/logx"
)

type ExpandHandler2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandHandler2Logic(ctx context.Context, svcCtx *svc.ServiceContext) ExpandHandler2Logic {
	return ExpandHandler2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandHandler2Logic) ExpandHandler2(req types.ExpandReq) (*types.ExpandResp, error) {
	err := mr.Finish(func() error {
		return nil
	}, func() error {
		return nil
	}, func() error {
		return errors.New("parallel err")
	})
	if err != nil {
		logx.Errorf("something wrong:%s", err.Error())
	} else {
		logx.Infof("something right:%s", "hello")
	}
	rsp, err := mr.MapReduce(func(source chan<- interface{}) {
		for i := 0; i <= 10; i++ { //数据源
			source <- i
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) { //map函数，多协程并发执行,所以不可能一个cancel掉后，其他都不会write的
		err := check(item.(int))
		if err != nil {
			logx.Errorf("check fail:%v", item)
			cancel(err)
		} else {
			writer.Write(item)
		}
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) { //reduce函数，单协程执行
		var uids []int
		for p := range pipe {
			uids = append(uids, p.(int))
		}
		logx.Infof("cancel 后的uids:%+v", uids)
		uidStr := rpcCheck(uids)
		if len(uidStr) <= 1 {
			cancel(errors.New("uid str no long"))
			logx.Infof("已经cancel了,cancel就是代表有err了")
		}
		logx.Infof("cancel了还会来到，但是不会执行write()")
		writer.Write(uidStr)
	})
	if err != nil {
		logx.Errorf("something wrong:%s", err.Error())
	} else {
		logx.Infof("something right:%+v", rsp)
	}

	return &types.ExpandResp{}, nil
}

func check(i int) error {
	if i == 4 {
		return errors.New("check err")
	}
	return nil
}

func rpcCheck(uids []int) (uidStr []string) {
	for _, v := range uids {
		if v == 0 {
			continue
		}
		uidStr = append(uidStr, strconv.Itoa(v))
	}
	return
}
