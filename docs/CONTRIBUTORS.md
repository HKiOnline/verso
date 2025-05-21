# CONTRIBUTORS

Thank you for considering contributing to the project! 🎉❤️

This project is a light weight hobby project with a low barrier for contributions. The repository's owner is committed to approach all contributions in good faith and respect. Ultimately, the repository owner maintains final say over what is included in the project.

Before you go too deep into the process, have a look at the [README-file](https://github.com/hkionline/verso.cli) to get a general idea of the project. And if you are conributing the first time, please read this document.

---

## How to

Here's a guide to the most common ways to contribute to this project. If you feel something is missing from the instructions, feel free to create an issue. If you are very new to coding, tools like git and contributing to project have a look at the Resources and Documentation -section below.

### Issues

An issue is where contributions start. Have a look at the "Issues" section of the repository and identify the issue you'd like to contribute to. **Contributions come in many forms**. They are not always code. A contribution can be as simple as a comment in an existing issue. If you however want to contribute code, there must be a issue for the contribution. If there isn't, you can create one. The **issue is then reviewed** by the author.

Make sure the description sufficiently **explains the issue**. At the moment there is no template for the issues.

**Not all issues are accepted** in the end. Don't feel bad about it. Your idea might have been good but it just didn't fit the idea of the project or the timing might have been wrong for instance. Issues which are not accepted will be closed.

**Issues should be labeled** for easier identification.
Issues labeled "help wanted" are free game for contributions. Comment in the issue, you'd like to work on the issue and get an acknowledgement from the repository owner. You can start working before the acknowledgement but its just good to have so others will know someone is already working on the issue and the repository owner is aware of that.

If you are looking for a simple issue to start with, look for the **"good first issue"** label. 

**Summary**:
- Anyone can create issues
- Make sure the issue describes the matter sufficiently
- Issues are reviewed by the repository owner or assigned contributor(s)
- Issues which are not accepted by the project are closed
- Issues should be labelled

### Pull Requests

New code, configurations and documentation are added to the project via [pull requests](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request). Pull request must be connected to an existing issue. This helps ensure that contributions are aligned with the project's goals and that discussions are centralized.

You'll work in your own branch and when your changes are completed you can make a pull request to the main branch of this repository.

If you add new code, it must include tests for the code and the existing tests must also pass. If the new code changes behaviour and existing tests are broken they must be fixed before a pull request can be merged.

In cases where the pull request does not add new code, adding tests is not needed unless the addition affects the code base in some way. All existing tests must still pass.  

Finally you must always update the CHANGELOG.md of the project. The format in the changelog is based on [Keep a Changelog -project](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

All pull requests will be reviewed by the repository owner or project maintainers. As part of the review, the reviewers may suggest or require changes before the pull request can beapproved and ultimately merged. An approval from the repository owner or a project maintainer is required before a merge.

Commits in the PR must be squashed when merged. The title of the merged commit message must be "Verso CLI {semver}", the name of the PR or the issue can be used in the body of the commit message.

**Summary**:
- PRs must be based on an issue
- PRs must include tests for new code
- All tests must pass
- Changelog must be updated with a new version
- PRs must be reviewed and approved before they can be merged
- PRs commits must be squashed when merged
- PRs commit message should use the template given above

---

## Conventions and Style

This section will describe the conventions, assumptions and the code style used in the project.

### GitHub Flow
Firstly, the project follows the [GitHub Flow](https://docs.github.com/en/get-started/using-github/github-flow) which is a simplified git branching technique. In principle this means there's only the main and the feature branches which are merged to the main branch via pull requests.

### Code Formatting and Style

Since this project is written exclusively in Go all code must be formatted using Go's own formatter utility [gofmt](https://pkg.go.dev/cmd/gofmt).
If your contribution is not code, please use a linter or a similar tool.

### Versioning

Everytime a new contribution is made, the version number of the project changes. It does not matter if the change is a code change or something else. The project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html). Changes are always documented to a [Changelog](https://keepachangelog.com/en/1.1.0/).

### Testing

All new contributions must be tested.

---

## Resources and Documentation

- [Finding ways to contribute to open source on GitHub](https://docs.github.com/en/get-started/exploring-projects-on-github/finding-ways-to-contribute-to-open-source-on-github)
- [Set up Git](https://docs.github.com/en/get-started/git-basics/set-up-git)
- [GitHub flow](https://docs.github.com/en/get-started/using-github/github-flow)
- [Collaborating with pull requests](https://docs.github.com/en/github/collaborating-with-pull-requests)
