package model

// BlockRef represents the relationship between block round and block hash.
type BlockRef struct {
	Round uint64 `pg:",pk"`
	Hash  string
}

// TransactionRef represents the relationship between ethereum tx and oasis tx.
type TransactionRef struct {
	EthTxHash string `pg:",pk"`
	Index     uint32 `pg:",use_zero"`
	Round     uint64 `pg:",use_zero"`
	BlockHash string
}

// AccessTuple is the element of an AccessList.
type AccessTuple struct {
	Address     string
	StorageKeys []string
}

// Continues is the latest Indexed Block Round.
const Continues string = "tip"

type ContinuesIndexedRound struct {
	Tip   string `pg:",pk"`
	Round uint64
}

type AccessList []AccessTuple

// Transaction is ethereum transaction.
type Transaction struct {
	Hash       string `pg:",pk"`
	Type       uint8  `pg:",use_zero"`
	ChainID    string
	Gas        uint64 `pg:",use_zero"`
	GasPrice   string
	GasTipCap  string
	GasFeeCap  string
	Nonce      uint64 `pg:",use_zero"`
	FromAddr   string
	ToAddr     string
	Value      string
	Data       string
	AccessList AccessList
	V, R, S    string
}

// Block represents ethereum block.
type Block struct {
	Round        uint64 `pg:",pk"`
	Header       *Header
	Uncles       []*Header
	Transactions []*Transaction
}

// Header represents ethereum block header.
type Header struct {
	ParentHash  string
	UncleHash   string
	Coinbase    string
	Root        string
	TxHash      string
	ReceiptHash string
	Bloom       string
	Difficulty  string
	Number      string
	GasLimit    uint64
	GasUsed     uint64
	Time        uint64
	Extra       string
	MixDigest   string
	Nonce       uint64
	BaseFee     string
}
