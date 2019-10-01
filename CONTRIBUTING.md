# Contributing

## How to Contribute

This repository is built for the purpose of creating a public collection of books. All changes are considered, as long as they are valid and follow the guidelines.

That said, suggested ways to contribute include:

* Fork the project in GitHub
* Clone your forked repo locally
* Add a file named `<book-name.json>` to the `books` directory. Make sure to include all details mentioned in the table below (a template has also been provided for reference).
* Commit and push to your forked repository
* Create a PR from your fork to the main repo in GitHub

### Book Details


Attribute | Details
----------------|--------------
Title | Title of the book
Storyline | Brief description of the book
Authors | Names of book author(s)
Genres | Related Genres
Characters | Main characters in the book
Pages | Number of pages in the book
ISBN | ISBN code of the book
Language | Orginal language of publication
Publisher | Names of book's publisher
Publication Date | Date of first publication

### Sample book JSON

```json
{
  "title": "The Outsider",
  "storyline": "An eleven-year-old boy’s violated corpse is found in a town park. Eyewitnesses and fingerprints point unmistakably to one of Flint City’s most popular citizens. He is Terry Maitland, Little League coach, English teacher, husband, and father of two girls. Detective Ralph Anderson, whose son Maitland once coached, orders a quick and very public arrest. Maitland has an alibi, but Anderson and the district attorney soon add DNA evidence to go with the fingerprints and witnesses. Their case seems ironclad. As the investigation expands and horrifying answers begin to emerge, King’s propulsive story kicks into high gear, generating strong tension and almost unbearable suspense. Terry Maitland seems like a nice guy, but is he wearing another face? When the answer comes, it will shock you as only Stephen King can.",
  "authors": ["Stephen King"],
  "genres": [
    "horror",
    "mystery",
    "fiction",
    "thriller"
  ],
  "characters": ["Holly Gibney"],
  "pages": 561,
  "ISBN": "1501180983",
  "language": "english",
  "publisher": "Scribner",
  "publication_date": "2018-05-22"
}
```

## Code of Conduct

Please note that this project is released with a [Contributor Code of Conduct](CODE_OF_CONDUCT.md). By participating in this project you agree to abide by its terms.
