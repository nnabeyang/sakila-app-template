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

// Rental holds the schema definition for the Rental entity.
type Rental struct {
	ent.Schema
}

func (Rental) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "rental"},
	}
}

func (Rental) Indexes() []ent.Index {
	return []ent.Index{
		// UNIQUE KEY `rental_date` (`rental_date`, `inventory_id`, `customer_id`),
		index.Fields("rental_date", "inventory_id", "customer_id").
			Unique().
			StorageKey("rental_date"),
		// KEY `idx_fk_inventory_id` (`inventory_id`),
		index.Fields("inventory_id").
			StorageKey("idx_fk_inventory_id"),
		// KEY `idx_fk_customer_id` (`customer_id`),
		index.Fields("customer_id").
			StorageKey("idx_fk_customer_id"),
		// KEY `idx_fk_staff_id` (`staff_id`),
		index.Fields("staff_id").
			StorageKey("idx_fk_staff_id"),
	}
}

// Fields of the Rental.
func (Rental) Fields() []ent.Field {
	return []ent.Field{
		// `rental_id` int NOT NULL AUTO_INCREMENT,
		// PRIMARY KEY (`rental_id`),
		field.Int32("id").
			StorageKey("rental_id"),
		//`rental_date` datetime NOT NULL,
		field.Time("rental_date").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		//`inventory_id` mediumint UNSIGNED NOT NULL,
		field.Uint32("inventory_id").SchemaType(map[string]string{
			dialect.MySQL: "mediumint unsigned",
		}),
		//`customer_id` smallint UNSIGNED NOT NULL,
		field.Uint16("customer_id"),
		//`return_date` datetime DEFAULT NULL,
		field.Time("return_date").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}).
			Nillable().
			Optional(),
		//`staff_id` tinyint UNSIGNED NOT NULL,
		field.Uint8("staff_id"),
		//`last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		field.Time("last_update").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			UpdateDefault(time.Now),
	}
}

// Edges of the Rental.
func (Rental) Edges() []ent.Edge {
	return []ent.Edge{
		// CONSTRAINT `fk_rental_customer` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`customer_id`) ON DELETE RESTRICT ON UPDATE CASCADE,
		edge.To("customer", Customer.Type).
			Field("customer_id").
			Unique().
			Required().
			StorageKey(edge.Symbol("fk_rental_customer")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
		// CONSTRAINT `fk_rental_inventory` FOREIGN KEY (`inventory_id`) REFERENCES `inventory` (`inventory_id`) ON DELETE RESTRICT ON UPDATE CASCADE,
		edge.To("inventory", Inventory.Type).
			Field("inventory_id").
			Unique().
			Required().
			StorageKey(edge.Symbol("fk_rental_inventory")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
		// CONSTRAINT `fk_rental_staff` FOREIGN KEY (`staff_id`) REFERENCES `staff` (`staff_id`) ON DELETE RESTRICT ON UPDATE CASCADE
		edge.To("staff", Staff.Type).
			Field("staff_id").
			Unique().
			Required().
			StorageKey(edge.Symbol("fk_rental_staff")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
	}
}
