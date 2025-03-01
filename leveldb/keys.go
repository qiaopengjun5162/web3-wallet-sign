package leveldb

import "github.com/ethereum/go-ethereum/log"

// Keys 包含一个指向LevelStore的指针，用于管理密钥。
type Keys struct {
	db *LevelStore
}

func NewKeyStore(path string) (*Keys, error) {
	db, err := NewLevelStore(path)
	if err != nil {
		log.Error("Could not create leveldb database.")
		return nil, err
	}
	return &Keys{
		db: db,
	}, nil
}

func (k *Keys) GetPrivateKey(publicKey string) (string, bool) {
	key := []byte(publicKey)
	data, err := k.db.Get(key)
	if err != nil {
		return "0x00", false
	}
	privateKey := toString(data)
	return privateKey, true
}

// StoreKeys 存储密钥列表到数据库中。
// 参数:
//
//	keyList - 一个Key对象的切片，每个Key对象包含一个公钥和一个私钥。
//
// 返回值:
//
//	成功存储所有密钥返回true，否则返回false。
func (k *Keys) StoreKeys(keyList []Key) bool {
	// 遍历密钥列表，将每个密钥对存储到数据库中。
	for _, item := range keyList {
		// 将公钥转换为字节切片，用作数据库中的键。
		key := []byte(item.Pubkey)
		// 将私钥转换为字节切片，用作数据库中的值。
		value := toBytes(item.PrivateKey)
		// 将键值对存储到数据库中。
		err := k.db.Put(key, value)
		// 如果存储过程中发生错误，记录错误日志并返回false。
		if err != nil {
			log.Error("store key value fail", "err", err, "key", key, "value", value)
			return false
		}
	}
	// 所有密钥对成功存储，返回true。
	return true
}
