import jwt from 'jsonwebtoken'
import { ObjectID } from 'typeorm'

const expiresIn = 24 * 3600 // 24h
export const jwtSecret = process.env.JWT_SECRET || 'HandNote!'

export function getToken(userId: string | ObjectID, openId: string, sessionKey: string) {
  return jwt.sign({ userId, openId, sessionKey }, jwtSecret, { expiresIn })
}
export function parseToken(token: string) {
  return jwt.verify(token, jwtSecret)
}
