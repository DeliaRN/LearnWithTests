package Maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is an improved test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is an improved test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertError(t, got, ErrorNotFound)
	})
}

//DO NEVER INITIALIZE A NIL MAP VARIABLE, as it works as a reference type,
//as their values are pointers to runtime.hmap structures,
//and trying to access a nil map will cause a runtime error
//	*NO* 	var m map[string]string
//	*YES* 	var dictionary = map[string]string{}
//	*YES* 	var dictionary = make(map[string]string)
//these will create an empty hash map and point dictionary at it to ensure
//you don't get a runtime error

func TestAdd(t *testing.T) {
	t.Run("Testing for new words", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Testing for adding existing words", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrorWordAlreadyExists)
		assertDefinition(t, dictionary, word, definition)

	})
}

func TestUpdate(t *testing.T) {

	t.Run("Existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("New word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrorWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Deleting existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}
		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assertError(t, err, ErrorNotFound)
	})

	t.Run("Deleting non existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word) //it will fail because it's an empty dictionary

		assertError(t, err, ErrorWordDoesNotExist)
	})
}

/* ______________________ Auxiliar functions ______________________ */

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("it's supposed to find the added word.", err)
	}
	assertStrings(t, got, definition)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
