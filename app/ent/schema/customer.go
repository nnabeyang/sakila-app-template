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

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

func (Customer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "customer"},
	}
}

func (Customer) Indexes() []ent.Index {
	return []ent.Index{
		//KEY `idx_fk_store_id` (`store_id`),
		index.Fields("store_id").
			StorageKey("idx_fk_store_id"),
		//KEY `idx_fk_address_id` (`address_id`),
		index.Fields("address_id").
			StorageKey("idx_fk_address_id"),
		//KEY `idx_last_name` (`last_name`),
		index.Fields("last_name").
			StorageKey("idx_last_name"),
	}
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		// `customer_id` smallint unsigned NOT NULL AUTO_INCREMENT,
		// PRIMARY KEY (`customer_id`),
		field.Uint16("id").
			StorageKey("customer_id"),
		//      `store_id` tinyint unsigned NOT NULL,
		field.Uint8("store_id"),
		// `first_name` varchar(45) NOT NULL,
		field.String("first_name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(45)",
			}),
		// `last_name` varchar(45) NOT NULL,
		field.String("last_name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(45)",
			}),
		// `email` varchar(50) DEFAULT NULL,
		field.String("email").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(50)",
			}).
			Nillable().
			Optional(),
		//  `address_id` smallint unsigned NOT NULL,
		field.Uint16("address_id"),
		// `active` tinyint(1) NOT NULL DEFAULT '1',
		field.Bool("active").
			Default(true),
		// `create_date` datetime NOT NULL,
		field.Time("create_date").SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		// `last_update` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		field.Time("last_update").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			UpdateDefault(time.Now),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		// CONSTRAINT `fk_customer_address` FOREIGN KEY (`address_id`) REFERENCES `address` (`address_id`) ON DELETE RESTRICT ON UPDATE CASCADE,
		edge.To("address", Address.Type).
			Field("address_id").
			Unique().
			Required().
			StorageKey(edge.Symbol("fk_customer_address")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
		// CONSTRAINT `fk_customer_store` FOREIGN KEY (`store_id`) REFERENCES `store` (`store_id`) ON DELETE RESTRICT ON UPDATE CASCADE
		edge.To("store", Store.Type).
			Field("store_id").
			Unique().
			Required().
			StorageKey(edge.Symbol("fk_customer_store")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
	}
}
