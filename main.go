package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"
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

// abigen --abi abi.json --pkg main --type MixinProcess --out abi.go

var (
	XINAssetId         = "c94ac88f-4671-3976-b60a-09064f1811e8"
	MVMProcessId       = "THE MIXIN APP UUID"
	MVMContractAddress = "THE CONTRACT ADDRESS"
	GethRPC            = "http://104.197.245.214:8545"
)

func main() {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(GethRPC)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	proc, err := NewMixinProcess(common.HexToAddress(MVMContractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	ctx := context.Background()
	worker := NewMessengerWorker(ctx, proc)
	worker.loop(ctx)
}

// Messenger is a simple MTG worker demo, it also sends some cmd to MM
type MessengerWorker struct {
	client *mixin.Client
	proc   *MixinProcess
}

func NewMessengerWorker(ctx context.Context, proc *MixinProcess) *MessengerWorker {
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
	rw := &MessengerWorker{
		client: client,
		proc:   proc,
	}
	return rw
}

func (rw *MessengerWorker) OnMessage(ctx context.Context, msg *mixin.MessageView, userId string) error {
	xin := "0x" + strings.Replace(XINAssetId, "-", "", -1)
	asset, _ := new(big.Int).SetString(xin, 0)
	pAmount, err := rw.proc.AMOUNT(nil)
	if err != nil {
		panic(err)
	}
	amount := decimal.NewFromBigInt(pAmount, -8)
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
	pPool, err := rw.proc.Custodian(nil, asset)
	if err != nil {
		panic(err)
	}
	pool := decimal.NewFromBigInt(pPool, -8)
	guide := fmt.Sprintf("回复任意消息都可以再次收到本条状态消息。\n\n⚠️⚠️⚠️ \n注意该游戏处于完全早期测试目的，加注后无法退回。\n\n📗📗📗\n任何人都可以注入 XIN 到奖池，如果 A 合法注入 XIN 后，24 小时内无人注入，那 A 将获得奖池所有 XIN。参与者每次注入的 XIN 都要比之前的多，才算合法注入。\n\n💸💸💸\n总奖金池：%sXIN\n上次注入：%sXIN\n开奖时间：%s\n\n🤑🤑🤑\n回复 GAO 可以获得一次合法注入的付款链接，大小写不重要。", pool, amount, duration)

	data, _ := base64.RawStdEncoding.DecodeString(msg.Data)
	cmd := strings.TrimSpace(strings.ToUpper(string(data)))
	if msg.Category == mixin.MessageCategoryPlainText && (cmd == "GAO" || cmd == "搞") {
		if amount.Sign() == 0 {
			amount, _ = decimal.NewFromString("0.0001")
		} else {
			amount = amount.Mul(decimal.NewFromFloat(1.05))
		}

		op := &encoding.Operation{
			Purpose: encoding.OperationPurposeGroupEvent,
			Process: MVMProcessId,
			Extra:   []byte("GAO"),
		}
		pr := mixin.TransferInput{
			AssetID: XINAssetId,
			Amount:  amount,
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
		guide = fmt.Sprintf("注入 %sXIN\n\n🤑🤑🤑\nmixin://codes/%s", amount, payment.CodeID)
	}

	mr := &mixin.MessageRequest{
		ConversationID: msg.ConversationID,
		Category:       mixin.MessageCategoryPlainText,
		MessageID:      mixin.UniqueConversationID(msg.MessageID, msg.MessageID),
		Data:           base64.RawURLEncoding.EncodeToString([]byte(guide)),
	}
	return rw.client.SendMessage(ctx, mr)
}

func (rw *MessengerWorker) OnAckReceipt(ctx context.Context, msg *mixin.MessageView, userId string) error {
	return nil
}

func (rw *MessengerWorker) loop(ctx context.Context) {
	for {
		err := rw.client.LoopBlaze(context.Background(), rw)
		logger.Printf("LoopBlaze() => %v\n", err)
		if ctx.Err() != nil {
			break
		}
		time.Sleep(3 * time.Second)
	}
}
