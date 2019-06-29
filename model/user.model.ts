import { prop, Typegoose, arrayProp, ModelType, staticMethod, pre } from 'typegoose'
import { Menstrual } from './menstrual.model'
import { Memorial } from './memorial.model'
import mongoose from 'mongoose'

const unique = true
const required = true

enum Gender {
  SECRET = 0,
  MALE,
  FEMALE,
}

@pre<User>('save', function(next) {
  this.updatedAt = new Date()
  next()
})
export class User extends Typegoose {
  @prop({ unique, required })
  openId!: string

  @prop()
  sessionKey?: string

  @prop()
  name?: string

  @prop()
  email?: string

  @prop()
  avatar?: string

  @prop({ enum: Gender, default: Gender.SECRET })
  gender: Gender = Gender.SECRET

  @prop()
  address?: string

  @prop({ _id: false })
  menstrual?: Menstrual

  @arrayProp({ items: Memorial, _id: false })
  memorials?: Memorial[]

  @prop({ required, default: new Date() })
  createdAt!: Date

  @prop({ required, default: new Date() })
  updatedAt!: Date

  @staticMethod
  static findByOpenId(this: ModelType<User> & typeof User, openId: string) {
    return this.findOne({ openId })
  }
}

export const UserModel = new User().getModelForClass(User, {
  existingConnection: mongoose.connection,
})
