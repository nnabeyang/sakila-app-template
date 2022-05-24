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

// Staff holds the schema definition for the Staff entity.
type Staff struct {
	ent.Schema
}

func (Staff) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "staff"},
	}
}

func (Staff) Indexes() []ent.Index {
	return []ent.Index{
		// KEY `idx_fk_store_id` (`store_id`),
		index.Fields("store_id").
			StorageKey("idx_fk_store_id"),
		// KEY `idx_fk_address_id` (`address_id`),
		index.Fields("address_id").
			StorageKey("idx_fk_address_id"),
	}
}

// Fields of the Staff.
func (Staff) Fields() []ent.Field {
	return []ent.Field{
		//   `staff_id` tinyint UNSIGNED NOT NULL AUTO_INCREMENT,
		field.Uint8("id").
			StorageKey("staff_id"),
		//	`first_name` varchar(45) NOT NULL,
		field.String("first_name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(45)",
			}),
		//	`last_name` varchar(45) NOT NULL,
		field.String("last_name").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(45)",
			}),
		//	`address_id` smallint UNSIGNED NOT NULL,
		field.Uint16("address_id"),
		//	`picture` blob,
		field.Bytes("picture").Nillable().Optional(),
		//	`email` varchar(50) DEFAULT NULL,
		field.String("email").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(50)",
			}).
			Nillable().
			Optional(),
		//	`store_id` tinyint UNSIGNED NOT NULL,
		field.Uint8("store_id"),
		//	`active` tinyint(1) NOT NULL DEFAULT '1',
		field.Bool("active").
			Default(true),
		//	`username` varchar(16) NOT NULL,
		field.String("username").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(16)",
			}),
		//	`password` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
		field.String("password").
			SchemaType(map[string]string{
				dialect.MySQL: "varchar(40)",
			}).
			Nillable().
			Optional(),
		//	`last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		field.Time("last_update").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			UpdateDefault(time.Now),
	}
}

// Edges of the Staff.
func (Staff) Edges() []ent.Edge {
	return []ent.Edge{
		// CONSTRAINT `fk_staff_address` FOREIGN KEY (`address_id`) REFERENCES `address` (`address_id`) ON DELETE RESTRICT ON UPDATE CASCADE,
		edge.To("address", Address.Type).
			Field("address_id").
			Unique().
			Required().
			StorageKey(edge.Symbol("fk_staff_address")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
		// CONSTRAINT `fk_staff_store` FOREIGN KEY (`store_id`) REFERENCES `store` (`store_id`) ON DELETE RESTRICT ON UPDATE CASCADE
		edge.To("store", Store.Type).
			Field("store_id").
			Unique().
			Required().
			StorageKey(edge.Symbol("fk_staff_store")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
	}
}
