package interest

import (
	"bytes"
	"encoding/gob"

	"github.com/ethereum/go-ethereum/common"
)

// 固定利率
// 半年期 1.5%,一年期2%,三年期2.5%,五年期3%,十年期3.5%,三十年期4%
const (
	InterestRateHalfYear   = 15
	InterestRateOneYear    = 20
	InterestRateThreeYear  = 25
	InterestRateFiveYear   = 30
	InterestRateTenYear    = 35
	InterestRateThirtyYear = 40
	// 利息差
	InterestRateDiff = 5
	// 利率除数
	InterestRateDivisor = 1000
)

func GetInterestRate(year int) int {
	switch year {
	case 1:
		return InterestRateHalfYear
	case 2:
		return InterestRateOneYear
	case 6:
		return InterestRateThreeYear
	case 10:
		return InterestRateFiveYear
	case 20:
		return InterestRateTenYear
	case 60:
		return InterestRateThirtyYear
	default:
		return 0
	}
}

type PledgeInfo struct {
	PledgeAmount uint64 //质押金额
	PledgeYear   int    //质押年限
	StartTime    uint64 //开始时间(区块高度)
	InterestRate int    //利率

	CurrentInterest uint64 //当前利息
	EarnInterest    uint64 //收益利息（利息差）

	AnnualFee         uint64 //合约部署年费
	LastAnnualFeeTime uint64 //上一次收取年费的时间

	ContractAddress    common.Address //合约地址
	DeployedAddress    common.Address //部署人地址
	InvestorAddress    common.Address //投资人地址
	BeneficiaryAddress common.Address //受益人地址

	StakeFlag bool //质押标志
}

// new pledge info
func NewPledgeInfo(pledgeAmount uint64, pledgeYear int, startTime uint64, annualFee uint64, ca common.Address, da common.Address, ia common.Address, ba common.Address) *PledgeInfo {
	interestRate := GetInterestRate(pledgeYear)
	return &PledgeInfo{
		PledgeAmount:       pledgeAmount,
		PledgeYear:         pledgeYear,
		InterestRate:       interestRate,
		StartTime:          startTime,
		CurrentInterest:    0,
		EarnInterest:       0,
		AnnualFee:          annualFee,
		LastAnnualFeeTime:  startTime,
		ContractAddress:    ca,
		DeployedAddress:    da,
		InvestorAddress:    ia,
		BeneficiaryAddress: ba,
		StakeFlag:          true,
	}
}

// EncodePledgeInfo 将 PledgeInfo 编码为字节数组
func EncodePledgeInfo(p *PledgeInfo) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecodePledgeInfo 将字节数组解码为 PledgeInfo
func DecodePledgeInfo(data []byte) (*PledgeInfo, error) {
	var p PledgeInfo
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func CalInterest(p *PledgeInfo) uint64 {
	// 计算利息
	// 利息 = 质押金额 * 利率 * 质押年限 / 利率除数
	interest := p.PledgeAmount * uint64(p.InterestRate) * uint64(p.PledgeYear) / InterestRateDivisor
	// 利息差 = 质押金额 * 利息差 * 质押年限 / 利率除数
	interestDiff := p.PledgeAmount * uint64(InterestRateDiff) * uint64(p.PledgeYear) / InterestRateDivisor

	// 扣除年费
	// 首先计算年费
	af := p.AnnualFee * uint64(p.PledgeYear) / 2
	// 判断是否大于年费
	bInterest := interest + interestDiff
	if bInterest > af {
		bInterest -= af
	} else {
		bInterest = 0
	}

	return bInterest
}
