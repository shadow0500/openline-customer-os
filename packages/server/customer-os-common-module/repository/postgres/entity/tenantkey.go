package entity

type TenantKey struct {
	ID     uint64 `gorm:"primary_key;autoIncrement:true" json:"id"`
	Tenant string `gorm:"column:tenant;type:varchar(255);NOT NULL" json:"tenant" binding:"required"`
	Key    string `gorm:"column:key;type:varchar(255);NOT NULL;index:idx_key,unique" json:"key" binding:"required"`
	Active bool   `gorm:"column:active;type:bool;NOT NULL" json:"active" binding:"required"`
}

func (TenantKey) TableName() string {
	return "tenant_keys"
}
