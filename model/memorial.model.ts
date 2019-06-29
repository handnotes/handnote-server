import { prop } from 'typegoose'

/** 纪念日类型 */
export type MemorialType = 'love' | 'birthday'

/** 保存纪念日的实体 */
export class Memorial {
  /** 纪念日类型 */
  @prop()
  type: MemorialType = 'love'

  /** 纪念日日期 */
  @prop()
  date: Date = new Date()

  /** 纪念日相关人物姓名 */
  @prop()
  person: String = ''

  constructor(type: MemorialType, date: Date, person: string) {
    this.type = type
    this.date = date
    this.person = person
  }
}
