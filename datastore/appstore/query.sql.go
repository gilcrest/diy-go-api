// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package appstore

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createApp = `-- name: CreateApp :execrows
INSERT INTO app (app_id, org_id, app_extl_id, app_name, app_description, create_app_id, create_user_id,
                 create_timestamp, update_app_id, update_user_id, update_timestamp)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
`

type CreateAppParams struct {
	AppID           uuid.UUID
	OrgID           uuid.UUID
	AppExtlID       string
	AppName         string
	AppDescription  string
	CreateAppID     uuid.UUID
	CreateUserID    uuid.NullUUID
	CreateTimestamp time.Time
	UpdateAppID     uuid.UUID
	UpdateUserID    uuid.NullUUID
	UpdateTimestamp time.Time
}

func (q *Queries) CreateApp(ctx context.Context, arg CreateAppParams) (int64, error) {
	result, err := q.db.Exec(ctx, createApp,
		arg.AppID,
		arg.OrgID,
		arg.AppExtlID,
		arg.AppName,
		arg.AppDescription,
		arg.CreateAppID,
		arg.CreateUserID,
		arg.CreateTimestamp,
		arg.UpdateAppID,
		arg.UpdateUserID,
		arg.UpdateTimestamp,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const createAppAPIKey = `-- name: CreateAppAPIKey :execrows
INSERT INTO app_api_key (api_key, app_id, deactv_date, create_app_id, create_user_id,
                         create_timestamp, update_app_id, update_user_id, update_timestamp)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
`

type CreateAppAPIKeyParams struct {
	ApiKey          string
	AppID           uuid.UUID
	DeactvDate      time.Time
	CreateAppID     uuid.UUID
	CreateUserID    uuid.NullUUID
	CreateTimestamp time.Time
	UpdateAppID     uuid.UUID
	UpdateUserID    uuid.NullUUID
	UpdateTimestamp time.Time
}

func (q *Queries) CreateAppAPIKey(ctx context.Context, arg CreateAppAPIKeyParams) (int64, error) {
	result, err := q.db.Exec(ctx, createAppAPIKey,
		arg.ApiKey,
		arg.AppID,
		arg.DeactvDate,
		arg.CreateAppID,
		arg.CreateUserID,
		arg.CreateTimestamp,
		arg.UpdateAppID,
		arg.UpdateUserID,
		arg.UpdateTimestamp,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const deleteApp = `-- name: DeleteApp :execrows
DELETE FROM app
WHERE app_id = $1
`

func (q *Queries) DeleteApp(ctx context.Context, appID uuid.UUID) (int64, error) {
	result, err := q.db.Exec(ctx, deleteApp, appID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const deleteAppAPIKey = `-- name: DeleteAppAPIKey :execrows
DELETE FROM app_api_key
WHERE api_key = $1
`

func (q *Queries) DeleteAppAPIKey(ctx context.Context, apiKey string) (int64, error) {
	result, err := q.db.Exec(ctx, deleteAppAPIKey, apiKey)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const deleteAppAPIKeys = `-- name: DeleteAppAPIKeys :execrows
DELETE FROM app_api_key
WHERE app_id = $1
`

func (q *Queries) DeleteAppAPIKeys(ctx context.Context, appID uuid.UUID) (int64, error) {
	result, err := q.db.Exec(ctx, deleteAppAPIKeys, appID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const findAPIKeysByAppID = `-- name: FindAPIKeysByAppID :many
SELECT api_key, app_id, deactv_date, create_app_id, create_user_id, create_timestamp, update_app_id, update_user_id, update_timestamp FROM app_api_key
WHERE app_id = $1
`

func (q *Queries) FindAPIKeysByAppID(ctx context.Context, appID uuid.UUID) ([]AppApiKey, error) {
	rows, err := q.db.Query(ctx, findAPIKeysByAppID, appID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AppApiKey
	for rows.Next() {
		var i AppApiKey
		if err := rows.Scan(
			&i.ApiKey,
			&i.AppID,
			&i.DeactvDate,
			&i.CreateAppID,
			&i.CreateUserID,
			&i.CreateTimestamp,
			&i.UpdateAppID,
			&i.UpdateUserID,
			&i.UpdateTimestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAppAPIKeysByAppExtlID = `-- name: FindAppAPIKeysByAppExtlID :many
select a.app_id,
       a.app_extl_id,
       a.app_name,
       a.app_description,
       o.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       aak.api_key,
       aak.deactv_date
from app a
         inner join org o on o.org_id = a.org_id
         inner join app_api_key aak on a.app_id = aak.app_id
where a.app_extl_id = $1
`

type FindAppAPIKeysByAppExtlIDRow struct {
	AppID          uuid.UUID
	AppExtlID      string
	AppName        string
	AppDescription string
	OrgID          uuid.UUID
	OrgExtlID      string
	OrgName        string
	OrgDescription string
	ApiKey         string
	DeactvDate     time.Time
}

func (q *Queries) FindAppAPIKeysByAppExtlID(ctx context.Context, appExtlID string) ([]FindAppAPIKeysByAppExtlIDRow, error) {
	rows, err := q.db.Query(ctx, findAppAPIKeysByAppExtlID, appExtlID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindAppAPIKeysByAppExtlIDRow
	for rows.Next() {
		var i FindAppAPIKeysByAppExtlIDRow
		if err := rows.Scan(
			&i.AppID,
			&i.AppExtlID,
			&i.AppName,
			&i.AppDescription,
			&i.OrgID,
			&i.OrgExtlID,
			&i.OrgName,
			&i.OrgDescription,
			&i.ApiKey,
			&i.DeactvDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findAppByExternalID = `-- name: FindAppByExternalID :one
SELECT a.app_id,
       a.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       ok.org_kind_id,
       ok.org_kind_extl_id,
       ok.org_kind_desc,
       a.app_extl_id,
       a.app_name,
       a.app_description
FROM app a
         INNER JOIN org o on o.org_id = a.org_id
         INNER JOIN org_kind ok on ok.org_kind_id = o.org_kind_id
WHERE a.app_extl_id = $1
`

type FindAppByExternalIDRow struct {
	AppID          uuid.UUID
	OrgID          uuid.UUID
	OrgExtlID      string
	OrgName        string
	OrgDescription string
	OrgKindID      uuid.UUID
	OrgKindExtlID  string
	OrgKindDesc    string
	AppExtlID      string
	AppName        string
	AppDescription string
}

func (q *Queries) FindAppByExternalID(ctx context.Context, appExtlID string) (FindAppByExternalIDRow, error) {
	row := q.db.QueryRow(ctx, findAppByExternalID, appExtlID)
	var i FindAppByExternalIDRow
	err := row.Scan(
		&i.AppID,
		&i.OrgID,
		&i.OrgExtlID,
		&i.OrgName,
		&i.OrgDescription,
		&i.OrgKindID,
		&i.OrgKindExtlID,
		&i.OrgKindDesc,
		&i.AppExtlID,
		&i.AppName,
		&i.AppDescription,
	)
	return i, err
}

const findAppByExternalIDWithAudit = `-- name: FindAppByExternalIDWithAudit :one
SELECT a.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       ok.org_kind_id,
       ok.org_kind_extl_id,
       ok.org_kind_desc,
       a.app_id,
       a.app_extl_id,
       a.app_name,
       a.app_description,
       a.create_app_id,
       ca.org_id          create_app_org_id,
       ca.app_extl_id     create_app_extl_id,
       ca.app_name        create_app_name,
       ca.app_description create_app_description,
       a.create_user_id,
       cu.username        create_username,
       cu.org_id          create_user_org_id,
       cup.first_name     create_user_first_name,
       cup.last_name      create_user_last_name,
       a.create_timestamp,
       a.update_app_id,
       ua.org_id          update_app_org_id,
       ua.app_extl_id     update_app_extl_id,
       ua.app_name        update_app_name,
       ua.app_description update_app_description,
       a.update_user_id,
       uu.username        update_username,
       uu.org_id          update_user_org_id,
       uup.first_name     update_user_first_name,
       uup.last_name      update_user_last_name,
       a.update_timestamp
FROM app a
         INNER JOIN org o on o.org_id = a.org_id
         INNER JOIN org_kind ok on ok.org_kind_id = o.org_kind_id
         INNER JOIN app ca on ca.app_id = a.create_app_id
         INNER JOIN app ua on ua.app_id = a.update_app_id
         LEFT JOIN org_user cu on cu.user_id = a.create_user_id
         INNER JOIN person_profile cup on cup.person_profile_id = cu.person_profile_id
         LEFT JOIN org_user uu on uu.user_id = a.update_user_id
         INNER JOIN person_profile uup on uup.person_profile_id = uu.person_profile_id
WHERE a.app_extl_id = $1
`

type FindAppByExternalIDWithAuditRow struct {
	OrgID                uuid.UUID
	OrgExtlID            string
	OrgName              string
	OrgDescription       string
	OrgKindID            uuid.UUID
	OrgKindExtlID        string
	OrgKindDesc          string
	AppID                uuid.UUID
	AppExtlID            string
	AppName              string
	AppDescription       string
	CreateAppID          uuid.UUID
	CreateAppOrgID       uuid.UUID
	CreateAppExtlID      string
	CreateAppName        string
	CreateAppDescription string
	CreateUserID         uuid.NullUUID
	CreateUsername       string
	CreateUserOrgID      uuid.UUID
	CreateUserFirstName  string
	CreateUserLastName   string
	CreateTimestamp      time.Time
	UpdateAppID          uuid.UUID
	UpdateAppOrgID       uuid.UUID
	UpdateAppExtlID      string
	UpdateAppName        string
	UpdateAppDescription string
	UpdateUserID         uuid.NullUUID
	UpdateUsername       string
	UpdateUserOrgID      uuid.UUID
	UpdateUserFirstName  string
	UpdateUserLastName   string
	UpdateTimestamp      time.Time
}

func (q *Queries) FindAppByExternalIDWithAudit(ctx context.Context, appExtlID string) (FindAppByExternalIDWithAuditRow, error) {
	row := q.db.QueryRow(ctx, findAppByExternalIDWithAudit, appExtlID)
	var i FindAppByExternalIDWithAuditRow
	err := row.Scan(
		&i.OrgID,
		&i.OrgExtlID,
		&i.OrgName,
		&i.OrgDescription,
		&i.OrgKindID,
		&i.OrgKindExtlID,
		&i.OrgKindDesc,
		&i.AppID,
		&i.AppExtlID,
		&i.AppName,
		&i.AppDescription,
		&i.CreateAppID,
		&i.CreateAppOrgID,
		&i.CreateAppExtlID,
		&i.CreateAppName,
		&i.CreateAppDescription,
		&i.CreateUserID,
		&i.CreateUsername,
		&i.CreateUserOrgID,
		&i.CreateUserFirstName,
		&i.CreateUserLastName,
		&i.CreateTimestamp,
		&i.UpdateAppID,
		&i.UpdateAppOrgID,
		&i.UpdateAppExtlID,
		&i.UpdateAppName,
		&i.UpdateAppDescription,
		&i.UpdateUserID,
		&i.UpdateUsername,
		&i.UpdateUserOrgID,
		&i.UpdateUserFirstName,
		&i.UpdateUserLastName,
		&i.UpdateTimestamp,
	)
	return i, err
}

const findAppByID = `-- name: FindAppByID :one
SELECT a.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       ok.org_kind_id,
       ok.org_kind_extl_id,
       ok.org_kind_desc,
       a.app_id,
       a.app_extl_id,
       a.app_name,
       a.app_description
FROM app a
         INNER JOIN org o on o.org_id = a.org_id
         INNER JOIN org_kind ok on ok.org_kind_id = o.org_kind_id
WHERE a.app_id = $1
`

type FindAppByIDRow struct {
	OrgID          uuid.UUID
	OrgExtlID      string
	OrgName        string
	OrgDescription string
	OrgKindID      uuid.UUID
	OrgKindExtlID  string
	OrgKindDesc    string
	AppID          uuid.UUID
	AppExtlID      string
	AppName        string
	AppDescription string
}

func (q *Queries) FindAppByID(ctx context.Context, appID uuid.UUID) (FindAppByIDRow, error) {
	row := q.db.QueryRow(ctx, findAppByID, appID)
	var i FindAppByIDRow
	err := row.Scan(
		&i.OrgID,
		&i.OrgExtlID,
		&i.OrgName,
		&i.OrgDescription,
		&i.OrgKindID,
		&i.OrgKindExtlID,
		&i.OrgKindDesc,
		&i.AppID,
		&i.AppExtlID,
		&i.AppName,
		&i.AppDescription,
	)
	return i, err
}

const findAppByIDWithAudit = `-- name: FindAppByIDWithAudit :one
SELECT a.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       ok.org_kind_id,
       ok.org_kind_extl_id,
       ok.org_kind_desc,
       a.app_id,
       a.app_extl_id,
       a.app_name,
       a.app_description,
       a.create_app_id,
       ca.org_id          create_app_org_id,
       ca.app_extl_id     create_app_extl_id,
       ca.app_name        create_app_name,
       ca.app_description create_app_description,
       a.create_user_id,
       cu.username        create_username,
       cu.org_id          create_user_org_id,
       cup.first_name     create_user_first_name,
       cup.last_name      create_user_last_name,
       a.create_timestamp,
       a.update_app_id,
       ua.org_id          update_app_org_id,
       ua.app_extl_id     update_app_extl_id,
       ua.app_name        update_app_name,
       ua.app_description update_app_description,
       a.update_user_id,
       uu.username        update_username,
       uu.org_id          update_user_org_id,
       uup.first_name     update_user_first_name,
       uup.last_name      update_user_last_name,
       a.update_timestamp
FROM app a
         INNER JOIN org o on o.org_id = a.org_id
         INNER JOIN org_kind ok on ok.org_kind_id = o.org_kind_id
         INNER JOIN app ca on ca.app_id = a.create_app_id
         INNER JOIN app ua on ua.app_id = a.update_app_id
         LEFT JOIN org_user cu on cu.user_id = a.create_user_id
         INNER JOIN person_profile cup on cup.person_profile_id = cu.person_profile_id
         LEFT JOIN org_user uu on uu.user_id = a.update_user_id
         INNER JOIN person_profile uup on uup.person_profile_id = uu.person_profile_id
WHERE a.app_id = $1
`

type FindAppByIDWithAuditRow struct {
	OrgID                uuid.UUID
	OrgExtlID            string
	OrgName              string
	OrgDescription       string
	OrgKindID            uuid.UUID
	OrgKindExtlID        string
	OrgKindDesc          string
	AppID                uuid.UUID
	AppExtlID            string
	AppName              string
	AppDescription       string
	CreateAppID          uuid.UUID
	CreateAppOrgID       uuid.UUID
	CreateAppExtlID      string
	CreateAppName        string
	CreateAppDescription string
	CreateUserID         uuid.NullUUID
	CreateUsername       string
	CreateUserOrgID      uuid.UUID
	CreateUserFirstName  string
	CreateUserLastName   string
	CreateTimestamp      time.Time
	UpdateAppID          uuid.UUID
	UpdateAppOrgID       uuid.UUID
	UpdateAppExtlID      string
	UpdateAppName        string
	UpdateAppDescription string
	UpdateUserID         uuid.NullUUID
	UpdateUsername       string
	UpdateUserOrgID      uuid.UUID
	UpdateUserFirstName  string
	UpdateUserLastName   string
	UpdateTimestamp      time.Time
}

func (q *Queries) FindAppByIDWithAudit(ctx context.Context, appID uuid.UUID) (FindAppByIDWithAuditRow, error) {
	row := q.db.QueryRow(ctx, findAppByIDWithAudit, appID)
	var i FindAppByIDWithAuditRow
	err := row.Scan(
		&i.OrgID,
		&i.OrgExtlID,
		&i.OrgName,
		&i.OrgDescription,
		&i.OrgKindID,
		&i.OrgKindExtlID,
		&i.OrgKindDesc,
		&i.AppID,
		&i.AppExtlID,
		&i.AppName,
		&i.AppDescription,
		&i.CreateAppID,
		&i.CreateAppOrgID,
		&i.CreateAppExtlID,
		&i.CreateAppName,
		&i.CreateAppDescription,
		&i.CreateUserID,
		&i.CreateUsername,
		&i.CreateUserOrgID,
		&i.CreateUserFirstName,
		&i.CreateUserLastName,
		&i.CreateTimestamp,
		&i.UpdateAppID,
		&i.UpdateAppOrgID,
		&i.UpdateAppExtlID,
		&i.UpdateAppName,
		&i.UpdateAppDescription,
		&i.UpdateUserID,
		&i.UpdateUsername,
		&i.UpdateUserOrgID,
		&i.UpdateUserFirstName,
		&i.UpdateUserLastName,
		&i.UpdateTimestamp,
	)
	return i, err
}

const findAppByName = `-- name: FindAppByName :one
SELECT a.app_id,
       a.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       ok.org_kind_id,
       ok.org_kind_extl_id,
       ok.org_kind_desc,
       a.app_extl_id,
       a.app_name,
       a.app_description
FROM app a
         INNER JOIN org o on o.org_id = a.org_id
         INNER JOIN org_kind ok on ok.org_kind_id = o.org_kind_id
WHERE o.org_id = $1
  AND a.app_name = $2
`

type FindAppByNameParams struct {
	OrgID   uuid.UUID
	AppName string
}

type FindAppByNameRow struct {
	AppID          uuid.UUID
	OrgID          uuid.UUID
	OrgExtlID      string
	OrgName        string
	OrgDescription string
	OrgKindID      uuid.UUID
	OrgKindExtlID  string
	OrgKindDesc    string
	AppExtlID      string
	AppName        string
	AppDescription string
}

func (q *Queries) FindAppByName(ctx context.Context, arg FindAppByNameParams) (FindAppByNameRow, error) {
	row := q.db.QueryRow(ctx, findAppByName, arg.OrgID, arg.AppName)
	var i FindAppByNameRow
	err := row.Scan(
		&i.AppID,
		&i.OrgID,
		&i.OrgExtlID,
		&i.OrgName,
		&i.OrgDescription,
		&i.OrgKindID,
		&i.OrgKindExtlID,
		&i.OrgKindDesc,
		&i.AppExtlID,
		&i.AppName,
		&i.AppDescription,
	)
	return i, err
}

const findApps = `-- name: FindApps :many
SELECT app_id, org_id, app_extl_id, app_name, app_description, create_app_id, create_user_id, create_timestamp, update_app_id, update_user_id, update_timestamp FROM app
ORDER BY app_name
`

func (q *Queries) FindApps(ctx context.Context) ([]App, error) {
	rows, err := q.db.Query(ctx, findApps)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []App
	for rows.Next() {
		var i App
		if err := rows.Scan(
			&i.AppID,
			&i.OrgID,
			&i.AppExtlID,
			&i.AppName,
			&i.AppDescription,
			&i.CreateAppID,
			&i.CreateUserID,
			&i.CreateTimestamp,
			&i.UpdateAppID,
			&i.UpdateUserID,
			&i.UpdateTimestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateApp = `-- name: UpdateApp :execrows
UPDATE app
SET app_name        = $1,
    app_description = $2,
    update_app_id    = $3,
    update_user_id   = $4,
    update_timestamp = $5
WHERE app_id = $6
`

type UpdateAppParams struct {
	AppName         string
	AppDescription  string
	UpdateAppID     uuid.UUID
	UpdateUserID    uuid.NullUUID
	UpdateTimestamp time.Time
	AppID           uuid.UUID
}

func (q *Queries) UpdateApp(ctx context.Context, arg UpdateAppParams) (int64, error) {
	result, err := q.db.Exec(ctx, updateApp,
		arg.AppName,
		arg.AppDescription,
		arg.UpdateAppID,
		arg.UpdateUserID,
		arg.UpdateTimestamp,
		arg.AppID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
