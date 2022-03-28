package database

import "vocabs-backend/api/model"

func AddVocab(toBeCreated model.Vocab) error {
	db := GetDatabase()
	return db.Create(toBeCreated).Error
}

func GetAllByUserId(userID string) (*[]model.Vocab, error) {
	var vocab []model.Vocab
	db := GetDatabase()
	err := db.Where("user_id = ?", userID).Find(&vocab).Error
	if err != nil {
		return nil, err
	}
	return &vocab, nil
}
