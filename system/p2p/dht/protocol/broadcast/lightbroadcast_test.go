// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package broadcast

import (
	"sync/atomic"
	"testing"

	"github.com/33cn/chain33/client/mocks"
	"github.com/33cn/chain33/queue"
	"github.com/33cn/chain33/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	payload = []byte("testpayload")
	tx      = &types.Transaction{Execer: []byte("coins"), Payload: payload, Fee: 4600, Expire: 2}
	tx1     = &types.Transaction{Execer: []byte("coins"), Payload: payload, Fee: 460000000, Expire: 0}
	tx2     = &types.Transaction{Execer: []byte("coins"), Payload: payload, Fee: 100, Expire: 1}
	minerTx = &types.Transaction{Execer: []byte("coins"), Payload: payload, Fee: 14600, Expire: 200}
)

func startHandleMempool(q queue.Queue, txList *[]*types.Transaction) chan struct{} {
	client := q.Client()
	client.Sub("mempool")
	done := make(chan struct{})
	go func() {
		close(done)
		for msg := range client.Recv() {
			msg.Reply(client.NewMessage("p2p", types.EventTxListByHash, &types.ReplyTxList{Txs: *txList}))
		}
	}()
	return done
}

func TestLightBroadcast(t *testing.T) {

	q := queue.New("test")
	proto, cancel := newTestProtocolWithQueue(q)
	defer cancel()
	proto.cfg.MinLtBlockSize = 0
	proto.cfg.LtBlockPendTimeout = 1
	defer q.Close()
	memTxList := []*types.Transaction{nil, tx1, tx2}
	done := startHandleMempool(q, &memTxList)
	<-done
	pid := proto.Host.ID()
	block := &types.Block{TxHash: []byte("test"), Txs: []*types.Transaction{minerTx, tx, tx1, tx2}, Height: 10}
	peerMsgCh := proto.ps.Sub(psBroadcast)
	proto.ltB.addLtBlock(proto.buildLtBlock(block), pid, pid)
	// 缺少tx2, 组装失败, 等待1s超时
	msg := <-peerMsgCh
	require.Equal(t, int32(blockReqMsgID), msg.(publishMsg).msg.(*types.PeerPubSubMsg).MsgID)
	require.Equal(t, msg.(publishMsg).topic, proto.getPeerTopic(pid))

	//交易组
	txGroup, _ := types.CreateTxGroup([]*types.Transaction{tx1, tx2}, proto.ChainCfg.GetMinTxFeeRate())
	gtx := txGroup.Tx()
	memTxList = []*types.Transaction{tx, gtx, nil}
	blcCli := q.Client()
	blcCli.Sub("blockchain")
	proto.ltB.addLtBlock(proto.buildLtBlock(block), pid, pid)
	// 组装成功, 发送到blockchain模块, 模拟blockchain接收数据
	msg1 := <-blcCli.Recv()
	require.Equal(t, types.EventBroadcastAddBlock, int(msg1.Ty))
	blc, ok := msg1.Data.(*types.BlockPid)
	require.True(t, ok)
	require.Equal(t, pid.Pretty(), blc.Pid)
}

func TestBlockRequest(t *testing.T) {

	q := queue.New("test")
	proto, cancel := newTestProtocolWithQueue(q)
	defer cancel()
	api := &mocks.QueueProtocolAPI{}
	proto.API = api
	pid := proto.Host.ID()
	// blockchain获取失败, 返回
	var height int64 = 10
	atomic.StoreInt64(&proto.currHeight, height)
	api.On("GetBlocks", mock.Anything).Return(nil, types.ErrActionNotSupport).Once()
	proto.ltB.addBlockRequest(height, pid)
	require.Equal(t, 0, proto.ltB.blockRequestList.Len())
	testBlock := &types.Block{Txs: []*types.Transaction{minerTx}, Height: 10}
	details := &types.BlockDetails{
		Items: []*types.BlockDetail{
			{
				Block: testBlock,
			},
		},
	}
	// 本地高度未达到, 加入到等待队列
	atomic.StoreInt64(&proto.currHeight, height-1)
	api.On("GetBlocks", mock.Anything).Return(details, nil)
	proto.ltB.addBlockRequest(10, pid)
	require.Equal(t, 1, proto.ltB.blockRequestList.Len())

	// 正确获取流程
	peerMsgCh := proto.ps.Sub(psBroadcast)
	atomic.StoreInt64(&proto.currHeight, height)
	msg := <-peerMsgCh
	require.Equal(t, int32(blockRespMsgID), msg.(publishMsg).msg.(*types.PeerPubSubMsg).MsgID)
	require.Equal(t, msg.(publishMsg).topic, proto.getPeerTopic(pid))
}
