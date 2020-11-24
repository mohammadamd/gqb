package gqb

import (
	"fmt"
	"strconv"
	"strings"
)

const PreDefined = "TRUE true FALSE false NULL null"

type QueryBuilder struct {
	from        string
	stmtBuilder func(*QueryBuilder) string
	selects     []string
	joins       []string
	wheres      []string
	orders      []string
	skip        string
	limit       string
}

func Select(columns ...string) *QueryBuilder {
	return &QueryBuilder{selects: columns, stmtBuilder: selectGenerator}
}

func (q *QueryBuilder) Table(tableName string) *QueryBuilder {
	q.from = tableName
	return q
}

func (q *QueryBuilder) Where(col string, operator string, value string) *QueryBuilder {
	s := "AND WHERE"

	if len(q.wheres) == 0 {
		s = "WHERE"
	}

	if _, err := strconv.Atoi(value); err != nil && !strings.Contains(PreDefined, value) {
		value = fmt.Sprintf("\"%s\"", value)
	}

	nw := fmt.Sprintf("%s %s %s %s", s, col, operator, value)
	q.wheres = append(q.wheres, nw)
	return q
}

func (q *QueryBuilder) OrWhere(col string, operator string, value string) *QueryBuilder {
	nw := fmt.Sprintf("OR WHERE %s %s %s", col, operator, value)
	q.wheres = append(q.wheres, nw)
	return q
}

func (q *QueryBuilder) OrderBy(col, order string) *QueryBuilder {
	no := fmt.Sprintf("ORDER BY %s %s", col, order)
	q.orders = append(q.orders, no)
	return q
}

func (q *QueryBuilder) Join(table, col1, operator, col2 string) *QueryBuilder {
	nj := fmt.Sprintf("JOIN %s ON %s %s %s", table, col1, operator, col2)
	q.joins = append(q.joins, nj)
	return q
}

func (q *QueryBuilder) Skip(s int) *QueryBuilder {
	q.skip = fmt.Sprintf("OFFSET %s", strconv.Itoa(s))
	return q
}

func (q *QueryBuilder) Limit(l int) *QueryBuilder {
	q.limit = fmt.Sprintf("LIMIT %s", strconv.Itoa(l))
	return q
}

func (q *QueryBuilder) Generate() string {
	return q.stmtBuilder(q)
}

func selectGenerator(q *QueryBuilder) string {
	s := strings.Join(q.selects, ",")
	j := strings.Join(q.joins, " ")
	w := strings.Join(q.wheres, " ")
	o := strings.Join(q.orders, " ")

	return fmt.Sprintf("SELECT %s FROM %s %s %s %s %s %s",
		s,
		q.from,
		j,
		w,
		o,
		q.limit,
		q.skip,
	)
}
