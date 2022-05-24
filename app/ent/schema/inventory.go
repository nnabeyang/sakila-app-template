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

// Inventory holds the schema definition for the Inventory entity.
type Inventory struct {
	ent.Schema
}

func (Inventory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "inventory"},
	}
}

func (Inventory) Indexes() []ent.Index {
	return []ent.Index{
		// KEY `idx_fk_film_id` (`film_id`),
		index.Fields("film_id").
			StorageKey("idx_fk_film_id"),
		// KEY `idx_store_id_film_id` (`store_id`, `film_id`),
		index.Fields("store_id", "film_id").
			StorageKey("idx_store_id_film_id"),
	}
}

// Fields of the Inventory.
func (Inventory) Fields() []ent.Field {
	return []ent.Field{
		// `inventory_id` mediumint UNSIGNED NOT NULL AUTO_INCREMENT,
		// PRIMARY KEY (`inventory_id`),
		field.Uint32("id").SchemaType(map[string]string{
			dialect.MySQL: "mediumint unsigned",
		}).
			StorageKey("inventory_id"),
		// `film_id` smallint UNSIGNED NOT NULL,
		field.Uint16("film_id"),
		// `store_id` tinyint UNSIGNED NOT NULL,
		field.Uint8("store_id"),
		// `last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		field.Time("last_update").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			UpdateDefault(time.Now),
	}
}

// Edges of the Inventory.
func (Inventory) Edges() []ent.Edge {
	return []ent.Edge{
		// CONSTRAINT `fk_inventory_film` FOREIGN KEY (`film_id`) REFERENCES `film` (`film_id`) ON DELETE RESTRICT ON UPDATE CASCADE,
		edge.To("film", Film.Type).
			Field("film_id").
			Unique().
			Required().
			StorageKey(edge.Symbol("fk_inventory_film")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
		// CONSTRAINT `fk_inventory_store` FOREIGN KEY (`store_id`) REFERENCES `store` (`store_id`) ON DELETE RESTRICT ON UPDATE CASCADE
		edge.To("store", Store.Type).
			Field("store_id").
			Unique().
			Required().
			StorageKey(edge.Symbol("fk_inventory_store")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
	}
}
