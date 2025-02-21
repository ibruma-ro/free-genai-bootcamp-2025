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

- GET /words
- GET /words/:id
- Get /groups
- Get /groups/:id
- Get /groups/:id/words
- GET /dashboard/last_study_session
- GET /dashboard/study_progress
- GET /dashboard/quick_stats
- GET /study_activities/:id
- GET /study_activities/:id/study_sessions
- POST /study_activities
  - required params: group_id, study_activity_id
- GET /words
 - pagination with 100 items per page
- GET /groups
 - pagination with 100 items per page
- GET /groups/:id/study_sessions
- GET /study_sessions
 - pagination with 100 items per page

- GET /study_sessions/:id
- GET /study_sessions/:id/words

- POST /settings/reset_history
- POST /settings/full_reset
- POST /study_sessions/:id/words/:word_id/review
 - required params: correct


