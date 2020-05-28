package proxy

import (
	"math/rand"
	"testing"
)

func TestUserListProxy(t *testing.T) {
	someDatabase := UserList{}
	rand.Seed(2342342)
	for i := 0; i < 1000000; i++ {
		n := rand.Int()
		someDatabase = append(someDatabase, &User{ID: n})
	}

	proxy := UserListProxy{
		SomeDatabase:  someDatabase,
		StackCache:    UserList{},
		StackCapacity: 2,
	}

	knownIds := [3]int{someDatabase[3].ID, someDatabase[4].ID, someDatabase[5].ID}
	t.Run("FindUser - Empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIds[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knownIds[0] {
			t.Error("user not match")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("must be one", proxy.StackCache)
		}

		if proxy.DidLastSearchUsedCache {
			t.Error("must be false")
		}
	})

	t.Run("FindUser - One user, ask for the same user", func(t *testing.T) {
		user, err := proxy.FindUser(knownIds[0])
		if err != nil {
			t.Fatal(err)
		}
		if user.ID != knownIds[0] {
			t.Error("user not match")
		}
		if len(proxy.StackCache) != 1 {
			t.Error("must be one", proxy.StackCache)
		}

		if !proxy.DidLastSearchUsedCache {
			t.Error("must be true")
		}

	})

	user1, err := proxy.FindUser(knownIds[0])
	if err != nil {
		t.Fatal(err)
	}
	user2, _:= proxy.FindUser(knownIds[1])
	if proxy.DidLastSearchUsedCache {
		t.Error("should be false")
	}

	user3, _:= proxy.FindUser(knownIds[2])
	if proxy.DidLastSearchUsedCache {
		t.Error("should be false")
	}

	for _, cache := range proxy.StackCache {
		if cache.ID == user1.ID {
			t.Error("should not be found")
		}
	}

	if len(proxy.StackCache) != 2 {
		t.Error("capacity is 2, expect 2, got:", len(proxy.StackCache))
	}

	for _, cache := range proxy.StackCache {
		if cache != user2 && cache != user3 {
			t.Error("found a non expected user", cache.ID)
		}
	}

	t.Log(proxy.StackCache)
}
