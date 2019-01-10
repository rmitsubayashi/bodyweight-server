# API Docs
Responses follow the [JSON API specifications](https://jsonapi.org/)

## Exercise
**id** int

**title** string

**level** int

**image_url** string

**description** string

**[measurement_type](#measurement-type)** string

**category_id** int

**target_sets** Array(Set)

**quantity** int //-1 = infinite

### GET /users/exercises/category_id=??
returns Array(Exercise)

## ExerciseProduct
**title** string

**exercises** Array(Exercise)

**price** int

### GET /shop/exercises
returns Array(ExerciseProduct)

### POST /shop/exercises/purchase
returns points

## User
**id** int

**firebase_uid** string

**points** int

### GET /users/points
returns User.points

### POST /users

## Experience
**category_id** int

**level** int

**next_level_current** int

**next_level_total** int

### GET /users/experiences
returns Array(Experience)

## Log
**id** int

**category_id** int

**date** string //YYYY-MM-DD

**sets** Array([Set](#set)) 

**memo** string

//Todo: get logs by month

### GET /users/logs
returns Array(Log) without memo

### GET /users/logs/log_id
returns Log with memo

### POST /users/logs
returns Feedback

## Feedback
**comment** string

**comment_highlight_spans** Array([int, int]) //inclusive-exclusive

**previous_experience** Experience

**after_experience** Experience

**previous_points** int

**after_points** int

**unlocked_exercises** Array(UnlockedExercise)

**dropped_exercises** Array(ExerciseTitle)

## UnlockedExercise
**exercise** ExerciseTitle

**level_unlocked** int

**other_exercises** Array(ExerciseTitle)

## ExerciseTitle
**exercise_id** int

**title** string

## Set

**id** int

**exercise_id** int

**exercise_title** string

**[measurement_type](#measurement-type)** string

**set_number** int

**value** int

## Measurement Type
REP, SECONDS
