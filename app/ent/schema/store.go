package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Store holds the schema definition for the Store entity.
type Store struct {
	ent.Schema
}

func (Store) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "store"},
	}
}

func (Store) Indexes() []ent.Index {
	return []ent.Index{
		// UNIQUE KEY `idx_unique_manager` (`manager_staff_id`),
		index.Fields("manager_staff_id").
			Unique().
			StorageKey("idx_unique_manager"),
		// KEY `idx_fk_address_id` (`address_id`),
		index.Fields("address_id").
			StorageKey("idx_fk_address_id"),
	}
}

// Fields of the Store.
func (Store) Fields() []ent.Field {
	return []ent.Field{
		// `store_id` tinyint unsigned NOT NULL AUTO_INCREMENT,
		// PRIMARY KEY (`store_id`),
		field.Uint8("id").
			StorageKey("store_id"),
		//	`manager_staff_id` tinyint unsigned NOT NULL,
		field.Uint8("manager_staff_id"),
		//	`address_id` smallint unsigned NOT NULL,
		field.Uint16("address_id"),
		//	`last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		field.Time("last_update").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			UpdateDefault(time.Now),
	}
}

// Edges of the Store.
func (Store) Edges() []ent.Edge {
	return []ent.Edge{
		//  CONSTRAINT `fk_store_address` FOREIGN KEY (`address_id`) REFERENCES `address` (`address_id`) ON DELETE RESTRICT ON UPDATE CASCADE,
		edge.To("address", Address.Type).
			Field("address_id").
			Unique().
			Required().
			StorageKey(edge.Symbol("fk_store_address")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
		//  CONSTRAINT `fk_store_staff` FOREIGN KEY (`manager_staff_id`) REFERENCES `staff` (`staff_id`) ON DELETE RESTRICT ON UPDATE CASCADE
		edge.To("staff", Staff.Type).
			Field("manager_staff_id").
			Unique().
			Required().
			StorageKey(edge.Symbol("fk_store_staff")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
	}
}
