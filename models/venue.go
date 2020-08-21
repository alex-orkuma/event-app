package models

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

//Venue struct represent the Venue data model
type Venue struct {
	gorm.Model
	Name        string `gorm:"size:100;not null;unique" json:"name"`
	Description string `gorm:"not null"                 json:"description"`
	Location    string `gorm:"size:100;not null"        json:"location"`
	Capacity    int    `gorm:"not null"                 json:"capacity"`
	Category    string `gorm:"size:100;not null"        json:"category"`
	CreatedBy   User   `gorm:"foreignKey:UserID;"       json:"-"`
	UserID      uint   `gorm:"not null"                 json:"user_id"`
}

//Prepare prepares the data
func (v *Venue) Prepare() {
	v.Name = strings.TrimSpace(v.Name)
	v.Description = strings.TrimSpace(v.Description)
	v.Location = strings.TrimSpace(v.Location)
	v.Category = strings.TrimSpace(v.Category)
	v.CreatedBy = User{}
}

//Validate validates the data
func (v *Venue) Validate() error {
	if v.Name == "" {
		return errors.New("Name is required")
	}
	if v.Description == "" {
		return errors.New("Description about venue is required")
	}
	if v.Location == "" {
		return errors.New("Location of venue is required")
	}
	if v.Category == "" {
		return errors.New("Category of venue is required")
	}
	if v.Capacity < 0 {
		return errors.New("Capacity of venue is invalid")
	}
	return nil
}

// Save saves the data into the database
func (v *Venue) Save(db *gorm.DB) (*Venue, error) {
	var err error

	// Debug a single operation, show detailed log for this operation
	err = db.Debug().Create(&v).Error
	if err != nil {
		return &Venue{}, err
	}
	return v, nil
}

//GetVenue returns a venue based on the name of the venue given
func (v *Venue) GetVenue(db *gorm.DB) (*Venue, error) {
	venue := &Venue{}
	if err := db.Debug().Table("venues").Where("name = ?", v.Name).First(venue).Error; err != nil {
		return nil, err
	}
	return venue, nil
}

//GetVenues returns all the venues
func GetVenues(db *gorm.DB) (*[]Venue, error) {
	venues := []Venue{}
	if err := db.Debug().Table("venues").Find(&venues).Error; err != nil {
		return &[]Venue{}, err
	}
	return &venues, nil
}

//GetVenueById returns the venue given the id
func GetVenueById(id int, db *gorm.DB) (*Venue, error) {
	venue := &Venue{}
	if err := db.Debug().Table("venues").Where("id = ?", id).First(venue).Error; err != nil {
		return nil, err
	}
	return venue, nil
}

//UpdateVenue updates the venue data in the database
func (v *Venue) UpdateVenue(id int, db *gorm.DB) (*Venue, error) {
	if err := db.Debug().Table("venues").Where("id = ?", id).Updates(Venue{
		Name:        v.Name,
		Description: v.Description,
		Location:    v.Location,
		Capacity:    v.Capacity,
		Category:    v.Category}).Error; err != nil {
		return &Venue{}, err
	}
	return v, nil
}

//DeleteVenue deletes the venue from the database
func DeleteVenue(id int, db *gorm.DB) error {
	if err := db.Debug().Table("venues").Where("id = ?", id).Delete(&Venue{}).Error; err != nil {
		return err
	}
	return nil
}
