export declare namespace wx {
  interface GetUserInfoSuccessCallbackResult {
    /** 包括敏感数据在内的完整用户信息的加密数据，详见 [用户数据的签名验证和加解密] */
    encryptedData: string
    /** 加密算法的初始向量，详见 [用户数据的签名验证和加解密] */
    iv: string
    /** 不包括敏感信息的原始数据字符串，用于计算签名 */
    rawData: string
    /** 使用 sha1( rawData + sessionkey ) 得到字符串，用于校验用户信息，详见 [用户数据的签名验证和加解密] */
    signature: string
    /** [UserInfo]
     *
     * 用户信息对象，不包含 openid 等敏感信息 */
    userInfo: UserInfo
  }

  /** 用户信息 */
  interface UserInfo {
    /** 用户头像图片的 URL。URL 最后一个数值代表正方形头像大小（有 0、46、64、96、132 数值可选，0 代表 640x640 的正方形头像，46 表示 46x46 的正方形头像，剩余数值以此类推。默认132），用户没有头像时该项为空。若用户更换头像，原有头像 URL 将失效。 */
    avatarUrl: string
    /** 用户所在城市 */
    city: string
    /** 用户所在国家 */
    country: string
    /** 用户性别
     *
     * 可选值：
     * - 0: 未知;
     * - 1: 男性;
     * - 2: 女性; */
    gender: 0 | 1 | 2
    /** 显示 country，province，city 所用的语言
     *
     * 可选值：
     * - 'en': 英文;
     * - 'zh_CN': 简体中文;
     * - 'zh_TW': 繁体中文; */
    language: 'en' | 'zh_CN' | 'zh_TW'
    /** 用户昵称 */
    nickName: string
    /** 用户所在省份 */
    province: string
  }
}
