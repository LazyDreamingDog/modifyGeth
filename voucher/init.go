package voucher

// Used to solve some problems caused by abi coding
import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

var (
	// Method declaration
	BalanceOf     *BoundMethod
	Buy           *BoundMethod
	Use           *BoundMethod
	CreateVoucher *BoundMethod
)

func init() {
	gas := uint64(100000000)

	// Load contract ABI
	MutiVoucherABI, err := abi.JSON(strings.NewReader(common.MutivoucherABI_json))
	if err != nil {
		log.Error("Load codestorage ABI err")
	}

	BalanceOf = NewBoundMethod(&common.MutiVoucherAddress, &MutiVoucherABI, "balanceOf", true, gas)
	Use = NewBoundMethod(&common.MutiVoucherAddress, &MutiVoucherABI, "use", false, gas)
	Buy = NewBoundMethod(&common.MutiVoucherAddress, &MutiVoucherABI, "buy", false, gas)
	CreateVoucher = NewBoundMethod(&common.MutiVoucherAddress, &MutiVoucherABI, "createVoucher", false, gas)
}
