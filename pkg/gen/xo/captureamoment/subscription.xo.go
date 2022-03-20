package captureamoment

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Subscription represents a row from 'captureamoment.subscription'.
type Subscription struct {
	UserID         int `json:"UserID"`         // UserID
	FollowedUserID int `json:"FollowedUserID"` // FollowedUserID
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Subscription exists in the database.
func (s *Subscription) Exists() bool {
	return s._exists
}

// Deleted returns true when the Subscription has been marked for deletion from
// the database.
func (s *Subscription) Deleted() bool {
	return s._deleted
}

// Insert inserts the Subscription to the database.
func (s *Subscription) Insert(ctx context.Context, db DB) error {
	switch {
	case s._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case s._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO captureamoment.subscription (` +
		`UserID, FollowedUserID` +
		`) VALUES (` +
		`?, ?` +
		`)`
	// run
	logf(sqlstr, s.UserID, s.FollowedUserID)
	if _, err := db.ExecContext(ctx, sqlstr, s.UserID, s.FollowedUserID); err != nil {
		return logerror(err)
	}
	// set exists
	s._exists = true
	return nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the Subscription from the database.
func (s *Subscription) Delete(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return nil
	case s._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	const sqlstr = `DELETE FROM captureamoment.subscription ` +
		`WHERE UserID = ? AND FollowedUserID = ?`
	// run
	logf(sqlstr, s.UserID, s.FollowedUserID)
	if _, err := db.ExecContext(ctx, sqlstr, s.UserID, s.FollowedUserID); err != nil {
		return logerror(err)
	}
	// set deleted
	s._deleted = true
	return nil
}

// SubscriptionByFollowedUserID retrieves a row from 'captureamoment.subscription' as a Subscription.
//
// Generated from index 'subscription_FollowedID_fk'.
func SubscriptionByFollowedUserID(ctx context.Context, db DB, followedUserID int) ([]*Subscription, error) {
	// query
	const sqlstr = `SELECT ` +
		`UserID, FollowedUserID ` +
		`FROM captureamoment.subscription ` +
		`WHERE FollowedUserID = ?`
	// run
	logf(sqlstr, followedUserID)
	rows, err := db.QueryContext(ctx, sqlstr, followedUserID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Subscription
	for rows.Next() {
		s := Subscription{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&s.UserID, &s.FollowedUserID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// SubscriptionByUserIDFollowedUserID retrieves a row from 'captureamoment.subscription' as a Subscription.
//
// Generated from index 'subscription_UserID_FollowedUserID_pkey'.
func SubscriptionByUserIDFollowedUserID(ctx context.Context, db DB, userID, followedUserID int) (*Subscription, error) {
	// query
	const sqlstr = `SELECT ` +
		`UserID, FollowedUserID ` +
		`FROM captureamoment.subscription ` +
		`WHERE UserID = ? AND FollowedUserID = ?`
	// run
	logf(sqlstr, userID, followedUserID)
	s := Subscription{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, userID, followedUserID).Scan(&s.UserID, &s.FollowedUserID); err != nil {
		return nil, logerror(err)
	}
	return &s, nil
}

// SubscriptionByUserID retrieves a row from 'captureamoment.subscription' as a Subscription.
//
// Generated from index 'subscription_UserID_fk'.
func SubscriptionByUserID(ctx context.Context, db DB, userID int) ([]*Subscription, error) {
	// query
	const sqlstr = `SELECT ` +
		`UserID, FollowedUserID ` +
		`FROM captureamoment.subscription ` +
		`WHERE UserID = ?`
	// run
	logf(sqlstr, userID)
	rows, err := db.QueryContext(ctx, sqlstr, userID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*Subscription
	for rows.Next() {
		s := Subscription{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&s.UserID, &s.FollowedUserID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// User returns the User associated with the Subscription's (FollowedUserID).
//
// Generated from foreign key 'subscription_FollowedID_fk'.
func (s *Subscription) FollowedUser(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, s.FollowedUserID)
}

// User returns the User associated with the Subscription's (UserID).
//
// Generated from foreign key 'subscription_UserID_fk'.
func (s *Subscription) User(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, s.UserID)
}
