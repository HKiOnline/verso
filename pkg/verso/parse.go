package verso

// Parse semantic versions from a CHANGELOG file.
// Takes in the file path as parameter.
// Returns a Changelog struct with a slice of versions and other metadata.
func Parse(filepath string) (Changelog, error) {

	fileBytes, err := getFile(filepath)

	if err != nil {
		return Changelog{}, err
	}

	return extract(string(fileBytes))

}
