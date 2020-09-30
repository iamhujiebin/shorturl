package logic

import (
	"context"
	"errors"
	"github.com/tal-tech/go-zero/core/mr"
	"shorturl/rpc/transform/transformer"
	"strconv"

	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type Expand2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpand2Logic(ctx context.Context, svcCtx *svc.ServiceContext) Expand2Logic {
	return Expand2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Expand2Logic) Expand2(req types.ExpandReq) (*types.ExpandResp, error) {
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
	resp, err := l.svcCtx.Transformer.Expand2(l.ctx, &transformer.Expand2Req{Shorten: req.Shorten})
	if err != nil {
		return nil, err
	}
	return &types.ExpandResp{
		Url: resp.Url,
	}, nil
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
