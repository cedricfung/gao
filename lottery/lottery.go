package lottery

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/MixinNetwork/mixin/logger"
	"github.com/MixinNetwork/trusted-group/mvm/encoding"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/shopspring/decimal"
)

// abigen --abi abi.json --pkg lottery --type MixinProcess -out abi.go

const (
	XINAssetId         = "c94ac88f-4671-3976-b60a-09064f1811e8"
	MVMProcessId       = "THE MIXIN APP UUID"
	MVMContractAddress = "THE CONTRACT ADDRESS"
)

type Worker struct {
	client *mixin.Client
	proc   *MixinProcess
}

func NewWorker(ctx context.Context, conn *ethclient.Client) *Worker {
	proc, err := NewMixinProcess(common.HexToAddress(MVMContractAddress), conn)
	if err != nil {
		panic(err)
	}
	s := &mixin.Keystore{
		ClientID:   "ADD THIS",
		SessionID:  "ADD THIS",
		PrivateKey: "ADD THIS",
		PinToken:   "ADD THIS",
	}
	client, err := mixin.NewFromKeystore(s)
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixNano())
	rw := &Worker{
		client: client,
		proc:   proc,
	}
	return rw
}

func (rw *Worker) OnMessage(ctx context.Context, msg *mixin.MessageView, userId string) error {
	pTimestamp, err := rw.proc.TIMESTAMP(nil)
	if err != nil {
		panic(err)
	}
	timestamp := time.Unix(0, int64(pTimestamp))
	if pTimestamp == 0 {
		timestamp = time.Now()
	}
	duration := 24*time.Hour - time.Now().Sub(timestamp)
	if duration < 0 {
		duration = 0
	}
	pPool, err := rw.proc.Stats(nil)
	if err != nil {
		panic(err)
	}
	pool := decimal.NewFromBigInt(pPool, 0).Mul(decimal.NewFromFloat(0.01))
	guide := fmt.Sprintf("回复任意消息都可以再次收到本条状态消息。\n\n⚠️⚠️⚠️ \n注意该游戏处于完全早期测试目的，加注后无法退回。\n\n📗📗📗\n任何人都可以注入 XIN 到奖池，每次数量必须是 0.01XIN 的整数倍。每 24 小时开奖一次，按照参与倍数获得抽奖机会，抽中一人得到奖池内全部资金。\n\n💸💸💸\n总奖金池：%sXIN\n开奖时间：%s\n\n🤑🤑🤑\n回复任意整数获得相应倍数的付款链接。", pool, duration)

	data, _ := base64.RawStdEncoding.DecodeString(msg.Data)
	cmd, _ := decimal.NewFromString(strings.TrimSpace(string(data)))
	if msg.Category == mixin.MessageCategoryPlainText && cmd.IntPart() > 0 {
		op := &encoding.Operation{
			Purpose: encoding.OperationPurposeGroupEvent,
			Process: MVMProcessId,
			Extra:   []byte("GAO"),
		}
		pr := mixin.TransferInput{
			AssetID: XINAssetId,
			Amount:  decimal.NewFromFloat(0.01).Mul(decimal.NewFromInt(cmd.IntPart())),
			TraceID: msg.MessageID,
			Memo:    base64.RawURLEncoding.EncodeToString(op.Encode()),
		}
		pr.OpponentMultisig.Receivers = []string{
			"a15e0b6d-76ed-4443-b83f-ade9eca2681a",
			"b9126674-b07d-49b6-bf4f-48d965b2242b",
			"15141fe4-1cfd-40f8-9819-71e453054639",
			"3e72ca0c-1bab-49ad-aa0a-4d8471d375e7",
		}
		pr.OpponentMultisig.Threshold = 3
		payment, err := rw.client.VerifyPayment(ctx, pr)
		if err != nil {
			panic(err)
		}
		guide = fmt.Sprintf("注入 %sXIN\n\n🤑🤑🤑\nmixin://codes/%s", pr.Amount, payment.CodeID)
	}

	mr := &mixin.MessageRequest{
		ConversationID: msg.ConversationID,
		Category:       mixin.MessageCategoryPlainText,
		MessageID:      mixin.UniqueConversationID(msg.MessageID, msg.MessageID),
		Data:           base64.RawURLEncoding.EncodeToString([]byte(guide)),
	}
	return rw.client.SendMessage(ctx, mr)
}

func (rw *Worker) OnAckReceipt(ctx context.Context, msg *mixin.MessageView, userId string) error {
	return nil
}

func (rw *Worker) Loop(ctx context.Context) {
	for {
		err := rw.client.LoopBlaze(context.Background(), rw)
		logger.Printf("LoopBlaze() => %v\n", err)
		if ctx.Err() != nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
}
