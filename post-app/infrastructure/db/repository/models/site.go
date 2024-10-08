// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Site is an object representing the database table.
type Site struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	IsActive  bool      `boil:"is_active" json:"is_active" toml:"is_active" yaml:"is_active"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *siteR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L siteL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SiteColumns = struct {
	ID        string
	Name      string
	IsActive  string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	IsActive:  "is_active",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var SiteTableColumns = struct {
	ID        string
	Name      string
	IsActive  string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "site.id",
	Name:      "site.name",
	IsActive:  "site.is_active",
	CreatedAt: "site.created_at",
	UpdatedAt: "site.updated_at",
}

// Generated where

var SiteWhere = struct {
	ID        whereHelperint64
	Name      whereHelperstring
	IsActive  whereHelperbool
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "\"site\".\"id\""},
	Name:      whereHelperstring{field: "\"site\".\"name\""},
	IsActive:  whereHelperbool{field: "\"site\".\"is_active\""},
	CreatedAt: whereHelpertime_Time{field: "\"site\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"site\".\"updated_at\""},
}

// SiteRels is where relationship names are stored.
var SiteRels = struct {
	SiteOwnerships string
	SNS            string
}{
	SiteOwnerships: "SiteOwnerships",
	SNS:            "SNS",
}

// siteR is where relationships are stored.
type siteR struct {
	SiteOwnerships SiteOwnershipSlice `boil:"SiteOwnerships" json:"SiteOwnerships" toml:"SiteOwnerships" yaml:"SiteOwnerships"`
	SNS            SNSlice            `boil:"SNS" json:"SNS" toml:"SNS" yaml:"SNS"`
}

// NewStruct creates a new relationship struct
func (*siteR) NewStruct() *siteR {
	return &siteR{}
}

func (r *siteR) GetSiteOwnerships() SiteOwnershipSlice {
	if r == nil {
		return nil
	}
	return r.SiteOwnerships
}

func (r *siteR) GetSNS() SNSlice {
	if r == nil {
		return nil
	}
	return r.SNS
}

// siteL is where Load methods for each relationship are stored.
type siteL struct{}

var (
	siteAllColumns            = []string{"id", "name", "is_active", "created_at", "updated_at"}
	siteColumnsWithoutDefault = []string{"name", "is_active", "created_at", "updated_at"}
	siteColumnsWithDefault    = []string{"id"}
	sitePrimaryKeyColumns     = []string{"id"}
	siteGeneratedColumns      = []string{}
)

type (
	// SiteSlice is an alias for a slice of pointers to Site.
	// This should almost always be used instead of []Site.
	SiteSlice []*Site
	// SiteHook is the signature for custom Site hook methods
	SiteHook func(context.Context, boil.ContextExecutor, *Site) error

	siteQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	siteType                 = reflect.TypeOf(&Site{})
	siteMapping              = queries.MakeStructMapping(siteType)
	sitePrimaryKeyMapping, _ = queries.BindMapping(siteType, siteMapping, sitePrimaryKeyColumns)
	siteInsertCacheMut       sync.RWMutex
	siteInsertCache          = make(map[string]insertCache)
	siteUpdateCacheMut       sync.RWMutex
	siteUpdateCache          = make(map[string]updateCache)
	siteUpsertCacheMut       sync.RWMutex
	siteUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var siteAfterSelectMu sync.Mutex
var siteAfterSelectHooks []SiteHook

var siteBeforeInsertMu sync.Mutex
var siteBeforeInsertHooks []SiteHook
var siteAfterInsertMu sync.Mutex
var siteAfterInsertHooks []SiteHook

var siteBeforeUpdateMu sync.Mutex
var siteBeforeUpdateHooks []SiteHook
var siteAfterUpdateMu sync.Mutex
var siteAfterUpdateHooks []SiteHook

var siteBeforeDeleteMu sync.Mutex
var siteBeforeDeleteHooks []SiteHook
var siteAfterDeleteMu sync.Mutex
var siteAfterDeleteHooks []SiteHook

var siteBeforeUpsertMu sync.Mutex
var siteBeforeUpsertHooks []SiteHook
var siteAfterUpsertMu sync.Mutex
var siteAfterUpsertHooks []SiteHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Site) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range siteAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Site) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range siteBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Site) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range siteAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Site) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range siteBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Site) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range siteAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Site) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range siteBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Site) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range siteAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Site) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range siteBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Site) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range siteAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSiteHook registers your hook function for all future operations.
func AddSiteHook(hookPoint boil.HookPoint, siteHook SiteHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		siteAfterSelectMu.Lock()
		siteAfterSelectHooks = append(siteAfterSelectHooks, siteHook)
		siteAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		siteBeforeInsertMu.Lock()
		siteBeforeInsertHooks = append(siteBeforeInsertHooks, siteHook)
		siteBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		siteAfterInsertMu.Lock()
		siteAfterInsertHooks = append(siteAfterInsertHooks, siteHook)
		siteAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		siteBeforeUpdateMu.Lock()
		siteBeforeUpdateHooks = append(siteBeforeUpdateHooks, siteHook)
		siteBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		siteAfterUpdateMu.Lock()
		siteAfterUpdateHooks = append(siteAfterUpdateHooks, siteHook)
		siteAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		siteBeforeDeleteMu.Lock()
		siteBeforeDeleteHooks = append(siteBeforeDeleteHooks, siteHook)
		siteBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		siteAfterDeleteMu.Lock()
		siteAfterDeleteHooks = append(siteAfterDeleteHooks, siteHook)
		siteAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		siteBeforeUpsertMu.Lock()
		siteBeforeUpsertHooks = append(siteBeforeUpsertHooks, siteHook)
		siteBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		siteAfterUpsertMu.Lock()
		siteAfterUpsertHooks = append(siteAfterUpsertHooks, siteHook)
		siteAfterUpsertMu.Unlock()
	}
}

// One returns a single site record from the query.
func (q siteQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Site, error) {
	o := &Site{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for site")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Site records from the query.
func (q siteQuery) All(ctx context.Context, exec boil.ContextExecutor) (SiteSlice, error) {
	var o []*Site

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Site slice")
	}

	if len(siteAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Site records in the query.
func (q siteQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count site rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q siteQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if site exists")
	}

	return count > 0, nil
}

// SiteOwnerships retrieves all the site_ownership's SiteOwnerships with an executor.
func (o *Site) SiteOwnerships(mods ...qm.QueryMod) siteOwnershipQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"site_ownership\".\"site_id\"=?", o.ID),
	)

	return SiteOwnerships(queryMods...)
}

// SNS retrieves all the sn's SNS with an executor.
func (o *Site) SNS(mods ...qm.QueryMod) snQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"sns\".\"site_id\"=?", o.ID),
	)

	return SNS(queryMods...)
}

// LoadSiteOwnerships allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (siteL) LoadSiteOwnerships(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSite interface{}, mods queries.Applicator) error {
	var slice []*Site
	var object *Site

	if singular {
		var ok bool
		object, ok = maybeSite.(*Site)
		if !ok {
			object = new(Site)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSite)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSite))
			}
		}
	} else {
		s, ok := maybeSite.(*[]*Site)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSite)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSite))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &siteR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &siteR{}
			}
			args[obj.ID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`site_ownership`),
		qm.WhereIn(`site_ownership.site_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load site_ownership")
	}

	var resultSlice []*SiteOwnership
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice site_ownership")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on site_ownership")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for site_ownership")
	}

	if len(siteOwnershipAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.SiteOwnerships = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &siteOwnershipR{}
			}
			foreign.R.Site = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.SiteID {
				local.R.SiteOwnerships = append(local.R.SiteOwnerships, foreign)
				if foreign.R == nil {
					foreign.R = &siteOwnershipR{}
				}
				foreign.R.Site = local
				break
			}
		}
	}

	return nil
}

