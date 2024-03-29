// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package tpportal

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// UserTestDate is an object representing the database table.
type UserTestDate struct {
	UserID        int64 `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	TestDateID    int64 `boil:"test_date_id" json:"test_date_id" toml:"test_date_id" yaml:"test_date_id"`
	EducationYear int16 `boil:"education_year" json:"education_year" toml:"education_year" yaml:"education_year"`
	IsAttended    bool  `boil:"is_attended" json:"is_attended" toml:"is_attended" yaml:"is_attended"`

	R *userTestDateR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userTestDateL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserTestDateColumns = struct {
	UserID        string
	TestDateID    string
	EducationYear string
	IsAttended    string
}{
	UserID:        "user_id",
	TestDateID:    "test_date_id",
	EducationYear: "education_year",
	IsAttended:    "is_attended",
}

var UserTestDateTableColumns = struct {
	UserID        string
	TestDateID    string
	EducationYear string
	IsAttended    string
}{
	UserID:        "user_test_dates.user_id",
	TestDateID:    "user_test_dates.test_date_id",
	EducationYear: "user_test_dates.education_year",
	IsAttended:    "user_test_dates.is_attended",
}

// Generated where

var UserTestDateWhere = struct {
	UserID        whereHelperint64
	TestDateID    whereHelperint64
	EducationYear whereHelperint16
	IsAttended    whereHelperbool
}{
	UserID:        whereHelperint64{field: "\"user_test_dates\".\"user_id\""},
	TestDateID:    whereHelperint64{field: "\"user_test_dates\".\"test_date_id\""},
	EducationYear: whereHelperint16{field: "\"user_test_dates\".\"education_year\""},
	IsAttended:    whereHelperbool{field: "\"user_test_dates\".\"is_attended\""},
}

// UserTestDateRels is where relationship names are stored.
var UserTestDateRels = struct {
	TestDate string
	User     string
}{
	TestDate: "TestDate",
	User:     "User",
}

// userTestDateR is where relationships are stored.
type userTestDateR struct {
	TestDate *TestDate `boil:"TestDate" json:"TestDate" toml:"TestDate" yaml:"TestDate"`
	User     *User     `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*userTestDateR) NewStruct() *userTestDateR {
	return &userTestDateR{}
}

func (r *userTestDateR) GetTestDate() *TestDate {
	if r == nil {
		return nil
	}
	return r.TestDate
}

func (r *userTestDateR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// userTestDateL is where Load methods for each relationship are stored.
type userTestDateL struct{}

var (
	userTestDateAllColumns            = []string{"user_id", "test_date_id", "education_year", "is_attended"}
	userTestDateColumnsWithoutDefault = []string{"user_id", "test_date_id", "education_year"}
	userTestDateColumnsWithDefault    = []string{"is_attended"}
	userTestDatePrimaryKeyColumns     = []string{"user_id", "test_date_id"}
	userTestDateGeneratedColumns      = []string{}
)

type (
	// UserTestDateSlice is an alias for a slice of pointers to UserTestDate.
	// This should almost always be used instead of []UserTestDate.
	UserTestDateSlice []*UserTestDate
	// UserTestDateHook is the signature for custom UserTestDate hook methods
	UserTestDateHook func(context.Context, boil.ContextExecutor, *UserTestDate) error

	userTestDateQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userTestDateType                 = reflect.TypeOf(&UserTestDate{})
	userTestDateMapping              = queries.MakeStructMapping(userTestDateType)
	userTestDatePrimaryKeyMapping, _ = queries.BindMapping(userTestDateType, userTestDateMapping, userTestDatePrimaryKeyColumns)
	userTestDateInsertCacheMut       sync.RWMutex
	userTestDateInsertCache          = make(map[string]insertCache)
	userTestDateUpdateCacheMut       sync.RWMutex
	userTestDateUpdateCache          = make(map[string]updateCache)
	userTestDateUpsertCacheMut       sync.RWMutex
	userTestDateUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userTestDateAfterSelectHooks []UserTestDateHook

var userTestDateBeforeInsertHooks []UserTestDateHook
var userTestDateAfterInsertHooks []UserTestDateHook

var userTestDateBeforeUpdateHooks []UserTestDateHook
var userTestDateAfterUpdateHooks []UserTestDateHook

var userTestDateBeforeDeleteHooks []UserTestDateHook
var userTestDateAfterDeleteHooks []UserTestDateHook

var userTestDateBeforeUpsertHooks []UserTestDateHook
var userTestDateAfterUpsertHooks []UserTestDateHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserTestDate) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userTestDateAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserTestDate) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userTestDateBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserTestDate) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userTestDateAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserTestDate) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userTestDateBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserTestDate) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userTestDateAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserTestDate) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userTestDateBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserTestDate) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userTestDateAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserTestDate) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userTestDateBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserTestDate) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userTestDateAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserTestDateHook registers your hook function for all future operations.
func AddUserTestDateHook(hookPoint boil.HookPoint, userTestDateHook UserTestDateHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userTestDateAfterSelectHooks = append(userTestDateAfterSelectHooks, userTestDateHook)
	case boil.BeforeInsertHook:
		userTestDateBeforeInsertHooks = append(userTestDateBeforeInsertHooks, userTestDateHook)
	case boil.AfterInsertHook:
		userTestDateAfterInsertHooks = append(userTestDateAfterInsertHooks, userTestDateHook)
	case boil.BeforeUpdateHook:
		userTestDateBeforeUpdateHooks = append(userTestDateBeforeUpdateHooks, userTestDateHook)
	case boil.AfterUpdateHook:
		userTestDateAfterUpdateHooks = append(userTestDateAfterUpdateHooks, userTestDateHook)
	case boil.BeforeDeleteHook:
		userTestDateBeforeDeleteHooks = append(userTestDateBeforeDeleteHooks, userTestDateHook)
	case boil.AfterDeleteHook:
		userTestDateAfterDeleteHooks = append(userTestDateAfterDeleteHooks, userTestDateHook)
	case boil.BeforeUpsertHook:
		userTestDateBeforeUpsertHooks = append(userTestDateBeforeUpsertHooks, userTestDateHook)
	case boil.AfterUpsertHook:
		userTestDateAfterUpsertHooks = append(userTestDateAfterUpsertHooks, userTestDateHook)
	}
}

