package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Genre struct {
	GenreID   bson.ObjectID `bson:"genre_id" json:"genre_id" validade:"required"`
	GenreName string        `bson:"genre_name" json:"genre_name" validade:"required,min=2,max=100"`
}

type Ranking struct {
	RankingValue int    `bson:"ranking_value" json:"ranking_value" validade:"required"`
	RankingName  string `bson:"ranking_name" json:"ranking_name" validade:"oneof=Excellent Good Okay Bad Terrible"`
}

type Movie struct {
	ID          bson.ObjectID `bson:"_id" json:"_id"`
	ImdbID      string        `bson:"imdb_id" json:"imdb_id" validade:"required"`
	Title       string        `bson:"title" json:"title" validade:"required,min=2,max=500"`
	PosterPath  string        `bson:"poster_path" json:"poster_path" validade:"required,url"`
	YoutubeID   string        `bson:"youtube_id" json:"youtube_id" validade:"required"`
	Genre       []Genre       `bson:"genre" json:"genre" validade:"required,dive"`
	AdminReview string        `bson:"admin_review" json:"admin_review" validade:"required"`
	Ranking     Ranking       `bson:"ranking" json:"ranking" validade:"required"`
}
