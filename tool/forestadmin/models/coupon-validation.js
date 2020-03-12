// This model was generated by Lumber. However, you remain in control of your models.
// Learn how here: https://docs.forestadmin.com/documentation/v/v5/reference-guide/models/enrich-your-models
module.exports = (sequelize, DataTypes) => {
  const { Sequelize } = sequelize;
  // This section contains the fields of your model, mapped to your table's columns.
  // Learn more here: https://docs.forestadmin.com/documentation/v/v5/reference-guide/models/enrich-your-models#declaring-a-new-field-in-a-model
  const CouponValidation = sequelize.define('coupon_validation', {
    createdAt: {
      type: DataTypes.DATE,
    },
    updatedAt: {
      type: DataTypes.DATE,
    },
    comment: {
      type: DataTypes.STRING,
    },
  }, {
    tableName: 'coupon_validation',
    underscored: true,
  });

  // This section contains the relationships for this model. See: https://docs.forestadmin.com/documentation/v/v5/reference-guide/relationships#adding-relationships.
  CouponValidation.associate = (models) => {
    CouponValidation.belongsTo(models.user, {
      foreignKey: {
        name: 'authorIdKey',
        field: 'author_id',
      },
      as: 'author',
    });
    CouponValidation.belongsTo(models.team, {
      foreignKey: {
        name: 'teamIdKey',
        field: 'team_id',
      },
      as: 'team',
    });
    CouponValidation.belongsTo(models.coupon, {
      foreignKey: {
        name: 'couponIdKey',
        field: 'coupon_id',
      },
      as: 'coupon',
    });
  };

  return CouponValidation;
};