// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// AutomodViolation is an object representing the database table.
type AutomodViolation struct {
	ID        int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID   int64      `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	UserID    int64      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	RuleID    null.Int64 `boil:"rule_id" json:"rule_id,omitempty" toml:"rule_id" yaml:"rule_id,omitempty"`
	CreatedAt time.Time  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	Name      string     `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *automodViolationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L automodViolationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AutomodViolationColumns = struct {
	ID        string
	GuildID   string
	UserID    string
	RuleID    string
	CreatedAt string
	Name      string
}{
	ID:        "id",
	GuildID:   "guild_id",
	UserID:    "user_id",
	RuleID:    "rule_id",
	CreatedAt: "created_at",
	Name:      "name",
}

// Generated where

var AutomodViolationWhere = struct {
	ID        whereHelperint64
	GuildID   whereHelperint64
	UserID    whereHelperint64
	RuleID    whereHelpernull_Int64
	CreatedAt whereHelpertime_Time
	Name      whereHelperstring
}{
	ID:        whereHelperint64{field: "\"automod_violations\".\"id\""},
	GuildID:   whereHelperint64{field: "\"automod_violations\".\"guild_id\""},
	UserID:    whereHelperint64{field: "\"automod_violations\".\"user_id\""},
	RuleID:    whereHelpernull_Int64{field: "\"automod_violations\".\"rule_id\""},
	CreatedAt: whereHelpertime_Time{field: "\"automod_violations\".\"created_at\""},
	Name:      whereHelperstring{field: "\"automod_violations\".\"name\""},
}

// AutomodViolationRels is where relationship names are stored.
var AutomodViolationRels = struct {
	Rule string
}{
	Rule: "Rule",
}

// automodViolationR is where relationships are stored.
type automodViolationR struct {
	Rule *AutomodRule
}

// NewStruct creates a new relationship struct
func (*automodViolationR) NewStruct() *automodViolationR {
	return &automodViolationR{}
}

// automodViolationL is where Load methods for each relationship are stored.
type automodViolationL struct{}

var (
	automodViolationAllColumns            = []string{"id", "guild_id", "user_id", "rule_id", "created_at", "name"}
	automodViolationColumnsWithoutDefault = []string{"guild_id", "user_id", "rule_id", "created_at", "name"}
	automodViolationColumnsWithDefault    = []string{"id"}
	automodViolationPrimaryKeyColumns     = []string{"id"}
)

type (
	// AutomodViolationSlice is an alias for a slice of pointers to AutomodViolation.
	// This should generally be used opposed to []AutomodViolation.
	AutomodViolationSlice []*AutomodViolation

	automodViolationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	automodViolationType                 = reflect.TypeOf(&AutomodViolation{})
	automodViolationMapping              = queries.MakeStructMapping(automodViolationType)
	automodViolationPrimaryKeyMapping, _ = queries.BindMapping(automodViolationType, automodViolationMapping, automodViolationPrimaryKeyColumns)
	automodViolationInsertCacheMut       sync.RWMutex
	automodViolationInsertCache          = make(map[string]insertCache)
	automodViolationUpdateCacheMut       sync.RWMutex
	automodViolationUpdateCache          = make(map[string]updateCache)
	automodViolationUpsertCacheMut       sync.RWMutex
	automodViolationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single automodViolation record from the query using the global executor.
func (q automodViolationQuery) OneG(ctx context.Context) (*AutomodViolation, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single automodViolation record from the query.
func (q automodViolationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AutomodViolation, error) {
	o := &AutomodViolation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for automod_violations")
	}

	return o, nil
}

// AllG returns all AutomodViolation records from the query using the global executor.
func (q automodViolationQuery) AllG(ctx context.Context) (AutomodViolationSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all AutomodViolation records from the query.
func (q automodViolationQuery) All(ctx context.Context, exec boil.ContextExecutor) (AutomodViolationSlice, error) {
	var o []*AutomodViolation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AutomodViolation slice")
	}

	return o, nil
}

// CountG returns the count of all AutomodViolation records in the query, and panics on error.
func (q automodViolationQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all AutomodViolation records in the query.
func (q automodViolationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count automod_violations rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q automodViolationQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q automodViolationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if automod_violations exists")
	}

	return count > 0, nil
}

// Rule pointed to by the foreign key.
func (o *AutomodViolation) Rule(mods ...qm.QueryMod) automodRuleQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.RuleID),
	}

	queryMods = append(queryMods, mods...)

	query := AutomodRules(queryMods...)
	queries.SetFrom(query.Query, "\"automod_rules\"")

	return query
}

