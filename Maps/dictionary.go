package Maps

//__________ BASIC WAY __________
/*
	func Search(dictionary map[string]string, word string) string {
		return dictionary[word]
	}*/

//__________ IMPROVED  WAY __________

type Dictionary map[string]string
type DictionaryError string

const (
	ErrorNotFound          = DictionaryError("We couln't find that word here")
	ErrorWordAlreadyExists = DictionaryError("We can't create a key that already exists")
	ErrorWordDoesNotExist  = DictionaryError("Can't update a word that doesn't exist")
)

func (e DictionaryError) Error() string {
	return string(e)
}

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, ok := dictionary[word]
	if !ok {
		return "", ErrorNotFound
	}
	return definition, nil
}

func (dictionary Dictionary) Add(word, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case ErrorNotFound:
		dictionary[word] = definition //assigning definition string to the word key
	case nil:
		return ErrorWordAlreadyExists
	default:
		return err
	}
	return nil
}

func (dictionary Dictionary) Update(word, definition string) error {
	//first call search to see if the word exists already
	_, err := dictionary.Search(word)

	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	case nil:
		dictionary[word] = definition
	default:
		return err
	}

	return nil
}

func (dictionary Dictionary) Delete(word string) error {

	_, err := dictionary.Search(word)

	switch err {
	case ErrorNotFound:
		return ErrorWordAlreadyExists
	case nil:
		delete(dictionary, word)
	default:
		return err
	}
	return nil
}
