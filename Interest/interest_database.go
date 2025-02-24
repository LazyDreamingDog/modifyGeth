package interest

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
)

const (
	pledgePrefix = "pledge-" // 质押信息前缀
)

type PledgeDB struct {
	db ethdb.Database
	mu sync.RWMutex
}

func NewPledgeDB(db ethdb.Database) *PledgeDB {
	return &PledgeDB{
		db: db,
	}
}

// 生成存储键
func generateKey(addr common.Address) []byte {
	return append([]byte(pledgePrefix), addr.Bytes()...)
}

// SavePledgeInfo 保存质押信息
func (pdb *PledgeDB) SavePledgeInfo(info *PledgeInfo) error {
	pdb.mu.Lock()
	defer pdb.mu.Unlock()

	data, err := EncodePledgeInfo(info)
	if err != nil {
		return err
	}

	key := generateKey(info.ContractAddress)
	return pdb.db.Put(key, data)
}

// GetPledgeInfo 获取质押信息
func (pdb *PledgeDB) GetPledgeInfo(contractAddr common.Address) (*PledgeInfo, error) {
	pdb.mu.RLock()
	defer pdb.mu.RUnlock()

	key := generateKey(contractAddr)
	data, err := pdb.db.Get(key)
	if err != nil {
		return nil, err
	}

	return DecodePledgeInfo(data)
}

// DeletePledgeInfo 删除质押信息
func (pdb *PledgeDB) DeletePledgeInfo(contractAddr common.Address) error {
	pdb.mu.Lock()
	defer pdb.mu.Unlock()

	key := generateKey(contractAddr)
	return pdb.db.Delete(key)
}

// UpdatePledgeInfo 更新质押信息
func (pdb *PledgeDB) UpdatePledgeInfo(info *PledgeInfo) error {
	return pdb.SavePledgeInfo(info) // 直接覆盖现有数据
}

// ListPledgeInfos 列出所有质押信息
func (pdb *PledgeDB) ListPledgeInfos() ([]*PledgeInfo, error) {
	pdb.mu.RLock()
	defer pdb.mu.RUnlock()

	var pledges []*PledgeInfo
	iter := pdb.db.NewIterator([]byte(pledgePrefix), nil)
	defer iter.Release()

	for iter.Next() {
		pledge, err := DecodePledgeInfo(iter.Value())
		if err != nil {
			return nil, err
		}
		pledges = append(pledges, pledge)
	}

	return pledges, iter.Error()
}
