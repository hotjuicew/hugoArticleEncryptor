console.log('js导入成功')
function AESDecrypt(ciphertext, password) {
    // 将密文分割成nonce和密文部分
    const parts = ciphertext.split("|");
    if (parts.length !== 2) {
        throw new Error("invalid ciphertext format");
    }

    // 解码nonce和密文
    const decodedNonce = CryptoJS.enc.Hex.parse(parts[0]);
    const decodedCiphertext = CryptoJS.enc.Hex.parse(parts[1]);

    // 创建AES的BlockCipher
    const block = CryptoJS.AES.createCipher(password);

    // 创建AES-GCM
    const aesgcm = CryptoJS.mode.GCM.createDecipher(block, decodedNonce);

    // 解密
    const decrypted = CryptoJS.enc.Utf8.stringify(
        aesgcm.decrypt({ ciphertext: decodedCiphertext })
    );

    return decrypted;
}
// 示例用法
const ciphertext = "0123456789abcdef|0123456789abcdef";
const password = CryptoJS.enc.Hex.parse("0123456789abcdef0123456789abcdef");
const plaintext = AESDecrypt(ciphertext, password);
console.log(plaintext);