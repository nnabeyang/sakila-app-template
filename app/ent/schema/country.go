package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Country holds the schema definition for the Country entity.
type Country struct {
	ent.Schema
}

func (Country) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "country"},
	}
}

// Fields of the Country.
func (Country) Fields() []ent.Field {
	return []ent.Field{
		// `country_id` smallint UNSIGNED NOT NULL AUTO_INCREMENT,
		//   PRIMARY KEY (`country_id`)
		field.Uint16("id").
			StorageKey("country_id"),
		//	`country` varchar(50) NOT NULL,
		field.String("country").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)",
		}),
		//	`last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		field.Time("last_update").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			UpdateDefault(time.Now),
	}
}

// Edges of the Country.
func (Country) Edges() []ent.Edge {
	return nil
}
