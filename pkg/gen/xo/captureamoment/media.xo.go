package captureamoment

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Media represents a row from 'captureamoment.media'.
type Media struct {
	MediaID   int            `json:"MediaID"`   // MediaID
	MediaLink sql.NullString `json:"MediaLink"` // MediaLink
	Postid    int            `json:"postid"`    // postid
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Media exists in the database.
func (m *Media) Exists() bool {
	return m._exists
}

// Deleted returns true when the Media has been marked for deletion from
// the database.
func (m *Media) Deleted() bool {
	return m._deleted
}

// Insert inserts the Media to the database.
func (m *Media) Insert(ctx context.Context, db DB) error {
	switch {
	case m._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case m._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO captureamoment.media (` +
		`MediaLink, postid` +
		`) VALUES (` +
		`?, ?` +
		`)`
	// run
	logf(sqlstr, m.MediaLink, m.Postid)
	res, err := db.ExecContext(ctx, sqlstr, m.MediaLink, m.Postid)
	if err != nil {
		return logerror(err)
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return logerror(err)
	} // set primary key
	m.MediaID = int(id)
	// set exists
	m._exists = true
	return nil
}

// Update updates a Media in the database.
func (m *Media) Update(ctx context.Context, db DB) error {
	switch {
	case !m._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case m._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE captureamoment.media SET ` +
		`MediaLink = ?, postid = ? ` +
		`WHERE MediaID = ?`
	// run
	logf(sqlstr, m.MediaLink, m.Postid, m.MediaID)
	if _, err := db.ExecContext(ctx, sqlstr, m.MediaLink, m.Postid, m.MediaID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Media to the database.
func (m *Media) Save(ctx context.Context, db DB) error {
	if m.Exists() {
		return m.Update(ctx, db)
	}
	return m.Insert(ctx, db)
}

// Upsert performs an upsert for Media.
func (m *Media) Upsert(ctx context.Context, db DB) error {
	switch {
	case m._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO captureamoment.media (` +
		`MediaID, MediaLink, postid` +
		`) VALUES (` +
		`?, ?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`MediaLink = VALUES(MediaLink), postid = VALUES(postid)`
	// run
	logf(sqlstr, m.MediaID, m.MediaLink, m.Postid)
	if _, err := db.ExecContext(ctx, sqlstr, m.MediaID, m.MediaLink, m.Postid); err != nil {
		return logerror(err)
	}
	// set exists
	m._exists = true
	return nil
}

// Delete deletes the Media from the database.
func (m *Media) Delete(ctx context.Context, db DB) error {
	switch {
	case !m._exists: // doesn't exist
		return nil
	case m._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM captureamoment.media ` +
		`WHERE MediaID = ?`
	// run
	logf(sqlstr, m.MediaID)
	if _, err := db.ExecContext(ctx, sqlstr, m.MediaID); err != nil {
		return logerror(err)
	}
	// set deleted
	m._deleted = true
	return nil
}

// MediaByMediaID retrieves a row from 'captureamoment.media' as a Media.
//
// Generated from index 'media_MediaID_pkey'.
func MediaByMediaID(ctx context.Context, db DB, mediaID int) (*Media, error) {
	// query
	const sqlstr = `SELECT ` +
		`MediaID, MediaLink, postid ` +
		`FROM captureamoment.media ` +
		`WHERE MediaID = ?`
	// run
	logf(sqlstr, mediaID)
	m := Media{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, mediaID).Scan(&m.MediaID, &m.MediaLink, &m.Postid); err != nil {
		return nil, logerror(err)
	}
	return &m, nil
}

// MediaByPostid retrieves a row from 'captureamoment.media' as a Media.
//
// Generated from index 'mediapost_idx'.
func MediaByPostid(ctx context.Context, db DB, postid int) ([]*Media, error) {
	// query
	const sqlstr = `SELECT ` +
		`MediaID, MediaLink, postid ` +
		`FROM captureamoment.media ` +
		`WHERE postid = ?`
	// run
	logf(sqlstr, postid)
	rows, err := db.QueryContext(ctx, sqlstr, postid)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Media
	for rows.Next() {
		m := Media{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&m.MediaID, &m.MediaLink, &m.Postid); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &m)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Post returns the Post associated with the Media's (Postid).
//
// Generated from foreign key 'mediapost'.
func (m *Media) Post(ctx context.Context, db DB) (*Post, error) {
	return PostByPostID(ctx, db, m.Postid)
}
