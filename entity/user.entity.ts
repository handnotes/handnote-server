import {
  Column,
  Entity,
  CreateDateColumn,
  UpdateDateColumn,
  ObjectIdColumn,
  ObjectID,
} from 'typeorm'
import { Menstrual } from './menstrual.entity'

const nullable = true
const unique = true

@Entity()
export class User {
  @ObjectIdColumn()
  id!: ObjectID

  @Column({ unique, select: false })
  openId!: string

  @Column({ nullable, select: false })
  sessionKey?: string

  @Column({ nullable })
  name?: string

  @Column({ nullable })
  email?: string

  @Column({ nullable })
  avatar?: string

  @Column({ type: 'tinyint', nullable })
  gender?: number

  @Column({ nullable })
  address?: string

  @Column(() => Menstrual)
  menstrual?: Menstrual

  @CreateDateColumn()
  createdAt!: Date

  @UpdateDateColumn()
  updatedAt!: Date
}
