# API Docs

## Exercise
**title** string

**level** int

**image_url** string

**description** string

**measurement_type** string

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
returns *experience* Array(Experience)

## Log
**id** int

**category** int

**sets** Array(Set) 

**memo** string

//Todo: get logs by month

### GET /user/logs
returns Array(Log) without memo

### GET /user/logs/log_id
returns Log with memo

### POST /logs
returns Experience, Feedback, optional Exercise (unlocakble)

## Feedback
**comment** string

**highlight_spans** Array([int, int]) //inclusive-exclusice

## Set
**measurement_type** string

**set_number** int

**value** int