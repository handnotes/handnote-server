import crypto from 'crypto'

export class WXBizDataCrypt {
  appId: string
  sessionKey: string

  constructor(appId: string, sessionKey: string) {
    this.appId = appId
    this.sessionKey = sessionKey
  }

  decryptData(encryptedData: string, iv: string) {
    // base64 decode
    const sessionKey = Buffer.from(this.sessionKey, 'base64')
    // const encryptedDataBuffer = Buffer.from(encryptedData, 'base64')
    const ivBuffer = Buffer.from(iv, 'base64')

    let decoded
    try {
      // 解密
      const decipher = crypto.createDecipheriv('aes-128-cbc', sessionKey, ivBuffer)
      // 设置自动 padding 为 true，删除填充补位
      decipher.setAutoPadding(true)
      decoded = decipher.update(encryptedData, 'base64', 'utf8')
      decoded += decipher.final('utf8')

      decoded = JSON.parse(decoded)
    } catch (err) {
      console.error(err)
      throw new Error('Illegal Buffer')
    }

    if (decoded.watermark.appid !== this.appId) {
      throw new Error('Illegal Buffer')
    }

    return decoded
  }
}
