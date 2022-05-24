package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// FilmText holds the schema definition for the FilmText entity.
type FilmText struct {
	ent.Schema
}

func (FilmText) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "film_text"},
	}
}

// func (FilmText) Indexes() []ent.Index {
//	return []ent.Index{
//		// FULLTEXT KEY `idx_title_description` (`title`, `description`)
//		index.Fields("title", "description").
//			Annotations(
//				entsql.IndexType("FULLTEXT"),
//			).
//			StorageKey("idx_title_description"),
//	}
// }

// Fields of the FilmText.
func (FilmText) Fields() []ent.Field {
	return []ent.Field{
		// `film_id` smallint NOT NULL,
		// PRIMARY KEY (`film_id`),
		field.Int16("id").
			StorageKey("film_id"),
		// `title` varchar(255) NOT NULL,
		field.String("title"),
		// `description` text,
		field.Text("description").SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),
	}
}

// Edges of the FilmText.
func (FilmText) Edges() []ent.Edge {
	return nil
}
