// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Loan is an object representing the database table.
type Loan struct {
	ID         string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	BookID     string    `boil:"book_id" json:"book_id" toml:"book_id" yaml:"book_id"`
	UserID     string    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	LoanDate   time.Time `boil:"loan_date" json:"loan_date" toml:"loan_date" yaml:"loan_date"`
	DueDate    time.Time `boil:"due_date" json:"due_date" toml:"due_date" yaml:"due_date"`
	ReturnDate time.Time `boil:"return_date" json:"return_date" toml:"return_date" yaml:"return_date"`
	CreatedAt  null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt  null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *loanR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L loanL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LoanColumns = struct {
	ID         string
	BookID     string
	UserID     string
	LoanDate   string
	DueDate    string
	ReturnDate string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "id",
	BookID:     "book_id",
	UserID:     "user_id",
	LoanDate:   "loan_date",
	DueDate:    "due_date",
	ReturnDate: "return_date",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

var LoanTableColumns = struct {
	ID         string
	BookID     string
	UserID     string
	LoanDate   string
	DueDate    string
	ReturnDate string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "loans.id",
	BookID:     "loans.book_id",
	UserID:     "loans.user_id",
	LoanDate:   "loans.loan_date",
	DueDate:    "loans.due_date",
	ReturnDate: "loans.return_date",
	CreatedAt:  "loans.created_at",
	UpdatedAt:  "loans.updated_at",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var LoanWhere = struct {
	ID         whereHelperstring
	BookID     whereHelperstring
	UserID     whereHelperstring
	LoanDate   whereHelpertime_Time
	DueDate    whereHelpertime_Time
	ReturnDate whereHelpertime_Time
	CreatedAt  whereHelpernull_Time
	UpdatedAt  whereHelpernull_Time
}{
	ID:         whereHelperstring{field: "`loans`.`id`"},
	BookID:     whereHelperstring{field: "`loans`.`book_id`"},
	UserID:     whereHelperstring{field: "`loans`.`user_id`"},
	LoanDate:   whereHelpertime_Time{field: "`loans`.`loan_date`"},
	DueDate:    whereHelpertime_Time{field: "`loans`.`due_date`"},
	ReturnDate: whereHelpertime_Time{field: "`loans`.`return_date`"},
	CreatedAt:  whereHelpernull_Time{field: "`loans`.`created_at`"},
	UpdatedAt:  whereHelpernull_Time{field: "`loans`.`updated_at`"},
}

// LoanRels is where relationship names are stored.
var LoanRels = struct {
	Book string
	User string
}{
	Book: "Book",
	User: "User",
}

// loanR is where relationships are stored.
type loanR struct {
	Book *Book `boil:"Book" json:"Book" toml:"Book" yaml:"Book"`
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*loanR) NewStruct() *loanR {
	return &loanR{}
}

// loanL is where Load methods for each relationship are stored.
type loanL struct{}

var (
	loanAllColumns            = []string{"id", "book_id", "user_id", "loan_date", "due_date", "return_date", "created_at", "updated_at"}
	loanColumnsWithoutDefault = []string{"id", "book_id", "user_id", "loan_date", "due_date", "return_date"}
	loanColumnsWithDefault    = []string{"created_at", "updated_at"}
	loanPrimaryKeyColumns     = []string{"id"}
)

type (
	// LoanSlice is an alias for a slice of pointers to Loan.
	// This should almost always be used instead of []Loan.
	LoanSlice []*Loan
	// LoanHook is the signature for custom Loan hook methods
	LoanHook func(context.Context, boil.ContextExecutor, *Loan) error

	loanQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	loanType                 = reflect.TypeOf(&Loan{})
	loanMapping              = queries.MakeStructMapping(loanType)
	loanPrimaryKeyMapping, _ = queries.BindMapping(loanType, loanMapping, loanPrimaryKeyColumns)
	loanInsertCacheMut       sync.RWMutex
	loanInsertCache          = make(map[string]insertCache)
	loanUpdateCacheMut       sync.RWMutex
	loanUpdateCache          = make(map[string]updateCache)
	loanUpsertCacheMut       sync.RWMutex
	loanUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var loanBeforeInsertHooks []LoanHook
var loanBeforeUpdateHooks []LoanHook
var loanBeforeDeleteHooks []LoanHook
var loanBeforeUpsertHooks []LoanHook

var loanAfterInsertHooks []LoanHook
var loanAfterSelectHooks []LoanHook
var loanAfterUpdateHooks []LoanHook
var loanAfterDeleteHooks []LoanHook
var loanAfterUpsertHooks []LoanHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Loan) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loanBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Loan) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loanBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Loan) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loanBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Loan) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loanBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Loan) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loanAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Loan) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loanAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Loan) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loanAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Loan) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loanAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Loan) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loanAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLoanHook registers your hook function for all future operations.
func AddLoanHook(hookPoint boil.HookPoint, loanHook LoanHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		loanBeforeInsertHooks = append(loanBeforeInsertHooks, loanHook)
	case boil.BeforeUpdateHook:
		loanBeforeUpdateHooks = append(loanBeforeUpdateHooks, loanHook)
	case boil.BeforeDeleteHook:
		loanBeforeDeleteHooks = append(loanBeforeDeleteHooks, loanHook)
	case boil.BeforeUpsertHook:
		loanBeforeUpsertHooks = append(loanBeforeUpsertHooks, loanHook)
	case boil.AfterInsertHook:
		loanAfterInsertHooks = append(loanAfterInsertHooks, loanHook)
	case boil.AfterSelectHook:
		loanAfterSelectHooks = append(loanAfterSelectHooks, loanHook)
	case boil.AfterUpdateHook:
		loanAfterUpdateHooks = append(loanAfterUpdateHooks, loanHook)
	case boil.AfterDeleteHook:
		loanAfterDeleteHooks = append(loanAfterDeleteHooks, loanHook)
	case boil.AfterUpsertHook:
		loanAfterUpsertHooks = append(loanAfterUpsertHooks, loanHook)
	}
}

