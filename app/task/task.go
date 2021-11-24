package task

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

const (
	AppName = "Task"
)

var (
	validate = validator.New()
)

func NewDefaultTask() *Task {
	return &Task{
		Details: []*Detail{},
	}
}

func NewTaskFromReq(req *CreateTaskRequst) (*Task, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate sync request error, %s", err)
	}

	return &Task{
		Id:           xid.New().String(),
		Region:       req.Region,
		ResourceType: req.ResourceType,
		SecretId:     req.SecretId,
		Details:      []*Detail{},
	}, nil
}

func (s *Task) Run() {
	s.StartAt = ftime.Now().Timestamp()
	s.Status = Status_RUNNING
}

func (s *Task) UpdateSecretDesc(desc string) {
	s.SecretDescription = desc
}

func (s *Task) Completed() {
	s.EndAt = ftime.Now().Timestamp()
	if s.Status != Status_FAILED {
		if s.TotalFailed == 0 {
			s.Status = Status_SUCCESS
		} else {
			s.Status = Status_WARNING
		}
	}
}

func (s *Task) Failed(message string) {
	s.Status = Status_FAILED
	s.Message = message
}

func (s *Task) AddDetailFailed(name, message string) {
	s.Details = append(s.Details, &Detail{
		IsSuccess: false,
		Name:      name,
		Message:   message,
	})
	s.TotalFailed++
}

func (s *Task) AddDetailSucceed(name, message string) {
	s.Details = append(s.Details, &Detail{
		IsSuccess: true,
		Name:      name,
		Message:   message,
	})
	s.TotalSucceed++
}

func NewTaskSet() *TaskSet {
	return &TaskSet{
		Items: []*Task{},
	}
}

func (r *TaskSet) Add(item *Task) {
	r.Items = append(r.Items, item)
}

func NewCreateTaskRequst() *CreateTaskRequst {
	return &CreateTaskRequst{}
}

func (req *CreateTaskRequst) Validate() error {
	return validate.Struct(req)
}

func NewQueryTaskRequestFromHTTP(r *http.Request) *QueryTaskRequest {
	qs := r.URL.Query()

	ps := qs.Get("page_size")
	pn := qs.Get("page_number")
	kw := qs.Get("keywords")

	psUint64, _ := strconv.ParseUint(ps, 10, 64)
	pnUint64, _ := strconv.ParseUint(pn, 10, 64)

	if psUint64 == 0 {
		psUint64 = 20
	}
	if pnUint64 == 0 {
		pnUint64 = 1
	}
	return &QueryTaskRequest{
		PageSize:   psUint64,
		PageNumber: pnUint64,
		Keywords:   kw,
	}
}

func (req *QueryTaskRequest) OffSet() int64 {
	return int64(req.PageSize) * int64(req.PageNumber-1)
}
