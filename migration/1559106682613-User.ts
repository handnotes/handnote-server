import { MigrationInterface, QueryRunner, Table } from 'typeorm'

export class User1559106682613 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<any> {
    await queryRunner.createTable(
      new Table({
        name: 'user',
        columns: [
          { name: 'id', type: 'int', isPrimary: true, isGenerated: true },
          { name: 'openId', type: 'varchar(32)', isUnique: true },
          { name: 'sessionKey', type: 'varchar(32)', isNullable: true },
          { name: 'email', type: 'varchar', isUnique: true, isNullable: true },
          { name: 'name', type: 'varchar(32)', isNullable: true },
          { name: 'avatar', type: 'varchar', isNullable: true },
          { name: 'gender', type: 'tinyint(1)', isNullable: true },
          { name: 'location', type: 'varchar', isNullable: true },
          { name: 'createdAt', type: 'datetime' },
          { name: 'updatedAt', type: 'datetime' },
        ],
      }),
      true,
    )
  }

  public async down(queryRunner: QueryRunner): Promise<any> {
    await queryRunner.dropTable('user')
  }
}
