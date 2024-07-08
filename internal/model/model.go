package model

import (
    "time"
)

type MyModel struct {
    ID uint
    Time time.Time
    Unit string
    Value float64 
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (m *MyModel) SetIDTo(id uint) {
    m.ID = id
}

func (m *MyModel) SetTimeTo(time time.Time) {
    m.Time = time
}

func (m *MyModel) SetUnitTo(unit string) {
    m.Unit = unit
}

func (m *MyModel) SetValueTo(value float64) {
    m.Value = value
}
