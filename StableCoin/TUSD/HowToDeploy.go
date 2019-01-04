package main_test

import (
	"StableCoin/tusd"
	"UniqueDice/ethutils"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	KRinkebyUrl      = "https://rinkeby.infura.io"
	KOwnerPrivateKey = "09A4E515FD15398BC70030................"
)

func StringToPrivateKey(hex string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(hex)
}

var (
	contract = common.HexToAddress("0x88a2a2e63dacf7ee801b627d7fed65844598f0bb")
)

func Mint(contractaddress *common.Address) error {
	conn, e := ethclient.Dial(KRinkebyUrl)
	if e != nil {
		return e
	}
	defer conn.Close()
	token, ex := tusd.NewTrueUSD(*contractaddress, conn)
	if ex != nil {
		return ex
	}

	privkey, err := ethutils.StringToPrivateKey(KOwnerPrivateKey)
	if err != nil {
		return err
	}
	auth := bind.NewKeyedTransactor(privkey)
	auth.GasLimit = 1000000

	bi := big.NewInt(100)
	bi = bi.Mul(bi, big.NewInt(1000000000))
	bi = bi.Mul(bi, big.NewInt(10000000000))
	tx, ex2 := token.Mint(auth, common.HexToAddress("0xe14c47861b9f20a6BAb730f10e8BB5d4aB420cD4"), bi)
	if ex2 == nil {
		fmt.Println("ok, transfer done")
	}

	fmt.Println(tx, ex2)
	return nil
}

func SetLists() error {

	conn, e := ethclient.Dial(KRinkebyUrl)
	if e != nil {
		return e
	}
	defer conn.Close()

	privkey, err := ethutils.StringToPrivateKey(KOwnerPrivateKey)
	if err != nil {
		return err
	}
	auth := bind.NewKeyedTransactor(privkey)
	auth.GasLimit = 1000000
	token, ex := tusd.NewTrueUSD(contract, conn)
	if ex != nil {
		return ex
	}

	//addressList
	owner := common.HexToAddress("0xce64cac4b9e5793d2e29f4a9a3f9275e7522113b")

	_, et := token.SetLists(auth, owner, owner, common.Address{}, owner)
	if et != nil {
		return et
	}

	return nil
}

func SetBalanceSheet() error {

	conn, e := ethclient.Dial(KRinkebyUrl)
	if e != nil {
		return e
	}
	defer conn.Close()

	privkey, err := ethutils.StringToPrivateKey(KOwnerPrivateKey)
	if err != nil {
		return err
	}
	auth := bind.NewKeyedTransactor(privkey)
	auth.GasLimit = 1000000
	token, ex := tusd.NewTrueUSD(contract, conn)
	if ex != nil {
		return ex
	}

	//addressList
	balanceSheetAddr := common.HexToAddress("0x9e6e2b632dfb359262328d4303b324f7cc6ec219")
	_, et := token.SetBalanceSheet(auth, balanceSheetAddr)
	if et != nil {
		return et
	}
	return nil
}

func DeployContract() (*common.Address, error) {
	conn, e := ethclient.Dial(KRinkebyUrl)
	if e != nil {
		return nil, e
	}
	defer conn.Close()
	privkey, e2 := StringToPrivateKey(KOwnerPrivateKey)
	if e2 != nil {
		return nil, e2
	}
	auth := bind.NewKeyedTransactor(privkey)

	// Deploy a new awesome contract for the binding demo
	address, tx, token, err := tusd.DeployTrueUSD(auth, conn)
	if err != nil {
		return nil, err
	}
	fmt.Printf("tusd Contract pending deploy: 0x%x\n", address)
	fmt.Printf("tusd Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	fmt.Println("tusd Token: ", token)

	return &address, nil
}

func DeployAddressList() (*common.Address, error) {
	conn, e := ethclient.Dial(KRinkebyUrl)
	if e != nil {
		return nil, e
	}
	defer conn.Close()
	privkey, e2 := StringToPrivateKey(KOwnerPrivateKey)
	if e2 != nil {
		return nil, e2
	}
	auth := bind.NewKeyedTransactor(privkey)

	// Deploy a new awesome contract for the binding demo
	address, tx, token, err := tusd.DeployAddressList(auth, conn, "0xe14c47861b9f20a6BAb730f10e8BB5d4aB420cD4", true)
	if err != nil {
		return nil, err
	}

	fmt.Printf("addresslist Contract pending deploy: 0x%x\n", address)
	fmt.Printf("addresslist Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	fmt.Println("addresslist Token: ", token)

	return &address, nil
}

func DeployBalanceSheet() (*common.Address, error) {
	conn, e := ethclient.Dial(KRinkebyUrl)
	if e != nil {
		return nil, e
	}
	defer conn.Close()
	privkey, e2 := StringToPrivateKey(KOwnerPrivateKey)
	if e2 != nil {
		return nil, e2
	}
	auth := bind.NewKeyedTransactor(privkey)

	// Deploy a new awesome contract for the binding demo
	address, tx, token, err := tusd.DeployBalanceSheet(auth, conn)
	if err != nil {
		return nil, err
	}

	fmt.Printf("balanceSheet Contract pending deploy: 0x%x\n", address)
	fmt.Printf("balanceSheet Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	fmt.Println("balanceSheet Token: ", token)

	return &address, nil
}

func TransferBalanceSheet() error {
	conn, e := ethclient.Dial(KRinkebyUrl)
	if e != nil {
		return e
	}
	defer conn.Close()
	privkey, e2 := StringToPrivateKey(KOwnerPrivateKey)
	if e2 != nil {
		return e2
	}
	auth := bind.NewKeyedTransactor(privkey)

	balancesheetAddr := common.HexToAddress("0x9e6e2b632dfb359262328d4303b324f7cc6ec219")
	// Deploy a new awesome contract for the binding demo
	token, err := tusd.NewBalanceSheet(balancesheetAddr, conn)
	if err != nil {
		return err
	}
	_, e2 = token.TransferOwnership(auth, contract)
	if e2 != nil {
		return e2
	}

	return nil
}

func Test_deployAddressList(t *testing.T) {
	t.Error(DeployAddressList())
}

func Test_deployBalanceSheet(t *testing.T) {
	t.Error(DeployBalanceSheet())
}

func Test_transferBalanceSheet(t *testing.T) {
	t.Error(TransferBalanceSheet())
}

func Test_deploy(t *testing.T) {
	t.Error(DeployContract())
}

func Test_setBalanceSheet(t *testing.T) {
	t.Error(SetBalanceSheet())
}
func Test_setList(t *testing.T) {
	t.Error(SetLists())
}

func Test_mint(t *testing.T) {
	t.Error(Mint(&contract))
}

func Test_all(t *testing.T) {
	addr, ex := DeployContract()
	if ex != nil {
		t.Error(ex)
	}

	time.Sleep(10 * time.Second)
	fmt.Println(addr)
	t.Error(Mint(addr))
}
