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
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// AutomodRuleDatum is an object representing the database table.
type AutomodRuleDatum struct {
	ID       int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID  int64      `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	RuleID   int64      `boil:"rule_id" json:"rule_id" toml:"rule_id" yaml:"rule_id"`
	Kind     int        `boil:"kind" json:"kind" toml:"kind" yaml:"kind"`
	TypeID   int        `boil:"type_id" json:"type_id" toml:"type_id" yaml:"type_id"`
	Settings types.JSON `boil:"settings" json:"settings" toml:"settings" yaml:"settings"`

	R *automodRuleDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L automodRuleDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AutomodRuleDatumColumns = struct {
	ID       string
	GuildID  string
	RuleID   string
	Kind     string
	TypeID   string
	Settings string
}{
	ID:       "id",
	GuildID:  "guild_id",
	RuleID:   "rule_id",
	Kind:     "kind",
	TypeID:   "type_id",
	Settings: "settings",
}

// Generated where

type whereHelpertypes_JSON struct{ field string }

func (w whereHelpertypes_JSON) EQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_JSON) NEQ(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_JSON) LT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_JSON) LTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_JSON) GT(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_JSON) GTE(x types.JSON) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AutomodRuleDatumWhere = struct {
	ID       whereHelperint64
	GuildID  whereHelperint64
	RuleID   whereHelperint64
	Kind     whereHelperint
	TypeID   whereHelperint
	Settings whereHelpertypes_JSON
}{
	ID:       whereHelperint64{field: "\"automod_rule_data\".\"id\""},
	GuildID:  whereHelperint64{field: "\"automod_rule_data\".\"guild_id\""},
	RuleID:   whereHelperint64{field: "\"automod_rule_data\".\"rule_id\""},
	Kind:     whereHelperint{field: "\"automod_rule_data\".\"kind\""},
	TypeID:   whereHelperint{field: "\"automod_rule_data\".\"type_id\""},
	Settings: whereHelpertypes_JSON{field: "\"automod_rule_data\".\"settings\""},
}

// AutomodRuleDatumRels is where relationship names are stored.
var AutomodRuleDatumRels = struct {
	Rule                         string
	TriggerAutomodTriggeredRules string
}{
	Rule:                         "Rule",
	TriggerAutomodTriggeredRules: "TriggerAutomodTriggeredRules",
}

// automodRuleDatumR is where relationships are stored.
type automodRuleDatumR struct {
	Rule                         *AutomodRule
	TriggerAutomodTriggeredRules AutomodTriggeredRuleSlice
}

// NewStruct creates a new relationship struct
func (*automodRuleDatumR) NewStruct() *automodRuleDatumR {
	return &automodRuleDatumR{}
}

// automodRuleDatumL is where Load methods for each relationship are stored.
type automodRuleDatumL struct{}

var (
	automodRuleDatumAllColumns            = []string{"id", "guild_id", "rule_id", "kind", "type_id", "settings"}
	automodRuleDatumColumnsWithoutDefault = []string{"guild_id", "rule_id", "kind", "type_id", "settings"}
	automodRuleDatumColumnsWithDefault    = []string{"id"}
	automodRuleDatumPrimaryKeyColumns     = []string{"id"}
)

type (
	// AutomodRuleDatumSlice is an alias for a slice of pointers to AutomodRuleDatum.
	// This should generally be used opposed to []AutomodRuleDatum.
	AutomodRuleDatumSlice []*AutomodRuleDatum

	automodRuleDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	automodRuleDatumType                 = reflect.TypeOf(&AutomodRuleDatum{})
	automodRuleDatumMapping              = queries.MakeStructMapping(automodRuleDatumType)
	automodRuleDatumPrimaryKeyMapping, _ = queries.BindMapping(automodRuleDatumType, automodRuleDatumMapping, automodRuleDatumPrimaryKeyColumns)
	automodRuleDatumInsertCacheMut       sync.RWMutex
	automodRuleDatumInsertCache          = make(map[string]insertCache)
	automodRuleDatumUpdateCacheMut       sync.RWMutex
	automodRuleDatumUpdateCache          = make(map[string]updateCache)
	automodRuleDatumUpsertCacheMut       sync.RWMutex
	automodRuleDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single automodRuleDatum record from the query using the global executor.
func (q automodRuleDatumQuery) OneG(ctx context.Context) (*AutomodRuleDatum, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single automodRuleDatum record from the query.
func (q automodRuleDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AutomodRuleDatum, error) {
	o := &AutomodRuleDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for automod_rule_data")
	}

	return o, nil
}

// AllG returns all AutomodRuleDatum records from the query using the global executor.
func (q automodRuleDatumQuery) AllG(ctx context.Context) (AutomodRuleDatumSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all AutomodRuleDatum records from the query.
func (q automodRuleDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (AutomodRuleDatumSlice, error) {
	var o []*AutomodRuleDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AutomodRuleDatum slice")
	}

	return o, nil
}

// CountG returns the count of all AutomodRuleDatum records in the query, and panics on error.
func (q automodRuleDatumQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all AutomodRuleDatum records in the query.
func (q automodRuleDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count automod_rule_data rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q automodRuleDatumQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q automodRuleDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if automod_rule_data exists")
	}

	return count > 0, nil
}

// Rule pointed to by the foreign key.
func (o *AutomodRuleDatum) Rule(mods ...qm.QueryMod) automodRuleQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.RuleID),
	}

	queryMods = append(queryMods, mods...)

	query := AutomodRules(queryMods...)
	queries.SetFrom(query.Query, "\"automod_rules\"")

	return query
}

// TriggerAutomodTriggeredRules retrieves all the automod_triggered_rule's AutomodTriggeredRules with an executor via trigger_id column.
func (o *AutomodRuleDatum) TriggerAutomodTriggeredRules(mods ...qm.QueryMod) automodTriggeredRuleQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"automod_triggered_rules\".\"trigger_id\"=?", o.ID),
	)

	query := AutomodTriggeredRules(queryMods...)
	queries.SetFrom(query.Query, "\"automod_triggered_rules\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"automod_triggered_rules\".*"})
	}

	return query
}

// LoadRule allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (automodRuleDatumL) LoadRule(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAutomodRuleDatum interface{}, mods queries.Applicator) error {
	var slice []*AutomodRuleDatum
	var object *AutomodRuleDatum

	if singular {
		object = maybeAutomodRuleDatum.(*AutomodRuleDatum)
	} else {
		slice = *maybeAutomodRuleDatum.(*[]*AutomodRuleDatum)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &automodRuleDatumR{}
		}
		args = append(args, object.RuleID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &automodRuleDatumR{}
			}

			for _, a := range args {
				if a == obj.RuleID {
					continue Outer
				}
			}

			args = append(args, obj.RuleID)

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
		foreign.R.RuleAutomodRuleData = append(foreign.R.RuleAutomodRuleData, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RuleID == foreign.ID {
				local.R.Rule = foreign
				if foreign.R == nil {
					foreign.R = &automodRuleR{}
				}
				foreign.R.RuleAutomodRuleData = append(foreign.R.RuleAutomodRuleData, local)
				break
			}
		}
	}

	return nil
}

// LoadTriggerAutomodTriggeredRules allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (automodRuleDatumL) LoadTriggerAutomodTriggeredRules(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAutomodRuleDatum interface{}, mods queries.Applicator) error {
	var slice []*AutomodRuleDatum
	var object *AutomodRuleDatum

	if singular {
		object = maybeAutomodRuleDatum.(*AutomodRuleDatum)
	} else {
		slice = *maybeAutomodRuleDatum.(*[]*AutomodRuleDatum)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &automodRuleDatumR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &automodRuleDatumR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`automod_triggered_rules`), qm.WhereIn(`automod_triggered_rules.trigger_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load automod_triggered_rules")
	}

	var resultSlice []*AutomodTriggeredRule
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice automod_triggered_rules")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on automod_triggered_rules")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for automod_triggered_rules")
	}

	if singular {
		object.R.TriggerAutomodTriggeredRules = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &automodTriggeredRuleR{}
			}
			foreign.R.Trigger = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.TriggerID) {
				local.R.TriggerAutomodTriggeredRules = append(local.R.TriggerAutomodTriggeredRules, foreign)
				if foreign.R == nil {
					foreign.R = &automodTriggeredRuleR{}
				}
				foreign.R.Trigger = local
				break
			}
		}
	}

	return nil
}

