# API Docs

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

### GET /exercises/shop
returns Array(ExerciseProduct)

## User
**id** int

**firebase_uid** string

**points** int

## GET /user/points
returns User.points

## Experience
**category_id** int

**level** int

**next_level_current** int

**next_level_total** int

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

**previous_points** int

**after_points** int

**unlocked_exercises** Array(Exercise)

## Set

**id** int

**exercise_id** int

**exercise_title** string

**[measurement_type](#measurement-type)** string

**set_number** int

**value** int

## Measurement Type
REP, SECONDS
