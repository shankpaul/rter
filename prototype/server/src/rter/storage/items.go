package storage

import (
	"fmt"
	"rter/data"
	"time"
)

func InsertItem(item *data.Item) error {
	ID, err := InsertEntry(
		"INSERT INTO Items (Type, AuthorID, ThumbnailURI, ContentURI, UploadURI, HasGeo, Heading, Lat, Lng, StartTime, StopTime) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		item.Type,
		item.AuthorID,
		item.ThumbnailURI,
		item.ContentURI,
		item.UploadURI,
		item.HasGeo,
		item.Heading,
		item.Lat,
		item.Lng,
		item.StartTime.UTC(),
		item.StopTime.UTC(),
	)

	if err != nil {
		return err
	}

	item.ID = ID

	return nil
}

func SelectItem(ID int64) (*data.Item, error) {
	rows, err := Query("SELECT * FROM Items WHERE ID=?", ID)

	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, fmt.Errorf("Select Failed, no Item in storage where ID=%v", ID)
	}

	item := new(data.Item)

	var startTimeString, stopTimeString string

	err = rows.Scan(
		&item.ID,
		&item.Type,
		&item.AuthorID,
		&item.ThumbnailURI,
		&item.ContentURI,
		&item.UploadURI,
		&item.HasGeo,
		&item.Heading,
		&item.Lat,
		&item.Lng,
		&startTimeString,
		&stopTimeString,
	)

	if err != nil {
		return nil, err
	}

	startTime, err := time.Parse("2006-01-02 15:04:05", startTimeString) // this assumes UTC as timezone

	if err != nil {
		return nil, err
	}

	item.StartTime = startTime

	stopTime, err := time.Parse("2006-01-02 15:04:05", stopTimeString) // this assumes UTC as timezone

	if err != nil {
		return nil, err
	}

	item.StopTime = stopTime

	return item, nil
}

func DeleteItem(item *data.Item) error {
	return DeleteEntry("DELETE FROM Items WHERE ID=?", item.ID)
}

func InsertItemComment(comment *data.ItemComment) error {
	ID, err := InsertEntry(
		"INSERT INTO ItemComments (ItemID, AuthorID, Body, CreateTime) VALUES (?, ?, ?, ?)",
		comment.ItemID,
		comment.AuthorID,
		comment.Body,
		comment.CreateTime,
	)

	if err != nil {
		return err
	}

	comment.ID = ID

	return nil
}

func SelectItemComment(ID int64) (*data.ItemComment, error) {
	rows, err := Query("SELECT * FROM ItemComments WHERE ID=?", ID)

	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, fmt.Errorf("Select Failed, no ItemComment in storage where ID=%v", ID)
	}

	comment := new(data.ItemComment)

	var createTimeString string

	err = rows.Scan(
		&comment.ID,
		&comment.ItemID,
		&comment.AuthorID,
		&comment.Body,
		&createTimeString,
	)

	if err != nil {
		return nil, err
	}

	createTime, err := time.Parse("2006-01-02 15:04:05", createTimeString) // this assumes UTC as timezone

	if err != nil {
		return nil, err
	}

	comment.CreateTime = createTime

	return comment, nil
}

func DeleteItemComment(comment *data.ItemComment) error {
	return DeleteEntry("DELETE FROM ItemComments WHERE ID=?", comment.ID)
}
