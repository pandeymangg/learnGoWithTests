package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrExistingFound    = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return val, nil
}

func (d Dictionary) Add(k, v string) error {
	_, err := d.Search(k)

	switch err {
	case ErrNotFound:
		d[k] = v

	case nil:
		return ErrExistingFound

	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(k, v string) error {

	_, err := d.Search(k)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[k] = v
	default:
		return err

	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
