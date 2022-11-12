package bluetooth

import (
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

type Scan struct {
	Id        uint64    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Address   string    `json:"address"`
	Name      string    `json:"name"`
	RSSI      int16     `json:"rssi"`
	CreatedAt time.Time `json:"created_at"`
}

func StartScanning() {
	if err := adapter.Enable(); err != nil {
		panic(err)
	}

	dsn := os.Getenv("DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&Scan{}); err != nil {
		panic(err)
	}

	if err := adapter.Scan(func(adapter *bluetooth.Adapter, event bluetooth.ScanResult) {
		scan := &Scan{
			Address:   event.Address.String(),
			Name:      event.LocalName(),
			RSSI:      event.RSSI,
			CreatedAt: time.Now(),
		}

		db.Create(scan)

	}); err != nil {
		panic(err)
	}
}
