package cmdbox_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox"
)

func ExampleNewCommandMap() {
	m := cmdbox.NewCommandMap()
	m.Print()
	// Output:
	// {}
}

func ExampleCommandMap_Set() {
	m := cmdbox.NewCommandMap()
	c := cmdbox.NewCommand("foo")

	m.Set("foo", c)
	m.Print()

	// Output:
	// {
	//     "foo": {
	//       "name": "foo",
	//       "commands": {}
	//     }
	//   }

}

func ExampleCommandMap_Get() {
	m := cmdbox.NewCommandMap()
	c := cmdbox.NewCommand("foo")

	m.Set("foo", c)
	m.Get("foo").Print()

	// Output:
	// {
	//     "name": "foo",
	//     "commands": {}
	//   }

}

func ExampleCommandMap_Init() {
	m := cmdbox.NewCommandMap()
	c := cmdbox.NewCommand("foo")

	m.Set("foo", c)
	m.Print()
	m.Init()
	m.Print()

	// Output:
	// {
	//     "foo": {
	//       "name": "foo",
	//       "commands": {}
	//     }
	//   }
	// {}

}

func ExampleCommandMap_Delete() {
	m := cmdbox.NewCommandMap()
	c := cmdbox.NewCommand("foo")

	m.Set("foo", c)
	m.Print()
	m.Delete("foo")
	m.Print()

	// Output:
	// {
	//     "foo": {
	//       "name": "foo",
	//       "commands": {}
	//     }
	//   }
	// {}

}

func ExampleCommandMap_Rename() {
	m := cmdbox.NewCommandMap()
	c := cmdbox.NewCommand("foo")

	m.Set("foo", c)
	m.Print()
	m.Rename("foo", "bar")
	m.Print()

	// Output:
	// {
	//     "foo": {
	//       "name": "foo",
	//       "commands": {}
	//     }
	//   }
	// {
	//     "bar": {
	//       "name": "bar",
	//       "commands": {}
	//     }
	//   }

}

func ExampleCommandMap_Names() {
	m := cmdbox.NewCommandMap()
	c1 := cmdbox.NewCommand("foo")
	c2 := cmdbox.NewCommand("bar")

	m.Set("foo", c1)
	m.Set("bar", c2)
	fmt.Println(m.Names())

	// Output:
	// [bar foo]

}

func ExampleCommandMap_Dups() {
	m := cmdbox.NewCommandMap()
	c1 := cmdbox.NewCommand("foo")
	c2 := cmdbox.NewCommand("foo_")
	c3 := cmdbox.NewCommand("bar")
	c4 := cmdbox.NewCommand("boo")

	m.Set("foo", c1)
	m.Set("foo_", c2)
	m.Set("bar", c3)
	m.Set("boo", c4)
	sl := m.Slice("boo", "bar")
	fmt.Println(len(sl))
	fmt.Println(sl[0].Name)
	fmt.Println(sl[1].Name)

	// Output:
	// 2
	// boo
	// bar

}

func ExampleCommandMap_RawJSON() {
	m := cmdbox.NewCommandMap()
	c1 := cmdbox.NewCommand("foo")
	m.Set("foo", c1)
	fmt.Println(m.RawJSON())
	// Output:
	// {"foo":{"name":"foo","commands":{}}}
}

func ExampleCommandMap_JSON() {
	m := cmdbox.NewCommandMap()
	c1 := cmdbox.NewCommand("foo")
	m.Set("foo", c1)
	fmt.Println(m.JSON())
	// Output:
	// {
	//     "foo": {
	//       "name": "foo",
	//       "commands": {}
	//     }
	//   }
}

func ExampleCommandMap_String() {
	m := cmdbox.NewCommandMap()
	c1 := cmdbox.NewCommand("foo")
	m.Set("foo", c1)
	fmt.Println(m.JSON())
	fmt.Println(m)
	// Output:
	// {
	//     "foo": {
	//       "name": "foo",
	//       "commands": {}
	//     }
	//   }
	// {
	//     "foo": {
	//       "name": "foo",
	//       "commands": {}
	//     }
	//   }
}
