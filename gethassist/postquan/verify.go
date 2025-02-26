package postquan

import (
	"fmt"
	"sync"

	"teddycode/pqcgo"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
)

// @state.StateDB implement this interface
type postQuanBackend interface {
	GetLastPostQuanPub(common.Address) []byte
	SetLastPostQuanPub(common.Address, []byte)
}

// When the executor is active, continuously listening post-quantum transaction events.
// If executor close, this function will return
func BindPostQuantumEvents(sub *event.TypeMuxSubscription, exitCh <-chan struct{}) {
	for {
		select {
		case postQuanEvent := <-sub.Chan():
			fmt.Println("Catch txs from miner")
			// Convert typemux event(wrapped) to event struct
			event := postQuanEvent.Data.(core.PostQuanEvent)

			// Set up a goroutine for events and control concurrency through wg
			var verfiedTxs types.Transactions
			var wg sync.WaitGroup
			for _, tx := range event.Txs {
				wg.Add(1)
				go handleEvent(tx, event.Signer, &verfiedTxs, event.State, &wg)
			}
			wg.Wait()

			// Must return even if it is empty, otherwise wg will block the main process
			event.VerifiedTxsCh <- verfiedTxs
		case <-exitCh:
			fmt.Println("Shutting down event listener...")
			return
		}
	}
}

// Verify post-quantum tx, attention @verfiedTxs is address passing instead of value passing
func handleEvent(tx *types.Transaction, signer types.Signer, verfiedTxs *types.Transactions, backend postQuanBackend, wg *sync.WaitGroup) {
	defer wg.Done()
	sheme := pqcgo.PQCSignType[string(tx.CryptoType())]
	// Type legitimacy check
	if tx.Type() != types.DynamicCryptoTxType {
		return
	}

	message, err := tx.JsonExcludePostQumSign()
	if err != nil {
		fmt.Printf("get message err: %v\n", err)
	}

	// Get tx sender's address
	sender, err := signer.Sender(tx)
	if err != nil {
		log.Error(fmt.Sprintf("Parse sender from tx(%s) error", tx.Hash()))
		return
	}

	// Verify1: the validity of public key.
	// If the address is firstly send post-quantum transaction (len(prePk)=0),the check is skipped
	prePk := backend.GetLastPostQuanPub(sender)
	if len(prePk) != 0 {
		pqcgo.VerifyPubKeyIndex(sheme, prePk, tx.PublicKey())
	}

	// Verify2: the validity and correctness of signature
	flag, err := pqcgo.Verify(sheme, tx.SignatureData(), message, tx.PublicKey())
	if err != nil {
		log.Error(fmt.Sprintf("Verify post-quantum tx(%s) error:%v", tx.Hash(), err))
	} else if !flag {
		log.Error(fmt.Sprintf("Verify post-quantum tx(%s) failed", tx.Hash()))
	} else {
		// Successfully verify
		*verfiedTxs = append(*verfiedTxs, tx)
		// Update public key
		backend.SetLastPostQuanPub(sender, tx.PublicKey())

		fmt.Println("Successful verify post-quantum signature")
	}

}
