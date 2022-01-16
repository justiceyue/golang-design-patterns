package flyweight

import (
	"sync"
)

type player struct {
	identity string  // CT or T
	arsenal  arsenal //武器库
}

// CT和T的武器都是固定的，因此可以独立出来
type arsenal struct {
	gun []string
}

type arsenalFactory struct {
	arsenalMap map[string]arsenal
}

var (
	arsenalFactoryInstance *arsenalFactory
	once                   sync.Once
	ctGun                  = []string{"m4a4", "aug"}
	tGun                   = []string{"ak-47", "sg-553"}
)

func arsenalGetter() func(identity string) arsenal {
	once.Do(func() {
		arsenalFactoryInstance = &arsenalFactory{
			arsenalMap: map[string]arsenal{
				"CT": {
					gun: ctGun,
				},
				"T": {
					gun: tGun,
				},
			},
		}
	})
	return func(identity string) arsenal {
		return arsenalFactoryInstance.arsenalMap[identity]
	}
}
