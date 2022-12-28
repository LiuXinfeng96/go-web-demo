package db

type DebrisType int32

const (
	SMALL DebrisType = iota + 1

	MEDIUM

	LARGE
)

const (
	DebrisTypeSmall = "Small"

	DebrisTypeMedium = "Medium"

	DebrisTypeLarge = "Large"
)

const DEBRIS_TABLE_NAME = "debris"

type Debris struct {
	GeneralField
	DebrisId     string `gorm:"index"`
	DebrisName   string `gorm:"index"`
	DebrisSource string
	Angle        float64
	Speed        float64
	Height       float64
	Volunme      float64
	Type         DebrisType
}

var DebrisTypeName = map[DebrisType]string{
	SMALL:  DebrisTypeSmall,
	MEDIUM: DebrisTypeMedium,
	LARGE:  DebrisTypeLarge,
}

var DebrisTypeValue = map[string]DebrisType{
	DebrisTypeSmall:  SMALL,
	DebrisTypeMedium: MEDIUM,
	DebrisTypeLarge:  LARGE,
}

func (d *Debris) TableName() string {
	return DEBRIS_TABLE_NAME
}

func init() {
	debris := new(Debris)
	TableSlice = append(TableSlice, &debris)
}
