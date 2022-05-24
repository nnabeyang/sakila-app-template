package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// City holds the schema definition for the City entity.
type City struct {
	ent.Schema
}

func (City) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "city"},
	}
}

func (City) Indexes() []ent.Index {
	return []ent.Index{
		//   KEY `idx_fk_country_id` (`country_id`),
		index.Fields("country_id").
			StorageKey("idx_fk_country_id"),
	}
}

// Fields of the City.
func (City) Fields() []ent.Field {
	return []ent.Field{
		//`city_id` smallint UNSIGNED NOT NULL AUTO_INCREMENT,
		//PRIMARY KEY (`city_id`),
		field.Uint16("id").
			StorageKey("city_id"),
		//`city` varchar(50) NOT NULL,
		field.String("city").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)",
		}),
		//`country_id` smallint UNSIGNED NOT NULL,
		field.Uint16("country_id"),
		//`last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		field.Time("last_update").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			UpdateDefault(time.Now),
	}
}

// Edges of the City.
func (City) Edges() []ent.Edge {
	return []ent.Edge{
		// CONSTRAINT `fk_city_country` FOREIGN KEY (`country_id`) REFERENCES `country` (`country_id`) ON DELETE RESTRICT ON UPDATE CASCADE
		edge.To("country", Country.Type).
			Field("country_id").
			Unique().
			Required().
			StorageKey(edge.Symbol("fk_city_country")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
	}
}
