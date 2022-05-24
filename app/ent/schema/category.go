package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "category"},
	}
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		//`category_id` tinyint UNSIGNED NOT NULL AUTO_INCREMENT,
		//PRIMARY KEY (`category_id`)
		field.Uint8("id").
			StorageKey("category_id"),
		//`name` varchar(25) NOT NULL,
		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(25)",
		}),
		//`last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		field.Time("last_update").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			UpdateDefault(time.Now),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("films", Film.Type).Ref("categories"),
	}
}
