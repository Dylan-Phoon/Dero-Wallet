package main

import (
	"fmt"
	"sync"
)
import "github.com/deroproject/derosuite/walletapi"

var command_line string = `dero-wallet-gui
DERO Wallet gui: A secure, private blockchain with smart-contracts
Usage:
  dero-wallet-gui [--help] [--version] [--debug] [--testnet]  [--noopengl] [--vmmode]
  dero-wallet-gui -h | --help
  dero-wallet-gui --version
Options:
  -h --help     Show this screen.
  --version     Show version.
  --debug       Debug mode enabled, print log messages
  --testnet     Enable testnet mode
  --noopengl    Enable minimal UI using software rendering
  --vmmode      Enable minimal UI using software rendering`

type CtxObject struct {

	walletptr *walletapi.Wallet

	_ string `property:"version"`

	_ func() `constructor:"init"`

	remote_server string `property:"remote_server"` // remote server
	_             bool   `property:"wallet_online"` // remote server

	_ func(string) `signal:"setwalletonline,auto"`  // used to set wallet online
	_ func()       `signal:"setwalletoffline,auto"` // used to set wallet offline

	_ bool   `property:"wallet_valid"`   // is wallet valid and has been successfully opened
	_ string `property:"wallet_address"` // wallet address
	_ string `property:"someString"`

	_ string `property:"initerr"` // used to track error when initially opening or creating database

	_ func(string)         `signal:"checkpassword,auto"` // used to check password
	_ func(string, string) `signal:"setpassword,auto"`   // used to set new password

	// property related to to outgoing tx
	_ string `property:"tx_hex"`   // wallet address
	_ string `property:"txid_hex"` // wallet address
	_ string `property:"tx_total"`
	_ string `property:"tx_transfer_amount"`
	_ string `property:"tx_change"`
	_ string `property:"tx_fees"`
	_ string `property:"tx_relayed"`

	_ func(string, string, string) `signal:"build_tx,auto"` // used to build up tx
	_ func(string)                 `signal:"relay_tx,auto"` // used to relay tx

	_ string       `property:"seed"`             // seed in localised language
	_ func(string) `signal:"seed_language,auto"` // used to request seed in language

	_ func(string, string)         `signal:"openwallet,auto"`            // used to openwallet
	_ func(string, string)         `signal:"createnewwallet,auto"`       // used to openwallet
	_ func(string, string, string) `signal:"recoverusingseedwords,auto"` // used to recover using seed words
	_ func(string, string, string) `signal:"recoverusingkey,auto"`       // used to recoverkey
	_ func()                       `signal:"closewallet,auto"`           // used to closewallet

	height       int64  `property:"height"`         // wallet height
	topoheight   int64  `property:"topoheight"`     // wallet topoheight
	nwheight     int64  `property:"nwheight"`       // network height
	nwtopoheight int64  `property:"nwtopoheight"`   // network topoheight
	_            string `property:"height_str"`     // height localised string
	_            string `property:"topoheight_str"` // topoheight localised string

	_ string `property:"total_balance"`    // wallet total balance
	_ string `property:"unlocked_balance"` // wallet unlocked balance
	_ string `property:"locked_balance"`   // network locked balance

	_ func()       `signal:"clicked,auto"`
	_ func(string) `signal:"sendString,auto"`

	_ func(string) `signal:"addressVerify,auto"`  // used to verify address
	_ bool         `property:"addressverified"`   // is address verfied
	_ bool         `property:"addressintegrated"` // is address integrated
	_ string       `property:"addressipaymentid"` // integrated payment ID in  hex form

	_ func(string) `signal:"paymentidVerify,auto"` // used to verify payment id
	_ bool         `property:"paymentidverified"`  // is payment id  verfied

	_ func(string) `signal:"amountVerify,auto"` // used to verify amount
	_ bool         `property:"amountverified"`  // is amount  verified

	_ func() `signal:"genintegratedaddress,auto"`         // used to verify amount
	_ string `property:"integrated_32_address"`           // integrated i32 address
	_ string `property:"integrated_32_address_paymentid"` // integrated i32 address
	_ string `property:"integrated_8_address"`            // integrated i8 address
	_ string `property:"integrated_8_address_paymentid"`  // integrated i32 address

	_ func(bool, bool, bool, int64) `signal:"reloadhistory,auto"` // used to reload history, available,in,out

	_ []string `property:"historyListHeight"`
	_ []string `property:"historyListTopoHeight"`
	_ []string `property:"historyListTXID"`
	_ []string `property:"historyListAmount"`
	_ []string `property:"historyListPaymentID"`
	_ []string `property:"historyListStatus"`
	_ []string `property:"historyListUnlockTime"`
	_ []string `property:"historyListOutDetails"` // contains json string

	sync.Mutex
}

var count int

func (t *CtxObject) init() {
	global_object = t // capture reference to original object

	/*
	    var err error
	   global_object.walletptr ,err  = walletapi.Open_Encrypted_Wallet("/tmp/tmp2.db", "")
	   if err != nil {
	           fmt.Printf("Wallet opened successfully")
	   }
	   addr := global_object.walletptr.GetAddress()
	   global_object.SetWallet_address(addr.String())
	   global_object.SetWallet_valid(true)  // mark wallet as valid
	*/

}

func (t *CtxObject) clicked() {
	count++
	fmt.Printf("clicked qml button\n")
}

func (t *CtxObject) sendString(a string) {
	fmt.Println("sendString:", a)
}

var global_object *CtxObject
//var global_gui *gui.QGuiApplication

func main() {

}