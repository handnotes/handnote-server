import { PrimaryGeneratedColumn, Column, Entity, CreateDateColumn, UpdateDateColumn } from 'typeorm'

const nullable = true
const unique = true

@Entity()
export class User {
  @PrimaryGeneratedColumn()
  id!: number

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

  @CreateDateColumn()
  createdAt!: Date

  @UpdateDateColumn()
  updatedAt!: Date
}
