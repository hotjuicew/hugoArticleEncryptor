
async function aesDecrypt(ciphertext, password,nonse) {
    const ciphertextBuffer = hexToBytes(ciphertext)
    // 计算密码的 SHA-256作为密钥
    // 将密码转换为 Uint8Array
    var encoder = new TextEncoder();
    var data = encoder.encode(password);

    // 计算 SHA-256 哈希
    var hashBuffer = await crypto.subtle.digest('SHA-256', data);

    // 将哈希结果转换为 Uint8Array
    var hash = Array.from(new Uint8Array(hashBuffer));

    const hashKey = new Uint8Array(hash);
    console.log("hashKey",hashKey)
    const key = await window.crypto.subtle.importKey(
        'raw',
        hashKey, {
            name: 'AES-GCM',
        },
        false,
        ['decrypt']
    )
    console.log("key",key)
    console.log('nonse',nonse)
    let iv=hexToBytes(nonse)
    console.log("iv",iv)
    const decrypted = await window.crypto.subtle.decrypt({
            name: 'AES-GCM',
            iv: iv,
            tagLength: 128,
        },
        key,
        new Uint8Array(ciphertextBuffer)
    )
    console.log("new Uint8Array(ciphertextBuffer)",new Uint8Array(ciphertextBuffer))
    console.log("decrypted",decrypted)
    return new TextDecoder('utf-8').decode(new Uint8Array(decrypted))
}
function hexToBytes(hexString) {
    // 去除可能存在的前缀 "0x" 或 "0X"
    if (hexString.startsWith("0x") || hexString.startsWith("0X")) {
        hexString = hexString.slice(2);
    }

    // 将十六进制字符串转换为 Uint8Array
    const bytes = new Uint8Array(hexString.length / 2);
    for (let i = 0; i < hexString.length; i += 2) {
        bytes[i / 2] = parseInt(hexString.substr(i, 2), 16);
    }

    return bytes;
}
console.log('js导入成功')

// 示例用法
const ciphertext ="18eb2d87f88705e1f2252e6a426dc96ca6f6b4f8bb1ee9b9c32fab63b105e51665266fe40720daa0bad1cf49a8bb64cdd9471fddfa1a63a6cd4c511c9fb8ec42dca02072a58d8908adfed346564208ed3c2fb956642aeb0df8bde8f923885c49ee5eb31eaa3ada304de25d377431c4da7437e2d20e07e1bd2969951f9675d411c48965"
const nonse='f5bb872a08ef929e6744d117'
const password = "111111"


const plaintext=aesDecrypt(ciphertext,password,nonse)

console.log("plaintext:",plaintext)