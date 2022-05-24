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

// Film holds the schema definition for the Film entity.
type Film struct {
	ent.Schema
}

func (Film) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "film"},
	}
}

func (Film) Indexes() []ent.Index {
	return []ent.Index{
		// KEY `idx_title` (`title`),
		index.Fields("title").
			StorageKey("idx_title"),
		// KEY `idx_fk_language_id` (`language_id`),
		index.Fields("language_id").
			StorageKey("idx_fk_language_id"),
		// KEY `idx_fk_original_language_id` (`original_language_id`),
		index.Fields("original_language_id").
			StorageKey("idx_fk_original_language_id"),
	}
}

// Fields of the Film.
func (Film) Fields() []ent.Field {
	return []ent.Field{
		// `film_id` smallint UNSIGNED NOT NULL AUTO_INCREMENT,
		// PRIMARY KEY (`film_id`),
		field.Uint16("id").
			StorageKey("film_id"),
		// `title` varchar(128) NOT NULL,
		field.String("title").SchemaType(map[string]string{
			dialect.MySQL: "varchar(128)",
		}),
		// `description` text,
		field.Text("description").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),
		// `release_year` year DEFAULT NULL,
		field.Uint16("release_year").
			Nillable().
			Optional(),
		// `language_id` tinyint UNSIGNED NOT NULL,
		field.Uint8("language_id"),
		// `original_language_id` tinyint UNSIGNED DEFAULT NULL,
		field.Uint8("original_language_id").Nillable().Optional(),
		// `rental_duration` tinyint UNSIGNED NOT NULL DEFAULT '3',
		field.Uint8("rental_duration").
			Default(3),
		// `rental_rate` decimal(4, 2) NOT NULL DEFAULT '4.99',
		field.Float("rental_rate").SchemaType(map[string]string{
			dialect.MySQL: "decimal(4, 2)",
		}).
			Default(4.99),
		// `length` smallint UNSIGNED DEFAULT NULL,
		field.Uint16("length").Nillable().Optional(),
		// `replacement_cost` decimal(5, 2) NOT NULL DEFAULT '19.99',
		field.Float("replacement_cost").SchemaType(map[string]string{
			dialect.MySQL: "decimal(5, 2)",
		}).
			Default(19.99),
		// `rating` enum('G', 'PG', 'PG-13', 'R', 'NC-17') DEFAULT 'G',
		field.Enum("rating").
			Values("G", "PG", "PG-13", "R", "NC-17").
			Default("G"),
		// `special_features`
		// SET('Trailers', 'Commentaries', 'Deleted Scenes', 'Behind the Scenes') DEFAULT NULL,
		field.String("special_features").
			Nillable().
			Optional(),
		//  `last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		field.Time("last_update").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			UpdateDefault(time.Now),
	}
}

// Edges of the Film.
func (Film) Edges() []ent.Edge {
	return []ent.Edge{
		// CONSTRAINT `fk_film_language` FOREIGN KEY (`language_id`) REFERENCES `language` (`language_id`) ON DELETE RESTRICT,
		edge.To("language", Language.Type).
			Field("language_id").
			Unique().
			Required().
			StorageKey(edge.Symbol("fk_film_language")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
		// CONSTRAINT `fk_film_language_original` FOREIGN KEY (`original_language_id`) REFERENCES `language` (`language_id`) ON DELETE RESTRICT
		edge.To("original_language", Language.Type).
			Field("original_language_id").
			Unique().
			StorageKey(edge.Symbol("fk_film_language_original")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
		edge.To("actors", Actor.Type).StorageKey(edge.Table("film_actor")),
		edge.To("categories", Category.Type).StorageKey(edge.Table("film_category")),
	}
}
