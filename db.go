package main

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type dbMatch struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	MatchID int64         `bson:"match_id,omitempty"`
	PostID  int           `bson:"post_id,omitempty"`
	IsShort bool          `bson:"is_short,omitempty"`
	IsFull  bool          `bson:"is_full,omitempty"`
}

var db *mgo.Database

type dbDao struct {
	Server   string
	Database string
}

func (m *dbDao) connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *dbDao) insert(what interface{}) error {
	err := db.C("matches").Insert(&what)
	return err
}

func (m *dbDao) get(what interface{}) (dbMatch, error) {
	var result dbMatch
	err := db.C("matches").Find(what).One(&result)
	return result, err
}

func (m *dbDao) edit(what interface{}, with interface{}) error {
	change := mgo.Change{
		Update:    bson.M{"$set": with},
		ReturnNew: true,
	}
	_, err := db.C("matches").Find(what).Apply(change, &with)
	return err
}

func addMatch(matchID int64) {
	what := dbMatch{
		MatchID: matchID,
	}
	dao.insert(what)
}

func markShort(matchID int64) {
	what := dbMatch{
		MatchID: matchID,
	}
	with := dbMatch{
		IsShort: true,
	}
	dao.edit(what, with)
}

func markFull(matchID int64) {
	what := dbMatch{
		MatchID: matchID,
	}
	with := dbMatch{
		IsFull: true,
	}
	dao.edit(what, with)
}

func addPostID(matchID int64, postID int) {
	what := dbMatch{
		MatchID: matchID,
	}
	with := dbMatch{
		PostID: postID,
	}
	dao.edit(what, with)
}
