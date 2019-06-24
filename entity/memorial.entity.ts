import { Entity, Column, ObjectID, ObjectIdColumn } from 'typeorm'

/** 纪念日类型 */
export type MemorialType = 'love' | 'birthday'

/** 保存纪念日的实体 */
@Entity()
export class Memorial {
  @ObjectIdColumn()
  id?: ObjectID

  /** 纪念日类型 */
  @Column({ default: 'birthday' })
  type: MemorialType = 'love'

  /** 纪念日日期 */
  @Column({ default: () => 'CURRENT_TIMESTAMP' })
  date: Date = new Date()

  /** 纪念日相关人物姓名 */
  @Column()
  person: String = ''

  constructor(type: MemorialType, date: Date, person: string) {
    this.type = type
    this.date = date
    this.person = person
  }
}
