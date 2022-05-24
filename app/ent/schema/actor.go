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

// Actor holds the schema definition for the Actor entity.
type Actor struct {
	ent.Schema
}

func (Actor) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "actor"},
	}
}

func (Actor) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("last_name").
			StorageKey("idx_actor_last_name"),
	}
}

// Fields of the Actor.
func (Actor) Fields() []ent.Field {
	return []ent.Field{
		//`actor_id` smallint UNSIGNED NOT NULL AUTO_INCREMENT,
		//PRIMARY KEY (`actor_id`),
		field.Uint16("id").
			StorageKey("actor_id"),
		//`first_name` varchar(45) NOT NULL,
		field.String("first_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(45)",
		}),
		//`last_name` varchar(45) NOT NULL,
		field.String("last_name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(45)",
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

// Edges of the Actor.
func (Actor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("films", Film.Type).Ref("actors"),
	}
}
