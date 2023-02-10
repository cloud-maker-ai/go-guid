package guid

import (
	"strconv"
	"testing"
	"github.com/google/uuid"
)

func TestGenerateUniqueString(t *testing.T) {
	prev := GenerateGuid("test")
	if prev != "d6a5baf1-8b9b-542a-8525-2982f1f98a0c" {
		t.Fatalf("GenerateGuid(\"test\") should = \"d6a5baf1-8b9b-542a-8525-2982f1f98a0c\"")
	}

	prev = GenerateGuid("test", "test2", "test3")
	if prev != "ba91b5ff-f126-51c5-a93f-d37f8292ee79" {
		t.Fatalf("GenerateGuid(\"test\", \"test2\", \"test3\") should = \"ba91b5ff-f126-51c5-a93f-d37f8292ee79\"")
	}

	for i := 0; i < 100; i++ {
		id := GenerateGuid("test", strconv.Itoa(i))
		if prev == id {
			t.Fatalf("Should get a new Guid!")
		}
		prev = id

		_, err := uuid.Parse(id)
		if err != nil {
			t.Fatalf("Bad Guid format %s %s", id, err)
		}
	}
}

func BenchmarkGenerateUniqueString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = GenerateGuid("test")
	}
}