// LoadRule allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (automodViolationL) LoadRule(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAutomodViolation interface{}, mods queries.Applicator) error {
	var slice []*AutomodViolation
	var object *AutomodViolation

	if singular {
		object = maybeAutomodViolation.(*AutomodViolation)
	} else {
		slice = *maybeAutomodViolation.(*[]*AutomodViolation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &automodViolationR{}
		}
		if !queries.IsNil(object.RuleID) {
			args = append(args, object.RuleID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &automodViolationR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.RuleID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.RuleID) {
				args = append(args, obj.RuleID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`automod_rules`), qm.WhereIn(`automod_rules.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AutomodRule")
	}

	var resultSlice []*AutomodRule
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AutomodRule")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for automod_rules")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for automod_rules")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Rule = foreign
		if foreign.R == nil {
			foreign.R = &automodRuleR{}
		}
		foreign.R.RuleAutomodViolations = append(foreign.R.RuleAutomodViolations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.RuleID, foreign.ID) {
				local.R.Rule = foreign
				if foreign.R == nil {
					foreign.R = &automodRuleR{}
				}
				foreign.R.RuleAutomodViolations = append(foreign.R.RuleAutomodViolations, local)
				break
			}
		}
	}

	return nil
}

// SetRuleG of the automodViolation to the related item.
// Sets o.R.Rule to related.
// Adds o to related.R.RuleAutomodViolations.
// Uses the global database handle.
func (o *AutomodViolation) SetRuleG(ctx context.Context, insert bool, related *AutomodRule) error {
	return o.SetRule(ctx, boil.GetContextDB(), insert, related)
}

// SetRule of the automodViolation to the related item.
// Sets o.R.Rule to related.
// Adds o to related.R.RuleAutomodViolations.
func (o *AutomodViolation) SetRule(ctx context.Context, exec boil.ContextExecutor, insert bool, related *AutomodRule) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"automod_violations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"rule_id"}),
		strmangle.WhereClause("\"", "\"", 2, automodViolationPrimaryKeyColumns),
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

	queries.Assign(&o.RuleID, related.ID)
	if o.R == nil {
		o.R = &automodViolationR{
			Rule: related,
		}
	} else {
		o.R.Rule = related
	}

	if related.R == nil {
		related.R = &automodRuleR{
			RuleAutomodViolations: AutomodViolationSlice{o},
		}
	} else {
		related.R.RuleAutomodViolations = append(related.R.RuleAutomodViolations, o)
	}

	return nil
}

// RemoveRuleG relationship.
// Sets o.R.Rule to nil.
// Removes o from all passed in related items' relationships struct (Optional).
// Uses the global database handle.
func (o *AutomodViolation) RemoveRuleG(ctx context.Context, related *AutomodRule) error {
	return o.RemoveRule(ctx, boil.GetContextDB(), related)
}

// RemoveRule relationship.
// Sets o.R.Rule to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *AutomodViolation) RemoveRule(ctx context.Context, exec boil.ContextExecutor, related *AutomodRule) error {
	var err error

	queries.SetScanner(&o.RuleID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("rule_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Rule = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.RuleAutomodViolations {
		if queries.Equal(o.RuleID, ri.RuleID) {
			continue
		}

		ln := len(related.R.RuleAutomodViolations)
		if ln > 1 && i < ln-1 {
			related.R.RuleAutomodViolations[i] = related.R.RuleAutomodViolations[ln-1]
		}
		related.R.RuleAutomodViolations = related.R.RuleAutomodViolations[:ln-1]
		break
	}
	return nil
}

// AutomodViolations retrieves all the records using an executor.
func AutomodViolations(mods ...qm.QueryMod) automodViolationQuery {
	mods = append(mods, qm.From("\"automod_violations\""))
	return automodViolationQuery{NewQuery(mods...)}
}

// FindAutomodViolationG retrieves a single record by ID.
func FindAutomodViolationG(ctx context.Context, iD int64, selectCols ...string) (*AutomodViolation, error) {
	return FindAutomodViolation(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindAutomodViolation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAutomodViolation(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*AutomodViolation, error) {
	automodViolationObj := &AutomodViolation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"automod_violations\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, automodViolationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from automod_violations")
	}

	return automodViolationObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AutomodViolation) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AutomodViolation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no automod_violations provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(automodViolationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	automodViolationInsertCacheMut.RLock()
	cache, cached := automodViolationInsertCache[key]
	automodViolationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			automodViolationAllColumns,
			automodViolationColumnsWithDefault,
			automodViolationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(automodViolationType, automodViolationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(automodViolationType, automodViolationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"automod_violations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"automod_violations\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into automod_violations")
	}

	if !cached {
		automodViolationInsertCacheMut.Lock()
		automodViolationInsertCache[key] = cache
		automodViolationInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single AutomodViolation record using the global executor.
// See Update for more documentation.
func (o *AutomodViolation) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the AutomodViolation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AutomodViolation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	automodViolationUpdateCacheMut.RLock()
	cache, cached := automodViolationUpdateCache[key]
	automodViolationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			automodViolationAllColumns,
			automodViolationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update automod_violations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"automod_violations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, automodViolationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(automodViolationType, automodViolationMapping, append(wl, automodViolationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update automod_violations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for automod_violations")
	}

	if !cached {
		automodViolationUpdateCacheMut.Lock()
		automodViolationUpdateCache[key] = cache
		automodViolationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q automodViolationQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q automodViolationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for automod_violations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for automod_violations")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AutomodViolationSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AutomodViolationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), automodViolationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"automod_violations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, automodViolationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in automodViolation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all automodViolation")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AutomodViolation) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AutomodViolation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no automod_violations provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(automodViolationColumnsWithDefault, o)

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

	automodViolationUpsertCacheMut.RLock()
	cache, cached := automodViolationUpsertCache[key]
	automodViolationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			automodViolationAllColumns,
			automodViolationColumnsWithDefault,
			automodViolationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			automodViolationAllColumns,
			automodViolationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert automod_violations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(automodViolationPrimaryKeyColumns))
			copy(conflict, automodViolationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"automod_violations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(automodViolationType, automodViolationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(automodViolationType, automodViolationMapping, ret)
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
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert automod_violations")
	}

	if !cached {
		automodViolationUpsertCacheMut.Lock()
		automodViolationUpsertCache[key] = cache
		automodViolationUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single AutomodViolation record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AutomodViolation) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single AutomodViolation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AutomodViolation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AutomodViolation provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), automodViolationPrimaryKeyMapping)
	sql := "DELETE FROM \"automod_violations\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from automod_violations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for automod_violations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q automodViolationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no automodViolationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from automod_violations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for automod_violations")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o AutomodViolationSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AutomodViolationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), automodViolationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"automod_violations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, automodViolationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from automodViolation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for automod_violations")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AutomodViolation) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no AutomodViolation provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AutomodViolation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAutomodViolation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AutomodViolationSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty AutomodViolationSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AutomodViolationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AutomodViolationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), automodViolationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"automod_violations\".* FROM \"automod_violations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, automodViolationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AutomodViolationSlice")
	}

	*o = slice

	return nil
}

// AutomodViolationExistsG checks if the AutomodViolation row exists.
func AutomodViolationExistsG(ctx context.Context, iD int64) (bool, error) {
	return AutomodViolationExists(ctx, boil.GetContextDB(), iD)
}

// AutomodViolationExists checks if the AutomodViolation row exists.
func AutomodViolationExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"automod_violations\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if automod_violations exists")
	}

	return exists, nil
}
