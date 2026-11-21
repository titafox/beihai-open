package main

import (
	"os"
	dog "coltstail-server/services/dog_contract"
	"coltstail-server/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func main() {
	conn, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545/")
	if err != nil {
		fmt.Println("链接失败:", err)
		return
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("ETH_PRIVATE_KEY"))

	auth, errAuth := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(97))
	if errAuth != nil {
		fmt.Println("Failed to create new transactor : %v", errAuth)
		return
	}
	//fmt.Println(auth)

	//加载合约
	contract := "0x08e0739F94759eD467569e5B1519a266158b2b79"
	//fmt.Println(contract)
	tokenAddress := common.HexToAddress(contract)
	i, e := dog.NewDog(tokenAddress, conn)
	if e != nil {
		fmt.Println("error:", e)
		return
	}
	//fmt.Println(tokenAddress)
	//
	//bol,err := i.IsApprovedForAll(nil,common.HexToAddress("0xAab9645C46991Ef576910E44fdB8CA430039615e"),tokenAddress)
	//fmt.Println(bol,err)

	//num,err := i.BalanceOf(nil,common.HexToAddress("0xAab9645C46991Ef576910E44fdB8CA430039615e"))
	//fmt.Println(num,err)
	//i.SetApprovalForAll()
	_to := common.HexToAddress("0x15d7CED0678694054c4de9288afE4D00A6683796")
	_tokenId := big.NewInt(1484051613961166848)
	_tokenURI := "ipfs://dog"
	_amount := utils.ToWei(250, 18)
	_royaltiesReceipientAddress := common.HexToAddress("0xAab9645C46991Ef576910E44fdB8CA430039615e")
	data, err := i.Mint(auth, _to, _tokenId, _tokenURI, _amount, _royaltiesReceipientAddress)
	fmt.Println(data.Hash(), err)

	//data,err := i.SafeTransferFrom(auth,_to,_royaltiesReceipientAddress,_tokenId)
	//fmt.Println(data.Hash(),err)

}
