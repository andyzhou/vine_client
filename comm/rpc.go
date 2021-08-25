package comm

import "github.com/andyzhou/vine_client/define"

type (
	//for error
	Error struct {
		Code define.ErrorCode
		Err  string
	}

	//for file info
	OneFileInfo struct {
		ShortUrl define.ShortUrl
		Name define.FileName
		Type define.FileType
		Length define.Length
		CreateAt int64
	}

	//////////////////
	//for master server
	//////////////////

	//for heart beat
	HeartbeatArg struct {
		Node define.Node //server node
		Address	define.ServerAddress //server address
	}
	HeartbeatReply struct {
		Time int64
	}

	//for write new file
	WriteFileArg struct {
		File define.FileName
		Type define.FileType
		Data []byte
	}
	WriteFileReply struct {
		ShortUrl define.ShortUrl
		ErrorCode define.ErrorCode
	}

	//for read old file
	ReadFileArg struct {
		ShortUrl define.ShortUrl
		Offset define.Offset //option
		Length define.Length //option
	}
	ReadFileReply struct {
		Type define.FileType
		Data []byte
		ErrorCode define.ErrorCode
	}

	//for delete file
	DelFileArg struct {
		ShortUrl define.ShortUrl
		Token string
	}
	DelFileReply struct {
		ErrorCode define.ErrorCode
	}

	//for list file
	ListFileArg struct {
		Page int
		PageSize int
	}
	ListFileReply struct {
		RecSize int
		Files []*OneFileInfo
	}

	//////////////////
	//for chunk server
	//////////////////

	//for write new chunk
	WriteChunkArg struct {
		Handle define.ChunkHandle
		Offset define.Offset
		Length define.Length
		Data []byte
	}
	WriteChunkReply struct {
		Handle define.ChunkHandle
		Offset define.Offset
		Length define.Length
		ErrorCode define.ErrorCode
	}

	//for read old chunk
	ReadChunkArg struct {
		File define.FileName
		Handle define.ChunkHandle
		Offset define.Offset //optional
		Length define.Length //optional
	}
	ReadChunkReply struct {
		Data      []byte
		Length    define.Length
		ErrorCode define.ErrorCode
	}

	//for add master node
	AddMasterArg struct {
		Node define.ServerAddress
	}
	AddMasterReply struct {
		ErrorCode define.ErrorCode
	}
)