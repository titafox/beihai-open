package utils

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/rand"
	"strings"
	"time"
)

// 生成随机数字字符串
func GenerateRandomString(length int) string {
	const charset = "1234567890"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	randomStrBytes := make([]byte, length)
	for i := range randomStrBytes {
		randomStrBytes[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(randomStrBytes)
}

// 登录验证
func CheckSignature(sign, address, nonce string) error {
	decodedSig, err := hexutil.Decode(sign)
	if err != nil {
		return err
	}

	if decodedSig[64] != 27 && decodedSig[64] != 28 {
		return errors.New("签名数据错误")
	}

	decodedSig[64] -= 27
	str := "Welcome to Coltstail!\n" +
		"\n" +
		"Click \"Sign\" to sign in. No password needed!\n" +
		"This request will not trigger a blockchain transaction or cost any gas fees.\n" +
		"\n" +
		"Your authentication status will be reset after 24 hours.\n" +
		"\n" +
		"Wallet address:\n" + address +
		"\n" +
		"\n" +
		"Nonce:\n" + nonce +
		"\n"
	nonceStr := []byte(str)
	prefixedNonce := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(nonceStr), nonceStr)
	hash := crypto.Keccak256Hash([]byte(prefixedNonce))
	recoveredPublicKey, err := crypto.Ecrecover(hash.Bytes(), decodedSig)
	if err != nil {
		return err
	}

	secp256k1RecoveredPublicKey, err := crypto.UnmarshalPubkey(recoveredPublicKey)
	if err != nil {
		return err
	}
	recoveredAddress := crypto.PubkeyToAddress(*secp256k1RecoveredPublicKey).Hex()
	isClientAddressEqualToRecoveredAddress := strings.ToLower(address) == strings.ToLower(recoveredAddress)
	if !isClientAddressEqualToRecoveredAddress {
		return errors.New("签名验证失败")
	}

	return nil
}

// 创建/出售 签名验证
func CheckBuyOrSellSignature(sign, address, nonce string) error {
	decodedSig, err := hexutil.Decode(sign)
	if err != nil {
		return err
	}

	if decodedSig[64] != 27 && decodedSig[64] != 28 {
		return errors.New("签名数据错误")
	}

	decodedSig[64] -= 27
	str := "Nonce:" + nonce
	nonceStr := []byte(str)
	prefixedNonce := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(nonceStr), nonceStr)
	hash := crypto.Keccak256Hash([]byte(prefixedNonce))
	recoveredPublicKey, err := crypto.Ecrecover(hash.Bytes(), decodedSig)
	if err != nil {
		return err
	}

	secp256k1RecoveredPublicKey, err := crypto.UnmarshalPubkey(recoveredPublicKey)
	if err != nil {
		return err
	}
	recoveredAddress := crypto.PubkeyToAddress(*secp256k1RecoveredPublicKey).Hex()
	fmt.Println("xxxxx:", recoveredAddress)
	isClientAddressEqualToRecoveredAddress := strings.ToLower(address) == strings.ToLower(recoveredAddress)
	if !isClientAddressEqualToRecoveredAddress {
		return errors.New("签名验证失败")
	}

	return nil
}
