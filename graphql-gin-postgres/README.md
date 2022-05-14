#### APIs
1. Mutation query to create question
````
mutation {
  createQuestion(
    input: { id: "0001", question_text: "Q1. How are you ?", pub_date: "12-09-2022" }
  ) {
    id
    question_text
    pub_date
    choices {
      id
      question_id
      choice_text
    }
  }
}
````
Response

```
{
  "data": {
    "createQuestion": {
      "id": "0001",
      "question_text": "Q1. How are you ?",
      "pub_date": "12-09-2022",
      "choices": null
    }
  }
}
```