// One returns a single userTestDate record from the query.
func (q userTestDateQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserTestDate, error) {
	o := &UserTestDate{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "tpportal: failed to execute a one query for user_test_dates")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserTestDate records from the query.
func (q userTestDateQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserTestDateSlice, error) {
	var o []*UserTestDate

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "tpportal: failed to assign all query results to UserTestDate slice")
	}

	if len(userTestDateAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserTestDate records in the query.
func (q userTestDateQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "tpportal: failed to count user_test_dates rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userTestDateQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "tpportal: failed to check if user_test_dates exists")
	}

	return count > 0, nil
}

// TestDate pointed to by the foreign key.
func (o *UserTestDate) TestDate(mods ...qm.QueryMod) testDateQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.TestDateID),
	}

	queryMods = append(queryMods, mods...)

	return TestDates(queryMods...)
}

// User pointed to by the foreign key.
func (o *UserTestDate) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadTestDate allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userTestDateL) LoadTestDate(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserTestDate interface{}, mods queries.Applicator) error {
	var slice []*UserTestDate
	var object *UserTestDate

	if singular {
		var ok bool
		object, ok = maybeUserTestDate.(*UserTestDate)
		if !ok {
			object = new(UserTestDate)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserTestDate)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserTestDate))
			}
		}
	} else {
		s, ok := maybeUserTestDate.(*[]*UserTestDate)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserTestDate)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserTestDate))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userTestDateR{}
		}
		args = append(args, object.TestDateID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userTestDateR{}
			}

			for _, a := range args {
				if a == obj.TestDateID {
					continue Outer
				}
			}

			args = append(args, obj.TestDateID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`test_dates`),
		qm.WhereIn(`test_dates.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load TestDate")
	}

	var resultSlice []*TestDate
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice TestDate")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for test_dates")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for test_dates")
	}

	if len(testDateAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.TestDate = foreign
		if foreign.R == nil {
			foreign.R = &testDateR{}
		}
		foreign.R.UserTestDates = append(foreign.R.UserTestDates, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TestDateID == foreign.ID {
				local.R.TestDate = foreign
				if foreign.R == nil {
					foreign.R = &testDateR{}
				}
				foreign.R.UserTestDates = append(foreign.R.UserTestDates, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (userTestDateL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeUserTestDate interface{}, mods queries.Applicator) error {
	var slice []*UserTestDate
	var object *UserTestDate

	if singular {
		var ok bool
		object, ok = maybeUserTestDate.(*UserTestDate)
		if !ok {
			object = new(UserTestDate)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeUserTestDate)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeUserTestDate))
			}
		}
	} else {
		s, ok := maybeUserTestDate.(*[]*UserTestDate)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeUserTestDate)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeUserTestDate))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &userTestDateR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &userTestDateR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.UserTestDates = append(foreign.R.UserTestDates, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.UserTestDates = append(foreign.R.UserTestDates, local)
				break
			}
		}
	}

	return nil
}

// SetTestDate of the userTestDate to the related item.
// Sets o.R.TestDate to related.
// Adds o to related.R.UserTestDates.
func (o *UserTestDate) SetTestDate(ctx context.Context, exec boil.ContextExecutor, insert bool, related *TestDate) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_test_dates\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"test_date_id"}),
		strmangle.WhereClause("\"", "\"", 2, userTestDatePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.TestDateID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TestDateID = related.ID
	if o.R == nil {
		o.R = &userTestDateR{
			TestDate: related,
		}
	} else {
		o.R.TestDate = related
	}

	if related.R == nil {
		related.R = &testDateR{
			UserTestDates: UserTestDateSlice{o},
		}
	} else {
		related.R.UserTestDates = append(related.R.UserTestDates, o)
	}

	return nil
}

// SetUser of the userTestDate to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserTestDates.
func (o *UserTestDate) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"user_test_dates\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, userTestDatePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.UserID, o.TestDateID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &userTestDateR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			UserTestDates: UserTestDateSlice{o},
		}
	} else {
		related.R.UserTestDates = append(related.R.UserTestDates, o)
	}

	return nil
}

// UserTestDates retrieves all the records using an executor.
func UserTestDates(mods ...qm.QueryMod) userTestDateQuery {
	mods = append(mods, qm.From("\"user_test_dates\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"user_test_dates\".*"})
	}

	return userTestDateQuery{q}
}

// FindUserTestDate retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserTestDate(ctx context.Context, exec boil.ContextExecutor, userID int64, testDateID int64, selectCols ...string) (*UserTestDate, error) {
	userTestDateObj := &UserTestDate{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"user_test_dates\" where \"user_id\"=$1 AND \"test_date_id\"=$2", sel,
	)

	q := queries.Raw(query, userID, testDateID)

	err := q.Bind(ctx, exec, userTestDateObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "tpportal: unable to select from user_test_dates")
	}

	if err = userTestDateObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userTestDateObj, err
	}

	return userTestDateObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserTestDate) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("tpportal: no user_test_dates provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userTestDateColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userTestDateInsertCacheMut.RLock()
	cache, cached := userTestDateInsertCache[key]
	userTestDateInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userTestDateAllColumns,
			userTestDateColumnsWithDefault,
			userTestDateColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userTestDateType, userTestDateMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userTestDateType, userTestDateMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"user_test_dates\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"user_test_dates\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "tpportal: unable to insert into user_test_dates")
	}

	if !cached {
		userTestDateInsertCacheMut.Lock()
		userTestDateInsertCache[key] = cache
		userTestDateInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserTestDate.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserTestDate) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userTestDateUpdateCacheMut.RLock()
	cache, cached := userTestDateUpdateCache[key]
	userTestDateUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userTestDateAllColumns,
			userTestDatePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("tpportal: unable to update user_test_dates, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"user_test_dates\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, userTestDatePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userTestDateType, userTestDateMapping, append(wl, userTestDatePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "tpportal: unable to update user_test_dates row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "tpportal: failed to get rows affected by update for user_test_dates")
	}

	if !cached {
		userTestDateUpdateCacheMut.Lock()
		userTestDateUpdateCache[key] = cache
		userTestDateUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userTestDateQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "tpportal: unable to update all for user_test_dates")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "tpportal: unable to retrieve rows affected for user_test_dates")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserTestDateSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("tpportal: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userTestDatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"user_test_dates\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, userTestDatePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "tpportal: unable to update all in userTestDate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "tpportal: unable to retrieve rows affected all in update all userTestDate")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserTestDate) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("tpportal: no user_test_dates provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userTestDateColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	userTestDateUpsertCacheMut.RLock()
	cache, cached := userTestDateUpsertCache[key]
	userTestDateUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userTestDateAllColumns,
			userTestDateColumnsWithDefault,
			userTestDateColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userTestDateAllColumns,
			userTestDatePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("tpportal: unable to upsert user_test_dates, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(userTestDatePrimaryKeyColumns))
			copy(conflict, userTestDatePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"user_test_dates\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(userTestDateType, userTestDateMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userTestDateType, userTestDateMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "tpportal: unable to upsert user_test_dates")
	}

	if !cached {
		userTestDateUpsertCacheMut.Lock()
		userTestDateUpsertCache[key] = cache
		userTestDateUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserTestDate record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserTestDate) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("tpportal: no UserTestDate provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userTestDatePrimaryKeyMapping)
	sql := "DELETE FROM \"user_test_dates\" WHERE \"user_id\"=$1 AND \"test_date_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "tpportal: unable to delete from user_test_dates")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "tpportal: failed to get rows affected by delete for user_test_dates")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userTestDateQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("tpportal: no userTestDateQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "tpportal: unable to delete all from user_test_dates")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "tpportal: failed to get rows affected by deleteall for user_test_dates")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserTestDateSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userTestDateBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userTestDatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"user_test_dates\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userTestDatePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "tpportal: unable to delete all from userTestDate slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "tpportal: failed to get rows affected by deleteall for user_test_dates")
	}

	if len(userTestDateAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *UserTestDate) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserTestDate(ctx, exec, o.UserID, o.TestDateID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserTestDateSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserTestDateSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userTestDatePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"user_test_dates\".* FROM \"user_test_dates\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, userTestDatePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "tpportal: unable to reload all in UserTestDateSlice")
	}

	*o = slice

	return nil
}

// UserTestDateExists checks if the UserTestDate row exists.
func UserTestDateExists(ctx context.Context, exec boil.ContextExecutor, userID int64, testDateID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"user_test_dates\" where \"user_id\"=$1 AND \"test_date_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, userID, testDateID)
	}
	row := exec.QueryRowContext(ctx, sql, userID, testDateID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "tpportal: unable to check if user_test_dates exists")
	}

	return exists, nil
}

// Exists checks if the UserTestDate row exists.
func (o *UserTestDate) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return UserTestDateExists(ctx, exec, o.UserID, o.TestDateID)
}
