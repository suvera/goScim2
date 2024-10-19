package mysql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterConverter(t *testing.T) {
	c := NewMysqlFilterConverter()

	c.SetBindCounter(1)
	c.Convert("userName eq \"bjensen\"", map[string]string{"userName": "user_name"})
	clause := c.GetClause()

	assert.Equal(t, "user_name = :svr_1", clause.WhereClause)
	assert.Equal(t, 1, len(clause.Binds))

	c.SetBindCounter(1)
	c.Convert("meta.lastModified gt \"2011-05-13T04:42:34Z\"", map[string]string{"meta.lastModified": "last_modified"})
	clause = c.GetClause()

	assert.Equal(t, "last_modified > :svr_1", clause.WhereClause)
	assert.Equal(t, 1, len(clause.Binds))

	c.SetBindCounter(1)
	c.Convert("emails[type eq \"work\" and value co \"@example.com\"] or ims[type eq \"xmpp\" and value co \"@foo.com\"]",
		map[string]string{
			"emails.value": "email_address",
			"emails.type":  "email_type",
			"ims.type":     "im_name",
			"ims.value":    "im_value",
		})
	clause = c.GetClause()

	assert.NotNil(t, clause.WhereClause)
	assert.Equal(t, " ( email_type = :svr_1 AND email_address LIKE :svr_2 )  OR  ( im_name = :svr_3 AND im_value LIKE :svr_4 ) ", clause.WhereClause)
	assert.Equal(t, 4, len(clause.Binds))
}
