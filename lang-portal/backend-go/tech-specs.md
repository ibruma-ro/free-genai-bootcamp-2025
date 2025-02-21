## Backend server tech specs

## Business Goal

-A language learning school wants to build a prototype of learning portal which will act as three things:
-Inventory of possible vocabulary that can be learned
-Act as a  Learning record store (LRS), providing correct and wrong score on practice vocabulary
-A unified launchpad to launch different learning apps

## Technical requirements

- the backend will be written in Go
- the database will be SQLite3
- the API will be built using Gin framework
- the API will allways return JSON
- there will bo no authentication and authorization
- everything treated as a single user


## database schema

Database schema will be called "words_db" that will be in the root of the project folder of the backend-go.	

We have the following tables:
- words -> list of words
	- id: integer
	- french: string
	- english: string
	- parts: json
- word groups -> join table between words and groups
	- id: integer
	- word_id: integer
	- group_id: integer
- groups -> thematic groups of words
	- id: integer
	- name: string
- study sessions -> record of a study session grouping word review items
	- id: integer
	- group_id: integer
    - created_at: timestamp
    - study_activity_id: integer
- study activities -> links goups to study sessions
    - id: integer
    - study_session_id: integer
    - group_id: integer
    - created_at: timestamp
- word review items -> a record of word practice deterining if the correct is k or not
    - word_id: integer
    - study_session_id: integer
    - correct: boolean
    - created_at: timestamp
    
## APi Endpoints

- GET /api/words
```json
{
  "items": [
    {
      "french": "bonjour",
      "english": "hello",
      "correct_count": 25,
      "wrong_count": 5
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 10,
    "items_per_page": 100,
    "total_items": 950
  }
}
```

- GET /api/words/:id
```json
{
  "id": 1,
  "french": "bonjour",
  "english": "hello",
  "parts": {
    "pronunciation": "bɔ̃.ʒuʁ",
    "part_of_speech": "interjection"
  },
  "correct_count": 25,
  "wrong_count": 5,
  "groups": [
    {
      "id": 1,
      "name": "Greetings"
    }
  ]
}

```

- Get /api/groups
```json
{
  "items": [
    {
      "id": 1,
      "name": "Greetings",
      "word_count": 25
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "items_per_page": 100,
    "total_items": 450
  }
}

- Get /api/groups/:id
```json
{
"id":1,
"name":"Greetings",
"stats":{
  "total_word_count":25
}
}
````
- Get /api/groups/:id/words
```json
{
  "items": [
    {
      "id": 1,
      "french": "bonjour",
      "english": "hello",
      "parts": {
        "pronunciation": "bɔ̃.ʒuʁ"
      }
    },
    {
      "id": 2,
      "french": "au revoir",
      "english": "goodbye",
      "parts": {
        "pronunciation": "o ʁə.vwaʁ"
      }
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "items_per_page": 100,
    "total_items": 2
  }
}
```
- GET /dashboard/last_study_session
```json
{
  "id": 123,
  "group_id": 1,
  "group_name": "Greetings",
  "start_time": "2024-03-20T14:30:00Z",
  "study_activity_id": 1,
}

```
- GET /dashboard/study_progress
```json
{
  "total_words_studied": 500,
  "total_available_words": 1000,
}

```
- GET /dashboard/quick_stats
```json
{
	"study_streak_days":4
}

```
- GET /study_activities/:id
```json
{
  "id": 1,
  "name": "Vocabulary Quiz",
  "description": "Practice your vocabulary",
  "thumbnail_url": "https://example.com/thumbnails/vocab-quiz.jpg",
}
````
- GET /study_activities/:id/study_sessions
```json
{
  "items": [
    {
      "id": 123,
      "activity_name": "Vocabulary Quiz",
      "start_time": "2024-03-20T14:30:00Z",
      "end_time": "2024-03-20T15:00:00Z",
      "group_name": "Greetings",
      "words_reviewed": 25
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "items_per_page": 100,
    "total_items": 450
  }
}
```
- POST /api/study_activities
  - required params: group_id, study_activity_id
  - returns:
  ```json
  {
    "id": 123,
    "group_id": 1
  }
  ```

- GET /api/groups/:id/study_sessions
```json
{
  "group_id": 1,
  "group_name": "Greetings",
  "study_sessions": [
    {
      "id": 123,
      "activity_name": "Vocabulary Quiz",
      "start_time": "2024-03-20T14:30:00Z",
      "end_time": "2024-03-20T15:00:00Z",
      "words_reviewed": 25
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 3,
    "items_per_page": 100,
    "total_items": 250
  }
}
```

- GET /api/study_sessions
```json
{
  "items": [
    {
      "id": 123,
      "activity_name": "Vocabulary Quiz",
      "group_name": "Greetings",
      "start_time": "2024-03-20T14:30:00Z",
      "end_time": "2024-03-20T15:00:00Z",
      "words_reviewed": 25
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 5,
    "items_per_page": 100,
    "total_items": 450
  }
}
```

- GET /api/study_sessions/:id
```json
{
  "id": 123,
  "activity_name": "Vocabulary Quiz",
  "group_name": "Greetings",
  "start_time": "2024-03-20T14:30:00Z",
  "end_time": "2024-03-20T15:00:00Z",
  "words_reviewed": 25
}
```

- GET /api/study_sessions/:id/words
```json
{
  "study_session_id": 123,
  "activity_name": "Vocabulary Quiz",
  "words": [
    {
      "french": "bonjour",
      "english": "hello",
      "correct_count": 25,
      "wrong_count": 5
    },
    {
      "french": "au revoir",
      "english": "goodbye",
      "correct_count": 15,
      "wrong_count": 10
    }
  ],
  "pagination": {
    "current_page": 1,
    "total_pages": 1,
    "items_per_page": 100,
    "total_items": 2
  }
}
```

- POST /api/settings/reset_history
```json
{
  "success": true,
  "message": "Study history has been reset",
  "deleted_records": {
    "study_sessions": 45,
    "word_reviews": 1250
  }
}
```
- POST /api/settings/full_reset
```json
{
  "success": true,
  "message": "Study history has been reset",
  "deleted_records": {
    "study_sessions": 45,
    "word_reviews": 1250
  }
}
````

- POST /api/study_sessions/:id/words/:word_id/review
 - required params: correct, id (study_session_id), word_id
 - returns:
 ```json
 {
  "success": true,
  "message": "Word review has been saved"
 }
 ```

 ## Mage tasks 

 mage is a task runner for Go

 ### Initialize database

 This will initialize the sqlite database called "words_db" in the root of the project folder of the backend-go.

 ### migrate database

 Will run a series of migration sql files o the database

 ### Seed data
 This task will import json files and transofrm them into target data for our database	

All seed files live in the seed folder
All seeds file should be loaded
IN our task we should have a DSL to specify each seed files and its expected  word group name
The file names should look like this
```sql
0001_init.sql
002_create_words_table.sql
```

 ````json
 [
  {
    "id": 1,
    "french": "bonjour",
    "english": "hello"
  },
  {
    "id": 2,
    "french": "au revoir",
    "english": "goodbye"
  }
 ]


