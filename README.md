# languange-dictionary

# Documentation for Programming Language API

This API provides endpoints for managing programming languages data.

## Endpoints

### 1. Add Language
- **URL:** `/language`
- **Method:** `POST`
- **Description:** Adds a new programming language to the database.
- **Request Body:**
  ```json
  {
      "language": "Go",
      "appeared": 2009,
      "created": [
          "Robert Griesemer",
          "Rob Pike",
          "Ken Thompson"
      ],
      "functional": true,
      "object-oriented": false,
      "relation": {
          "influenced-by": [
              "C",
              "Python",
              "Pascal",
              "Smalltalk",
              "Modula"
          ],
          "influences": [
              "Odin",
              "Crystal",
              "Zig"
          ]
      }
  }
Response: Returns the added programming language with its ID.
## 2. Get Language by ID
- **URL:** `/language/:id`
- **Method:** `GET`
- **Description:** `Retrieves a programming language by its ID.`
- **Response:** `Returns the programming language data with the specified ID.`

## 3. Get All Languages
- **URL:** `/languages`
- **Method:** `GET`
- **Description:** `Retrieves all programming languages stored in the database.`
- **Response:** `Returns an array of all programming languages.`

## 4. Update Language
- **URL:** `/language/:id`
- **Method:** `PATCH`
- **Description:** `Updates an existing programming language by its ID.`
- **Request Body:** `Same as the "Add Language" endpoint.`
- **Response:** `Returns the updated programming language data.`

## 5. Delete Language
- **URL:** `/language/:id`
- **Method:** `DELETE`
- **Description:** `Deletes a programming language by its ID.`
- **Response:** `Returns a message confirming the deletion of the programming language.`

## 6. Check Palindrome
- **URL:** `/palindrome?text=<text>`
- **Method:** `GET`
- **Description:** `Checks if the given text is a palindrome.`
- **Query Parameter:** `text - The text to be checked.`
- **Response:** `Returns "Palindrome" if the text is a palindrome, otherwise returns "Not palindrome".`