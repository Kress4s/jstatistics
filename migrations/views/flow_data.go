package views

import (
	"fmt"
	"js_statistics/app/models/tables"
	"js_statistics/app/models/views"

	"gorm.io/gorm"
)

func CreateFlowDataView(db *gorm.DB) error {
	sql := fmt.Sprintf(`CREATE OR REPLACE VIEW "%s" AS
	WITH ip_search as (
SELECT
	js_id,
	category_id,
	primary_id,
	visit_time,
	count(*) as count
FROM
	%s
GROUP BY
	primary_id, js_id, category_id, visit_time
	),
	uv_search as (
SELECT
	js_id,
	category_id,
	primary_id,
	visit_time,
	count(*) as count
FROM
	%s
GROUP BY
	js_id, category_id, primary_id, visit_time
	)
SELECT 
	ip.js_id as js_id,
	ip.category_id as category_id,
	ip.primary_id as primary_id,
	ip.count as ip_count,
	ip.visit_time as ip_time,
	uv.count as uv_count,
	uv.visit_time as uv_time,
	jm.title as title 
FROM  ip_search AS ip 
INNER JOIN uv_search as uv on ip.js_id = uv.js_id
INNER JOIN %s as jm on jm.id = uv.js_id;`,
		views.FlowDataView, tables.IPStatistics, tables.UVStatistics, tables.JsManage)
	return db.Exec(sql).Error
}
