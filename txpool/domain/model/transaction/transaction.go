package transaction

import (
	"time"

	"encoding/json"

	"github.com/it-chain/it-chain-Engine/common"
	"github.com/rs/xid"
)

const (
	VALID   TransactionStatus = 0
	INVALID TransactionStatus = 1

	General TransactionType = 0 + iota
)

type TransactionId string

func (tId TransactionId) ToString() string {
	return string(tId)
}

type TransactionStatus int
type TransactionType int
type Transaction struct {
	TxId          TransactionId
	PublishPeerId string
	TxStatus      TransactionStatus
	TxType        TxDataType
	TxHash        string
	TimeStamp     time.Time
	TxData        *TxData
}

func NewTransaction(publishPeerId string, txType TxDataType, txData *TxData) *Transaction {
	tx := Transaction{
		TxId:          TransactionId(xid.New().String()),
		PublishPeerId: publishPeerId,
		TxStatus:      VALID,
		TxType:        txType,
		TxHash:        "",
		TimeStamp:     time.Now(),
		TxData:        txData,
	}
	hashArgs := []string{txData.Jsonrpc, string(txData.Method), string(txData.Params.Function), txData.ICodeID, publishPeerId, tx.TimeStamp.String(), string(tx.TxId), string(tx.TxType)}
	for _, str := range txData.Params.Args {
		hashArgs = append(hashArgs, str)
	}
	tx.TxHash = common.ComputeSHA256(hashArgs)
	return &tx
}

func (t Transaction) Serialize() ([]byte, error) {
	return common.Serialize(t)
}

func (t Transaction) GetID() string {
	return string(t.TxId)
}

func Deserialize(b []byte, transaction *Transaction) error {
	err := json.Unmarshal(b, transaction)

	if err != nil {
		return err
	}

	return nil
}