// One returns a single loan record from the query.
func (q loanQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Loan, error) {
	o := &Loan{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for loans")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Loan records from the query.
func (q loanQuery) All(ctx context.Context, exec boil.ContextExecutor) (LoanSlice, error) {
	var o []*Loan

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Loan slice")
	}

	if len(loanAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Loan records in the query.
func (q loanQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count loans rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q loanQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if loans exists")
	}

	return count > 0, nil
}

// Book pointed to by the foreign key.
func (o *Loan) Book(mods ...qm.QueryMod) bookQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.BookID),
	}

	queryMods = append(queryMods, mods...)

	query := Books(queryMods...)
	queries.SetFrom(query.Query, "`books`")

	return query
}

// User pointed to by the foreign key.
func (o *Loan) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "`users`")

	return query
}

// LoadBook allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (loanL) LoadBook(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLoan interface{}, mods queries.Applicator) error {
	var slice []*Loan
	var object *Loan

	if singular {
		object = maybeLoan.(*Loan)
	} else {
		slice = *maybeLoan.(*[]*Loan)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &loanR{}
		}
		args = append(args, object.BookID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &loanR{}
			}

			for _, a := range args {
				if a == obj.BookID {
					continue Outer
				}
			}

			args = append(args, obj.BookID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`books`),
		qm.WhereIn(`books.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Book")
	}

	var resultSlice []*Book
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Book")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for books")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for books")
	}

	if len(loanAfterSelectHooks) != 0 {
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
		object.R.Book = foreign
		if foreign.R == nil {
			foreign.R = &bookR{}
		}
		foreign.R.Loans = append(foreign.R.Loans, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.BookID == foreign.ID {
				local.R.Book = foreign
				if foreign.R == nil {
					foreign.R = &bookR{}
				}
				foreign.R.Loans = append(foreign.R.Loans, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (loanL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLoan interface{}, mods queries.Applicator) error {
	var slice []*Loan
	var object *Loan

	if singular {
		object = maybeLoan.(*Loan)
	} else {
		slice = *maybeLoan.(*[]*Loan)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &loanR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &loanR{}
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

	if len(loanAfterSelectHooks) != 0 {
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
		foreign.R.Loans = append(foreign.R.Loans, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Loans = append(foreign.R.Loans, local)
				break
			}
		}
	}

	return nil
}

// SetBook of the loan to the related item.
// Sets o.R.Book to related.
// Adds o to related.R.Loans.
func (o *Loan) SetBook(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Book) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `loans` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"book_id"}),
		strmangle.WhereClause("`", "`", 0, loanPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.BookID = related.ID
	if o.R == nil {
		o.R = &loanR{
			Book: related,
		}
	} else {
		o.R.Book = related
	}

	if related.R == nil {
		related.R = &bookR{
			Loans: LoanSlice{o},
		}
	} else {
		related.R.Loans = append(related.R.Loans, o)
	}

	return nil
}

// SetUser of the loan to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Loans.
func (o *Loan) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `loans` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, loanPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

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
		o.R = &loanR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Loans: LoanSlice{o},
		}
	} else {
		related.R.Loans = append(related.R.Loans, o)
	}

	return nil
}

// Loans retrieves all the records using an executor.
func Loans(mods ...qm.QueryMod) loanQuery {
	mods = append(mods, qm.From("`loans`"))
	return loanQuery{NewQuery(mods...)}
}

// FindLoan retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLoan(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Loan, error) {
	loanObj := &Loan{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `loans` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, loanObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from loans")
	}

	if err = loanObj.doAfterSelectHooks(ctx, exec); err != nil {
		return loanObj, err
	}

	return loanObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Loan) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no loans provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(loanColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	loanInsertCacheMut.RLock()
	cache, cached := loanInsertCache[key]
	loanInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			loanAllColumns,
			loanColumnsWithDefault,
			loanColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(loanType, loanMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(loanType, loanMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `loans` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `loans` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `loans` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, loanPrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into loans")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for loans")
	}

CacheNoHooks:
	if !cached {
		loanInsertCacheMut.Lock()
		loanInsertCache[key] = cache
		loanInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Loan.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Loan) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	loanUpdateCacheMut.RLock()
	cache, cached := loanUpdateCache[key]
	loanUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			loanAllColumns,
			loanPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update loans, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `loans` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, loanPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(loanType, loanMapping, append(wl, loanPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update loans row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for loans")
	}

	if !cached {
		loanUpdateCacheMut.Lock()
		loanUpdateCache[key] = cache
		loanUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q loanQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for loans")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for loans")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LoanSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), loanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `loans` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, loanPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in loan slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all loan")
	}
	return rowsAff, nil
}

var mySQLLoanUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Loan) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no loans provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(loanColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLLoanUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	loanUpsertCacheMut.RLock()
	cache, cached := loanUpsertCache[key]
	loanUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			loanAllColumns,
			loanColumnsWithDefault,
			loanColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			loanAllColumns,
			loanPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert loans, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`loans`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `loans` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(loanType, loanMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(loanType, loanMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for loans")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(loanType, loanMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for loans")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for loans")
	}

CacheNoHooks:
	if !cached {
		loanUpsertCacheMut.Lock()
		loanUpsertCache[key] = cache
		loanUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Loan record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Loan) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Loan provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), loanPrimaryKeyMapping)
	sql := "DELETE FROM `loans` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from loans")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for loans")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q loanQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no loanQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from loans")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for loans")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LoanSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(loanBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), loanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `loans` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, loanPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from loan slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for loans")
	}

	if len(loanAfterDeleteHooks) != 0 {
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
func (o *Loan) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLoan(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LoanSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LoanSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), loanPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `loans`.* FROM `loans` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, loanPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LoanSlice")
	}

	*o = slice

	return nil
}

// LoanExists checks if the Loan row exists.
func LoanExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `loans` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if loans exists")
	}

	return exists, nil
}
