package test_reflect

const (
	FromQueryTag = "search"
	Mysql        = "mysql"
	Postgres     = "postgres"
)

type resolveSearchTag struct {
	Type   string
	Column string
	Table  string
	On     []string
	Join   string
}

type Condition interface {
	SetWhere(k string, v []interface{})
	SetOr(k string, v []interface{})
	SetOrder(k string)
	SetJoinOn(t, on string) Condition
}

type GormCondition struct {
	GormPublic
	Join []*GormJoin
}

type GormPublic struct {
	Where map[string]interface{}
	Order []string
	Or    map[string]interface{}
}

type GormJoin struct {
	Type   string
	JoinOn string
	GormPublic
}

func (e *GormJoin) SetJoinOn(t, on string) Condition {
	return nil
}

func (e *GormPublic) SetWhere(k string, v []interface{}) {
	if e.Where == nil {
		e.Where = make(map[string]interface{})
	}
	e.Where[k] = v
}

func (e *GormPublic) SetOr(k string, v []interface{}) {
	if e.Or == nil {
		e.Or = make(map[string]interface{})
	}
	e.Or[k] = v
}

func (e *GormPublic) SetOrder(k string) {
	if e.Order == nil {
		e.Order = make([]string, 0)
	}
	e.Order = append(e.Order, k)
}

func (e *GormCondition) SetJoinOn(t, on string) Condition{
	if e.Join == nil{
		e.Join = make([]*GormJoin,0)
	}
	join := &GormJoin{
		Type:       t,
		JoinOn:     on,
		GormPublic: GormPublic{},
	}
	e.Join = append(e.Join, join)
	return join
}

func ResolveSearchQuery(driver string, q interface{},condition Condition) {

}
