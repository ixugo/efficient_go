package orm

import (
	"context"
	"fmt"
	"testing"

	"entgo.io/ent/dialect/sql"
	"github.com/ixugo/efficient_go/demo/orm/ent"
	"github.com/ixugo/efficient_go/demo/orm/ent/user"
	_ "github.com/lib/pq"
)

func TestAABC(t *testing.T) {
	cli, err := ent.Open("postgres", "postgresql://postgres:123456789@127.0.0.1:1557/test?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	if err := cli.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	if err := cli.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	u, err := cli.User.Create().SetName("test").Save(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Name)

	u, err = cli.User.Query().Where(user.Name("test")).Order(user.ByID(sql.OrderDesc())).First(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(u.ID)

}
