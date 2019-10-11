// Code generated by shiftgen at shift.go:9. DO NOT EDIT.

package rounds

import (
	"context"
	"database/sql"
	"strings"
	"time"
	"github.com/luno/jettison/errors"
	"github.com/luno/jettison/j"
	"github.com/luno/shift"
)

// Insert inserts a new player_rounds table entity. All the fields of the 
// joining receiver are set, as well as status, created_at and updated_at. 
// The newly created entity id is returned on success or an error.
func (一 joining) Insert(ctx context.Context, tx *sql.Tx,st shift.Status) (int64, error) {
	var (
		q    strings.Builder
		args []interface{}
	)

	q.WriteString("insert into player_rounds set `status`=?, `created_at`=?, `updated_at`=? ")
	args = append(args, st.Enum(), time.Now(), time.Now())

	res, err := tx.ExecContext(ctx, q.String(), args...)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Update updates the status of a player_rounds table entity. All the fields of the
// joined receiver are updated, as well as status and updated_at. 
// The entity id is returned on success or an error.
func (一 joined) Update(ctx context.Context, tx *sql.Tx,from shift.Status, 
	to shift.Status) (int64, error) {
	var (
		q    strings.Builder
		args []interface{}
	)

	q.WriteString("update player_rounds set `status`=?, `updated_at`=? ")
	args = append(args, to.Enum(), time.Now())

	q.WriteString(" where `id`=? and `status`=?")
	args = append(args, 一.ID, from.Enum())

	res, err := tx.ExecContext(ctx, q.String(), args...)
	if err != nil {
		return 0, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, errors.Wrap(shift.ErrRowCount, "joined", j.KV("count", n))
	}

	return 一.ID, nil
}

// Update updates the status of a player_rounds table entity. All the fields of the
// excluded receiver are updated, as well as status and updated_at. 
// The entity id is returned on success or an error.
func (一 excluded) Update(ctx context.Context, tx *sql.Tx,from shift.Status, 
	to shift.Status) (int64, error) {
	var (
		q    strings.Builder
		args []interface{}
	)

	q.WriteString("update player_rounds set `status`=?, `updated_at`=? ")
	args = append(args, to.Enum(), time.Now())

	q.WriteString(" where `id`=? and `status`=?")
	args = append(args, 一.ID, from.Enum())

	res, err := tx.ExecContext(ctx, q.String(), args...)
	if err != nil {
		return 0, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, errors.Wrap(shift.ErrRowCount, "excluded", j.KV("count", n))
	}

	return 一.ID, nil
}

// Update updates the status of a player_rounds table entity. All the fields of the
// collecting receiver are updated, as well as status and updated_at. 
// The entity id is returned on success or an error.
func (一 collecting) Update(ctx context.Context, tx *sql.Tx,from shift.Status, 
	to shift.Status) (int64, error) {
	var (
		q    strings.Builder
		args []interface{}
	)

	q.WriteString("update player_rounds set `status`=?, `updated_at`=? ")
	args = append(args, to.Enum(), time.Now())

	q.WriteString(" where `id`=? and `status`=?")
	args = append(args, 一.ID, from.Enum())

	res, err := tx.ExecContext(ctx, q.String(), args...)
	if err != nil {
		return 0, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, errors.Wrap(shift.ErrRowCount, "collecting", j.KV("count", n))
	}

	return 一.ID, nil
}

// Update updates the status of a player_rounds table entity. All the fields of the
// collected receiver are updated, as well as status and updated_at. 
// The entity id is returned on success or an error.
func (一 collected) Update(ctx context.Context, tx *sql.Tx,from shift.Status, 
	to shift.Status) (int64, error) {
	var (
		q    strings.Builder
		args []interface{}
	)

	q.WriteString("update player_rounds set `status`=?, `updated_at`=? ")
	args = append(args, to.Enum(), time.Now())

	q.WriteString(" where `id`=? and `status`=?")
	args = append(args, 一.ID, from.Enum())

	res, err := tx.ExecContext(ctx, q.String(), args...)
	if err != nil {
		return 0, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, errors.Wrap(shift.ErrRowCount, "collected", j.KV("count", n))
	}

	return 一.ID, nil
}

// Update updates the status of a player_rounds table entity. All the fields of the
// submitting receiver are updated, as well as status and updated_at. 
// The entity id is returned on success or an error.
func (一 submitting) Update(ctx context.Context, tx *sql.Tx,from shift.Status, 
	to shift.Status) (int64, error) {
	var (
		q    strings.Builder
		args []interface{}
	)

	q.WriteString("update player_rounds set `status`=?, `updated_at`=? ")
	args = append(args, to.Enum(), time.Now())

	q.WriteString(" where `id`=? and `status`=?")
	args = append(args, 一.ID, from.Enum())

	res, err := tx.ExecContext(ctx, q.String(), args...)
	if err != nil {
		return 0, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, errors.Wrap(shift.ErrRowCount, "submitting", j.KV("count", n))
	}

	return 一.ID, nil
}

// Update updates the status of a player_rounds table entity. All the fields of the
// submitted receiver are updated, as well as status and updated_at. 
// The entity id is returned on success or an error.
func (一 submitted) Update(ctx context.Context, tx *sql.Tx,from shift.Status, 
	to shift.Status) (int64, error) {
	var (
		q    strings.Builder
		args []interface{}
	)

	q.WriteString("update player_rounds set `status`=?, `updated_at`=? ")
	args = append(args, to.Enum(), time.Now())

	q.WriteString(" where `id`=? and `status`=?")
	args = append(args, 一.ID, from.Enum())

	res, err := tx.ExecContext(ctx, q.String(), args...)
	if err != nil {
		return 0, err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if n != 1 {
		return 0, errors.Wrap(shift.ErrRowCount, "submitted", j.KV("count", n))
	}

	return 一.ID, nil
}
