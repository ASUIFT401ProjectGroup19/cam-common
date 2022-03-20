package captureamoment

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Tag represents a row from 'captureamoment.tag'.
type Tag struct {
	TagID  int    `json:"TagID"`  // TagID
	TagMsg string `json:"TagMsg"` // TagMsg
	PostID int    `json:"PostID"` // PostID
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Tag exists in the database.
func (t *Tag) Exists() bool {
	return t._exists
}

// Deleted returns true when the Tag has been marked for deletion from
// the database.
func (t *Tag) Deleted() bool {
	return t._deleted
}

// Insert inserts the Tag to the database.
func (t *Tag) Insert(ctx context.Context, db DB) error {
	switch {
	case t._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case t._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO captureamoment.tag (` +
		`TagMsg, PostID` +
		`) VALUES (` +
		`?, ?` +
		`)`
	// run
	logf(sqlstr, t.TagMsg, t.PostID)
	res, err := db.ExecContext(ctx, sqlstr, t.TagMsg, t.PostID)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	t.TagID = int(id)
	// set exists
	t._exists = true
	return nil
}

// Update updates a Tag in the database.
func (t *Tag) Update(ctx context.Context, db DB) error {
	switch {
	case !t._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case t._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE captureamoment.tag SET ` +
		`TagMsg = ?, PostID = ? ` +
		`WHERE TagID = ?`
	// run
	logf(sqlstr, t.TagMsg, t.PostID, t.TagID)
	if _, err := db.ExecContext(ctx, sqlstr, t.TagMsg, t.PostID, t.TagID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Tag to the database.
func (t *Tag) Save(ctx context.Context, db DB) error {
	if t.Exists() {
		return t.Update(ctx, db)
	}
	return t.Insert(ctx, db)
}

// Upsert performs an upsert for Tag.
func (t *Tag) Upsert(ctx context.Context, db DB) error {
	switch {
	case t._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO captureamoment.tag (` +
		`TagID, TagMsg, PostID` +
		`) VALUES (` +
		`?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`TagMsg = VALUES(TagMsg), PostID = VALUES(PostID)`
	// run
	logf(sqlstr, t.TagID, t.TagMsg, t.PostID)
	if _, err := db.ExecContext(ctx, sqlstr, t.TagID, t.TagMsg, t.PostID); err != nil {
		return logerror(err)
	}
	// set exists
	t._exists = true
	return nil
}

// Delete deletes the Tag from the database.
func (t *Tag) Delete(ctx context.Context, db DB) error {
	switch {
	case !t._exists: // doesn't exist
		return nil
	case t._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM captureamoment.tag ` +
		`WHERE TagID = ?`
	// run
	logf(sqlstr, t.TagID)
	if _, err := db.ExecContext(ctx, sqlstr, t.TagID); err != nil {
		return logerror(err)
	}
	// set deleted
	t._deleted = true
	return nil
}

// TagByPostID retrieves a row from 'captureamoment.tag' as a Tag.
//
// Generated from index 'tag_PostID_fk'.
func TagByPostID(ctx context.Context, db DB, postID int) ([]*Tag, error) {
	// query
	const sqlstr = `SELECT ` +
		`TagID, TagMsg, PostID ` +
		`FROM captureamoment.tag ` +
		`WHERE PostID = ?`
	// run
	logf(sqlstr, postID)
	rows, err := db.QueryContext(ctx, sqlstr, postID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Tag
	for rows.Next() {
		t := Tag{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&t.TagID, &t.TagMsg, &t.PostID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// TagByTagID retrieves a row from 'captureamoment.tag' as a Tag.
//
// Generated from index 'tag_TagID_pkey'.
func TagByTagID(ctx context.Context, db DB, tagID int) (*Tag, error) {
	// query
	const sqlstr = `SELECT ` +
		`TagID, TagMsg, PostID ` +
		`FROM captureamoment.tag ` +
		`WHERE TagID = ?`
	// run
	logf(sqlstr, tagID)
	t := Tag{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, tagID).Scan(&t.TagID, &t.TagMsg, &t.PostID); err != nil {
		return nil, logerror(err)
	}
	return &t, nil
}

// Post returns the Post associated with the Tag's (PostID).
//
// Generated from foreign key 'tag_PostID_fk'.
func (t *Tag) Post(ctx context.Context, db DB) (*Post, error) {
	return PostByPostID(ctx, db, t.PostID)
}
