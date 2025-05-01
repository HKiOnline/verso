package verso

// Parse semantic versions from a CHANGELOG file in a filepath.
// Takes in the file path as parameter.
// Returns a Changelog struct with a slice of versions and other metadata.
func ParsePath(filepath string) (Changelog, error) {

	fileBytes, err := getFile(filepath)

	if err != nil {
		return Changelog{}, err
	}

	return extract(string(fileBytes))

}

// Parse semantic versions from a CHANGELOG file in a filepath.
// Takes in the file path as parameter.
// Returns a Changelog struct with a slice of versions and other metadata.
func ParseBytes(bytes []byte) (Changelog, error) {

	return extract(string(bytes))

}
