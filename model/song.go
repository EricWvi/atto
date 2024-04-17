package model

import (
	"errors"
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	SongField
}

const (
	Song_Table    = "songs"
	Song_Sid      = "sid"
	Song_AlbumId  = "album_id"
	Song_ArtistId = "artist_id"
	Song_Title    = "title"
	Song_Comment  = "comment"
	Song_Duration = "duration"
	Song_Track    = "track"
	Song_Disc     = "disc"
	Song_Link     = "link"
	Song_Cover    = "cover"
)

type SongField struct {
	Sid      Sid     `json:"sid"`
	AlbumId  Sid     `json:"album_id"`
	ArtistId Sid     `json:"artist_id"`
	Title    string  `json:"title"`
	Comment  string  `json:"comment"`
	Duration float64 `json:"duration"`
	Track    int     `json:"track"`
	Disc     int     `json:"disc"`
	Link     string  `json:"link"`
	Cover    string  `json:"cover"`
}

func (s *SongField) Album() Sid {
	return s.AlbumId
}

func (s *SongField) Artist() Sid {
	return s.ArtistId
}

func (s *Song) Create(db *gorm.DB) error {
	return db.Create(s).Error
}

func GetSong(db *gorm.DB, where map[string]any) (s *Song, err error) {
	var songs []Song
	db.Table(Song_Table).Where(where).Order("created_at DESC").Find(&songs)
	if len(songs) == 0 {
		err = errors.New("song not found")
		return
	}
	s = &songs[0]
	return
}

type SongView struct {
	Sid      string `json:"sid"`
	AlbumId  string `json:"album_id"`
	ArtistId string `json:"artist_id"`
	SongField
}

func GetSongView(s *Song) *SongView {
	return &SongView{
		Sid:       s.Sid.String(),
		AlbumId:   s.AlbumId.String(),
		ArtistId:  s.ArtistId.String(),
		SongField: s.SongField,
	}
}
