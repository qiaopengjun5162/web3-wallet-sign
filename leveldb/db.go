// Package leveldb provides a leveldb store.
package leveldb

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/log"
	"github.com/syndtr/goleveldb/leveldb"
)

// LevelStore 是一个封装了LevelDB数据库的结构体。
// 它提供了一种方式，使得在Go代码中可以更容易地与LevelDB数据库进行交互。
// 通过继承leveldb.DB，LevelStore可以直接使用leveldb.DB的所有方法和属性，
// 同时也可以在其上添加更多的功能或进行定制化，以满足特定的需求。
type LevelStore struct {
	*leveldb.DB
}

// NewLevelStore 创建并初始化一个新的 LevelStore 实例。
// 该函数接受一个路径字符串作为参数，用于指定 LevelDB 数据库文件的路径。
// 成功时，返回一个指向 LevelStore 实例的指针；失败时，返回 nil 和一个错误对象。
func NewLevelStore(path string) (*LevelStore, error) {
	// 尝试打开指定路径的 LevelDB 数据库文件。
	// 如果数据库文件不存在，将会创建一个新的数据库文件。
	handle, err := leveldb.OpenFile(path, nil)
	if err != nil {
		// 如果打开数据库文件失败，记录错误日志并返回错误。
		log.Error("open level db file fail", "err", err)
		return nil, err
	}
	// 如果成功打开数据库文件，创建并返回一个 LevelStore 实例。
	return &LevelStore{handle}, nil
}

// 向LevelStore中插入键值对
func (db *LevelStore) Put(key []byte, value []byte) error {
	// 调用LevelDB的Put方法，将键值对插入到数据库中
	return db.DB.Put(key, value, nil)
}

// Get 从 LevelStore 数据库中获取与指定键关联的值。
// 参数:
//
//	key []byte: 要检索的值的唯一键。
//
// 返回值:
//
//	[]byte: 与键关联的值，如果未找到则返回nil。
//	error: 如果获取值时发生错误，则返回错误信息。
func (db *LevelStore) Get(key []byte) ([]byte, error) {
	// 调用底层DB的Get方法来检索值，传入key和nil作为选项。
	// 这里使用nil作为选项是因为LevelDB的Get方法需要两个参数，
	// 第二个参数是可选的读取选项，此处不需要自定义选项。
	return db.DB.Get(key, nil)
}

func (db *LevelStore) Delete(key []byte) error {
	return db.DB.Delete(key, nil)
}

func toBytes(dataStr string) []byte {
	dataBytes, _ := hex.DecodeString(dataStr)
	return dataBytes
}

func toString(byteArr []byte) string {
	return hex.EncodeToString(byteArr)
}
