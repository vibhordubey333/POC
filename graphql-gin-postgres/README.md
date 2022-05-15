##Introduction

POC to demonstrate fetching of data using `Graphql`,`Golang` and `Postgresql`.

## Pre-requisite
0. Go
1. Postgresql
2. Make [For windows users]

## Run

0. Go to path `POC/graphql-gin-postgres/graphql_poc/api`
1. Fire `make run`

## APIs
>http://localhost:8888/query

0. Mutation query to create question
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
1. Fetch question list:
```
{
  question {
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
```
Response
```
{
  "data": {
    "question": [
      {
        "id": "0001",
        "question_text": "Q1. How are you ?",
        "pub_date": "12-09-2022",
        "choices": []
      },
      {
        "id": "0002",
        "question_text": "Q2. How are you ?",
        "pub_date": "12-09-2022",
        "choices": []
      }
    ]
  }
}
```
2. Mutation to create choices.
```
mutation {
  createChoice(
    input: { question_id: "0001", choice_text: "Choice_1", choice_id: "C_001" }
  ) {
    id
    question {
      id
      question_text
      pub_date
      choices {
        id
        question {
          id
          question_text
          pub_date
    
        }
        question_id
        choice_text
      }
    }
    question_id
    choice_text
  }
}
```
Response:
```
{
  "data": {
    "createChoice": {
      "id": "C_001",
      "question": {
        "id": "0001",
        "question_text": "Q1. How are you ?",
        "pub_date": "12-09-2022",
        "choices": null
      },
      "question_id": "0001",
      "choice_text": "Choice_1"
    }
  }
}
```
3. Query fetch all choices.
```
{
  choices {
    id
    question_id
    choice_text
  }
}
```
Response
```
{
  "data": {
    "choices": [
      {
        "id": "C_001",
        "question_id": "0001",
        "choice_text": "Choice_1"
      }
    ]
  }
}
```
Reference:
> https://www.agiliq.com/blog/2020/04/graphql-api-server-using-golang-gin-framework/