// SetRuleG of the automodRuleDatum to the related item.
// Sets o.R.Rule to related.
// Adds o to related.R.RuleAutomodRuleData.
// Uses the global database handle.
func (o *AutomodRuleDatum) SetRuleG(ctx context.Context, insert bool, related *AutomodRule) error {
	return o.SetRule(ctx, boil.GetContextDB(), insert, related)
}

// SetRule of the automodRuleDatum to the related item.
// Sets o.R.Rule to related.
// Adds o to related.R.RuleAutomodRuleData.
func (o *AutomodRuleDatum) SetRule(ctx context.Context, exec boil.ContextExecutor, insert bool, related *AutomodRule) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"automod_rule_data\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"rule_id"}),
		strmangle.WhereClause("\"", "\"", 2, automodRuleDatumPrimaryKeyColumns),
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

	o.RuleID = related.ID
	if o.R == nil {
		o.R = &automodRuleDatumR{
			Rule: related,
		}
	} else {
		o.R.Rule = related
	}

	if related.R == nil {
		related.R = &automodRuleR{
			RuleAutomodRuleData: AutomodRuleDatumSlice{o},
		}
	} else {
		related.R.RuleAutomodRuleData = append(related.R.RuleAutomodRuleData, o)
	}

	return nil
}

// AddTriggerAutomodTriggeredRulesG adds the given related objects to the existing relationships
// of the automod_rule_datum, optionally inserting them as new records.
// Appends related to o.R.TriggerAutomodTriggeredRules.
// Sets related.R.Trigger appropriately.
// Uses the global database handle.
func (o *AutomodRuleDatum) AddTriggerAutomodTriggeredRulesG(ctx context.Context, insert bool, related ...*AutomodTriggeredRule) error {
	return o.AddTriggerAutomodTriggeredRules(ctx, boil.GetContextDB(), insert, related...)
}

