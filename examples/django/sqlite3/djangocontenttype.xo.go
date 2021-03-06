// Package sqlite3 contains the types for schema ''.
package sqlite3

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
)

// DjangoContentType represents a row from 'django_content_type'.
type DjangoContentType struct {
	ID       int    `json:"id"`        // id
	AppLabel string `json:"app_label"` // app_label
	Model    string `json:"model"`     // model

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DjangoContentType exists in the database.
func (dct *DjangoContentType) Exists() bool {
	return dct._exists
}

// Deleted provides information if the DjangoContentType has been deleted from the database.
func (dct *DjangoContentType) Deleted() bool {
	return dct._deleted
}

// Insert inserts the DjangoContentType to the database.
func (dct *DjangoContentType) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if dct._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO django_content_type (` +
		`app_label, model` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, dct.AppLabel, dct.Model)
	res, err := db.Exec(sqlstr, dct.AppLabel, dct.Model)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	dct.ID = int(id)
	dct._exists = true

	return nil
}

// Update updates the DjangoContentType in the database.
func (dct *DjangoContentType) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !dct._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if dct._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE django_content_type SET ` +
		`app_label = ?, model = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, dct.AppLabel, dct.Model, dct.ID)
	_, err = db.Exec(sqlstr, dct.AppLabel, dct.Model, dct.ID)
	return err
}

// Save saves the DjangoContentType to the database.
func (dct *DjangoContentType) Save(db XODB) error {
	if dct.Exists() {
		return dct.Update(db)
	}

	return dct.Insert(db)
}

// Delete deletes the DjangoContentType from the database.
func (dct *DjangoContentType) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !dct._exists {
		return nil
	}

	// if deleted, bail
	if dct._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM django_content_type WHERE id = ?`

	// run query
	XOLog(sqlstr, dct.ID)
	_, err = db.Exec(sqlstr, dct.ID)
	if err != nil {
		return err
	}

	// set deleted
	dct._deleted = true

	return nil
}

// DjangoContentTypeByAppLabelModel retrieves a row from 'django_content_type' as a DjangoContentType.
//
// Generated from index 'django_content_type_app_label_76bd3d3b_uniq'.
func DjangoContentTypeByAppLabelModel(db XODB, appLabel string, model string) (*DjangoContentType, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, app_label, model ` +
		`FROM django_content_type ` +
		`WHERE app_label = ? AND model = ?`

	// run query
	XOLog(sqlstr, appLabel, model)
	dct := DjangoContentType{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, appLabel, model).Scan(&dct.ID, &dct.AppLabel, &dct.Model)
	if err != nil {
		return nil, err
	}

	return &dct, nil
}

// DjangoContentTypeByID retrieves a row from 'django_content_type' as a DjangoContentType.
//
// Generated from index 'django_content_type_id_pkey'.
func DjangoContentTypeByID(db XODB, id int) (*DjangoContentType, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, app_label, model ` +
		`FROM django_content_type ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	dct := DjangoContentType{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&dct.ID, &dct.AppLabel, &dct.Model)
	if err != nil {
		return nil, err
	}

	return &dct, nil
}
