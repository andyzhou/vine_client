package define

import "time"

type (
	Node string
	ServerAddress string
	RpcNode string
	RpcPort int
	Path string
	ShortUrl string
	FileMd5 string
	FileName string
	FileType string
	Owner int64
	OwnerStr string
	Offset int64
	Length int64
	ChunkIndex int
	ChunkHandle int64
	ChunkVersion int64
	Checksum int64
	MutationType int
	ErrorCode int
)

const (
	RpcClientCheckTick = 3 * time.Second
)