// AddTriggerAutomodTriggeredRules adds the given related objects to the existing relationships
// of the automod_rule_datum, optionally inserting them as new records.
// Appends related to o.R.TriggerAutomodTriggeredRules.
// Sets related.R.Trigger appropriately.
func (o *AutomodRuleDatum) AddTriggerAutomodTriggeredRules(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*AutomodTriggeredRule) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.TriggerID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"automod_triggered_rules\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"trigger_id"}),
				strmangle.WhereClause("\"", "\"", 2, automodTriggeredRulePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.TriggerID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &automodRuleDatumR{
			TriggerAutomodTriggeredRules: related,
		}
	} else {
		o.R.TriggerAutomodTriggeredRules = append(o.R.TriggerAutomodTriggeredRules, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &automodTriggeredRuleR{
				Trigger: o,
			}
		} else {
			rel.R.Trigger = o
		}
	}
	return nil
}

// SetTriggerAutomodTriggeredRulesG removes all previously related items of the
// automod_rule_datum replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Trigger's TriggerAutomodTriggeredRules accordingly.
// Replaces o.R.TriggerAutomodTriggeredRules with related.
// Sets related.R.Trigger's TriggerAutomodTriggeredRules accordingly.
// Uses the global database handle.
func (o *AutomodRuleDatum) SetTriggerAutomodTriggeredRulesG(ctx context.Context, insert bool, related ...*AutomodTriggeredRule) error {
	return o.SetTriggerAutomodTriggeredRules(ctx, boil.GetContextDB(), insert, related...)
}

// SetTriggerAutomodTriggeredRules removes all previously related items of the
// automod_rule_datum replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Trigger's TriggerAutomodTriggeredRules accordingly.
// Replaces o.R.TriggerAutomodTriggeredRules with related.
// Sets related.R.Trigger's TriggerAutomodTriggeredRules accordingly.
func (o *AutomodRuleDatum) SetTriggerAutomodTriggeredRules(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*AutomodTriggeredRule) error {
	query := "update \"automod_triggered_rules\" set \"trigger_id\" = null where \"trigger_id\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.TriggerAutomodTriggeredRules {
			queries.SetScanner(&rel.TriggerID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Trigger = nil
		}

		o.R.TriggerAutomodTriggeredRules = nil
	}
	return o.AddTriggerAutomodTriggeredRules(ctx, exec, insert, related...)
}

// RemoveTriggerAutomodTriggeredRulesG relationships from objects passed in.
// Removes related items from R.TriggerAutomodTriggeredRules (uses pointer comparison, removal does not keep order)
// Sets related.R.Trigger.
// Uses the global database handle.
func (o *AutomodRuleDatum) RemoveTriggerAutomodTriggeredRulesG(ctx context.Context, related ...*AutomodTriggeredRule) error {
	return o.RemoveTriggerAutomodTriggeredRules(ctx, boil.GetContextDB(), related...)
}

// RemoveTriggerAutomodTriggeredRules relationships from objects passed in.
// Removes related items from R.TriggerAutomodTriggeredRules (uses pointer comparison, removal does not keep order)
// Sets related.R.Trigger.
func (o *AutomodRuleDatum) RemoveTriggerAutomodTriggeredRules(ctx context.Context, exec boil.ContextExecutor, related ...*AutomodTriggeredRule) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.TriggerID, nil)
		if rel.R != nil {
			rel.R.Trigger = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("trigger_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.TriggerAutomodTriggeredRules {
			if rel != ri {
				continue
			}

			ln := len(o.R.TriggerAutomodTriggeredRules)
			if ln > 1 && i < ln-1 {
				o.R.TriggerAutomodTriggeredRules[i] = o.R.TriggerAutomodTriggeredRules[ln-1]
			}
			o.R.TriggerAutomodTriggeredRules = o.R.TriggerAutomodTriggeredRules[:ln-1]
			break
		}
	}

	return nil
}

