package schema

import (
	"database/sql/driver"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
)

// A Point consists of (X,Y) or (Lat, Lon) coordinates
// and it is stored in MySQL the POINT spatial data type.
type Point [2]float64

// Scan implements the Scanner interface.
func (p *Point) Scan(value interface{}) error {
	bin, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("invalid binary value for point")
	}
	var op orb.Point
	if err := wkb.Scanner(&op).Scan(bin[4:]); err != nil {
		return err
	}
	p[0], p[1] = op.X(), op.Y()
	return nil
}

// Value implements the driver Valuer interface.
func (p Point) Value() (driver.Value, error) {
	op := orb.Point{p[0], p[1]}
	return wkb.Value(op).Value()
}

// SchemaType defines the schema-type of the Point object.
func (Point) SchemaType() map[string]string {
	return map[string]string{
		dialect.MySQL: "geometry",
	}
}

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

func (Address) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "address", Collation: "utf8mb4_0900_ai_ci"},
	}
}

func (Address) Indexes() []ent.Index {
	return []ent.Index{
		//   KEY `idx_fk_city_id` (`city_id`),
		index.Fields("city_id").
			StorageKey("idx_fk_city_id"),
		//  SPATIAL KEY `idx_location` (`location`),
		index.Fields("location").
			Annotations(
				entsql.IndexTypes(map[string]string{
					dialect.MySQL: "SPATIAL",
				}),
			).
			StorageKey("idx_location"),
	}
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		//	`address_id` smallint UNSIGNED NOT NULL AUTO_INCREMENT,
		//	PRIMARY KEY (`address_id`),
		field.Uint16("id").
			StorageKey("address_id"),
		//	`address` varchar(50) NOT NULL,
		field.String("address").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)",
		}),
		//	`address2` varchar(50) DEFAULT NULL,
		field.String("address2").SchemaType(map[string]string{
			dialect.MySQL: "varchar(50)",
		}).
			Nillable().
			Optional(),
		//	`district` varchar(20) NOT NULL,
		field.String("district").SchemaType(map[string]string{
			dialect.MySQL: "varchar(20)",
		}),
		//	`city_id` smallint UNSIGNED NOT NULL,
		field.Uint16("city_id"),
		//	`postal_code` varchar(10) DEFAULT NULL,
		field.String("postal_code").SchemaType(map[string]string{
			dialect.MySQL: "varchar(10)",
		}).
			Nillable().
			Optional(),
		//	`phone` varchar(20) NOT NULL,
		field.String("phone").SchemaType(map[string]string{
			dialect.MySQL: "varchar(20)",
		}),
		//	`location` geometry NOT NULL /*!80003 SRID 0 */,
		field.Other("location", &Point{}).
			SchemaType(Point{}.SchemaType()),
		//	`last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		field.Time("last_update").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}).
			UpdateDefault(time.Now),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		//	CONSTRAINT `fk_address_city` FOREIGN KEY (`city_id`) REFERENCES `city` (`city_id`) ON DELETE RESTRICT ON UPDATE CASCADE
		edge.To("city", City.Type).
			Field("city_id").
			Unique().
			Required().
			StorageKey(edge.Symbol("fk_address_city")).
			Annotations(&entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
	}
}
