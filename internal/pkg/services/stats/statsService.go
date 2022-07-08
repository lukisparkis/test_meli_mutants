package stats

import (
	"meli_test/pkg/datalayers/models"
)

//GetAllChecksPrevious get previous validations stored in db
func GetAllChecksPrevious() (*models.Stat, error) {
	var stat models.Stat
	err := stat.GetStats()
	if err != nil {
		return nil, err
	}
	return &stat, nil
}
