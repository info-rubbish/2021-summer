package tokens

import (
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	store := NewStore(time.Second*5, time.Second*10, 10)
	defer store.Close()
	token, err := store.NewToken("paula")
	if err != nil {
		t.Fatal(err)
	}
	gctoken, err := store.NewToken("paula")
	if err != nil {
		t.Fatal(err)
	}
	data, err := store.GetToken(token)
	if err != nil {
		t.Fatal(err)
	}
	if data.(string) != "paula" {
		t.Fatal("not the same")
	}
	time.Sleep(time.Second * 6)
	if _, err := store.GetToken(token); err == nil {
		t.Fatal("should expired")
	}
	time.Sleep(time.Second * 6)
	if _, err := store.GetToken(gctoken); err == nil {
		t.Fatal("should gc")
	}
}

func TestRandomID(t *testing.T) {
	s := RandomID(10)
	if len(s) != 10 {
		t.Fatal("len should be 10")
	}
	t.Log(s)
}
