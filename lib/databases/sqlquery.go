package databases

import "fmt"

type SQLFilterQuery struct {
	query      string
	args       []any
	argCounter int
}

func NewSQLFilterQuery(query string) *SQLFilterQuery {
	return &SQLFilterQuery{query: query, argCounter: 0}
}

func (q *SQLFilterQuery) WhereEq(column string, arg any) {
	if arg == nil {
		return
	}

	q.args = append(q.args, arg)
	q.argCounter++
	q.query = fmt.Sprintf("%s %s %s = $%d ", q.query, "WHERE", column, q.argCounter)
}

func (q *SQLFilterQuery) WhereLike(column string, arg any) {
	if arg == nil {
		return
	}

	q.args = append(q.args, arg)
	q.argCounter++
	q.query = fmt.Sprintf("%s %s %s LIKE $%d ", q.query, "WHERE", column, q.argCounter)
}

func (q *SQLFilterQuery) Query() string {
	return q.query
}

func (q *SQLFilterQuery) Args() []any {
	return q.args
}
