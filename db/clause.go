package db

import (
	"gorm.io/gorm/clause"
)

type (
	AndConditions   = clause.AndConditions
	Assignment      = clause.Assignment
	Clause          = clause.Clause
	Column          = clause.Column
	CommaExpression = clause.CommaExpression
	Delete          = clause.Delete
	Eq              = clause.Eq
	Expr            = clause.Expr
	From            = clause.From
	GroupBy         = clause.GroupBy
	Gt              = clause.Gt
	Gte             = clause.Gte
	IN              = clause.IN
	Insert          = clause.Insert
	Join            = clause.Join
	JoinType        = clause.JoinType
	Like            = clause.Like
	Limit           = clause.Limit
	Locking         = clause.Locking
	Lt              = clause.Lt
	Lte             = clause.Lte
	NamedExpr       = clause.NamedExpr
	Neq             = clause.Neq
	NotConditions   = clause.NotConditions
	OnConflict      = clause.OnConflict
	OrConditions    = clause.OrConditions
	OrderBy         = clause.OrderBy
	OrderByColumn   = clause.OrderByColumn
	Returning       = clause.Returning
	Select          = clause.Select
	Set             = clause.Set
	Table           = clause.Table
	Update          = clause.Update
	Values          = clause.Values
	Where           = clause.Where
	With            = clause.With
)

var (
	And = clause.And
	Not = clause.Not
	Or  = clause.Or

	Assignments       = clause.Assignments
	AssignmentColumns = clause.AssignmentColumns
)
