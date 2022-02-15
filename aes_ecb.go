package fun

import "crypto/aes"

// =================== ECB(电码本模式) ======================
func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipherBlock, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipherBlock.BlockSize(); bs <= len(origData); bs, be = bs+cipherBlock.BlockSize(), be+cipherBlock.BlockSize() {
		cipherBlock.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipherBlock, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipherBlock.BlockSize(); bs < len(encrypted); bs, be = bs+cipherBlock.BlockSize(), be+cipherBlock.BlockSize() {
		cipherBlock.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	bEnd := SearchByteSliceIndex(decrypted, 0)

	return decrypted[:bEnd]
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// []byte 字节切片 循环查找
func SearchByteSliceIndex(bSrc []byte, b byte) int {
	for i := 0; i < len(bSrc); i++ {
		if bSrc[i] == b {
			return i
		}
	}

	return -1
}