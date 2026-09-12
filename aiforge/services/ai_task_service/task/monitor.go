package task

import (
	"encoding/csv"
	"strconv"
	"time"

	"code.gitea.io/gitea/modules/context"

	"code.gitea.io/gitea/models"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/util"
)

func MonitorTaskFile(status string, duration int, days int, ctx *context.Context) error {

	total := monitorTasksCount(status, duration, days)
	pageSize := 300
	totalPage := util.GetTotalPage(total, pageSize)

	f := csv.NewWriter(ctx.Resp)

	err := f.Write(allTaskHeader())
	if err != nil {
		return err
	}

	for i := 0; i <= totalPage; i++ {

		pageRecords, err := monitorTasks(status, duration, days, i+1, pageSize)
		if err != nil {
			log.Warn("Get monitor page task err", err)
			continue

		}
		for _, record := range pageRecords {

			err = f.Write(allTaskValues(record))
			if err != nil {
				return err
			}

		}

	}
	f.Flush()
	return nil
}

func allTaskValues(record map[string]string) []string {
	return []string{record["id"], record["job_id"], record["job_type"], record["job_name"], record["compute_resource"],
		record["status"], record["type"], record["work_server_number"], record["owner_name"], record["repo_name"],
		record["train_job_duration"], record["email"], record["wechat_open_id"], record["name"]}
}

func allTaskHeader() []string {
	return []string{"id", "job_id", "job_type", "job_name",
		"compute_resource", "status", "type",
		"work_server_number",
		"owner_name", "repo_name", "train_job_duration", "email", "wechat_open_id", "user_name"}
}

func monitorTasksCount(status string, duration int, days int) int64 {

	countSql := "select count(*) from cloudbrain a, repository b where a.repo_id=b.id and (job_type='TRAIN' or job_type='ONLINEINFERENCE') and type!=0 and type!=3 "
	if status != "" {
		countSql += " and a.status= '" + status + "' "
	}
	if duration > 0 {
		countSql += " and a.duration >=" + strconv.Itoa(duration)
	}
	if days > 0 {
		countSql += " and a.created_unix >=" + strconv.FormatInt(time.Now().Unix()-int64(days*24*60*60), 10)
	}

	count, err := models.CountByRawSql(countSql)

	if err != nil {
		log.Error("Get monitor task count error", err)
	}
	return count
}

func monitorTasks(status string, duration int, days int, page, pageSize int) ([]map[string]string, error) {

	sql := "select a.id, job_id,job_type,job_name,compute_resource, a.status,a.type,work_server_number,owner_name, b.name as repo_name,train_job_duration, c.email,c.wechat_open_id,c.name from cloudbrain a, repository b,public.user c where a.repo_id=b.id and a.user_id=c.id and (job_type='TRAIN' or job_type='ONLINEINFERENCE') and a.type!=0 and a.type!=3 "
	if status != "" {
		sql += " and a.status= '" + status + "' "
	}
	if duration > 0 {
		sql += " and a.duration >=" + strconv.Itoa(duration)
	}
	if days > 0 {
		sql += " and a.created_unix >=" + strconv.FormatInt(time.Now().Unix()-int64(days*24*60*60), 10)
	}
	sql += " order by a.id desc"

	sql += " limit " + strconv.Itoa(pageSize) + " offset " + strconv.Itoa((page-1)*pageSize)

	return models.QueryByRawSql(sql)

}