// AutomodRuleData retrieves all the records using an executor.
func AutomodRuleData(mods ...qm.QueryMod) automodRuleDatumQuery {
	mods = append(mods, qm.From("\"automod_rule_data\""))
	return automodRuleDatumQuery{NewQuery(mods...)}
}

// FindAutomodRuleDatumG retrieves a single record by ID.
func FindAutomodRuleDatumG(ctx context.Context, iD int64, selectCols ...string) (*AutomodRuleDatum, error) {
	return FindAutomodRuleDatum(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindAutomodRuleDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAutomodRuleDatum(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*AutomodRuleDatum, error) {
	automodRuleDatumObj := &AutomodRuleDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"automod_rule_data\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, automodRuleDatumObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from automod_rule_data")
	}

	return automodRuleDatumObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *AutomodRuleDatum) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AutomodRuleDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no automod_rule_data provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(automodRuleDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	automodRuleDatumInsertCacheMut.RLock()
	cache, cached := automodRuleDatumInsertCache[key]
	automodRuleDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			automodRuleDatumAllColumns,
			automodRuleDatumColumnsWithDefault,
			automodRuleDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(automodRuleDatumType, automodRuleDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(automodRuleDatumType, automodRuleDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"automod_rule_data\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"automod_rule_data\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into automod_rule_data")
	}

	if !cached {
		automodRuleDatumInsertCacheMut.Lock()
		automodRuleDatumInsertCache[key] = cache
		automodRuleDatumInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single AutomodRuleDatum record using the global executor.
// See Update for more documentation.
func (o *AutomodRuleDatum) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the AutomodRuleDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AutomodRuleDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	automodRuleDatumUpdateCacheMut.RLock()
	cache, cached := automodRuleDatumUpdateCache[key]
	automodRuleDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			automodRuleDatumAllColumns,
			automodRuleDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update automod_rule_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"automod_rule_data\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, automodRuleDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(automodRuleDatumType, automodRuleDatumMapping, append(wl, automodRuleDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update automod_rule_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for automod_rule_data")
	}

	if !cached {
		automodRuleDatumUpdateCacheMut.Lock()
		automodRuleDatumUpdateCache[key] = cache
		automodRuleDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q automodRuleDatumQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q automodRuleDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for automod_rule_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for automod_rule_data")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AutomodRuleDatumSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AutomodRuleDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), automodRuleDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"automod_rule_data\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, automodRuleDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in automodRuleDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all automodRuleDatum")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *AutomodRuleDatum) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AutomodRuleDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no automod_rule_data provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(automodRuleDatumColumnsWithDefault, o)

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

	automodRuleDatumUpsertCacheMut.RLock()
	cache, cached := automodRuleDatumUpsertCache[key]
	automodRuleDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			automodRuleDatumAllColumns,
			automodRuleDatumColumnsWithDefault,
			automodRuleDatumColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			automodRuleDatumAllColumns,
			automodRuleDatumPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert automod_rule_data, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(automodRuleDatumPrimaryKeyColumns))
			copy(conflict, automodRuleDatumPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"automod_rule_data\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(automodRuleDatumType, automodRuleDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(automodRuleDatumType, automodRuleDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert automod_rule_data")
	}

	if !cached {
		automodRuleDatumUpsertCacheMut.Lock()
		automodRuleDatumUpsertCache[key] = cache
		automodRuleDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single AutomodRuleDatum record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *AutomodRuleDatum) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single AutomodRuleDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AutomodRuleDatum) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AutomodRuleDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), automodRuleDatumPrimaryKeyMapping)
	sql := "DELETE FROM \"automod_rule_data\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from automod_rule_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for automod_rule_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q automodRuleDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no automodRuleDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from automod_rule_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for automod_rule_data")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o AutomodRuleDatumSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AutomodRuleDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), automodRuleDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"automod_rule_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, automodRuleDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from automodRuleDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for automod_rule_data")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *AutomodRuleDatum) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no AutomodRuleDatum provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AutomodRuleDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAutomodRuleDatum(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AutomodRuleDatumSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty AutomodRuleDatumSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AutomodRuleDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AutomodRuleDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), automodRuleDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"automod_rule_data\".* FROM \"automod_rule_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, automodRuleDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AutomodRuleDatumSlice")
	}

	*o = slice

	return nil
}

// AutomodRuleDatumExistsG checks if the AutomodRuleDatum row exists.
func AutomodRuleDatumExistsG(ctx context.Context, iD int64) (bool, error) {
	return AutomodRuleDatumExists(ctx, boil.GetContextDB(), iD)
}

// AutomodRuleDatumExists checks if the AutomodRuleDatum row exists.
func AutomodRuleDatumExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"automod_rule_data\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if automod_rule_data exists")
	}

	return exists, nil
}
