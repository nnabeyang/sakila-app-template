package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Language holds the schema definition for the Language entity.
type Language struct {
	ent.Schema
}

func (Language) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "language"},
	}
}

// Fields of the Language.
func (Language) Fields() []ent.Field {
	return []ent.Field{
		// `language_id` tinyint UNSIGNED NOT NULL AUTO_INCREMENT,
		//  PRIMARY KEY (`language_id`)
		field.Uint8("id").
			StorageKey("language_id"),
		// `name` char(20) NOT NULL,
		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "char(20)",
		}),
		// `last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		field.Time("last_update").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			UpdateDefault(time.Now),
	}
}

// Edges of the Language.
func (Language) Edges() []ent.Edge {
	return nil
}
