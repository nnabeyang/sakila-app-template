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

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

func (Payment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "payment"},
	}
}

func (Payment) Indexes() []ent.Index {
	return []ent.Index{
		// KEY `idx_fk_staff_id` (`staff_id`),
		index.Fields("staff_id").
			StorageKey("idx_fk_staff_id"),
		// KEY `idx_fk_customer_id` (`customer_id`),
		index.Fields("customer_id").
			StorageKey("idx_fk_customer_id"),
		// KEY `fk_payment_rental` (`rental_id`),
		index.Fields("rental_id").
			StorageKey("fk_payment_rental"),
	}
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		//   `payment_id` smallint UNSIGNED NOT NULL AUTO_INCREMENT,
		//   PRIMARY KEY (`payment_id`),
		field.Uint16("id").
			StorageKey("payment_id"),
		//  `customer_id` smallint UNSIGNED NOT NULL,
		field.Uint16("customer_id"),
		//  `staff_id` tinyint UNSIGNED NOT NULL,
		field.Uint8("staff_id"),
		//  `rental_id` int DEFAULT NULL,
		field.Int32("rental_id").Nillable().Optional(),
		//  `amount` decimal(5, 2) NOT NULL,
		field.Float("amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(5, 2)",
		}),
		//  `payment_date` datetime NOT NULL,
		field.Time("payment_date").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		//  `last_update` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		field.Time("last_update").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			UpdateDefault(time.Now),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		// CONSTRAINT `fk_payment_customer` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`customer_id`) ON DELETE RESTRICT ON UPDATE CASCADE,
		edge.To("customer", Customer.Type).
			Field("customer_id").
			Unique().
			Required().
			StorageKey(edge.Symbol("fk_payment_customer")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
		// CONSTRAINT `fk_payment_rental` FOREIGN KEY (`rental_id`) REFERENCES `rental` (`rental_id`) ON DELETE SET NULL ON UPDATE CASCADE,
		edge.To("rental", Rental.Type).
			Field("rental_id").
			Unique().
			StorageKey(edge.Symbol("fk_payment_rental")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.SetNull,
			}),
		// CONSTRAINT `fk_payment_staff` FOREIGN KEY (`staff_id`) REFERENCES `staff` (`staff_id`) ON DELETE RESTRICT ON UPDATE CASCADE
		edge.To("staff", Staff.Type).
			Field("staff_id").
			Unique().
			Required().
			StorageKey(edge.Symbol("fk_payment_staff")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
	}
}
