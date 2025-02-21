# Frontend Technical Specifications

## Pages

### Dashboard
The purpose of this page is to give the user a quick overview of their learning progress.It will be the default page that loads when the user logs in the app.


- Last Study Session
- Study Progress
    - total words studied
    - display mastery percentage
- Quick Stats
    - success rate
    - total study sessions
    - total active groups
    - study streak
- Start Studying button
    - goes to study activity page

### Endpoints

- GET /dashboard/last_study_session
- GET /dashboard/study_progress
- GET /dashboard/quick_stats

### Study Activity

The purpose of this page is to show a collection of study activities with a thumbnail and its title.

#### Components
The page contains the following components:
- Study Activity Card
    - Thumbnail
    - Name
    - Laucn button
    - view page to view more information about past study sessions
    - Start Studying button

#### Needed API endpoints

- GET /study_activities


### Study Activity Show 

#### Purpose

The purpose of this page is to show the details of a study activity.

#### Components
- Name of study activity
- Thumbnail of study activity
- Description of study activity
- Launch button
- Study activity Paginated List
    - id
    - activity Name
    - group Name
    - start time
    - end time (inffered by the last word_review_item submitted)
    - number of review items

#### Needed endpoints

- GET /study_activities/:id
- GET /study_activities/:id/study_sessions

### Study Activities Launch

#### Purpose

The purpose of this page is to launch a study activity.

#### Components

- name of the study activity
- launch form
    - select field for group
    - launch now button
## behavior

- when the form is submitted, a new tab opens with the study activity based on its URL provided in the database.
 - Also when the form is submitted the page will redirect to the study session show page

#### Needed endpoints
- POST /study_activities/:id/launch


### Words

#### Purpose

The purpose of this page is to show all words  in our database

#### Components

- paginated word list
    - Fields
      - French
      - Pronunciation
      - English
      - Correct count
      - Wrong count
    - Pagination with 100 items per page
    - Clicking the french work will take us to the word show page    

#### Needed endpoints

- GET /words

### Word Show

#### Purpose

The purpose of this page is to show the details of a word.

#### Components

- French
- Pronunciation
- Study statistics
    - Correct count
    - Wrong count
- Word groups
    - show a series of pills eg tags
    - when a word is clicked it will take us to the word show page

#### Needed endpoints

- GET /words/:id

### Word Groups

#### Purpose

The purpose of this page is to show all word groups in our database

#### Components

- paginated word group list
    - Fields
      - Name
      - Word count
    - Pagination with 100 items per page
    - Clicking the name of the word group will take us to the word group show page

 #### Needed endpoints

 - GET /groups

### Word Group Show

#### Purpose

The purpose of this page is to show the details of a word group.

#### Components
- Group Name
- Group statistics
 - Word count
- Words in group (paginated list of words)
 - should use the same component as the word index page

 - Study sessions (paginated list of study sessions)
  - should use the same component as the study activity index page

#### Needed endpoints

- GET /groups/:id (the name and group stats)
- GET /groups/:id/words
- GET /groups/:id/study_sessions

## Study session Index

#### Purpose

The purpose of this page is to show all study sessions in our database

#### Components

- paginated study session list
    - Fields
      - ID
      - Activity name
      - Group Name
      - Start time
      - End time
      - number of words reviewed

#### Needed endpoints

- GET /study_sessions


## Study Session Show

#### Purpose

The purpose of this page is to show the details of a study session.

#### Components

- Study Session Details
    - Activity name
    - Group Name
    - Start time
    - End time
    - number of words reviewed
- Words review items (paginated list of words review items)
    - should use the same component as the word show page

#### Needed endpoints

- GET /study_sessions/:id
- GET /study_sessions/:id/words

### Settings

#### Purpose

The purpose of this page is to show the settings of the user.

#### Components
- Theme Selection
- Language Selection
- reset history - delete all study session and word review items
- Load Seed data
- Reload seed data - this will drop all tables and recreate

#### Needed endpoints

- POST /settings/reset_history
- POST /settings/full_reset

