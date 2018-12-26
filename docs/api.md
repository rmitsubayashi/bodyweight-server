# API Docs

## Exercise
**id** int

**title** string

**level** int

**image_url** string

**description** string

**[measurement_type](#measurement-type)** string

**category_id** int

### GET /users/exercises/category_id=??
returns Array(Exercise)

## User
**id** int

**firebase_uid** string

## Experience
**level** int

**next_level_total** int

**next_level_current** int

### GET /user/experiences
returns Array(Experience)

## Log
**id** int

**category_id** int

**date** string //YYYY-MM-DD

**sets** Array([Set](#set)) 

**memo** string

//Todo: get logs by month

### GET /user/logs
returns Array(Log) without memo

### GET /user/logs/log_id
returns Log with memo

### POST /logs
returns Feedback

## Feedback
**comment** string

**comment_highlight_spans** Array([int, int]) //inclusive-exclusive

**previous_experience** Experience

**after_experience** Experience

**unlocked_exercise** Exercise

## Set

**id** int

**exercise_id** int

**exercise_title** string

**[measurement_type](#measurement-type)** string

**set_number** int

**value** int

## Measurement Type
REP, SECONDS
