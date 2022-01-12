package impl

const (
	insertTaskSQL = `INSERT INTO task (
		id,region,resource_type,secret_id,secret_desc,timeout,status,
		message,start_at,end_at,total_succeed,total_failed
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?);`
	updateTaskSQL = `UPDATE task SET status=?,message=?,end_at=?,
		total_succeed=?,total_failed=? 
	WHERE id = ?`
	updateTaskRecordSQL = `INSERT INTO task_record (
		instance_id,instance_name,is_success,message,task_id,create_at
	) VALUES (?,?,?,?,?,?);`

	queryTaskSQL = `SELECT * FROM task`
	queryTaskRecordSQL = `SELECT * FROM task_record`
)