// LoadSNS allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (siteL) LoadSNS(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSite interface{}, mods queries.Applicator) error {
	var slice []*Site
	var object *Site

	if singular {
		var ok bool
		object, ok = maybeSite.(*Site)
		if !ok {
			object = new(Site)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSite)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSite))
			}
		}
	} else {
		s, ok := maybeSite.(*[]*Site)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSite)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSite))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &siteR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &siteR{}
			}
			args[obj.ID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`sns`),
		qm.WhereIn(`sns.site_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load sns")
	}

	var resultSlice []*SN
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice sns")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on sns")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for sns")
	}

	if len(snAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.SNS = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &snR{}
			}
			foreign.R.Site = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.SiteID {
				local.R.SNS = append(local.R.SNS, foreign)
				if foreign.R == nil {
					foreign.R = &snR{}
				}
				foreign.R.Site = local
				break
			}
		}
	}

	return nil
}

// AddSiteOwnerships adds the given related objects to the existing relationships
// of the site, optionally inserting them as new records.
// Appends related to o.R.SiteOwnerships.
// Sets related.R.Site appropriately.
func (o *Site) AddSiteOwnerships(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*SiteOwnership) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.SiteID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"site_ownership\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"site_id"}),
				strmangle.WhereClause("\"", "\"", 2, siteOwnershipPrimaryKeyColumns),
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

			rel.SiteID = o.ID
		}
	}

	if o.R == nil {
		o.R = &siteR{
			SiteOwnerships: related,
		}
	} else {
		o.R.SiteOwnerships = append(o.R.SiteOwnerships, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &siteOwnershipR{
				Site: o,
			}
		} else {
			rel.R.Site = o
		}
	}
	return nil
}

// AddSNS adds the given related objects to the existing relationships
// of the site, optionally inserting them as new records.
// Appends related to o.R.SNS.
// Sets related.R.Site appropriately.
func (o *Site) AddSNS(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*SN) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.SiteID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"sns\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"site_id"}),
				strmangle.WhereClause("\"", "\"", 2, snPrimaryKeyColumns),
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

			rel.SiteID = o.ID
		}
	}

	if o.R == nil {
		o.R = &siteR{
			SNS: related,
		}
	} else {
		o.R.SNS = append(o.R.SNS, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &snR{
				Site: o,
			}
		} else {
			rel.R.Site = o
		}
	}
	return nil
}

// Sites retrieves all the records using an executor.
func Sites(mods ...qm.QueryMod) siteQuery {
	mods = append(mods, qm.From("\"site\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"site\".*"})
	}

	return siteQuery{q}
}

// FindSite retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSite(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Site, error) {
	siteObj := &Site{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"site\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, siteObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from site")
	}

	if err = siteObj.doAfterSelectHooks(ctx, exec); err != nil {
		return siteObj, err
	}

	return siteObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Site) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no site provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(siteColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	siteInsertCacheMut.RLock()
	cache, cached := siteInsertCache[key]
	siteInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			siteAllColumns,
			siteColumnsWithDefault,
			siteColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(siteType, siteMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(siteType, siteMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"site\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"site\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into site")
	}

	if !cached {
		siteInsertCacheMut.Lock()
		siteInsertCache[key] = cache
		siteInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Site.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Site) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	siteUpdateCacheMut.RLock()
	cache, cached := siteUpdateCache[key]
	siteUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			siteAllColumns,
			sitePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update site, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"site\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, sitePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(siteType, siteMapping, append(wl, sitePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update site row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for site")
	}

	if !cached {
		siteUpdateCacheMut.Lock()
		siteUpdateCache[key] = cache
		siteUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q siteQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for site")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for site")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SiteSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sitePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"site\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, sitePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in site slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all site")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Site) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no site provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(siteColumnsWithDefault, o)

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

	siteUpsertCacheMut.RLock()
	cache, cached := siteUpsertCache[key]
	siteUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			siteAllColumns,
			siteColumnsWithDefault,
			siteColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			siteAllColumns,
			sitePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert site, could not build update column list")
		}

		ret := strmangle.SetComplement(siteAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(sitePrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert site, could not build conflict column list")
			}

			conflict = make([]string, len(sitePrimaryKeyColumns))
			copy(conflict, sitePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"site\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(siteType, siteMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(siteType, siteMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert site")
	}

	if !cached {
		siteUpsertCacheMut.Lock()
		siteUpsertCache[key] = cache
		siteUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Site record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Site) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Site provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), sitePrimaryKeyMapping)
	sql := "DELETE FROM \"site\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from site")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for site")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q siteQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no siteQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from site")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for site")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SiteSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(siteBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sitePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"site\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, sitePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from site slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for site")
	}

	if len(siteAfterDeleteHooks) != 0 {
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
func (o *Site) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSite(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SiteSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SiteSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sitePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"site\".* FROM \"site\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, sitePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SiteSlice")
	}

	*o = slice

	return nil
}

// SiteExists checks if the Site row exists.
func SiteExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"site\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if site exists")
	}

	return exists, nil
}

// Exists checks if the Site row exists.
func (o *Site) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SiteExists(ctx, exec, o.ID)
}
