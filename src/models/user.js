'use strict'
module.exports = (sequelize, DataTypes) => {
  const user = sequelize.define(
    'user',
    {
      open_id: DataTypes.STRING,
      name: DataTypes.STRING,
      avatar: DataTypes.STRING,
      gender: DataTypes.TINYINT,
      location: DataTypes.STRING,
    },
    {
      underscored: true,
    },
  )
  user.associate = function(models) {
    // associations can be defined here
  }
  return user
}
